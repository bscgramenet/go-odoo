package odoo

// MailChannel represents mail.channel model.
type MailChannel struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omitempty"`
	Active                    *Bool      `xmlrpc:"active,omitempty"`
	Avatar128                 *String    `xmlrpc:"avatar_128,omitempty"`
	ChannelMemberIds          *Relation  `xmlrpc:"channel_member_ids,omitempty"`
	ChannelPartnerIds         *Relation  `xmlrpc:"channel_partner_ids,omitempty"`
	ChannelType               *Selection `xmlrpc:"channel_type,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	DefaultDisplayMode        *Selection `xmlrpc:"default_display_mode,omitempty"`
	Description               *String    `xmlrpc:"description,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	GroupIds                  *Relation  `xmlrpc:"group_ids,omitempty"`
	GroupPublicId             *Many2One  `xmlrpc:"group_public_id,omitempty"`
	HasMessage                *Bool      `xmlrpc:"has_message,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	Image128                  *String    `xmlrpc:"image_128,omitempty"`
	InvitationUrl             *String    `xmlrpc:"invitation_url,omitempty"`
	IsChat                    *Bool      `xmlrpc:"is_chat,omitempty"`
	IsMember                  *Bool      `xmlrpc:"is_member,omitempty"`
	MemberCount               *Int       `xmlrpc:"member_count,omitempty"`
	MessageAttachmentCount    *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds        *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError           *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter    *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError        *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower         *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId   *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction         *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter  *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds         *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	RtcSessionIds             *Relation  `xmlrpc:"rtc_session_ids,omitempty"`
	SubscriptionDepartmentIds *Relation  `xmlrpc:"subscription_department_ids,omitempty"`
	Uuid                      *String    `xmlrpc:"uuid,omitempty"`
	WebsiteMessageIds         *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailChannels represents array of mail.channel model.
type MailChannels []MailChannel

// MailChannelModel is the odoo model name.
const MailChannelModel = "mail.channel"

// Many2One convert MailChannel to *Many2One.
func (mc *MailChannel) Many2One() *Many2One {
	return NewMany2One(mc.Id.Get(), "")
}

// CreateMailChannel creates a new mail.channel model and returns its id.
func (c *Client) CreateMailChannel(mc *MailChannel) (int64, error) {
	ids, err := c.CreateMailChannels([]*MailChannel{mc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailChannel creates a new mail.channel model and returns its id.
func (c *Client) CreateMailChannels(mcs []*MailChannel) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcs {
		vv = append(vv, v)
	}
	return c.Create(MailChannelModel, vv, nil)
}

// UpdateMailChannel updates an existing mail.channel record.
func (c *Client) UpdateMailChannel(mc *MailChannel) error {
	return c.UpdateMailChannels([]int64{mc.Id.Get()}, mc)
}

// UpdateMailChannels updates existing mail.channel records.
// All records (represented by ids) will be updated by mc values.
func (c *Client) UpdateMailChannels(ids []int64, mc *MailChannel) error {
	return c.Update(MailChannelModel, ids, mc, nil)
}

// DeleteMailChannel deletes an existing mail.channel record.
func (c *Client) DeleteMailChannel(id int64) error {
	return c.DeleteMailChannels([]int64{id})
}

// DeleteMailChannels deletes existing mail.channel records.
func (c *Client) DeleteMailChannels(ids []int64) error {
	return c.Delete(MailChannelModel, ids)
}

// GetMailChannel gets mail.channel existing record.
func (c *Client) GetMailChannel(id int64) (*MailChannel, error) {
	mcs, err := c.GetMailChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcs)[0]), nil
}

// GetMailChannels gets mail.channel existing records.
func (c *Client) GetMailChannels(ids []int64) (*MailChannels, error) {
	mcs := &MailChannels{}
	if err := c.Read(MailChannelModel, ids, nil, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMailChannel finds mail.channel record by querying it with criteria.
func (c *Client) FindMailChannel(criteria *Criteria) (*MailChannel, error) {
	mcs := &MailChannels{}
	if err := c.SearchRead(MailChannelModel, criteria, NewOptions().Limit(1), mcs); err != nil {
		return nil, err
	}
	return &((*mcs)[0]), nil
}

// FindMailChannels finds mail.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannels(criteria *Criteria, options *Options) (*MailChannels, error) {
	mcs := &MailChannels{}
	if err := c.SearchRead(MailChannelModel, criteria, options, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMailChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailChannelModel, criteria, options)
}

// FindMailChannelId finds record id by querying it with criteria.
func (c *Client) FindMailChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
