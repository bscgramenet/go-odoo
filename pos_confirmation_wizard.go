package odoo

// PosConfirmationWizard represents pos.confirmation.wizard model.
type PosConfirmationWizard struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Message     *String   `xmlrpc:"message,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosConfirmationWizards represents array of pos.confirmation.wizard model.
type PosConfirmationWizards []PosConfirmationWizard

// PosConfirmationWizardModel is the odoo model name.
const PosConfirmationWizardModel = "pos.confirmation.wizard"

// Many2One convert PosConfirmationWizard to *Many2One.
func (pcw *PosConfirmationWizard) Many2One() *Many2One {
	return NewMany2One(pcw.Id.Get(), "")
}

// CreatePosConfirmationWizard creates a new pos.confirmation.wizard model and returns its id.
func (c *Client) CreatePosConfirmationWizard(pcw *PosConfirmationWizard) (int64, error) {
	ids, err := c.CreatePosConfirmationWizards([]*PosConfirmationWizard{pcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosConfirmationWizard creates a new pos.confirmation.wizard model and returns its id.
func (c *Client) CreatePosConfirmationWizards(pcws []*PosConfirmationWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcws {
		vv = append(vv, v)
	}
	return c.Create(PosConfirmationWizardModel, vv, nil)
}

// UpdatePosConfirmationWizard updates an existing pos.confirmation.wizard record.
func (c *Client) UpdatePosConfirmationWizard(pcw *PosConfirmationWizard) error {
	return c.UpdatePosConfirmationWizards([]int64{pcw.Id.Get()}, pcw)
}

// UpdatePosConfirmationWizards updates existing pos.confirmation.wizard records.
// All records (represented by ids) will be updated by pcw values.
func (c *Client) UpdatePosConfirmationWizards(ids []int64, pcw *PosConfirmationWizard) error {
	return c.Update(PosConfirmationWizardModel, ids, pcw, nil)
}

// DeletePosConfirmationWizard deletes an existing pos.confirmation.wizard record.
func (c *Client) DeletePosConfirmationWizard(id int64) error {
	return c.DeletePosConfirmationWizards([]int64{id})
}

// DeletePosConfirmationWizards deletes existing pos.confirmation.wizard records.
func (c *Client) DeletePosConfirmationWizards(ids []int64) error {
	return c.Delete(PosConfirmationWizardModel, ids)
}

// GetPosConfirmationWizard gets pos.confirmation.wizard existing record.
func (c *Client) GetPosConfirmationWizard(id int64) (*PosConfirmationWizard, error) {
	pcws, err := c.GetPosConfirmationWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcws)[0]), nil
}

// GetPosConfirmationWizards gets pos.confirmation.wizard existing records.
func (c *Client) GetPosConfirmationWizards(ids []int64) (*PosConfirmationWizards, error) {
	pcws := &PosConfirmationWizards{}
	if err := c.Read(PosConfirmationWizardModel, ids, nil, pcws); err != nil {
		return nil, err
	}
	return pcws, nil
}

// FindPosConfirmationWizard finds pos.confirmation.wizard record by querying it with criteria.
func (c *Client) FindPosConfirmationWizard(criteria *Criteria) (*PosConfirmationWizard, error) {
	pcws := &PosConfirmationWizards{}
	if err := c.SearchRead(PosConfirmationWizardModel, criteria, NewOptions().Limit(1), pcws); err != nil {
		return nil, err
	}
	return &((*pcws)[0]), nil
}

// FindPosConfirmationWizards finds pos.confirmation.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosConfirmationWizards(criteria *Criteria, options *Options) (*PosConfirmationWizards, error) {
	pcws := &PosConfirmationWizards{}
	if err := c.SearchRead(PosConfirmationWizardModel, criteria, options, pcws); err != nil {
		return nil, err
	}
	return pcws, nil
}

// FindPosConfirmationWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosConfirmationWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosConfirmationWizardModel, criteria, options)
}

// FindPosConfirmationWizardId finds record id by querying it with criteria.
func (c *Client) FindPosConfirmationWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosConfirmationWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
