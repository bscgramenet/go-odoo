package odoo

// UtmCampaign represents utm.campaign model.
type UtmCampaign struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omitempty"`
	AbTestingCompleted         *Bool      `xmlrpc:"ab_testing_completed,omitempty"`
	AbTestingMailingsCount     *Int       `xmlrpc:"ab_testing_mailings_count,omitempty"`
	AbTestingScheduleDatetime  *Time      `xmlrpc:"ab_testing_schedule_datetime,omitempty"`
	AbTestingTotalPc           *Int       `xmlrpc:"ab_testing_total_pc,omitempty"`
	AbTestingWinnerSelection   *Selection `xmlrpc:"ab_testing_winner_selection,omitempty"`
	BouncedRatio               *Int       `xmlrpc:"bounced_ratio,omitempty"`
	ClickCount                 *Int       `xmlrpc:"click_count,omitempty"`
	Color                      *Int       `xmlrpc:"color,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmLeadCount               *Int       `xmlrpc:"crm_lead_count,omitempty"`
	CurrencyId                 *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	InvoicedAmount             *Int       `xmlrpc:"invoiced_amount,omitempty"`
	IsAutoCampaign             *Bool      `xmlrpc:"is_auto_campaign,omitempty"`
	IsMailingCampaignActivated *Bool      `xmlrpc:"is_mailing_campaign_activated,omitempty"`
	MailingMailCount           *Int       `xmlrpc:"mailing_mail_count,omitempty"`
	MailingMailIds             *Relation  `xmlrpc:"mailing_mail_ids,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	OpenedRatio                *Int       `xmlrpc:"opened_ratio,omitempty"`
	QuotationCount             *Int       `xmlrpc:"quotation_count,omitempty"`
	ReceivedRatio              *Int       `xmlrpc:"received_ratio,omitempty"`
	RepliedRatio               *Int       `xmlrpc:"replied_ratio,omitempty"`
	StageId                    *Many2One  `xmlrpc:"stage_id,omitempty"`
	TagIds                     *Relation  `xmlrpc:"tag_ids,omitempty"`
	Title                      *String    `xmlrpc:"title,omitempty"`
	UseLeads                   *Bool      `xmlrpc:"use_leads,omitempty"`
	UserId                     *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// UtmCampaigns represents array of utm.campaign model.
type UtmCampaigns []UtmCampaign

// UtmCampaignModel is the odoo model name.
const UtmCampaignModel = "utm.campaign"

// Many2One convert UtmCampaign to *Many2One.
func (uc *UtmCampaign) Many2One() *Many2One {
	return NewMany2One(uc.Id.Get(), "")
}

// CreateUtmCampaign creates a new utm.campaign model and returns its id.
func (c *Client) CreateUtmCampaign(uc *UtmCampaign) (int64, error) {
	ids, err := c.CreateUtmCampaigns([]*UtmCampaign{uc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateUtmCampaign creates a new utm.campaign model and returns its id.
func (c *Client) CreateUtmCampaigns(ucs []*UtmCampaign) ([]int64, error) {
	var vv []interface{}
	for _, v := range ucs {
		vv = append(vv, v)
	}
	return c.Create(UtmCampaignModel, vv, nil)
}

// UpdateUtmCampaign updates an existing utm.campaign record.
func (c *Client) UpdateUtmCampaign(uc *UtmCampaign) error {
	return c.UpdateUtmCampaigns([]int64{uc.Id.Get()}, uc)
}

// UpdateUtmCampaigns updates existing utm.campaign records.
// All records (represented by ids) will be updated by uc values.
func (c *Client) UpdateUtmCampaigns(ids []int64, uc *UtmCampaign) error {
	return c.Update(UtmCampaignModel, ids, uc, nil)
}

// DeleteUtmCampaign deletes an existing utm.campaign record.
func (c *Client) DeleteUtmCampaign(id int64) error {
	return c.DeleteUtmCampaigns([]int64{id})
}

// DeleteUtmCampaigns deletes existing utm.campaign records.
func (c *Client) DeleteUtmCampaigns(ids []int64) error {
	return c.Delete(UtmCampaignModel, ids)
}

// GetUtmCampaign gets utm.campaign existing record.
func (c *Client) GetUtmCampaign(id int64) (*UtmCampaign, error) {
	ucs, err := c.GetUtmCampaigns([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ucs)[0]), nil
}

// GetUtmCampaigns gets utm.campaign existing records.
func (c *Client) GetUtmCampaigns(ids []int64) (*UtmCampaigns, error) {
	ucs := &UtmCampaigns{}
	if err := c.Read(UtmCampaignModel, ids, nil, ucs); err != nil {
		return nil, err
	}
	return ucs, nil
}

// FindUtmCampaign finds utm.campaign record by querying it with criteria.
func (c *Client) FindUtmCampaign(criteria *Criteria) (*UtmCampaign, error) {
	ucs := &UtmCampaigns{}
	if err := c.SearchRead(UtmCampaignModel, criteria, NewOptions().Limit(1), ucs); err != nil {
		return nil, err
	}
	return &((*ucs)[0]), nil
}

// FindUtmCampaigns finds utm.campaign records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmCampaigns(criteria *Criteria, options *Options) (*UtmCampaigns, error) {
	ucs := &UtmCampaigns{}
	if err := c.SearchRead(UtmCampaignModel, criteria, options, ucs); err != nil {
		return nil, err
	}
	return ucs, nil
}

// FindUtmCampaignIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmCampaignIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(UtmCampaignModel, criteria, options)
}

// FindUtmCampaignId finds record id by querying it with criteria.
func (c *Client) FindUtmCampaignId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UtmCampaignModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
