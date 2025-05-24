package odoo

// ComplianceLetterWizard represents compliance.letter.wizard model.
type ComplianceLetterWizard struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ComplianceLetterWizards represents array of compliance.letter.wizard model.
type ComplianceLetterWizards []ComplianceLetterWizard

// ComplianceLetterWizardModel is the odoo model name.
const ComplianceLetterWizardModel = "compliance.letter.wizard"

// Many2One convert ComplianceLetterWizard to *Many2One.
func (clw *ComplianceLetterWizard) Many2One() *Many2One {
	return NewMany2One(clw.Id.Get(), "")
}

// CreateComplianceLetterWizard creates a new compliance.letter.wizard model and returns its id.
func (c *Client) CreateComplianceLetterWizard(clw *ComplianceLetterWizard) (int64, error) {
	ids, err := c.CreateComplianceLetterWizards([]*ComplianceLetterWizard{clw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateComplianceLetterWizard creates a new compliance.letter.wizard model and returns its id.
func (c *Client) CreateComplianceLetterWizards(clws []*ComplianceLetterWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range clws {
		vv = append(vv, v)
	}
	return c.Create(ComplianceLetterWizardModel, vv, nil)
}

// UpdateComplianceLetterWizard updates an existing compliance.letter.wizard record.
func (c *Client) UpdateComplianceLetterWizard(clw *ComplianceLetterWizard) error {
	return c.UpdateComplianceLetterWizards([]int64{clw.Id.Get()}, clw)
}

// UpdateComplianceLetterWizards updates existing compliance.letter.wizard records.
// All records (represented by ids) will be updated by clw values.
func (c *Client) UpdateComplianceLetterWizards(ids []int64, clw *ComplianceLetterWizard) error {
	return c.Update(ComplianceLetterWizardModel, ids, clw, nil)
}

// DeleteComplianceLetterWizard deletes an existing compliance.letter.wizard record.
func (c *Client) DeleteComplianceLetterWizard(id int64) error {
	return c.DeleteComplianceLetterWizards([]int64{id})
}

// DeleteComplianceLetterWizards deletes existing compliance.letter.wizard records.
func (c *Client) DeleteComplianceLetterWizards(ids []int64) error {
	return c.Delete(ComplianceLetterWizardModel, ids)
}

// GetComplianceLetterWizard gets compliance.letter.wizard existing record.
func (c *Client) GetComplianceLetterWizard(id int64) (*ComplianceLetterWizard, error) {
	clws, err := c.GetComplianceLetterWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*clws)[0]), nil
}

// GetComplianceLetterWizards gets compliance.letter.wizard existing records.
func (c *Client) GetComplianceLetterWizards(ids []int64) (*ComplianceLetterWizards, error) {
	clws := &ComplianceLetterWizards{}
	if err := c.Read(ComplianceLetterWizardModel, ids, nil, clws); err != nil {
		return nil, err
	}
	return clws, nil
}

// FindComplianceLetterWizard finds compliance.letter.wizard record by querying it with criteria.
func (c *Client) FindComplianceLetterWizard(criteria *Criteria) (*ComplianceLetterWizard, error) {
	clws := &ComplianceLetterWizards{}
	if err := c.SearchRead(ComplianceLetterWizardModel, criteria, NewOptions().Limit(1), clws); err != nil {
		return nil, err
	}
	return &((*clws)[0]), nil
}

// FindComplianceLetterWizards finds compliance.letter.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindComplianceLetterWizards(criteria *Criteria, options *Options) (*ComplianceLetterWizards, error) {
	clws := &ComplianceLetterWizards{}
	if err := c.SearchRead(ComplianceLetterWizardModel, criteria, options, clws); err != nil {
		return nil, err
	}
	return clws, nil
}

// FindComplianceLetterWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindComplianceLetterWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ComplianceLetterWizardModel, criteria, options)
}

// FindComplianceLetterWizardId finds record id by querying it with criteria.
func (c *Client) FindComplianceLetterWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ComplianceLetterWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
