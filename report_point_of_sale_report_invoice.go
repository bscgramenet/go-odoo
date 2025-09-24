package odoo

// ReportPointOfSaleReportInvoice represents report.point_of_sale.report_invoice model.
type ReportPointOfSaleReportInvoice struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportPointOfSaleReportInvoices represents array of report.point_of_sale.report_invoice model.
type ReportPointOfSaleReportInvoices []ReportPointOfSaleReportInvoice

// ReportPointOfSaleReportInvoiceModel is the odoo model name.
const ReportPointOfSaleReportInvoiceModel = "report.point_of_sale.report_invoice"

// Many2One convert ReportPointOfSaleReportInvoice to *Many2One.
func (rpr *ReportPointOfSaleReportInvoice) Many2One() *Many2One {
	return NewMany2One(rpr.Id.Get(), "")
}

// CreateReportPointOfSaleReportInvoice creates a new report.point_of_sale.report_invoice model and returns its id.
func (c *Client) CreateReportPointOfSaleReportInvoice(rpr *ReportPointOfSaleReportInvoice) (int64, error) {
	ids, err := c.CreateReportPointOfSaleReportInvoices([]*ReportPointOfSaleReportInvoice{rpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportPointOfSaleReportInvoice creates a new report.point_of_sale.report_invoice model and returns its id.
func (c *Client) CreateReportPointOfSaleReportInvoices(rprs []*ReportPointOfSaleReportInvoice) ([]int64, error) {
	var vv []interface{}
	for _, v := range rprs {
		vv = append(vv, v)
	}
	return c.Create(ReportPointOfSaleReportInvoiceModel, vv, nil)
}

// UpdateReportPointOfSaleReportInvoice updates an existing report.point_of_sale.report_invoice record.
func (c *Client) UpdateReportPointOfSaleReportInvoice(rpr *ReportPointOfSaleReportInvoice) error {
	return c.UpdateReportPointOfSaleReportInvoices([]int64{rpr.Id.Get()}, rpr)
}

// UpdateReportPointOfSaleReportInvoices updates existing report.point_of_sale.report_invoice records.
// All records (represented by ids) will be updated by rpr values.
func (c *Client) UpdateReportPointOfSaleReportInvoices(ids []int64, rpr *ReportPointOfSaleReportInvoice) error {
	return c.Update(ReportPointOfSaleReportInvoiceModel, ids, rpr, nil)
}

// DeleteReportPointOfSaleReportInvoice deletes an existing report.point_of_sale.report_invoice record.
func (c *Client) DeleteReportPointOfSaleReportInvoice(id int64) error {
	return c.DeleteReportPointOfSaleReportInvoices([]int64{id})
}

// DeleteReportPointOfSaleReportInvoices deletes existing report.point_of_sale.report_invoice records.
func (c *Client) DeleteReportPointOfSaleReportInvoices(ids []int64) error {
	return c.Delete(ReportPointOfSaleReportInvoiceModel, ids)
}

// GetReportPointOfSaleReportInvoice gets report.point_of_sale.report_invoice existing record.
func (c *Client) GetReportPointOfSaleReportInvoice(id int64) (*ReportPointOfSaleReportInvoice, error) {
	rprs, err := c.GetReportPointOfSaleReportInvoices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// GetReportPointOfSaleReportInvoices gets report.point_of_sale.report_invoice existing records.
func (c *Client) GetReportPointOfSaleReportInvoices(ids []int64) (*ReportPointOfSaleReportInvoices, error) {
	rprs := &ReportPointOfSaleReportInvoices{}
	if err := c.Read(ReportPointOfSaleReportInvoiceModel, ids, nil, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportPointOfSaleReportInvoice finds report.point_of_sale.report_invoice record by querying it with criteria.
func (c *Client) FindReportPointOfSaleReportInvoice(criteria *Criteria) (*ReportPointOfSaleReportInvoice, error) {
	rprs := &ReportPointOfSaleReportInvoices{}
	if err := c.SearchRead(ReportPointOfSaleReportInvoiceModel, criteria, NewOptions().Limit(1), rprs); err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// FindReportPointOfSaleReportInvoices finds report.point_of_sale.report_invoice records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPointOfSaleReportInvoices(criteria *Criteria, options *Options) (*ReportPointOfSaleReportInvoices, error) {
	rprs := &ReportPointOfSaleReportInvoices{}
	if err := c.SearchRead(ReportPointOfSaleReportInvoiceModel, criteria, options, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportPointOfSaleReportInvoiceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPointOfSaleReportInvoiceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportPointOfSaleReportInvoiceModel, criteria, options)
}

// FindReportPointOfSaleReportInvoiceId finds record id by querying it with criteria.
func (c *Client) FindReportPointOfSaleReportInvoiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportPointOfSaleReportInvoiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
