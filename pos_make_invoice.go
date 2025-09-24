package odoo

// PosMakeInvoice represents pos.make.invoice model.
type PosMakeInvoice struct {
	ConsolidatedBilling *Bool     `xmlrpc:"consolidated_billing,omitempty"`
	Count               *Int      `xmlrpc:"count,omitempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosMakeInvoices represents array of pos.make.invoice model.
type PosMakeInvoices []PosMakeInvoice

// PosMakeInvoiceModel is the odoo model name.
const PosMakeInvoiceModel = "pos.make.invoice"

// Many2One convert PosMakeInvoice to *Many2One.
func (pmi *PosMakeInvoice) Many2One() *Many2One {
	return NewMany2One(pmi.Id.Get(), "")
}

// CreatePosMakeInvoice creates a new pos.make.invoice model and returns its id.
func (c *Client) CreatePosMakeInvoice(pmi *PosMakeInvoice) (int64, error) {
	ids, err := c.CreatePosMakeInvoices([]*PosMakeInvoice{pmi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosMakeInvoice creates a new pos.make.invoice model and returns its id.
func (c *Client) CreatePosMakeInvoices(pmis []*PosMakeInvoice) ([]int64, error) {
	var vv []interface{}
	for _, v := range pmis {
		vv = append(vv, v)
	}
	return c.Create(PosMakeInvoiceModel, vv, nil)
}

// UpdatePosMakeInvoice updates an existing pos.make.invoice record.
func (c *Client) UpdatePosMakeInvoice(pmi *PosMakeInvoice) error {
	return c.UpdatePosMakeInvoices([]int64{pmi.Id.Get()}, pmi)
}

// UpdatePosMakeInvoices updates existing pos.make.invoice records.
// All records (represented by ids) will be updated by pmi values.
func (c *Client) UpdatePosMakeInvoices(ids []int64, pmi *PosMakeInvoice) error {
	return c.Update(PosMakeInvoiceModel, ids, pmi, nil)
}

// DeletePosMakeInvoice deletes an existing pos.make.invoice record.
func (c *Client) DeletePosMakeInvoice(id int64) error {
	return c.DeletePosMakeInvoices([]int64{id})
}

// DeletePosMakeInvoices deletes existing pos.make.invoice records.
func (c *Client) DeletePosMakeInvoices(ids []int64) error {
	return c.Delete(PosMakeInvoiceModel, ids)
}

// GetPosMakeInvoice gets pos.make.invoice existing record.
func (c *Client) GetPosMakeInvoice(id int64) (*PosMakeInvoice, error) {
	pmis, err := c.GetPosMakeInvoices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pmis)[0]), nil
}

// GetPosMakeInvoices gets pos.make.invoice existing records.
func (c *Client) GetPosMakeInvoices(ids []int64) (*PosMakeInvoices, error) {
	pmis := &PosMakeInvoices{}
	if err := c.Read(PosMakeInvoiceModel, ids, nil, pmis); err != nil {
		return nil, err
	}
	return pmis, nil
}

// FindPosMakeInvoice finds pos.make.invoice record by querying it with criteria.
func (c *Client) FindPosMakeInvoice(criteria *Criteria) (*PosMakeInvoice, error) {
	pmis := &PosMakeInvoices{}
	if err := c.SearchRead(PosMakeInvoiceModel, criteria, NewOptions().Limit(1), pmis); err != nil {
		return nil, err
	}
	return &((*pmis)[0]), nil
}

// FindPosMakeInvoices finds pos.make.invoice records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosMakeInvoices(criteria *Criteria, options *Options) (*PosMakeInvoices, error) {
	pmis := &PosMakeInvoices{}
	if err := c.SearchRead(PosMakeInvoiceModel, criteria, options, pmis); err != nil {
		return nil, err
	}
	return pmis, nil
}

// FindPosMakeInvoiceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosMakeInvoiceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosMakeInvoiceModel, criteria, options)
}

// FindPosMakeInvoiceId finds record id by querying it with criteria.
func (c *Client) FindPosMakeInvoiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosMakeInvoiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
