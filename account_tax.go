package odoo

// AccountTax represents account.tax model.
type AccountTax struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omitempty"`
	Active                       *Bool      `xmlrpc:"active,omitempty"`
	Amount                       *Float     `xmlrpc:"amount,omitempty"`
	AmountType                   *Selection `xmlrpc:"amount_type,omitempty"`
	Analytic                     *Bool      `xmlrpc:"analytic,omitempty"`
	CashBasisTransitionAccountId *Many2One  `xmlrpc:"cash_basis_transition_account_id,omitempty"`
	ChildrenTaxIds               *Relation  `xmlrpc:"children_tax_ids,omitempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryCode                  *String    `xmlrpc:"country_code,omitempty"`
	CountryId                    *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description                  *String    `xmlrpc:"description,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	HideTaxExigibility           *Bool      `xmlrpc:"hide_tax_exigibility,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	IncludeBaseAmount            *Bool      `xmlrpc:"include_base_amount,omitempty"`
	InvoiceRepartitionLineIds    *Relation  `xmlrpc:"invoice_repartition_line_ids,omitempty"`
	IsBaseAffected               *Bool      `xmlrpc:"is_base_affected,omitempty"`
	Name                         *String    `xmlrpc:"name,omitempty"`
	NameSearchable               *String    `xmlrpc:"name_searchable,omitempty"`
	PriceInclude                 *Bool      `xmlrpc:"price_include,omitempty"`
	RealAmount                   *Float     `xmlrpc:"real_amount,omitempty"`
	RefundRepartitionLineIds     *Relation  `xmlrpc:"refund_repartition_line_ids,omitempty"`
	Sequence                     *Int       `xmlrpc:"sequence,omitempty"`
	TaxExigibility               *Selection `xmlrpc:"tax_exigibility,omitempty"`
	TaxGroupId                   *Many2One  `xmlrpc:"tax_group_id,omitempty"`
	TaxScope                     *Selection `xmlrpc:"tax_scope,omitempty"`
	TypeTaxUse                   *Selection `xmlrpc:"type_tax_use,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
	XClaseid                     *Int       `xmlrpc:"x_claseid,omitempty"`
}

// AccountTaxs represents array of account.tax model.
type AccountTaxs []AccountTax

// AccountTaxModel is the odoo model name.
const AccountTaxModel = "account.tax"

// Many2One convert AccountTax to *Many2One.
func (at *AccountTax) Many2One() *Many2One {
	return NewMany2One(at.Id.Get(), "")
}

// CreateAccountTax creates a new account.tax model and returns its id.
func (c *Client) CreateAccountTax(at *AccountTax) (int64, error) {
	ids, err := c.CreateAccountTaxs([]*AccountTax{at})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountTax creates a new account.tax model and returns its id.
func (c *Client) CreateAccountTaxs(ats []*AccountTax) ([]int64, error) {
	var vv []interface{}
	for _, v := range ats {
		vv = append(vv, v)
	}
	return c.Create(AccountTaxModel, vv, nil)
}

// UpdateAccountTax updates an existing account.tax record.
func (c *Client) UpdateAccountTax(at *AccountTax) error {
	return c.UpdateAccountTaxs([]int64{at.Id.Get()}, at)
}

// UpdateAccountTaxs updates existing account.tax records.
// All records (represented by ids) will be updated by at values.
func (c *Client) UpdateAccountTaxs(ids []int64, at *AccountTax) error {
	return c.Update(AccountTaxModel, ids, at, nil)
}

// DeleteAccountTax deletes an existing account.tax record.
func (c *Client) DeleteAccountTax(id int64) error {
	return c.DeleteAccountTaxs([]int64{id})
}

// DeleteAccountTaxs deletes existing account.tax records.
func (c *Client) DeleteAccountTaxs(ids []int64) error {
	return c.Delete(AccountTaxModel, ids)
}

// GetAccountTax gets account.tax existing record.
func (c *Client) GetAccountTax(id int64) (*AccountTax, error) {
	ats, err := c.GetAccountTaxs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ats)[0]), nil
}

// GetAccountTaxs gets account.tax existing records.
func (c *Client) GetAccountTaxs(ids []int64) (*AccountTaxs, error) {
	ats := &AccountTaxs{}
	if err := c.Read(AccountTaxModel, ids, nil, ats); err != nil {
		return nil, err
	}
	return ats, nil
}

// FindAccountTax finds account.tax record by querying it with criteria.
func (c *Client) FindAccountTax(criteria *Criteria) (*AccountTax, error) {
	ats := &AccountTaxs{}
	if err := c.SearchRead(AccountTaxModel, criteria, NewOptions().Limit(1), ats); err != nil {
		return nil, err
	}
	return &((*ats)[0]), nil
}

// FindAccountTaxs finds account.tax records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxs(criteria *Criteria, options *Options) (*AccountTaxs, error) {
	ats := &AccountTaxs{}
	if err := c.SearchRead(AccountTaxModel, criteria, options, ats); err != nil {
		return nil, err
	}
	return ats, nil
}

// FindAccountTaxIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountTaxModel, criteria, options)
}

// FindAccountTaxId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
