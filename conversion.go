package odoo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	dateFormat     = "2006-01-02"
	datetimeFormat = "2006-01-02 15:04:05"
	tagName        = "xmlrpc"
)

func convertFromStaticToDynamic(static interface{}) map[string]interface{} {
	var dynamic = make(map[string]interface{})
	sv := reflect.ValueOf(static).Elem()
	st := reflect.TypeOf(static).Elem()
	for i := 0; i < sv.NumField(); i++ {
		field := sv.Field(i)
		if field.IsNil() {
			continue
		}
		key, _ := st.Field(i).Tag.Lookup(tagName)
		if dynamicValue := convertFromStaticToDynamicValue(field.Interface()); dynamicValue != nil {
			dynamic[strings.Split(key, ",")[0]] = dynamicValue
		}
	}
	return dynamic
}

func convertFromStaticToDynamicValue(staticValue interface{}) interface{} {
	var v interface{}
	switch sv := staticValue.(type) {
	case *String:
		v = sv.v
	case *Int:
		v = sv.v
	case *Bool:
		v = sv.v
	case *Selection:
		v = sv.v
	case *Time:
		v = sv.v.Format(datetimeFormat)
	case *Float:
		v = sv.v
	case *Many2One:
		if sv.ID == 0 {
			v = false
		} else {
			v = sv.ID
		}
	case *Relation:
		v = sv.v
	default:
		v = staticValue
	}
	return v
}

func convertFromDynamicToStatic(dynamic interface{}, static interface{}) error {
	model := reflect.TypeOf(static).Elem()
	var sv reflect.Value
	switch d := dynamic.(type) {
	case []interface{}:
		if model.Kind() != reflect.Slice {
			return fmt.Errorf("cannot convert dynamic model to static model %s", model.Name())
		}
		sv = convertFromDynamicToStaticSlice(d, model)
	case map[string]interface{}:
		if model.Kind() == reflect.Slice {
			return fmt.Errorf("cannot convert dynamic model to static model %s", model.Name())
		}
		sv = convertFromDynamicToStaticOne(d, model)
	default:
		return errors.New("cannot convert dynamic of this type")
	}
	reflect.ValueOf(static).Elem().Set(sv)
	return nil
}

func convertFromDynamicToStaticSlice(dynamic []interface{}, sliceModel reflect.Type) reflect.Value {
	lenSlice := len(dynamic)
	ss := reflect.MakeSlice(sliceModel, lenSlice, lenSlice)
	for i := 0; i < lenSlice; i++ {
		ss.Index(i).Set(convertFromDynamicToStaticOne(dynamic[i].(map[string]interface{}), sliceModel.Elem()))
	}
	return ss
}

func convertFromDynamicToStaticOne(dynamic map[string]interface{}, oneModel reflect.Type) reflect.Value {
	s := reflect.New(oneModel).Elem()
	staticValues := scanStaticModelValues(oneModel, s)
	for key, dynamicValue := range dynamic {
		if _, ok := staticValues[key]; ok {
			staticField := staticValues[key]
			staticValue := convertFromDynamicToStaticValue(staticField.Type(), dynamicValue)
			if staticValue != nil {
				staticField.Set(reflect.ValueOf(staticValue))
			}
		}
	}
	return s
}

func convertFromDynamicToStaticValue(staticType reflect.Type, dynamicValue interface{}) interface{} {
	var staticValue interface{}
	if staticType.Kind() == reflect.Ptr {
		staticType = staticType.Elem()
	}
	typeName := staticType.Name()
	if !(dynamicValue == nil || (reflect.ValueOf(dynamicValue).Kind() == reflect.Bool && typeName != "Bool")) {
		switch typeName {
		case "String":
			if strVal, ok := dynamicValue.(string); ok {
				staticValue = NewString(strVal)
			} else {
				// We use "String" for Odoo Binary field type also, which is used to store binary data.
				// However, in rare cases (compute fields), this field might return "[]interface{}" instead of "[]byte".
				// @TODO: It's important to handle this scenario as well.
			}
		case "Int":
			switch v := dynamicValue.(type) {
			case int64:
				staticValue = NewInt(v)
			case float64:
				staticValue = NewInt(int64(v))
			case int:
				staticValue = NewInt(int64(v))
			case string:
				if iv, err := strconv.ParseInt(v, 10, 64); err == nil {
					staticValue = NewInt(iv)
				} else if debugEnabled() {
					logTypeMismatch("Int", fmt.Sprintf("string(%s)", v), "int64", dynamicValue)
				}
			default:
				// 其它类型可选处理
			}
		case "Selection":
			staticValue = NewSelection(dynamicValue)
		case "Float":
			// 兼容 Odoo 返回 float64 或 int64
			switch v := dynamicValue.(type) {
			case float64:
				staticValue = NewFloat(v)
			case int64:
				staticValue = NewFloat(float64(v))
			case int:
				staticValue = NewFloat(float64(v))
			case string:
				if fv, err := strconv.ParseFloat(v, 64); err == nil {
					staticValue = NewFloat(fv)
				} else if debugEnabled() {
					logTypeMismatch("Float", fmt.Sprintf("string(%s)", v), "float64", dynamicValue)
				}
			default:
				// 其它类型直接丢弃或报错
			}
		case "Time":
			if dynamicValue == nil || dynamicValue == false {
				break
			}
			str, ok := dynamicValue.(string)
			if !ok {
				if debugEnabled() {
					logTypeMismatch("Time", fmt.Sprintf("%T", dynamicValue), "string(yyyy-mm-dd[ hh:mm:ss])", dynamicValue)
				}
				break
			}
			str = strings.TrimSpace(str)
			if str == "" {
				break
			}
			format := dateFormat
			if len(str) > 10 {
				format = datetimeFormat
			}
			if t, err := time.Parse(format, str); err == nil {
				staticValue = NewTime(t)
			} else if debugEnabled() {
				logTypeMismatch("Time", str, format, dynamicValue)
			}
		case "Many2One":
			if dynamicValue == nil || dynamicValue == false {
				staticValue = NewMany2One(0, "")
				break
			}
			switch v := dynamicValue.(type) {
			case int64:
				staticValue = NewMany2One(v, "")
			case []interface{}:
				if len(v) > 0 {
					idVal, _ := v[0].(int64)
					nameVal := ""
					if len(v) > 1 {
						nameVal, _ = v[1].(string)
					}
					staticValue = NewMany2One(idVal, nameVal)
				}
			default:
				if debugEnabled() {
					logTypeMismatch("Many2One", fmt.Sprintf("%T", dynamicValue), "int64|[id,name]", dynamicValue)
				}
			}
		case "Relation":
			rel := NewRelation()
			if dynamicValue == nil || dynamicValue == false {
				staticValue = rel
				break
			}
			if arr, ok := dynamicValue.([]interface{}); ok {
				rel.ids = sliceInterfaceToInt64Slice(arr)
				staticValue = rel
			} else if debugEnabled() {
				logTypeMismatch("Relation", fmt.Sprintf("%T", dynamicValue), "[]interface{}", dynamicValue)
			}
		case "Bool":
			if bv, ok := parseBoolLike(dynamicValue); ok {
				staticValue = NewBool(bv)
			} else {
				if debugEnabled() {
					logTypeMismatch("Bool", fmt.Sprintf("%T", dynamicValue), "bool", dynamicValue)
				}
				// 返回 nil 不设置该字段，保持兼容
			}
		default:
			staticValue = dynamicValue
		}
	}
	return staticValue
}

// ---- Debug & Parsing Helpers (added for Odoo 19 compatibility) ----

var (
	debugOnce    sync.Once
	mismatchMu   sync.Mutex
	mismatchSeen = map[string]struct{}{}
)

func debugEnabled() bool {
	v := os.Getenv("GO_ODOO_DEBUG_CONVERSION")
	if v == "" {
		return false
	}
	vv := strings.ToLower(strings.TrimSpace(v))
	switch vv {
	case "1", "true", "on", "yes", "y":
		return true
	default:
		return false
	}
}

// parseBoolLike attempts to interpret various dynamic forms as bool.
func parseBoolLike(v interface{}) (bool, bool) {
	switch b := v.(type) {
	case bool:
		return b, true
	case int:
		return b != 0, true
	case int64:
		return b != 0, true
	case float64:
		return b != 0, true
	case string:
		s := strings.ToLower(strings.TrimSpace(b))
		switch s {
		case "1", "true", "t", "y", "yes", "on":
			return true, true
		case "0", "false", "f", "n", "no", "off", "":
			return false, true
		default:
			return false, false
		}
	case []byte:
		s := strings.ToLower(strings.TrimSpace(string(b)))
		switch s {
		case "1", "true", "t", "y", "yes", "on":
			return true, true
		case "0", "false", "f", "n", "no", "off", "":
			return false, true
		default:
			return false, false
		}
	default:
		return false, false
	}
}

// logTypeMismatch prints a one-time debug line per unique key.
func logTypeMismatch(kind, gotRepr, expect string, raw interface{}) {
	key := kind + "|" + gotRepr + "|" + expect
	mismatchMu.Lock()
	if _, ok := mismatchSeen[key]; ok {
		mismatchMu.Unlock()
		return
	}
	mismatchSeen[key] = struct{}{}
	mismatchMu.Unlock()
	log.Printf("[go-odoo][debug] type mismatch converting %s expect=%s got=%s valueType=%T", kind, expect, gotRepr, raw)
}

func scanStaticModelValues(typ reflect.Type, s reflect.Value) map[string]reflect.Value {
	fields := make(map[string]reflect.Value)
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		key, _ := typ.Field(i).Tag.Lookup(tagName)
		fields[strings.Split(key, ",")[0]] = field
	}
	return fields
}

func sliceInterfaceToInt64Slice(s []interface{}) []int64 {
	i64 := make([]int64, len(s))
	for i := 0; i < len(s); i++ {
		i64[i] = s[i].(int64)
	}
	return i64
}
