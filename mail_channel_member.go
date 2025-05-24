package odoo

// MailChannelMember represents mail.channel.member model.
type MailChannelMember struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omitempty"`
	ChannelId            *Many2One  `xmlrpc:"channel_id,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	CustomChannelName    *String    `xmlrpc:"custom_channel_name,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	FetchedMessageId     *Many2One  `xmlrpc:"fetched_message_id,omitempty"`
	FoldState            *Selection `xmlrpc:"fold_state,omitempty"`
	GuestId              *Many2One  `xmlrpc:"guest_id,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	IsMinimized          *Bool      `xmlrpc:"is_minimized,omitempty"`
	IsPinned             *Bool      `xmlrpc:"is_pinned,omitempty"`
	LastInterestDt       *Time      `xmlrpc:"last_interest_dt,omitempty"`
	LastSeenDt           *Time      `xmlrpc:"last_seen_dt,omitempty"`
	MessageUnreadCounter *Int       `xmlrpc:"message_unread_counter,omitempty"`
	PartnerEmail         *String    `xmlrpc:"partner_email,omitempty"`
	PartnerId            *Many2One  `xmlrpc:"partner_id,omitempty"`
	RtcInvitingSessionId *Many2One  `xmlrpc:"rtc_inviting_session_id,omitempty"`
	RtcSessionIds        *Relation  `xmlrpc:"rtc_session_ids,omitempty"`
	SeenMessageId        *Many2One  `xmlrpc:"seen_message_id,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailChannelMembers represents array of mail.channel.member model.
type MailChannelMembers []MailChannelMember

// MailChannelMemberModel is the odoo model name.
const MailChannelMemberModel = "mail.channel.member"

// Many2One convert MailChannelMember to *Many2One.
func (mcm *MailChannelMember) Many2One() *Many2One {
	return NewMany2One(mcm.Id.Get(), "")
}

// CreateMailChannelMember creates a new mail.channel.member model and returns its id.
func (c *Client) CreateMailChannelMember(mcm *MailChannelMember) (int64, error) {
	ids, err := c.CreateMailChannelMembers([]*MailChannelMember{mcm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailChannelMember creates a new mail.channel.member model and returns its id.
func (c *Client) CreateMailChannelMembers(mcms []*MailChannelMember) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcms {
		vv = append(vv, v)
	}
	return c.Create(MailChannelMemberModel, vv, nil)
}

// UpdateMailChannelMember updates an existing mail.channel.member record.
func (c *Client) UpdateMailChannelMember(mcm *MailChannelMember) error {
	return c.UpdateMailChannelMembers([]int64{mcm.Id.Get()}, mcm)
}

// UpdateMailChannelMembers updates existing mail.channel.member records.
// All records (represented by ids) will be updated by mcm values.
func (c *Client) UpdateMailChannelMembers(ids []int64, mcm *MailChannelMember) error {
	return c.Update(MailChannelMemberModel, ids, mcm, nil)
}

// DeleteMailChannelMember deletes an existing mail.channel.member record.
func (c *Client) DeleteMailChannelMember(id int64) error {
	return c.DeleteMailChannelMembers([]int64{id})
}

// DeleteMailChannelMembers deletes existing mail.channel.member records.
func (c *Client) DeleteMailChannelMembers(ids []int64) error {
	return c.Delete(MailChannelMemberModel, ids)
}

// GetMailChannelMember gets mail.channel.member existing record.
func (c *Client) GetMailChannelMember(id int64) (*MailChannelMember, error) {
	mcms, err := c.GetMailChannelMembers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcms)[0]), nil
}

// GetMailChannelMembers gets mail.channel.member existing records.
func (c *Client) GetMailChannelMembers(ids []int64) (*MailChannelMembers, error) {
	mcms := &MailChannelMembers{}
	if err := c.Read(MailChannelMemberModel, ids, nil, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailChannelMember finds mail.channel.member record by querying it with criteria.
func (c *Client) FindMailChannelMember(criteria *Criteria) (*MailChannelMember, error) {
	mcms := &MailChannelMembers{}
	if err := c.SearchRead(MailChannelMemberModel, criteria, NewOptions().Limit(1), mcms); err != nil {
		return nil, err
	}
	return &((*mcms)[0]), nil
}

// FindMailChannelMembers finds mail.channel.member records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelMembers(criteria *Criteria, options *Options) (*MailChannelMembers, error) {
	mcms := &MailChannelMembers{}
	if err := c.SearchRead(MailChannelMemberModel, criteria, options, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailChannelMemberIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelMemberIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailChannelMemberModel, criteria, options)
}

// FindMailChannelMemberId finds record id by querying it with criteria.
func (c *Client) FindMailChannelMemberId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailChannelMemberModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
