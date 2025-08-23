package odoo

// L10NEsEdiFacturaeAcRoleType represents l10n_es_edi_facturae.ac_role_type model.
type L10NEsEdiFacturaeAcRoleType struct {
	Code        *String   `xmlrpc:"code,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// L10NEsEdiFacturaeAcRoleTypes represents array of l10n_es_edi_facturae.ac_role_type model.
type L10NEsEdiFacturaeAcRoleTypes []L10NEsEdiFacturaeAcRoleType

// L10NEsEdiFacturaeAcRoleTypeModel is the odoo model name.
const L10NEsEdiFacturaeAcRoleTypeModel = "l10n_es_edi_facturae.ac_role_type"

// Many2One convert L10NEsEdiFacturaeAcRoleType to *Many2One.
func (la *L10NEsEdiFacturaeAcRoleType) Many2One() *Many2One {
	return NewMany2One(la.Id.Get(), "")
}

// CreateL10NEsEdiFacturaeAcRoleType creates a new l10n_es_edi_facturae.ac_role_type model and returns its id.
func (c *Client) CreateL10NEsEdiFacturaeAcRoleType(la *L10NEsEdiFacturaeAcRoleType) (int64, error) {
	ids, err := c.CreateL10NEsEdiFacturaeAcRoleTypes([]*L10NEsEdiFacturaeAcRoleType{la})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NEsEdiFacturaeAcRoleType creates a new l10n_es_edi_facturae.ac_role_type model and returns its id.
func (c *Client) CreateL10NEsEdiFacturaeAcRoleTypes(las []*L10NEsEdiFacturaeAcRoleType) ([]int64, error) {
	var vv []interface{}
	for _, v := range las {
		vv = append(vv, v)
	}
	return c.Create(L10NEsEdiFacturaeAcRoleTypeModel, vv, nil)
}

// UpdateL10NEsEdiFacturaeAcRoleType updates an existing l10n_es_edi_facturae.ac_role_type record.
func (c *Client) UpdateL10NEsEdiFacturaeAcRoleType(la *L10NEsEdiFacturaeAcRoleType) error {
	return c.UpdateL10NEsEdiFacturaeAcRoleTypes([]int64{la.Id.Get()}, la)
}

// UpdateL10NEsEdiFacturaeAcRoleTypes updates existing l10n_es_edi_facturae.ac_role_type records.
// All records (represented by ids) will be updated by la values.
func (c *Client) UpdateL10NEsEdiFacturaeAcRoleTypes(ids []int64, la *L10NEsEdiFacturaeAcRoleType) error {
	return c.Update(L10NEsEdiFacturaeAcRoleTypeModel, ids, la, nil)
}

// DeleteL10NEsEdiFacturaeAcRoleType deletes an existing l10n_es_edi_facturae.ac_role_type record.
func (c *Client) DeleteL10NEsEdiFacturaeAcRoleType(id int64) error {
	return c.DeleteL10NEsEdiFacturaeAcRoleTypes([]int64{id})
}

// DeleteL10NEsEdiFacturaeAcRoleTypes deletes existing l10n_es_edi_facturae.ac_role_type records.
func (c *Client) DeleteL10NEsEdiFacturaeAcRoleTypes(ids []int64) error {
	return c.Delete(L10NEsEdiFacturaeAcRoleTypeModel, ids)
}

// GetL10NEsEdiFacturaeAcRoleType gets l10n_es_edi_facturae.ac_role_type existing record.
func (c *Client) GetL10NEsEdiFacturaeAcRoleType(id int64) (*L10NEsEdiFacturaeAcRoleType, error) {
	las, err := c.GetL10NEsEdiFacturaeAcRoleTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*las)[0]), nil
}

// GetL10NEsEdiFacturaeAcRoleTypes gets l10n_es_edi_facturae.ac_role_type existing records.
func (c *Client) GetL10NEsEdiFacturaeAcRoleTypes(ids []int64) (*L10NEsEdiFacturaeAcRoleTypes, error) {
	las := &L10NEsEdiFacturaeAcRoleTypes{}
	if err := c.Read(L10NEsEdiFacturaeAcRoleTypeModel, ids, nil, las); err != nil {
		return nil, err
	}
	return las, nil
}

// FindL10NEsEdiFacturaeAcRoleType finds l10n_es_edi_facturae.ac_role_type record by querying it with criteria.
func (c *Client) FindL10NEsEdiFacturaeAcRoleType(criteria *Criteria) (*L10NEsEdiFacturaeAcRoleType, error) {
	las := &L10NEsEdiFacturaeAcRoleTypes{}
	if err := c.SearchRead(L10NEsEdiFacturaeAcRoleTypeModel, criteria, NewOptions().Limit(1), las); err != nil {
		return nil, err
	}
	return &((*las)[0]), nil
}

// FindL10NEsEdiFacturaeAcRoleTypes finds l10n_es_edi_facturae.ac_role_type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NEsEdiFacturaeAcRoleTypes(criteria *Criteria, options *Options) (*L10NEsEdiFacturaeAcRoleTypes, error) {
	las := &L10NEsEdiFacturaeAcRoleTypes{}
	if err := c.SearchRead(L10NEsEdiFacturaeAcRoleTypeModel, criteria, options, las); err != nil {
		return nil, err
	}
	return las, nil
}

// FindL10NEsEdiFacturaeAcRoleTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NEsEdiFacturaeAcRoleTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NEsEdiFacturaeAcRoleTypeModel, criteria, options)
}

// FindL10NEsEdiFacturaeAcRoleTypeId finds record id by querying it with criteria.
func (c *Client) FindL10NEsEdiFacturaeAcRoleTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NEsEdiFacturaeAcRoleTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
