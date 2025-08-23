package odoo

// PosSelfOrderCustomLink represents pos_self_order.custom_link model.
type PosSelfOrderCustomLink struct {
	CreateDate   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String    `xmlrpc:"display_name,omitempty"`
	Id           *Int       `xmlrpc:"id,omitempty"`
	LinkHtml     *String    `xmlrpc:"link_html,omitempty"`
	Name         *String    `xmlrpc:"name,omitempty"`
	PosConfigIds *Relation  `xmlrpc:"pos_config_ids,omitempty"`
	Sequence     *Int       `xmlrpc:"sequence,omitempty"`
	Style        *Selection `xmlrpc:"style,omitempty"`
	Url          *String    `xmlrpc:"url,omitempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PosSelfOrderCustomLinks represents array of pos_self_order.custom_link model.
type PosSelfOrderCustomLinks []PosSelfOrderCustomLink

// PosSelfOrderCustomLinkModel is the odoo model name.
const PosSelfOrderCustomLinkModel = "pos_self_order.custom_link"

// Many2One convert PosSelfOrderCustomLink to *Many2One.
func (pc *PosSelfOrderCustomLink) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreatePosSelfOrderCustomLink creates a new pos_self_order.custom_link model and returns its id.
func (c *Client) CreatePosSelfOrderCustomLink(pc *PosSelfOrderCustomLink) (int64, error) {
	ids, err := c.CreatePosSelfOrderCustomLinks([]*PosSelfOrderCustomLink{pc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosSelfOrderCustomLink creates a new pos_self_order.custom_link model and returns its id.
func (c *Client) CreatePosSelfOrderCustomLinks(pcs []*PosSelfOrderCustomLink) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcs {
		vv = append(vv, v)
	}
	return c.Create(PosSelfOrderCustomLinkModel, vv, nil)
}

// UpdatePosSelfOrderCustomLink updates an existing pos_self_order.custom_link record.
func (c *Client) UpdatePosSelfOrderCustomLink(pc *PosSelfOrderCustomLink) error {
	return c.UpdatePosSelfOrderCustomLinks([]int64{pc.Id.Get()}, pc)
}

// UpdatePosSelfOrderCustomLinks updates existing pos_self_order.custom_link records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdatePosSelfOrderCustomLinks(ids []int64, pc *PosSelfOrderCustomLink) error {
	return c.Update(PosSelfOrderCustomLinkModel, ids, pc, nil)
}

// DeletePosSelfOrderCustomLink deletes an existing pos_self_order.custom_link record.
func (c *Client) DeletePosSelfOrderCustomLink(id int64) error {
	return c.DeletePosSelfOrderCustomLinks([]int64{id})
}

// DeletePosSelfOrderCustomLinks deletes existing pos_self_order.custom_link records.
func (c *Client) DeletePosSelfOrderCustomLinks(ids []int64) error {
	return c.Delete(PosSelfOrderCustomLinkModel, ids)
}

// GetPosSelfOrderCustomLink gets pos_self_order.custom_link existing record.
func (c *Client) GetPosSelfOrderCustomLink(id int64) (*PosSelfOrderCustomLink, error) {
	pcs, err := c.GetPosSelfOrderCustomLinks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// GetPosSelfOrderCustomLinks gets pos_self_order.custom_link existing records.
func (c *Client) GetPosSelfOrderCustomLinks(ids []int64) (*PosSelfOrderCustomLinks, error) {
	pcs := &PosSelfOrderCustomLinks{}
	if err := c.Read(PosSelfOrderCustomLinkModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosSelfOrderCustomLink finds pos_self_order.custom_link record by querying it with criteria.
func (c *Client) FindPosSelfOrderCustomLink(criteria *Criteria) (*PosSelfOrderCustomLink, error) {
	pcs := &PosSelfOrderCustomLinks{}
	if err := c.SearchRead(PosSelfOrderCustomLinkModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// FindPosSelfOrderCustomLinks finds pos_self_order.custom_link records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosSelfOrderCustomLinks(criteria *Criteria, options *Options) (*PosSelfOrderCustomLinks, error) {
	pcs := &PosSelfOrderCustomLinks{}
	if err := c.SearchRead(PosSelfOrderCustomLinkModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosSelfOrderCustomLinkIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosSelfOrderCustomLinkIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosSelfOrderCustomLinkModel, criteria, options)
}

// FindPosSelfOrderCustomLinkId finds record id by querying it with criteria.
func (c *Client) FindPosSelfOrderCustomLinkId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosSelfOrderCustomLinkModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
