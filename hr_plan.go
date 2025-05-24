package odoo

// HrPlan represents hr.plan model.
type HrPlan struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omitempty"`
	Active              *Bool     `xmlrpc:"active,omitempty"`
	CompanyId           *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	DepartmentId        *Many2One `xmlrpc:"department_id,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	Name                *String   `xmlrpc:"name,omitempty"`
	PlanActivityTypeIds *Relation `xmlrpc:"plan_activity_type_ids,omitempty"`
	StepsCount          *Int      `xmlrpc:"steps_count,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrPlans represents array of hr.plan model.
type HrPlans []HrPlan

// HrPlanModel is the odoo model name.
const HrPlanModel = "hr.plan"

// Many2One convert HrPlan to *Many2One.
func (hp *HrPlan) Many2One() *Many2One {
	return NewMany2One(hp.Id.Get(), "")
}

// CreateHrPlan creates a new hr.plan model and returns its id.
func (c *Client) CreateHrPlan(hp *HrPlan) (int64, error) {
	ids, err := c.CreateHrPlans([]*HrPlan{hp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPlan creates a new hr.plan model and returns its id.
func (c *Client) CreateHrPlans(hps []*HrPlan) ([]int64, error) {
	var vv []interface{}
	for _, v := range hps {
		vv = append(vv, v)
	}
	return c.Create(HrPlanModel, vv, nil)
}

// UpdateHrPlan updates an existing hr.plan record.
func (c *Client) UpdateHrPlan(hp *HrPlan) error {
	return c.UpdateHrPlans([]int64{hp.Id.Get()}, hp)
}

// UpdateHrPlans updates existing hr.plan records.
// All records (represented by ids) will be updated by hp values.
func (c *Client) UpdateHrPlans(ids []int64, hp *HrPlan) error {
	return c.Update(HrPlanModel, ids, hp, nil)
}

// DeleteHrPlan deletes an existing hr.plan record.
func (c *Client) DeleteHrPlan(id int64) error {
	return c.DeleteHrPlans([]int64{id})
}

// DeleteHrPlans deletes existing hr.plan records.
func (c *Client) DeleteHrPlans(ids []int64) error {
	return c.Delete(HrPlanModel, ids)
}

// GetHrPlan gets hr.plan existing record.
func (c *Client) GetHrPlan(id int64) (*HrPlan, error) {
	hps, err := c.GetHrPlans([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hps)[0]), nil
}

// GetHrPlans gets hr.plan existing records.
func (c *Client) GetHrPlans(ids []int64) (*HrPlans, error) {
	hps := &HrPlans{}
	if err := c.Read(HrPlanModel, ids, nil, hps); err != nil {
		return nil, err
	}
	return hps, nil
}

// FindHrPlan finds hr.plan record by querying it with criteria.
func (c *Client) FindHrPlan(criteria *Criteria) (*HrPlan, error) {
	hps := &HrPlans{}
	if err := c.SearchRead(HrPlanModel, criteria, NewOptions().Limit(1), hps); err != nil {
		return nil, err
	}
	return &((*hps)[0]), nil
}

// FindHrPlans finds hr.plan records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPlans(criteria *Criteria, options *Options) (*HrPlans, error) {
	hps := &HrPlans{}
	if err := c.SearchRead(HrPlanModel, criteria, options, hps); err != nil {
		return nil, err
	}
	return hps, nil
}

// FindHrPlanIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPlanIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrPlanModel, criteria, options)
}

// FindHrPlanId finds record id by querying it with criteria.
func (c *Client) FindHrPlanId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPlanModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
