package odoo

// AccountReconcileModelTemplate represents account.reconcile.model.template model.
type AccountReconcileModelTemplate struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omitempty"`
	AllowPaymentTolerance      *Bool      `xmlrpc:"allow_payment_tolerance,omitempty"`
	AutoReconcile              *Bool      `xmlrpc:"auto_reconcile,omitempty"`
	ChartTemplateId            *Many2One  `xmlrpc:"chart_template_id,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	DecimalSeparator           *String    `xmlrpc:"decimal_separator,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	LineIds                    *Relation  `xmlrpc:"line_ids,omitempty"`
	MatchAmount                *Selection `xmlrpc:"match_amount,omitempty"`
	MatchAmountMax             *Float     `xmlrpc:"match_amount_max,omitempty"`
	MatchAmountMin             *Float     `xmlrpc:"match_amount_min,omitempty"`
	MatchJournalIds            *Relation  `xmlrpc:"match_journal_ids,omitempty"`
	MatchLabel                 *Selection `xmlrpc:"match_label,omitempty"`
	MatchLabelParam            *String    `xmlrpc:"match_label_param,omitempty"`
	MatchNature                *Selection `xmlrpc:"match_nature,omitempty"`
	MatchNote                  *Selection `xmlrpc:"match_note,omitempty"`
	MatchNoteParam             *String    `xmlrpc:"match_note_param,omitempty"`
	MatchPartner               *Bool      `xmlrpc:"match_partner,omitempty"`
	MatchPartnerCategoryIds    *Relation  `xmlrpc:"match_partner_category_ids,omitempty"`
	MatchPartnerIds            *Relation  `xmlrpc:"match_partner_ids,omitempty"`
	MatchSameCurrency          *Bool      `xmlrpc:"match_same_currency,omitempty"`
	MatchTextLocationLabel     *Bool      `xmlrpc:"match_text_location_label,omitempty"`
	MatchTextLocationNote      *Bool      `xmlrpc:"match_text_location_note,omitempty"`
	MatchTextLocationReference *Bool      `xmlrpc:"match_text_location_reference,omitempty"`
	MatchTransactionType       *Selection `xmlrpc:"match_transaction_type,omitempty"`
	MatchTransactionTypeParam  *String    `xmlrpc:"match_transaction_type_param,omitempty"`
	MatchingOrder              *Selection `xmlrpc:"matching_order,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	PaymentToleranceParam      *Float     `xmlrpc:"payment_tolerance_param,omitempty"`
	PaymentToleranceType       *Selection `xmlrpc:"payment_tolerance_type,omitempty"`
	RuleType                   *Selection `xmlrpc:"rule_type,omitempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omitempty"`
	ToCheck                    *Bool      `xmlrpc:"to_check,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReconcileModelTemplates represents array of account.reconcile.model.template model.
type AccountReconcileModelTemplates []AccountReconcileModelTemplate

// AccountReconcileModelTemplateModel is the odoo model name.
const AccountReconcileModelTemplateModel = "account.reconcile.model.template"

// Many2One convert AccountReconcileModelTemplate to *Many2One.
func (armt *AccountReconcileModelTemplate) Many2One() *Many2One {
	return NewMany2One(armt.Id.Get(), "")
}

// CreateAccountReconcileModelTemplate creates a new account.reconcile.model.template model and returns its id.
func (c *Client) CreateAccountReconcileModelTemplate(armt *AccountReconcileModelTemplate) (int64, error) {
	ids, err := c.CreateAccountReconcileModelTemplates([]*AccountReconcileModelTemplate{armt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReconcileModelTemplate creates a new account.reconcile.model.template model and returns its id.
func (c *Client) CreateAccountReconcileModelTemplates(armts []*AccountReconcileModelTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range armts {
		vv = append(vv, v)
	}
	return c.Create(AccountReconcileModelTemplateModel, vv, nil)
}

// UpdateAccountReconcileModelTemplate updates an existing account.reconcile.model.template record.
func (c *Client) UpdateAccountReconcileModelTemplate(armt *AccountReconcileModelTemplate) error {
	return c.UpdateAccountReconcileModelTemplates([]int64{armt.Id.Get()}, armt)
}

// UpdateAccountReconcileModelTemplates updates existing account.reconcile.model.template records.
// All records (represented by ids) will be updated by armt values.
func (c *Client) UpdateAccountReconcileModelTemplates(ids []int64, armt *AccountReconcileModelTemplate) error {
	return c.Update(AccountReconcileModelTemplateModel, ids, armt, nil)
}

// DeleteAccountReconcileModelTemplate deletes an existing account.reconcile.model.template record.
func (c *Client) DeleteAccountReconcileModelTemplate(id int64) error {
	return c.DeleteAccountReconcileModelTemplates([]int64{id})
}

// DeleteAccountReconcileModelTemplates deletes existing account.reconcile.model.template records.
func (c *Client) DeleteAccountReconcileModelTemplates(ids []int64) error {
	return c.Delete(AccountReconcileModelTemplateModel, ids)
}

// GetAccountReconcileModelTemplate gets account.reconcile.model.template existing record.
func (c *Client) GetAccountReconcileModelTemplate(id int64) (*AccountReconcileModelTemplate, error) {
	armts, err := c.GetAccountReconcileModelTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*armts)[0]), nil
}

// GetAccountReconcileModelTemplates gets account.reconcile.model.template existing records.
func (c *Client) GetAccountReconcileModelTemplates(ids []int64) (*AccountReconcileModelTemplates, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.Read(AccountReconcileModelTemplateModel, ids, nil, armts); err != nil {
		return nil, err
	}
	return armts, nil
}

// FindAccountReconcileModelTemplate finds account.reconcile.model.template record by querying it with criteria.
func (c *Client) FindAccountReconcileModelTemplate(criteria *Criteria) (*AccountReconcileModelTemplate, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.SearchRead(AccountReconcileModelTemplateModel, criteria, NewOptions().Limit(1), armts); err != nil {
		return nil, err
	}
	return &((*armts)[0]), nil
}

// FindAccountReconcileModelTemplates finds account.reconcile.model.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelTemplates(criteria *Criteria, options *Options) (*AccountReconcileModelTemplates, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.SearchRead(AccountReconcileModelTemplateModel, criteria, options, armts); err != nil {
		return nil, err
	}
	return armts, nil
}

// FindAccountReconcileModelTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReconcileModelTemplateModel, criteria, options)
}

// FindAccountReconcileModelTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
