package odoo

// ReportPointOfSaleReportSaledetails represents report.point_of_sale.report_saledetails model.
type ReportPointOfSaleReportSaledetails struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportPointOfSaleReportSaledetailss represents array of report.point_of_sale.report_saledetails model.
type ReportPointOfSaleReportSaledetailss []ReportPointOfSaleReportSaledetails

// ReportPointOfSaleReportSaledetailsModel is the odoo model name.
const ReportPointOfSaleReportSaledetailsModel = "report.point_of_sale.report_saledetails"

// Many2One convert ReportPointOfSaleReportSaledetails to *Many2One.
func (rpr *ReportPointOfSaleReportSaledetails) Many2One() *Many2One {
	return NewMany2One(rpr.Id.Get(), "")
}

// CreateReportPointOfSaleReportSaledetails creates a new report.point_of_sale.report_saledetails model and returns its id.
func (c *Client) CreateReportPointOfSaleReportSaledetails(rpr *ReportPointOfSaleReportSaledetails) (int64, error) {
	ids, err := c.CreateReportPointOfSaleReportSaledetailss([]*ReportPointOfSaleReportSaledetails{rpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportPointOfSaleReportSaledetails creates a new report.point_of_sale.report_saledetails model and returns its id.
func (c *Client) CreateReportPointOfSaleReportSaledetailss(rprs []*ReportPointOfSaleReportSaledetails) ([]int64, error) {
	var vv []interface{}
	for _, v := range rprs {
		vv = append(vv, v)
	}
	return c.Create(ReportPointOfSaleReportSaledetailsModel, vv, nil)
}

// UpdateReportPointOfSaleReportSaledetails updates an existing report.point_of_sale.report_saledetails record.
func (c *Client) UpdateReportPointOfSaleReportSaledetails(rpr *ReportPointOfSaleReportSaledetails) error {
	return c.UpdateReportPointOfSaleReportSaledetailss([]int64{rpr.Id.Get()}, rpr)
}

// UpdateReportPointOfSaleReportSaledetailss updates existing report.point_of_sale.report_saledetails records.
// All records (represented by ids) will be updated by rpr values.
func (c *Client) UpdateReportPointOfSaleReportSaledetailss(ids []int64, rpr *ReportPointOfSaleReportSaledetails) error {
	return c.Update(ReportPointOfSaleReportSaledetailsModel, ids, rpr, nil)
}

// DeleteReportPointOfSaleReportSaledetails deletes an existing report.point_of_sale.report_saledetails record.
func (c *Client) DeleteReportPointOfSaleReportSaledetails(id int64) error {
	return c.DeleteReportPointOfSaleReportSaledetailss([]int64{id})
}

// DeleteReportPointOfSaleReportSaledetailss deletes existing report.point_of_sale.report_saledetails records.
func (c *Client) DeleteReportPointOfSaleReportSaledetailss(ids []int64) error {
	return c.Delete(ReportPointOfSaleReportSaledetailsModel, ids)
}

// GetReportPointOfSaleReportSaledetails gets report.point_of_sale.report_saledetails existing record.
func (c *Client) GetReportPointOfSaleReportSaledetails(id int64) (*ReportPointOfSaleReportSaledetails, error) {
	rprs, err := c.GetReportPointOfSaleReportSaledetailss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// GetReportPointOfSaleReportSaledetailss gets report.point_of_sale.report_saledetails existing records.
func (c *Client) GetReportPointOfSaleReportSaledetailss(ids []int64) (*ReportPointOfSaleReportSaledetailss, error) {
	rprs := &ReportPointOfSaleReportSaledetailss{}
	if err := c.Read(ReportPointOfSaleReportSaledetailsModel, ids, nil, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportPointOfSaleReportSaledetails finds report.point_of_sale.report_saledetails record by querying it with criteria.
func (c *Client) FindReportPointOfSaleReportSaledetails(criteria *Criteria) (*ReportPointOfSaleReportSaledetails, error) {
	rprs := &ReportPointOfSaleReportSaledetailss{}
	if err := c.SearchRead(ReportPointOfSaleReportSaledetailsModel, criteria, NewOptions().Limit(1), rprs); err != nil {
		return nil, err
	}
	return &((*rprs)[0]), nil
}

// FindReportPointOfSaleReportSaledetailss finds report.point_of_sale.report_saledetails records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPointOfSaleReportSaledetailss(criteria *Criteria, options *Options) (*ReportPointOfSaleReportSaledetailss, error) {
	rprs := &ReportPointOfSaleReportSaledetailss{}
	if err := c.SearchRead(ReportPointOfSaleReportSaledetailsModel, criteria, options, rprs); err != nil {
		return nil, err
	}
	return rprs, nil
}

// FindReportPointOfSaleReportSaledetailsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPointOfSaleReportSaledetailsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportPointOfSaleReportSaledetailsModel, criteria, options)
}

// FindReportPointOfSaleReportSaledetailsId finds record id by querying it with criteria.
func (c *Client) FindReportPointOfSaleReportSaledetailsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportPointOfSaleReportSaledetailsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
