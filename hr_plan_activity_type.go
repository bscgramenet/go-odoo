package odoo

// HrPlanActivityType represents hr.plan.activity.type model.
type HrPlanActivityType struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityTypeId *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	Note           *String    `xmlrpc:"note,omitempty"`
	PlanId         *Many2One  `xmlrpc:"plan_id,omitempty"`
	Responsible    *Selection `xmlrpc:"responsible,omitempty"`
	ResponsibleId  *Many2One  `xmlrpc:"responsible_id,omitempty"`
	Summary        *String    `xmlrpc:"summary,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrPlanActivityTypes represents array of hr.plan.activity.type model.
type HrPlanActivityTypes []HrPlanActivityType

// HrPlanActivityTypeModel is the odoo model name.
const HrPlanActivityTypeModel = "hr.plan.activity.type"

// Many2One convert HrPlanActivityType to *Many2One.
func (hpat *HrPlanActivityType) Many2One() *Many2One {
	return NewMany2One(hpat.Id.Get(), "")
}

// CreateHrPlanActivityType creates a new hr.plan.activity.type model and returns its id.
func (c *Client) CreateHrPlanActivityType(hpat *HrPlanActivityType) (int64, error) {
	ids, err := c.CreateHrPlanActivityTypes([]*HrPlanActivityType{hpat})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPlanActivityType creates a new hr.plan.activity.type model and returns its id.
func (c *Client) CreateHrPlanActivityTypes(hpats []*HrPlanActivityType) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpats {
		vv = append(vv, v)
	}
	return c.Create(HrPlanActivityTypeModel, vv, nil)
}

// UpdateHrPlanActivityType updates an existing hr.plan.activity.type record.
func (c *Client) UpdateHrPlanActivityType(hpat *HrPlanActivityType) error {
	return c.UpdateHrPlanActivityTypes([]int64{hpat.Id.Get()}, hpat)
}

// UpdateHrPlanActivityTypes updates existing hr.plan.activity.type records.
// All records (represented by ids) will be updated by hpat values.
func (c *Client) UpdateHrPlanActivityTypes(ids []int64, hpat *HrPlanActivityType) error {
	return c.Update(HrPlanActivityTypeModel, ids, hpat, nil)
}

// DeleteHrPlanActivityType deletes an existing hr.plan.activity.type record.
func (c *Client) DeleteHrPlanActivityType(id int64) error {
	return c.DeleteHrPlanActivityTypes([]int64{id})
}

// DeleteHrPlanActivityTypes deletes existing hr.plan.activity.type records.
func (c *Client) DeleteHrPlanActivityTypes(ids []int64) error {
	return c.Delete(HrPlanActivityTypeModel, ids)
}

// GetHrPlanActivityType gets hr.plan.activity.type existing record.
func (c *Client) GetHrPlanActivityType(id int64) (*HrPlanActivityType, error) {
	hpats, err := c.GetHrPlanActivityTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hpats)[0]), nil
}

// GetHrPlanActivityTypes gets hr.plan.activity.type existing records.
func (c *Client) GetHrPlanActivityTypes(ids []int64) (*HrPlanActivityTypes, error) {
	hpats := &HrPlanActivityTypes{}
	if err := c.Read(HrPlanActivityTypeModel, ids, nil, hpats); err != nil {
		return nil, err
	}
	return hpats, nil
}

// FindHrPlanActivityType finds hr.plan.activity.type record by querying it with criteria.
func (c *Client) FindHrPlanActivityType(criteria *Criteria) (*HrPlanActivityType, error) {
	hpats := &HrPlanActivityTypes{}
	if err := c.SearchRead(HrPlanActivityTypeModel, criteria, NewOptions().Limit(1), hpats); err != nil {
		return nil, err
	}
	return &((*hpats)[0]), nil
}

// FindHrPlanActivityTypes finds hr.plan.activity.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPlanActivityTypes(criteria *Criteria, options *Options) (*HrPlanActivityTypes, error) {
	hpats := &HrPlanActivityTypes{}
	if err := c.SearchRead(HrPlanActivityTypeModel, criteria, options, hpats); err != nil {
		return nil, err
	}
	return hpats, nil
}

// FindHrPlanActivityTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPlanActivityTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrPlanActivityTypeModel, criteria, options)
}

// FindHrPlanActivityTypeId finds record id by querying it with criteria.
func (c *Client) FindHrPlanActivityTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPlanActivityTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
