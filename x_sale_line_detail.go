package odoo

// XSaleLineDetail represents x.sale.line.detail model.
type XSaleLineDetail struct {
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
	XAttributeIds               *Relation  `xmlrpc:"x_attribute_ids,omitempty"`
	XDescription                *String    `xmlrpc:"x_description,omitempty"`
	XOrderId                    *Many2One  `xmlrpc:"x_order_id,omitempty"`
	XPartnerId                  *Many2One  `xmlrpc:"x_partner_id,omitempty"`
	XProductId                  *Many2One  `xmlrpc:"x_product_id,omitempty"`
	XSaleOrderLineId            *Many2One  `xmlrpc:"x_sale_order_line_id,omitempty"`
	XSequence                   *Int       `xmlrpc:"x_sequence,omitempty"`
	XTotalAmount                *Float     `xmlrpc:"x_total_amount,omitempty"`
	XUnitPrice                  *Float     `xmlrpc:"x_unit_price,omitempty"`
	XVariantQuantity            *Float     `xmlrpc:"x_variant_quantity,omitempty"`
}

// XSaleLineDetails represents array of x.sale.line.detail model.
type XSaleLineDetails []XSaleLineDetail

// XSaleLineDetailModel is the odoo model name.
const XSaleLineDetailModel = "x.sale.line.detail"

// Many2One convert XSaleLineDetail to *Many2One.
func (xsld *XSaleLineDetail) Many2One() *Many2One {
	return NewMany2One(xsld.Id.Get(), "")
}

// CreateXSaleLineDetail creates a new x.sale.line.detail model and returns its id.
func (c *Client) CreateXSaleLineDetail(xsld *XSaleLineDetail) (int64, error) {
	ids, err := c.CreateXSaleLineDetails([]*XSaleLineDetail{xsld})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXSaleLineDetail creates a new x.sale.line.detail model and returns its id.
func (c *Client) CreateXSaleLineDetails(xslds []*XSaleLineDetail) ([]int64, error) {
	var vv []interface{}
	for _, v := range xslds {
		vv = append(vv, v)
	}
	return c.Create(XSaleLineDetailModel, vv, nil)
}

// UpdateXSaleLineDetail updates an existing x.sale.line.detail record.
func (c *Client) UpdateXSaleLineDetail(xsld *XSaleLineDetail) error {
	return c.UpdateXSaleLineDetails([]int64{xsld.Id.Get()}, xsld)
}

// UpdateXSaleLineDetails updates existing x.sale.line.detail records.
// All records (represented by ids) will be updated by xsld values.
func (c *Client) UpdateXSaleLineDetails(ids []int64, xsld *XSaleLineDetail) error {
	return c.Update(XSaleLineDetailModel, ids, xsld, nil)
}

// DeleteXSaleLineDetail deletes an existing x.sale.line.detail record.
func (c *Client) DeleteXSaleLineDetail(id int64) error {
	return c.DeleteXSaleLineDetails([]int64{id})
}

// DeleteXSaleLineDetails deletes existing x.sale.line.detail records.
func (c *Client) DeleteXSaleLineDetails(ids []int64) error {
	return c.Delete(XSaleLineDetailModel, ids)
}

// GetXSaleLineDetail gets x.sale.line.detail existing record.
func (c *Client) GetXSaleLineDetail(id int64) (*XSaleLineDetail, error) {
	xslds, err := c.GetXSaleLineDetails([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xslds)[0]), nil
}

// GetXSaleLineDetails gets x.sale.line.detail existing records.
func (c *Client) GetXSaleLineDetails(ids []int64) (*XSaleLineDetails, error) {
	xslds := &XSaleLineDetails{}
	if err := c.Read(XSaleLineDetailModel, ids, nil, xslds); err != nil {
		return nil, err
	}
	return xslds, nil
}

// FindXSaleLineDetail finds x.sale.line.detail record by querying it with criteria.
func (c *Client) FindXSaleLineDetail(criteria *Criteria) (*XSaleLineDetail, error) {
	xslds := &XSaleLineDetails{}
	if err := c.SearchRead(XSaleLineDetailModel, criteria, NewOptions().Limit(1), xslds); err != nil {
		return nil, err
	}
	return &((*xslds)[0]), nil
}

// FindXSaleLineDetails finds x.sale.line.detail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXSaleLineDetails(criteria *Criteria, options *Options) (*XSaleLineDetails, error) {
	xslds := &XSaleLineDetails{}
	if err := c.SearchRead(XSaleLineDetailModel, criteria, options, xslds); err != nil {
		return nil, err
	}
	return xslds, nil
}

// FindXSaleLineDetailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXSaleLineDetailIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XSaleLineDetailModel, criteria, options)
}

// FindXSaleLineDetailId finds record id by querying it with criteria.
func (c *Client) FindXSaleLineDetailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XSaleLineDetailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
