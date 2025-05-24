package odoo

// PosConfig represents pos.config model.
type PosConfig struct {
	LastUpdate                          *Time      `xmlrpc:"__last_update,omitempty"`
	Active                              *Bool      `xmlrpc:"active,omitempty"`
	AmountAuthorizedDiff                *Float     `xmlrpc:"amount_authorized_diff,omitempty"`
	AvailablePricelistIds               *Relation  `xmlrpc:"available_pricelist_ids,omitempty"`
	CashControl                         *Bool      `xmlrpc:"cash_control,omitempty"`
	CashRounding                        *Bool      `xmlrpc:"cash_rounding,omitempty"`
	CompanyHasTemplate                  *Bool      `xmlrpc:"company_has_template,omitempty"`
	CompanyId                           *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                           *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmTeamId                           *Many2One  `xmlrpc:"crm_team_id,omitempty"`
	CurrencyId                          *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrentSessionId                    *Many2One  `xmlrpc:"current_session_id,omitempty"`
	CurrentSessionState                 *String    `xmlrpc:"current_session_state,omitempty"`
	CurrentUserId                       *Many2One  `xmlrpc:"current_user_id,omitempty"`
	DefaultBillIds                      *Relation  `xmlrpc:"default_bill_ids,omitempty"`
	DefaultFiscalPositionId             *Many2One  `xmlrpc:"default_fiscal_position_id,omitempty"`
	DisplayName                         *String    `xmlrpc:"display_name,omitempty"`
	DownPaymentProductId                *Many2One  `xmlrpc:"down_payment_product_id,omitempty"`
	EmployeeIds                         *Relation  `xmlrpc:"employee_ids,omitempty"`
	EpsonPrinterIp                      *String    `xmlrpc:"epson_printer_ip,omitempty"`
	FiscalPositionIds                   *Relation  `xmlrpc:"fiscal_position_ids,omitempty"`
	GroupPosManagerId                   *Many2One  `xmlrpc:"group_pos_manager_id,omitempty"`
	GroupPosUserId                      *Many2One  `xmlrpc:"group_pos_user_id,omitempty"`
	HasActiveSession                    *Bool      `xmlrpc:"has_active_session,omitempty"`
	Id                                  *Int       `xmlrpc:"id,omitempty"`
	IfaceAvailableCategIds              *Relation  `xmlrpc:"iface_available_categ_ids,omitempty"`
	IfaceBigScrollbars                  *Bool      `xmlrpc:"iface_big_scrollbars,omitempty"`
	IfaceCashdrawer                     *Bool      `xmlrpc:"iface_cashdrawer,omitempty"`
	IfaceCustomerFacingDisplay          *Bool      `xmlrpc:"iface_customer_facing_display,omitempty"`
	IfaceCustomerFacingDisplayLocal     *Bool      `xmlrpc:"iface_customer_facing_display_local,omitempty"`
	IfaceCustomerFacingDisplayViaProxy  *Bool      `xmlrpc:"iface_customer_facing_display_via_proxy,omitempty"`
	IfaceElectronicScale                *Bool      `xmlrpc:"iface_electronic_scale,omitempty"`
	IfacePrintAuto                      *Bool      `xmlrpc:"iface_print_auto,omitempty"`
	IfacePrintSkipScreen                *Bool      `xmlrpc:"iface_print_skip_screen,omitempty"`
	IfacePrintViaProxy                  *Bool      `xmlrpc:"iface_print_via_proxy,omitempty"`
	IfaceScanViaProxy                   *Bool      `xmlrpc:"iface_scan_via_proxy,omitempty"`
	IfaceStartCategId                   *Many2One  `xmlrpc:"iface_start_categ_id,omitempty"`
	IfaceTaxIncluded                    *Selection `xmlrpc:"iface_tax_included,omitempty"`
	IfaceTipproduct                     *Bool      `xmlrpc:"iface_tipproduct,omitempty"`
	InvoiceJournalId                    *Many2One  `xmlrpc:"invoice_journal_id,omitempty"`
	IsHeaderOrFooter                    *Bool      `xmlrpc:"is_header_or_footer,omitempty"`
	IsInstalledAccountAccountant        *Bool      `xmlrpc:"is_installed_account_accountant,omitempty"`
	IsMarginsCostsAccessibleToEveryUser *Bool      `xmlrpc:"is_margins_costs_accessible_to_every_user,omitempty"`
	IsPosbox                            *Bool      `xmlrpc:"is_posbox,omitempty"`
	JournalId                           *Many2One  `xmlrpc:"journal_id,omitempty"`
	LastSessionClosingCash              *Float     `xmlrpc:"last_session_closing_cash,omitempty"`
	LastSessionClosingDate              *Time      `xmlrpc:"last_session_closing_date,omitempty"`
	LimitCategories                     *Bool      `xmlrpc:"limit_categories,omitempty"`
	LimitedPartnersAmount               *Int       `xmlrpc:"limited_partners_amount,omitempty"`
	LimitedPartnersLoading              *Bool      `xmlrpc:"limited_partners_loading,omitempty"`
	LimitedProductsAmount               *Int       `xmlrpc:"limited_products_amount,omitempty"`
	LimitedProductsLoading              *Bool      `xmlrpc:"limited_products_loading,omitempty"`
	ManualDiscount                      *Bool      `xmlrpc:"manual_discount,omitempty"`
	ModulePosDiscount                   *Bool      `xmlrpc:"module_pos_discount,omitempty"`
	ModulePosHr                         *Bool      `xmlrpc:"module_pos_hr,omitempty"`
	ModulePosMercury                    *Bool      `xmlrpc:"module_pos_mercury,omitempty"`
	ModulePosRestaurant                 *Bool      `xmlrpc:"module_pos_restaurant,omitempty"`
	Name                                *String    `xmlrpc:"name,omitempty"`
	NumberOfOpenedSession               *Int       `xmlrpc:"number_of_opened_session,omitempty"`
	OnlyRoundCashMethod                 *Bool      `xmlrpc:"only_round_cash_method,omitempty"`
	OtherDevices                        *Bool      `xmlrpc:"other_devices,omitempty"`
	PartnerLoadBackground               *Bool      `xmlrpc:"partner_load_background,omitempty"`
	PaymentMethodIds                    *Relation  `xmlrpc:"payment_method_ids,omitempty"`
	PickingPolicy                       *Selection `xmlrpc:"picking_policy,omitempty"`
	PickingTypeId                       *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	PosSessionDuration                  *String    `xmlrpc:"pos_session_duration,omitempty"`
	PosSessionState                     *String    `xmlrpc:"pos_session_state,omitempty"`
	PosSessionUsername                  *String    `xmlrpc:"pos_session_username,omitempty"`
	PricelistId                         *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	ProductLoadBackground               *Bool      `xmlrpc:"product_load_background,omitempty"`
	ProxyIp                             *String    `xmlrpc:"proxy_ip,omitempty"`
	ReceiptFooter                       *String    `xmlrpc:"receipt_footer,omitempty"`
	ReceiptHeader                       *String    `xmlrpc:"receipt_header,omitempty"`
	RestrictPriceControl                *Bool      `xmlrpc:"restrict_price_control,omitempty"`
	RoundingMethod                      *Many2One  `xmlrpc:"rounding_method,omitempty"`
	RouteId                             *Many2One  `xmlrpc:"route_id,omitempty"`
	SequenceId                          *Many2One  `xmlrpc:"sequence_id,omitempty"`
	SequenceLineId                      *Many2One  `xmlrpc:"sequence_line_id,omitempty"`
	SessionIds                          *Relation  `xmlrpc:"session_ids,omitempty"`
	SetMaximumDifference                *Bool      `xmlrpc:"set_maximum_difference,omitempty"`
	ShipLater                           *Bool      `xmlrpc:"ship_later,omitempty"`
	StartCategory                       *Bool      `xmlrpc:"start_category,omitempty"`
	TaxRegimeSelection                  *Bool      `xmlrpc:"tax_regime_selection,omitempty"`
	TipProductId                        *Many2One  `xmlrpc:"tip_product_id,omitempty"`
	UsePricelist                        *Bool      `xmlrpc:"use_pricelist,omitempty"`
	Uuid                                *String    `xmlrpc:"uuid,omitempty"`
	WarehouseId                         *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WriteDate                           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PosConfigs represents array of pos.config model.
type PosConfigs []PosConfig

// PosConfigModel is the odoo model name.
const PosConfigModel = "pos.config"

// Many2One convert PosConfig to *Many2One.
func (pc *PosConfig) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreatePosConfig creates a new pos.config model and returns its id.
func (c *Client) CreatePosConfig(pc *PosConfig) (int64, error) {
	ids, err := c.CreatePosConfigs([]*PosConfig{pc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosConfig creates a new pos.config model and returns its id.
func (c *Client) CreatePosConfigs(pcs []*PosConfig) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcs {
		vv = append(vv, v)
	}
	return c.Create(PosConfigModel, vv, nil)
}

// UpdatePosConfig updates an existing pos.config record.
func (c *Client) UpdatePosConfig(pc *PosConfig) error {
	return c.UpdatePosConfigs([]int64{pc.Id.Get()}, pc)
}

// UpdatePosConfigs updates existing pos.config records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdatePosConfigs(ids []int64, pc *PosConfig) error {
	return c.Update(PosConfigModel, ids, pc, nil)
}

// DeletePosConfig deletes an existing pos.config record.
func (c *Client) DeletePosConfig(id int64) error {
	return c.DeletePosConfigs([]int64{id})
}

// DeletePosConfigs deletes existing pos.config records.
func (c *Client) DeletePosConfigs(ids []int64) error {
	return c.Delete(PosConfigModel, ids)
}

// GetPosConfig gets pos.config existing record.
func (c *Client) GetPosConfig(id int64) (*PosConfig, error) {
	pcs, err := c.GetPosConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// GetPosConfigs gets pos.config existing records.
func (c *Client) GetPosConfigs(ids []int64) (*PosConfigs, error) {
	pcs := &PosConfigs{}
	if err := c.Read(PosConfigModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosConfig finds pos.config record by querying it with criteria.
func (c *Client) FindPosConfig(criteria *Criteria) (*PosConfig, error) {
	pcs := &PosConfigs{}
	if err := c.SearchRead(PosConfigModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// FindPosConfigs finds pos.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosConfigs(criteria *Criteria, options *Options) (*PosConfigs, error) {
	pcs := &PosConfigs{}
	if err := c.SearchRead(PosConfigModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosConfigModel, criteria, options)
}

// FindPosConfigId finds record id by querying it with criteria.
func (c *Client) FindPosConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
