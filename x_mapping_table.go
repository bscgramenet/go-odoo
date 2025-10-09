package odoo

// XMappingTable represents x.mapping.table model.
type XMappingTable struct {
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
	XLastSyncDate               *Time      `xmlrpc:"x_last_sync_date,omitempty"`
	XLocalId                    *String    `xmlrpc:"x_local_id,omitempty"`
	XLocalTable                 *String    `xmlrpc:"x_local_table,omitempty"`
	XNotes                      *String    `xmlrpc:"x_notes,omitempty"`
	XOdooId                     *String    `xmlrpc:"x_odoo_id,omitempty"`
	XOdooModel                  *String    `xmlrpc:"x_odoo_model,omitempty"`
	XSyncDirection              *Selection `xmlrpc:"x_sync_direction,omitempty"`
	XSyncStatus                 *Selection `xmlrpc:"x_sync_status,omitempty"`
}

// XMappingTables represents array of x.mapping.table model.
type XMappingTables []XMappingTable

// XMappingTableModel is the odoo model name.
const XMappingTableModel = "x.mapping.table"

// Many2One convert XMappingTable to *Many2One.
func (xmt *XMappingTable) Many2One() *Many2One {
	return NewMany2One(xmt.Id.Get(), "")
}

// CreateXMappingTable creates a new x.mapping.table model and returns its id.
func (c *Client) CreateXMappingTable(xmt *XMappingTable) (int64, error) {
	ids, err := c.CreateXMappingTables([]*XMappingTable{xmt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXMappingTable creates a new x.mapping.table model and returns its id.
func (c *Client) CreateXMappingTables(xmts []*XMappingTable) ([]int64, error) {
	var vv []interface{}
	for _, v := range xmts {
		vv = append(vv, v)
	}
	return c.Create(XMappingTableModel, vv, nil)
}

// UpdateXMappingTable updates an existing x.mapping.table record.
func (c *Client) UpdateXMappingTable(xmt *XMappingTable) error {
	return c.UpdateXMappingTables([]int64{xmt.Id.Get()}, xmt)
}

// UpdateXMappingTables updates existing x.mapping.table records.
// All records (represented by ids) will be updated by xmt values.
func (c *Client) UpdateXMappingTables(ids []int64, xmt *XMappingTable) error {
	return c.Update(XMappingTableModel, ids, xmt, nil)
}

// DeleteXMappingTable deletes an existing x.mapping.table record.
func (c *Client) DeleteXMappingTable(id int64) error {
	return c.DeleteXMappingTables([]int64{id})
}

// DeleteXMappingTables deletes existing x.mapping.table records.
func (c *Client) DeleteXMappingTables(ids []int64) error {
	return c.Delete(XMappingTableModel, ids)
}

// GetXMappingTable gets x.mapping.table existing record.
func (c *Client) GetXMappingTable(id int64) (*XMappingTable, error) {
	xmts, err := c.GetXMappingTables([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xmts)[0]), nil
}

// GetXMappingTables gets x.mapping.table existing records.
func (c *Client) GetXMappingTables(ids []int64) (*XMappingTables, error) {
	xmts := &XMappingTables{}
	if err := c.Read(XMappingTableModel, ids, nil, xmts); err != nil {
		return nil, err
	}
	return xmts, nil
}

// FindXMappingTable finds x.mapping.table record by querying it with criteria.
func (c *Client) FindXMappingTable(criteria *Criteria) (*XMappingTable, error) {
	xmts := &XMappingTables{}
	if err := c.SearchRead(XMappingTableModel, criteria, NewOptions().Limit(1), xmts); err != nil {
		return nil, err
	}
	return &((*xmts)[0]), nil
}

// FindXMappingTables finds x.mapping.table records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXMappingTables(criteria *Criteria, options *Options) (*XMappingTables, error) {
	xmts := &XMappingTables{}
	if err := c.SearchRead(XMappingTableModel, criteria, options, xmts); err != nil {
		return nil, err
	}
	return xmts, nil
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
