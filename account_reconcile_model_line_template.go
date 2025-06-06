package odoo

// AccountReconcileModelLineTemplate represents account.reconcile.model.line.template model.
type AccountReconcileModelLineTemplate struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omitempty"`
	AccountId        *Many2One  `xmlrpc:"account_id,omitempty"`
	AmountString     *String    `xmlrpc:"amount_string,omitempty"`
	AmountType       *Selection `xmlrpc:"amount_type,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	ForceTaxIncluded *Bool      `xmlrpc:"force_tax_included,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Label            *String    `xmlrpc:"label,omitempty"`
	ModelId          *Many2One  `xmlrpc:"model_id,omitempty"`
	Sequence         *Int       `xmlrpc:"sequence,omitempty"`
	TaxIds           *Relation  `xmlrpc:"tax_ids,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReconcileModelLineTemplates represents array of account.reconcile.model.line.template model.
type AccountReconcileModelLineTemplates []AccountReconcileModelLineTemplate

// AccountReconcileModelLineTemplateModel is the odoo model name.
const AccountReconcileModelLineTemplateModel = "account.reconcile.model.line.template"

// Many2One convert AccountReconcileModelLineTemplate to *Many2One.
func (armlt *AccountReconcileModelLineTemplate) Many2One() *Many2One {
	return NewMany2One(armlt.Id.Get(), "")
}

// CreateAccountReconcileModelLineTemplate creates a new account.reconcile.model.line.template model and returns its id.
func (c *Client) CreateAccountReconcileModelLineTemplate(armlt *AccountReconcileModelLineTemplate) (int64, error) {
	ids, err := c.CreateAccountReconcileModelLineTemplates([]*AccountReconcileModelLineTemplate{armlt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReconcileModelLineTemplate creates a new account.reconcile.model.line.template model and returns its id.
func (c *Client) CreateAccountReconcileModelLineTemplates(armlts []*AccountReconcileModelLineTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range armlts {
		vv = append(vv, v)
	}
	return c.Create(AccountReconcileModelLineTemplateModel, vv, nil)
}

// UpdateAccountReconcileModelLineTemplate updates an existing account.reconcile.model.line.template record.
func (c *Client) UpdateAccountReconcileModelLineTemplate(armlt *AccountReconcileModelLineTemplate) error {
	return c.UpdateAccountReconcileModelLineTemplates([]int64{armlt.Id.Get()}, armlt)
}

// UpdateAccountReconcileModelLineTemplates updates existing account.reconcile.model.line.template records.
// All records (represented by ids) will be updated by armlt values.
func (c *Client) UpdateAccountReconcileModelLineTemplates(ids []int64, armlt *AccountReconcileModelLineTemplate) error {
	return c.Update(AccountReconcileModelLineTemplateModel, ids, armlt, nil)
}

// DeleteAccountReconcileModelLineTemplate deletes an existing account.reconcile.model.line.template record.
func (c *Client) DeleteAccountReconcileModelLineTemplate(id int64) error {
	return c.DeleteAccountReconcileModelLineTemplates([]int64{id})
}

// DeleteAccountReconcileModelLineTemplates deletes existing account.reconcile.model.line.template records.
func (c *Client) DeleteAccountReconcileModelLineTemplates(ids []int64) error {
	return c.Delete(AccountReconcileModelLineTemplateModel, ids)
}

// GetAccountReconcileModelLineTemplate gets account.reconcile.model.line.template existing record.
func (c *Client) GetAccountReconcileModelLineTemplate(id int64) (*AccountReconcileModelLineTemplate, error) {
	armlts, err := c.GetAccountReconcileModelLineTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*armlts)[0]), nil
}

// GetAccountReconcileModelLineTemplates gets account.reconcile.model.line.template existing records.
func (c *Client) GetAccountReconcileModelLineTemplates(ids []int64) (*AccountReconcileModelLineTemplates, error) {
	armlts := &AccountReconcileModelLineTemplates{}
	if err := c.Read(AccountReconcileModelLineTemplateModel, ids, nil, armlts); err != nil {
		return nil, err
	}
	return armlts, nil
}

// FindAccountReconcileModelLineTemplate finds account.reconcile.model.line.template record by querying it with criteria.
func (c *Client) FindAccountReconcileModelLineTemplate(criteria *Criteria) (*AccountReconcileModelLineTemplate, error) {
	armlts := &AccountReconcileModelLineTemplates{}
	if err := c.SearchRead(AccountReconcileModelLineTemplateModel, criteria, NewOptions().Limit(1), armlts); err != nil {
		return nil, err
	}
	return &((*armlts)[0]), nil
}

// FindAccountReconcileModelLineTemplates finds account.reconcile.model.line.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelLineTemplates(criteria *Criteria, options *Options) (*AccountReconcileModelLineTemplates, error) {
	armlts := &AccountReconcileModelLineTemplates{}
	if err := c.SearchRead(AccountReconcileModelLineTemplateModel, criteria, options, armlts); err != nil {
		return nil, err
	}
	return armlts, nil
}

// FindAccountReconcileModelLineTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelLineTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReconcileModelLineTemplateModel, criteria, options)
}

// FindAccountReconcileModelLineTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelLineTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelLineTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
