package odoo

// AccountChartTemplate represents account.chart.template model.
type AccountChartTemplate struct {
	LastUpdate                                  *Time     `xmlrpc:"__last_update,omitempty"`
	AccountIds                                  *Relation `xmlrpc:"account_ids,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountJournalPaymentCreditAccountId        *Many2One `xmlrpc:"account_journal_payment_credit_account_id,omitempty"`
	AccountJournalPaymentDebitAccountId         *Many2One `xmlrpc:"account_journal_payment_debit_account_id,omitempty"`
	AccountJournalSuspenseAccountId             *Many2One `xmlrpc:"account_journal_suspense_account_id,omitempty"`
	BankAccountCodePrefix                       *String   `xmlrpc:"bank_account_code_prefix,omitempty"`
	CashAccountCodePrefix                       *String   `xmlrpc:"cash_account_code_prefix,omitempty"`
	CodeDigits                                  *Int      `xmlrpc:"code_digits,omitempty"`
	CountryId                                   *Many2One `xmlrpc:"country_id,omitempty"`
	CreateDate                                  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId                                  *Many2One `xmlrpc:"currency_id,omitempty"`
	DefaultCashDifferenceExpenseAccountId       *Many2One `xmlrpc:"default_cash_difference_expense_account_id,omitempty"`
	DefaultCashDifferenceIncomeAccountId        *Many2One `xmlrpc:"default_cash_difference_income_account_id,omitempty"`
	DefaultPosReceivableAccountId               *Many2One `xmlrpc:"default_pos_receivable_account_id,omitempty"`
	DisplayName                                 *String   `xmlrpc:"display_name,omitempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	Id                                          *Int      `xmlrpc:"id,omitempty"`
	IncomeCurrencyExchangeAccountId             *Many2One `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	Name                                        *String   `xmlrpc:"name,omitempty"`
	ParentId                                    *Many2One `xmlrpc:"parent_id,omitempty"`
	PropertyAccountExpenseCategId               *Many2One `xmlrpc:"property_account_expense_categ_id,omitempty"`
	PropertyAccountExpenseId                    *Many2One `xmlrpc:"property_account_expense_id,omitempty"`
	PropertyAccountIncomeCategId                *Many2One `xmlrpc:"property_account_income_categ_id,omitempty"`
	PropertyAccountIncomeId                     *Many2One `xmlrpc:"property_account_income_id,omitempty"`
	PropertyAccountPayableId                    *Many2One `xmlrpc:"property_account_payable_id,omitempty"`
	PropertyAccountReceivableId                 *Many2One `xmlrpc:"property_account_receivable_id,omitempty"`
	PropertyAdvanceTaxPaymentAccountId          *Many2One `xmlrpc:"property_advance_tax_payment_account_id,omitempty"`
	PropertyCashBasisBaseAccountId              *Many2One `xmlrpc:"property_cash_basis_base_account_id,omitempty"`
	PropertyStockAccountInputCategId            *Many2One `xmlrpc:"property_stock_account_input_categ_id,omitempty"`
	PropertyStockAccountOutputCategId           *Many2One `xmlrpc:"property_stock_account_output_categ_id,omitempty"`
	PropertyStockValuationAccountId             *Many2One `xmlrpc:"property_stock_valuation_account_id,omitempty"`
	PropertyTaxPayableAccountId                 *Many2One `xmlrpc:"property_tax_payable_account_id,omitempty"`
	PropertyTaxReceivableAccountId              *Many2One `xmlrpc:"property_tax_receivable_account_id,omitempty"`
	TaxTemplateIds                              *Relation `xmlrpc:"tax_template_ids,omitempty"`
	TransferAccountCodePrefix                   *String   `xmlrpc:"transfer_account_code_prefix,omitempty"`
	UseAngloSaxon                               *Bool     `xmlrpc:"use_anglo_saxon,omitempty"`
	UseStornoAccounting                         *Bool     `xmlrpc:"use_storno_accounting,omitempty"`
	Visible                                     *Bool     `xmlrpc:"visible,omitempty"`
	WriteDate                                   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountChartTemplates represents array of account.chart.template model.
type AccountChartTemplates []AccountChartTemplate

// AccountChartTemplateModel is the odoo model name.
const AccountChartTemplateModel = "account.chart.template"

// Many2One convert AccountChartTemplate to *Many2One.
func (act *AccountChartTemplate) Many2One() *Many2One {
	return NewMany2One(act.Id.Get(), "")
}

// CreateAccountChartTemplate creates a new account.chart.template model and returns its id.
func (c *Client) CreateAccountChartTemplate(act *AccountChartTemplate) (int64, error) {
	ids, err := c.CreateAccountChartTemplates([]*AccountChartTemplate{act})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountChartTemplate creates a new account.chart.template model and returns its id.
func (c *Client) CreateAccountChartTemplates(acts []*AccountChartTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range acts {
		vv = append(vv, v)
	}
	return c.Create(AccountChartTemplateModel, vv, nil)
}

// UpdateAccountChartTemplate updates an existing account.chart.template record.
func (c *Client) UpdateAccountChartTemplate(act *AccountChartTemplate) error {
	return c.UpdateAccountChartTemplates([]int64{act.Id.Get()}, act)
}

// UpdateAccountChartTemplates updates existing account.chart.template records.
// All records (represented by ids) will be updated by act values.
func (c *Client) UpdateAccountChartTemplates(ids []int64, act *AccountChartTemplate) error {
	return c.Update(AccountChartTemplateModel, ids, act, nil)
}

// DeleteAccountChartTemplate deletes an existing account.chart.template record.
func (c *Client) DeleteAccountChartTemplate(id int64) error {
	return c.DeleteAccountChartTemplates([]int64{id})
}

// DeleteAccountChartTemplates deletes existing account.chart.template records.
func (c *Client) DeleteAccountChartTemplates(ids []int64) error {
	return c.Delete(AccountChartTemplateModel, ids)
}

// GetAccountChartTemplate gets account.chart.template existing record.
func (c *Client) GetAccountChartTemplate(id int64) (*AccountChartTemplate, error) {
	acts, err := c.GetAccountChartTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acts)[0]), nil
}

// GetAccountChartTemplates gets account.chart.template existing records.
func (c *Client) GetAccountChartTemplates(ids []int64) (*AccountChartTemplates, error) {
	acts := &AccountChartTemplates{}
	if err := c.Read(AccountChartTemplateModel, ids, nil, acts); err != nil {
		return nil, err
	}
	return acts, nil
}

// FindAccountChartTemplate finds account.chart.template record by querying it with criteria.
func (c *Client) FindAccountChartTemplate(criteria *Criteria) (*AccountChartTemplate, error) {
	acts := &AccountChartTemplates{}
	if err := c.SearchRead(AccountChartTemplateModel, criteria, NewOptions().Limit(1), acts); err != nil {
		return nil, err
	}
	return &((*acts)[0]), nil
}

// FindAccountChartTemplates finds account.chart.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountChartTemplates(criteria *Criteria, options *Options) (*AccountChartTemplates, error) {
	acts := &AccountChartTemplates{}
	if err := c.SearchRead(AccountChartTemplateModel, criteria, options, acts); err != nil {
		return nil, err
	}
	return acts, nil
}

// FindAccountChartTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountChartTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountChartTemplateModel, criteria, options)
}

// FindAccountChartTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountChartTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountChartTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
