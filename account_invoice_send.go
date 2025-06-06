package odoo

// AccountInvoiceSend represents account.invoice.send model.
type AccountInvoiceSend struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omitempty"`
	ActiveDomain         *String    `xmlrpc:"active_domain,omitempty"`
	AttachmentIds        *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorId             *Many2One  `xmlrpc:"author_id,omitempty"`
	AutoDelete           *Bool      `xmlrpc:"auto_delete,omitempty"`
	AutoDeleteMessage    *Bool      `xmlrpc:"auto_delete_message,omitempty"`
	Body                 *String    `xmlrpc:"body,omitempty"`
	CampaignId           *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CanEditBody          *Bool      `xmlrpc:"can_edit_body,omitempty"`
	ComposerId           *Many2One  `xmlrpc:"composer_id,omitempty"`
	CompositionMode      *Selection `xmlrpc:"composition_mode,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	EmailAddSignature    *Bool      `xmlrpc:"email_add_signature,omitempty"`
	EmailFrom            *String    `xmlrpc:"email_from,omitempty"`
	EmailLayoutXmlid     *String    `xmlrpc:"email_layout_xmlid,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	InvalidAddresses     *Int       `xmlrpc:"invalid_addresses,omitempty"`
	InvalidInvoices      *Int       `xmlrpc:"invalid_invoices,omitempty"`
	InvalidPartnerIds    *Relation  `xmlrpc:"invalid_partner_ids,omitempty"`
	InvoiceIds           *Relation  `xmlrpc:"invoice_ids,omitempty"`
	InvoiceWithoutEmail  *String    `xmlrpc:"invoice_without_email,omitempty"`
	IsEmail              *Bool      `xmlrpc:"is_email,omitempty"`
	IsLog                *Bool      `xmlrpc:"is_log,omitempty"`
	IsMailTemplateEditor *Bool      `xmlrpc:"is_mail_template_editor,omitempty"`
	IsPrint              *Bool      `xmlrpc:"is_print,omitempty"`
	Lang                 *String    `xmlrpc:"lang,omitempty"`
	MailActivityTypeId   *Many2One  `xmlrpc:"mail_activity_type_id,omitempty"`
	MailServerId         *Many2One  `xmlrpc:"mail_server_id,omitempty"`
	MailingListIds       *Relation  `xmlrpc:"mailing_list_ids,omitempty"`
	MassMailingId        *Many2One  `xmlrpc:"mass_mailing_id,omitempty"`
	MassMailingName      *String    `xmlrpc:"mass_mailing_name,omitempty"`
	MessageType          *Selection `xmlrpc:"message_type,omitempty"`
	Model                *String    `xmlrpc:"model,omitempty"`
	MoveTypes            *String    `xmlrpc:"move_types,omitempty"`
	Notify               *Bool      `xmlrpc:"notify,omitempty"`
	ParentId             *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerId            *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerIds           *Relation  `xmlrpc:"partner_ids,omitempty"`
	Printed              *Bool      `xmlrpc:"printed,omitempty"`
	RecordName           *String    `xmlrpc:"record_name,omitempty"`
	RenderModel          *String    `xmlrpc:"render_model,omitempty"`
	ReplyTo              *String    `xmlrpc:"reply_to,omitempty"`
	ReplyToForceNew      *Bool      `xmlrpc:"reply_to_force_new,omitempty"`
	ReplyToMode          *Selection `xmlrpc:"reply_to_mode,omitempty"`
	ResId                *Int       `xmlrpc:"res_id,omitempty"`
	SnailmailCost        *Float     `xmlrpc:"snailmail_cost,omitempty"`
	SnailmailIsLetter    *Bool      `xmlrpc:"snailmail_is_letter,omitempty"`
	Subject              *String    `xmlrpc:"subject,omitempty"`
	SubtypeId            *Many2One  `xmlrpc:"subtype_id,omitempty"`
	TemplateId           *Many2One  `xmlrpc:"template_id,omitempty"`
	UseActiveDomain      *Bool      `xmlrpc:"use_active_domain,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountInvoiceSends represents array of account.invoice.send model.
type AccountInvoiceSends []AccountInvoiceSend

// AccountInvoiceSendModel is the odoo model name.
const AccountInvoiceSendModel = "account.invoice.send"

// Many2One convert AccountInvoiceSend to *Many2One.
func (ais *AccountInvoiceSend) Many2One() *Many2One {
	return NewMany2One(ais.Id.Get(), "")
}

// CreateAccountInvoiceSend creates a new account.invoice.send model and returns its id.
func (c *Client) CreateAccountInvoiceSend(ais *AccountInvoiceSend) (int64, error) {
	ids, err := c.CreateAccountInvoiceSends([]*AccountInvoiceSend{ais})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountInvoiceSend creates a new account.invoice.send model and returns its id.
func (c *Client) CreateAccountInvoiceSends(aiss []*AccountInvoiceSend) ([]int64, error) {
	var vv []interface{}
	for _, v := range aiss {
		vv = append(vv, v)
	}
	return c.Create(AccountInvoiceSendModel, vv, nil)
}

// UpdateAccountInvoiceSend updates an existing account.invoice.send record.
func (c *Client) UpdateAccountInvoiceSend(ais *AccountInvoiceSend) error {
	return c.UpdateAccountInvoiceSends([]int64{ais.Id.Get()}, ais)
}

// UpdateAccountInvoiceSends updates existing account.invoice.send records.
// All records (represented by ids) will be updated by ais values.
func (c *Client) UpdateAccountInvoiceSends(ids []int64, ais *AccountInvoiceSend) error {
	return c.Update(AccountInvoiceSendModel, ids, ais, nil)
}

// DeleteAccountInvoiceSend deletes an existing account.invoice.send record.
func (c *Client) DeleteAccountInvoiceSend(id int64) error {
	return c.DeleteAccountInvoiceSends([]int64{id})
}

// DeleteAccountInvoiceSends deletes existing account.invoice.send records.
func (c *Client) DeleteAccountInvoiceSends(ids []int64) error {
	return c.Delete(AccountInvoiceSendModel, ids)
}

// GetAccountInvoiceSend gets account.invoice.send existing record.
func (c *Client) GetAccountInvoiceSend(id int64) (*AccountInvoiceSend, error) {
	aiss, err := c.GetAccountInvoiceSends([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aiss)[0]), nil
}

// GetAccountInvoiceSends gets account.invoice.send existing records.
func (c *Client) GetAccountInvoiceSends(ids []int64) (*AccountInvoiceSends, error) {
	aiss := &AccountInvoiceSends{}
	if err := c.Read(AccountInvoiceSendModel, ids, nil, aiss); err != nil {
		return nil, err
	}
	return aiss, nil
}

// FindAccountInvoiceSend finds account.invoice.send record by querying it with criteria.
func (c *Client) FindAccountInvoiceSend(criteria *Criteria) (*AccountInvoiceSend, error) {
	aiss := &AccountInvoiceSends{}
	if err := c.SearchRead(AccountInvoiceSendModel, criteria, NewOptions().Limit(1), aiss); err != nil {
		return nil, err
	}
	return &((*aiss)[0]), nil
}

// FindAccountInvoiceSends finds account.invoice.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceSends(criteria *Criteria, options *Options) (*AccountInvoiceSends, error) {
	aiss := &AccountInvoiceSends{}
	if err := c.SearchRead(AccountInvoiceSendModel, criteria, options, aiss); err != nil {
		return nil, err
	}
	return aiss, nil
}

// FindAccountInvoiceSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountInvoiceSendModel, criteria, options)
}

// FindAccountInvoiceSendId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
