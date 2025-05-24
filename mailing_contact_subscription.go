package odoo

// MailingContactSubscription represents mailing.contact.subscription model.
type MailingContactSubscription struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omitempty"`
	ContactId          *Many2One `xmlrpc:"contact_id,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	IsBlacklisted      *Bool     `xmlrpc:"is_blacklisted,omitempty"`
	ListId             *Many2One `xmlrpc:"list_id,omitempty"`
	MessageBounce      *Int      `xmlrpc:"message_bounce,omitempty"`
	OptOut             *Bool     `xmlrpc:"opt_out,omitempty"`
	UnsubscriptionDate *Time     `xmlrpc:"unsubscription_date,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailingContactSubscriptions represents array of mailing.contact.subscription model.
type MailingContactSubscriptions []MailingContactSubscription

// MailingContactSubscriptionModel is the odoo model name.
const MailingContactSubscriptionModel = "mailing.contact.subscription"

// Many2One convert MailingContactSubscription to *Many2One.
func (mcs *MailingContactSubscription) Many2One() *Many2One {
	return NewMany2One(mcs.Id.Get(), "")
}

// CreateMailingContactSubscription creates a new mailing.contact.subscription model and returns its id.
func (c *Client) CreateMailingContactSubscription(mcs *MailingContactSubscription) (int64, error) {
	ids, err := c.CreateMailingContactSubscriptions([]*MailingContactSubscription{mcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingContactSubscription creates a new mailing.contact.subscription model and returns its id.
func (c *Client) CreateMailingContactSubscriptions(mcss []*MailingContactSubscription) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcss {
		vv = append(vv, v)
	}
	return c.Create(MailingContactSubscriptionModel, vv, nil)
}

// UpdateMailingContactSubscription updates an existing mailing.contact.subscription record.
func (c *Client) UpdateMailingContactSubscription(mcs *MailingContactSubscription) error {
	return c.UpdateMailingContactSubscriptions([]int64{mcs.Id.Get()}, mcs)
}

// UpdateMailingContactSubscriptions updates existing mailing.contact.subscription records.
// All records (represented by ids) will be updated by mcs values.
func (c *Client) UpdateMailingContactSubscriptions(ids []int64, mcs *MailingContactSubscription) error {
	return c.Update(MailingContactSubscriptionModel, ids, mcs, nil)
}

// DeleteMailingContactSubscription deletes an existing mailing.contact.subscription record.
func (c *Client) DeleteMailingContactSubscription(id int64) error {
	return c.DeleteMailingContactSubscriptions([]int64{id})
}

// DeleteMailingContactSubscriptions deletes existing mailing.contact.subscription records.
func (c *Client) DeleteMailingContactSubscriptions(ids []int64) error {
	return c.Delete(MailingContactSubscriptionModel, ids)
}

// GetMailingContactSubscription gets mailing.contact.subscription existing record.
func (c *Client) GetMailingContactSubscription(id int64) (*MailingContactSubscription, error) {
	mcss, err := c.GetMailingContactSubscriptions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcss)[0]), nil
}

// GetMailingContactSubscriptions gets mailing.contact.subscription existing records.
func (c *Client) GetMailingContactSubscriptions(ids []int64) (*MailingContactSubscriptions, error) {
	mcss := &MailingContactSubscriptions{}
	if err := c.Read(MailingContactSubscriptionModel, ids, nil, mcss); err != nil {
		return nil, err
	}
	return mcss, nil
}

// FindMailingContactSubscription finds mailing.contact.subscription record by querying it with criteria.
func (c *Client) FindMailingContactSubscription(criteria *Criteria) (*MailingContactSubscription, error) {
	mcss := &MailingContactSubscriptions{}
	if err := c.SearchRead(MailingContactSubscriptionModel, criteria, NewOptions().Limit(1), mcss); err != nil {
		return nil, err
	}
	return &((*mcss)[0]), nil
}

// FindMailingContactSubscriptions finds mailing.contact.subscription records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingContactSubscriptions(criteria *Criteria, options *Options) (*MailingContactSubscriptions, error) {
	mcss := &MailingContactSubscriptions{}
	if err := c.SearchRead(MailingContactSubscriptionModel, criteria, options, mcss); err != nil {
		return nil, err
	}
	return mcss, nil
}

// FindMailingContactSubscriptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingContactSubscriptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingContactSubscriptionModel, criteria, options)
}

// FindMailingContactSubscriptionId finds record id by querying it with criteria.
func (c *Client) FindMailingContactSubscriptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingContactSubscriptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
