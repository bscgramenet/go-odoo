package odoo_test

import (
	"fmt"
	"testing"

	"github.com/bscgramenet/go-odoo"
)

func TestFieldsGet(t *testing.T) {
	client, err := odoo.NewClient(&odoo.ClientConfig{
		URL:      "http://localhost:8069",                    // 修改为你的 Odoo 地址
		Database: "odoo16",                                   // 修改为你的数据库名
		Admin:    "cchuayan@gmail.com",                       // 修改为你的用户名
		Password: "a2100483ce54237e0dca3dc2de7c1e4577e3bb3c", // 修改为你的密码
	})
	if err != nil {
		t.Fatalf("连接 Odoo 失败: %v", err)
	}

	fields, err := client.FieldsGet("account.move", nil)
	if err != nil {
		t.Fatalf("fields_get 失败: %v", err)
	}

	for k := range fields {
		fmt.Println(k)
	}
}
