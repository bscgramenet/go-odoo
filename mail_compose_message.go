package odoo

// MailComposeMessage represents mail.compose.message model.
type MailComposeMessage struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omitempty"`
	ActiveDomain         *String    `xmlrpc:"active_domain,omitempty"`
	AttachmentIds        *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorId             *Many2One  `xmlrpc:"author_id,omitempty"`
	AutoDelete           *Bool      `xmlrpc:"auto_delete,omitempty"`
	AutoDeleteMessage    *Bool      `xmlrpc:"auto_delete_message,omitempty"`
	Body                 *String    `xmlrpc:"body,omitempty"`
	CampaignId           *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CanEditBody          *Bool      `xmlrpc:"can_edit_body,omitempty"`
	CompositionMode      *Selection `xmlrpc:"composition_mode,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	EmailAddSignature    *Bool      `xmlrpc:"email_add_signature,omitempty"`
	EmailFrom            *String    `xmlrpc:"email_from,omitempty"`
	EmailLayoutXmlid     *String    `xmlrpc:"email_layout_xmlid,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	IsLog                *Bool      `xmlrpc:"is_log,omitempty"`
	IsMailTemplateEditor *Bool      `xmlrpc:"is_mail_template_editor,omitempty"`
	Lang                 *String    `xmlrpc:"lang,omitempty"`
	MailActivityTypeId   *Many2One  `xmlrpc:"mail_activity_type_id,omitempty"`
	MailServerId         *Many2One  `xmlrpc:"mail_server_id,omitempty"`
	MailingListIds       *Relation  `xmlrpc:"mailing_list_ids,omitempty"`
	MassMailingId        *Many2One  `xmlrpc:"mass_mailing_id,omitempty"`
	MassMailingName      *String    `xmlrpc:"mass_mailing_name,omitempty"`
	MessageType          *Selection `xmlrpc:"message_type,omitempty"`
	Model                *String    `xmlrpc:"model,omitempty"`
	Notify               *Bool      `xmlrpc:"notify,omitempty"`
	ParentId             *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerIds           *Relation  `xmlrpc:"partner_ids,omitempty"`
	RecordName           *String    `xmlrpc:"record_name,omitempty"`
	RenderModel          *String    `xmlrpc:"render_model,omitempty"`
	ReplyTo              *String    `xmlrpc:"reply_to,omitempty"`
	ReplyToForceNew      *Bool      `xmlrpc:"reply_to_force_new,omitempty"`
	ReplyToMode          *Selection `xmlrpc:"reply_to_mode,omitempty"`
	ResId                *Int       `xmlrpc:"res_id,omitempty"`
	Subject              *String    `xmlrpc:"subject,omitempty"`
	SubtypeId            *Many2One  `xmlrpc:"subtype_id,omitempty"`
	TemplateId           *Many2One  `xmlrpc:"template_id,omitempty"`
	UseActiveDomain      *Bool      `xmlrpc:"use_active_domain,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailComposeMessages represents array of mail.compose.message model.
type MailComposeMessages []MailComposeMessage

// MailComposeMessageModel is the odoo model name.
const MailComposeMessageModel = "mail.compose.message"

// Many2One convert MailComposeMessage to *Many2One.
func (mcm *MailComposeMessage) Many2One() *Many2One {
	return NewMany2One(mcm.Id.Get(), "")
}

// CreateMailComposeMessage creates a new mail.compose.message model and returns its id.
func (c *Client) CreateMailComposeMessage(mcm *MailComposeMessage) (int64, error) {
	ids, err := c.CreateMailComposeMessages([]*MailComposeMessage{mcm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailComposeMessage creates a new mail.compose.message model and returns its id.
func (c *Client) CreateMailComposeMessages(mcms []*MailComposeMessage) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcms {
		vv = append(vv, v)
	}
	return c.Create(MailComposeMessageModel, vv, nil)
}

// UpdateMailComposeMessage updates an existing mail.compose.message record.
func (c *Client) UpdateMailComposeMessage(mcm *MailComposeMessage) error {
	return c.UpdateMailComposeMessages([]int64{mcm.Id.Get()}, mcm)
}

// UpdateMailComposeMessages updates existing mail.compose.message records.
// All records (represented by ids) will be updated by mcm values.
func (c *Client) UpdateMailComposeMessages(ids []int64, mcm *MailComposeMessage) error {
	return c.Update(MailComposeMessageModel, ids, mcm, nil)
}

// DeleteMailComposeMessage deletes an existing mail.compose.message record.
func (c *Client) DeleteMailComposeMessage(id int64) error {
	return c.DeleteMailComposeMessages([]int64{id})
}

// DeleteMailComposeMessages deletes existing mail.compose.message records.
func (c *Client) DeleteMailComposeMessages(ids []int64) error {
	return c.Delete(MailComposeMessageModel, ids)
}

// GetMailComposeMessage gets mail.compose.message existing record.
func (c *Client) GetMailComposeMessage(id int64) (*MailComposeMessage, error) {
	mcms, err := c.GetMailComposeMessages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcms)[0]), nil
}

// GetMailComposeMessages gets mail.compose.message existing records.
func (c *Client) GetMailComposeMessages(ids []int64) (*MailComposeMessages, error) {
	mcms := &MailComposeMessages{}
	if err := c.Read(MailComposeMessageModel, ids, nil, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposeMessage finds mail.compose.message record by querying it with criteria.
func (c *Client) FindMailComposeMessage(criteria *Criteria) (*MailComposeMessage, error) {
	mcms := &MailComposeMessages{}
	if err := c.SearchRead(MailComposeMessageModel, criteria, NewOptions().Limit(1), mcms); err != nil {
		return nil, err
	}
	return &((*mcms)[0]), nil
}

// FindMailComposeMessages finds mail.compose.message records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposeMessages(criteria *Criteria, options *Options) (*MailComposeMessages, error) {
	mcms := &MailComposeMessages{}
	if err := c.SearchRead(MailComposeMessageModel, criteria, options, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposeMessageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposeMessageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailComposeMessageModel, criteria, options)
}

// FindMailComposeMessageId finds record id by querying it with criteria.
func (c *Client) FindMailComposeMessageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailComposeMessageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
