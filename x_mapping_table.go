package odoo

// XMappingTable represents x_mapping_table model.
type XMappingTable struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	XLocalId    *String   `xmlrpc:"x_local_id,omitempty"`
	XLocalTable *String   `xmlrpc:"x_local_table,omitempty"`
	XName       *String   `xmlrpc:"x_name,omitempty"`
	XOdooId     *String   `xmlrpc:"x_odoo_id,omitempty"`
	XOdooModel  *String   `xmlrpc:"x_odoo_model,omitempty"`
}

// XMappingTables represents array of x_mapping_table model.
type XMappingTables []XMappingTable

// XMappingTableModel is the odoo model name.
const XMappingTableModel = "x_mapping_table"

// Many2One convert XMappingTable to *Many2One.
func (x *XMappingTable) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXMappingTable creates a new x_mapping_table model and returns its id.
func (c *Client) CreateXMappingTable(x *XMappingTable) (int64, error) {
	ids, err := c.CreateXMappingTables([]*XMappingTable{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXMappingTable creates a new x_mapping_table model and returns its id.
func (c *Client) CreateXMappingTables(xs []*XMappingTable) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XMappingTableModel, vv, nil)
}

// UpdateXMappingTable updates an existing x_mapping_table record.
func (c *Client) UpdateXMappingTable(x *XMappingTable) error {
	return c.UpdateXMappingTables([]int64{x.Id.Get()}, x)
}

// UpdateXMappingTables updates existing x_mapping_table records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXMappingTables(ids []int64, x *XMappingTable) error {
	return c.Update(XMappingTableModel, ids, x, nil)
}

// DeleteXMappingTable deletes an existing x_mapping_table record.
func (c *Client) DeleteXMappingTable(id int64) error {
	return c.DeleteXMappingTables([]int64{id})
}

// DeleteXMappingTables deletes existing x_mapping_table records.
func (c *Client) DeleteXMappingTables(ids []int64) error {
	return c.Delete(XMappingTableModel, ids)
}

// GetXMappingTable gets x_mapping_table existing record.
func (c *Client) GetXMappingTable(id int64) (*XMappingTable, error) {
	xs, err := c.GetXMappingTables([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXMappingTables gets x_mapping_table existing records.
func (c *Client) GetXMappingTables(ids []int64) (*XMappingTables, error) {
	xs := &XMappingTables{}
	if err := c.Read(XMappingTableModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXMappingTable finds x_mapping_table record by querying it with criteria.
func (c *Client) FindXMappingTable(criteria *Criteria) (*XMappingTable, error) {
	xs := &XMappingTables{}
	if err := c.SearchRead(XMappingTableModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXMappingTables finds x_mapping_table records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXMappingTables(criteria *Criteria, options *Options) (*XMappingTables, error) {
	xs := &XMappingTables{}
	if err := c.SearchRead(XMappingTableModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXMappingTableIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXMappingTableIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XMappingTableModel, criteria, options)
}

// FindXMappingTableId finds record id by querying it with criteria.
func (c *Client) FindXMappingTableId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XMappingTableModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
