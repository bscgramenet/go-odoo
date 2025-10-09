package odoo

// ReportPosHrSingleEmployeeSalesReport represents report.pos_hr.single_employee_sales_report model.
type ReportPosHrSingleEmployeeSalesReport struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportPosHrSingleEmployeeSalesReports represents array of report.pos_hr.single_employee_sales_report model.
type ReportPosHrSingleEmployeeSalesReports []ReportPosHrSingleEmployeeSalesReport

// ReportPosHrSingleEmployeeSalesReportModel is the odoo model name.
const ReportPosHrSingleEmployeeSalesReportModel = "report.pos_hr.single_employee_sales_report"

// Many2One convert ReportPosHrSingleEmployeeSalesReport to *Many2One.
func (rps *ReportPosHrSingleEmployeeSalesReport) Many2One() *Many2One {
	return NewMany2One(rps.Id.Get(), "")
}

// CreateReportPosHrSingleEmployeeSalesReport creates a new report.pos_hr.single_employee_sales_report model and returns its id.
func (c *Client) CreateReportPosHrSingleEmployeeSalesReport(rps *ReportPosHrSingleEmployeeSalesReport) (int64, error) {
	ids, err := c.CreateReportPosHrSingleEmployeeSalesReports([]*ReportPosHrSingleEmployeeSalesReport{rps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportPosHrSingleEmployeeSalesReport creates a new report.pos_hr.single_employee_sales_report model and returns its id.
func (c *Client) CreateReportPosHrSingleEmployeeSalesReports(rpss []*ReportPosHrSingleEmployeeSalesReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range rpss {
		vv = append(vv, v)
	}
	return c.Create(ReportPosHrSingleEmployeeSalesReportModel, vv, nil)
}

// UpdateReportPosHrSingleEmployeeSalesReport updates an existing report.pos_hr.single_employee_sales_report record.
func (c *Client) UpdateReportPosHrSingleEmployeeSalesReport(rps *ReportPosHrSingleEmployeeSalesReport) error {
	return c.UpdateReportPosHrSingleEmployeeSalesReports([]int64{rps.Id.Get()}, rps)
}

// UpdateReportPosHrSingleEmployeeSalesReports updates existing report.pos_hr.single_employee_sales_report records.
// All records (represented by ids) will be updated by rps values.
func (c *Client) UpdateReportPosHrSingleEmployeeSalesReports(ids []int64, rps *ReportPosHrSingleEmployeeSalesReport) error {
	return c.Update(ReportPosHrSingleEmployeeSalesReportModel, ids, rps, nil)
}

// DeleteReportPosHrSingleEmployeeSalesReport deletes an existing report.pos_hr.single_employee_sales_report record.
func (c *Client) DeleteReportPosHrSingleEmployeeSalesReport(id int64) error {
	return c.DeleteReportPosHrSingleEmployeeSalesReports([]int64{id})
}

// DeleteReportPosHrSingleEmployeeSalesReports deletes existing report.pos_hr.single_employee_sales_report records.
func (c *Client) DeleteReportPosHrSingleEmployeeSalesReports(ids []int64) error {
	return c.Delete(ReportPosHrSingleEmployeeSalesReportModel, ids)
}

// GetReportPosHrSingleEmployeeSalesReport gets report.pos_hr.single_employee_sales_report existing record.
func (c *Client) GetReportPosHrSingleEmployeeSalesReport(id int64) (*ReportPosHrSingleEmployeeSalesReport, error) {
	rpss, err := c.GetReportPosHrSingleEmployeeSalesReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rpss)[0]), nil
}

// GetReportPosHrSingleEmployeeSalesReports gets report.pos_hr.single_employee_sales_report existing records.
func (c *Client) GetReportPosHrSingleEmployeeSalesReports(ids []int64) (*ReportPosHrSingleEmployeeSalesReports, error) {
	rpss := &ReportPosHrSingleEmployeeSalesReports{}
	if err := c.Read(ReportPosHrSingleEmployeeSalesReportModel, ids, nil, rpss); err != nil {
		return nil, err
	}
	return rpss, nil
}

// FindReportPosHrSingleEmployeeSalesReport finds report.pos_hr.single_employee_sales_report record by querying it with criteria.
func (c *Client) FindReportPosHrSingleEmployeeSalesReport(criteria *Criteria) (*ReportPosHrSingleEmployeeSalesReport, error) {
	rpss := &ReportPosHrSingleEmployeeSalesReports{}
	if err := c.SearchRead(ReportPosHrSingleEmployeeSalesReportModel, criteria, NewOptions().Limit(1), rpss); err != nil {
		return nil, err
	}
	return &((*rpss)[0]), nil
}

// FindReportPosHrSingleEmployeeSalesReports finds report.pos_hr.single_employee_sales_report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPosHrSingleEmployeeSalesReports(criteria *Criteria, options *Options) (*ReportPosHrSingleEmployeeSalesReports, error) {
	rpss := &ReportPosHrSingleEmployeeSalesReports{}
	if err := c.SearchRead(ReportPosHrSingleEmployeeSalesReportModel, criteria, options, rpss); err != nil {
		return nil, err
	}
	return rpss, nil
}

// FindReportPosHrSingleEmployeeSalesReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPosHrSingleEmployeeSalesReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportPosHrSingleEmployeeSalesReportModel, criteria, options)
}

// FindReportPosHrSingleEmployeeSalesReportId finds record id by querying it with criteria.
func (c *Client) FindReportPosHrSingleEmployeeSalesReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportPosHrSingleEmployeeSalesReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
