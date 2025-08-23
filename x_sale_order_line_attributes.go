package odoo

// XSaleOrderLineAttributes represents x.sale.order.line.attributes model.
type XSaleOrderLineAttributes struct {
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
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
	XAddedDate                  *Time      `xmlrpc:"x_added_date,omitempty"`
	XAttributeName              *String    `xmlrpc:"x_attribute_name,omitempty"`
	XAttributeType              *Selection `xmlrpc:"x_attribute_type,omitempty"`
	XAttributeValue             *String    `xmlrpc:"x_attribute_value,omitempty"`
	XCustomAttributeName        *String    `xmlrpc:"x_custom_attribute_name,omitempty"`
	XCustomAttributeValue       *String    `xmlrpc:"x_custom_attribute_value,omitempty"`
	XIsStandardAttribute        *Bool      `xmlrpc:"x_is_standard_attribute,omitempty"`
	XMeasurementUnit            *String    `xmlrpc:"x_measurement_unit,omitempty"`
	XNotes                      *String    `xmlrpc:"x_notes,omitempty"`
	XPartnerId                  *Many2One  `xmlrpc:"x_partner_id,omitempty"`
	XProductAttributeId         *Many2One  `xmlrpc:"x_product_attribute_id,omitempty"`
	XProductAttributeValueId    *Many2One  `xmlrpc:"x_product_attribute_value_id,omitempty"`
	XProductCategoryId          *Many2One  `xmlrpc:"x_product_category_id,omitempty"`
	XProductId                  *Many2One  `xmlrpc:"x_product_id,omitempty"`
	XSaleLineDetailId           *Many2One  `xmlrpc:"x_sale_line_detail_id,omitempty"`
	XSaleOrderId                *Many2One  `xmlrpc:"x_sale_order_id,omitempty"`
	XSaleOrderLineId            *Many2One  `xmlrpc:"x_sale_order_line_id,omitempty"`
	XSourceSystem               *String    `xmlrpc:"x_source_system,omitempty"`
	XVerified                   *Bool      `xmlrpc:"x_verified,omitempty"`
}

// XSaleOrderLineAttributess represents array of x.sale.order.line.attributes model.
type XSaleOrderLineAttributess []XSaleOrderLineAttributes

// XSaleOrderLineAttributesModel is the odoo model name.
const XSaleOrderLineAttributesModel = "x.sale.order.line.attributes"

// Many2One convert XSaleOrderLineAttributes to *Many2One.
func (xsola *XSaleOrderLineAttributes) Many2One() *Many2One {
	return NewMany2One(xsola.Id.Get(), "")
}

// CreateXSaleOrderLineAttributes creates a new x.sale.order.line.attributes model and returns its id.
func (c *Client) CreateXSaleOrderLineAttributes(xsola *XSaleOrderLineAttributes) (int64, error) {
	ids, err := c.CreateXSaleOrderLineAttributess([]*XSaleOrderLineAttributes{xsola})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXSaleOrderLineAttributes creates a new x.sale.order.line.attributes model and returns its id.
func (c *Client) CreateXSaleOrderLineAttributess(xsolas []*XSaleOrderLineAttributes) ([]int64, error) {
	var vv []interface{}
	for _, v := range xsolas {
		vv = append(vv, v)
	}
	return c.Create(XSaleOrderLineAttributesModel, vv, nil)
}

// UpdateXSaleOrderLineAttributes updates an existing x.sale.order.line.attributes record.
func (c *Client) UpdateXSaleOrderLineAttributes(xsola *XSaleOrderLineAttributes) error {
	return c.UpdateXSaleOrderLineAttributess([]int64{xsola.Id.Get()}, xsola)
}

// UpdateXSaleOrderLineAttributess updates existing x.sale.order.line.attributes records.
// All records (represented by ids) will be updated by xsola values.
func (c *Client) UpdateXSaleOrderLineAttributess(ids []int64, xsola *XSaleOrderLineAttributes) error {
	return c.Update(XSaleOrderLineAttributesModel, ids, xsola, nil)
}

// DeleteXSaleOrderLineAttributes deletes an existing x.sale.order.line.attributes record.
func (c *Client) DeleteXSaleOrderLineAttributes(id int64) error {
	return c.DeleteXSaleOrderLineAttributess([]int64{id})
}

// DeleteXSaleOrderLineAttributess deletes existing x.sale.order.line.attributes records.
func (c *Client) DeleteXSaleOrderLineAttributess(ids []int64) error {
	return c.Delete(XSaleOrderLineAttributesModel, ids)
}

// GetXSaleOrderLineAttributes gets x.sale.order.line.attributes existing record.
func (c *Client) GetXSaleOrderLineAttributes(id int64) (*XSaleOrderLineAttributes, error) {
	xsolas, err := c.GetXSaleOrderLineAttributess([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xsolas)[0]), nil
}

// GetXSaleOrderLineAttributess gets x.sale.order.line.attributes existing records.
func (c *Client) GetXSaleOrderLineAttributess(ids []int64) (*XSaleOrderLineAttributess, error) {
	xsolas := &XSaleOrderLineAttributess{}
	if err := c.Read(XSaleOrderLineAttributesModel, ids, nil, xsolas); err != nil {
		return nil, err
	}
	return xsolas, nil
}

// FindXSaleOrderLineAttributes finds x.sale.order.line.attributes record by querying it with criteria.
func (c *Client) FindXSaleOrderLineAttributes(criteria *Criteria) (*XSaleOrderLineAttributes, error) {
	xsolas := &XSaleOrderLineAttributess{}
	if err := c.SearchRead(XSaleOrderLineAttributesModel, criteria, NewOptions().Limit(1), xsolas); err != nil {
		return nil, err
	}
	return &((*xsolas)[0]), nil
}

// FindXSaleOrderLineAttributess finds x.sale.order.line.attributes records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXSaleOrderLineAttributess(criteria *Criteria, options *Options) (*XSaleOrderLineAttributess, error) {
	xsolas := &XSaleOrderLineAttributess{}
	if err := c.SearchRead(XSaleOrderLineAttributesModel, criteria, options, xsolas); err != nil {
		return nil, err
	}
	return xsolas, nil
}

// FindXSaleOrderLineAttributesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXSaleOrderLineAttributesIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XSaleOrderLineAttributesModel, criteria, options)
}

// FindXSaleOrderLineAttributesId finds record id by querying it with criteria.
func (c *Client) FindXSaleOrderLineAttributesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XSaleOrderLineAttributesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
