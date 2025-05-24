package odoo

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	LastUpdate                                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccountCashBasisBaseAccountId               *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omitempty"`
	AccountDefaultCreditLimit                   *Float     `xmlrpc:"account_default_credit_limit,omitempty"`
	AccountDefaultPosReceivableAccountId        *Many2One  `xmlrpc:"account_default_pos_receivable_account_id,omitempty"`
	AccountFiscalCountryId                      *Many2One  `xmlrpc:"account_fiscal_country_id,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountJournalPaymentCreditAccountId        *Many2One  `xmlrpc:"account_journal_payment_credit_account_id,omitempty"`
	AccountJournalPaymentDebitAccountId         *Many2One  `xmlrpc:"account_journal_payment_debit_account_id,omitempty"`
	AccountJournalSuspenseAccountId             *Many2One  `xmlrpc:"account_journal_suspense_account_id,omitempty"`
	AccountStorno                               *Bool      `xmlrpc:"account_storno,omitempty"`
	AccountUseCreditLimit                       *Bool      `xmlrpc:"account_use_credit_limit,omitempty"`
	ActiveUserCount                             *Int       `xmlrpc:"active_user_count,omitempty"`
	AliasDomain                                 *String    `xmlrpc:"alias_domain,omitempty"`
	AnalyticPlanId                              *Many2One  `xmlrpc:"analytic_plan_id,omitempty"`
	AnnualInventoryDay                          *Int       `xmlrpc:"annual_inventory_day,omitempty"`
	AnnualInventoryMonth                        *Selection `xmlrpc:"annual_inventory_month,omitempty"`
	AuthSignupResetPassword                     *Bool      `xmlrpc:"auth_signup_reset_password,omitempty"`
	AuthSignupTemplateUserId                    *Many2One  `xmlrpc:"auth_signup_template_user_id,omitempty"`
	AuthSignupUninvited                         *Selection `xmlrpc:"auth_signup_uninvited,omitempty"`
	AutomaticInvoice                            *Bool      `xmlrpc:"automatic_invoice,omitempty"`
	BarcodeNomenclatureId                       *Many2One  `xmlrpc:"barcode_nomenclature_id,omitempty"`
	ChartTemplateId                             *Many2One  `xmlrpc:"chart_template_id,omitempty"`
	CompanyCount                                *Int       `xmlrpc:"company_count,omitempty"`
	CompanyId                                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyInformations                         *String    `xmlrpc:"company_informations,omitempty"`
	CompanyName                                 *String    `xmlrpc:"company_name,omitempty"`
	CompanySoTemplateId                         *Many2One  `xmlrpc:"company_so_template_id,omitempty"`
	CountryCode                                 *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmAutoAssignmentAction                     *Selection `xmlrpc:"crm_auto_assignment_action,omitempty"`
	CrmAutoAssignmentIntervalNumber             *Int       `xmlrpc:"crm_auto_assignment_interval_number,omitempty"`
	CrmAutoAssignmentIntervalType               *Selection `xmlrpc:"crm_auto_assignment_interval_type,omitempty"`
	CrmAutoAssignmentRunDatetime                *Time      `xmlrpc:"crm_auto_assignment_run_datetime,omitempty"`
	CrmUseAutoAssignment                        *Bool      `xmlrpc:"crm_use_auto_assignment,omitempty"`
	CurrencyExchangeJournalId                   *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	DefaultInvoicePolicy                        *Selection `xmlrpc:"default_invoice_policy,omitempty"`
	DefaultPickingPolicy                        *Selection `xmlrpc:"default_picking_policy,omitempty"`
	DepositDefaultProductId                     *Many2One  `xmlrpc:"deposit_default_product_id,omitempty"`
	DigestEmails                                *Bool      `xmlrpc:"digest_emails,omitempty"`
	DigestId                                    *Many2One  `xmlrpc:"digest_id,omitempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omitempty"`
	EarlyPayDiscountComputation                 *Selection `xmlrpc:"early_pay_discount_computation,omitempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	ExternalEmailServerDefault                  *Bool      `xmlrpc:"external_email_server_default,omitempty"`
	ExternalReportLayoutId                      *Many2One  `xmlrpc:"external_report_layout_id,omitempty"`
	FailCounter                                 *Int       `xmlrpc:"fail_counter,omitempty"`
	GoogleGmailClientIdentifier                 *String    `xmlrpc:"google_gmail_client_identifier,omitempty"`
	GoogleGmailClientSecret                     *String    `xmlrpc:"google_gmail_client_secret,omitempty"`
	GroupAnalyticAccounting                     *Bool      `xmlrpc:"group_analytic_accounting,omitempty"`
	GroupAutoDoneSetting                        *Bool      `xmlrpc:"group_auto_done_setting,omitempty"`
	GroupCashRounding                           *Bool      `xmlrpc:"group_cash_rounding,omitempty"`
	GroupDiscountPerSoLine                      *Bool      `xmlrpc:"group_discount_per_so_line,omitempty"`
	GroupDisplayIncoterm                        *Bool      `xmlrpc:"group_display_incoterm,omitempty"`
	GroupLotOnDeliverySlip                      *Bool      `xmlrpc:"group_lot_on_delivery_slip,omitempty"`
	GroupLotOnInvoice                           *Bool      `xmlrpc:"group_lot_on_invoice,omitempty"`
	GroupMassMailingCampaign                    *Bool      `xmlrpc:"group_mass_mailing_campaign,omitempty"`
	GroupMultiCurrency                          *Bool      `xmlrpc:"group_multi_currency,omitempty"`
	GroupProductPricelist                       *Bool      `xmlrpc:"group_product_pricelist,omitempty"`
	GroupProductVariant                         *Bool      `xmlrpc:"group_product_variant,omitempty"`
	GroupProformaSales                          *Bool      `xmlrpc:"group_proforma_sales,omitempty"`
	GroupProjectMilestone                       *Bool      `xmlrpc:"group_project_milestone,omitempty"`
	GroupProjectRating                          *Bool      `xmlrpc:"group_project_rating,omitempty"`
	GroupProjectRecurringTasks                  *Bool      `xmlrpc:"group_project_recurring_tasks,omitempty"`
	GroupProjectStages                          *Bool      `xmlrpc:"group_project_stages,omitempty"`
	GroupProjectTaskDependencies                *Bool      `xmlrpc:"group_project_task_dependencies,omitempty"`
	GroupSaleDeliveryAddress                    *Bool      `xmlrpc:"group_sale_delivery_address,omitempty"`
	GroupSaleOrderTemplate                      *Bool      `xmlrpc:"group_sale_order_template,omitempty"`
	GroupSalePricelist                          *Bool      `xmlrpc:"group_sale_pricelist,omitempty"`
	GroupShowLineSubtotalsTaxExcluded           *Bool      `xmlrpc:"group_show_line_subtotals_tax_excluded,omitempty"`
	GroupShowLineSubtotalsTaxIncluded           *Bool      `xmlrpc:"group_show_line_subtotals_tax_included,omitempty"`
	GroupShowPurchaseReceipts                   *Bool      `xmlrpc:"group_show_purchase_receipts,omitempty"`
	GroupShowSaleReceipts                       *Bool      `xmlrpc:"group_show_sale_receipts,omitempty"`
	GroupStockAdvLocation                       *Bool      `xmlrpc:"group_stock_adv_location,omitempty"`
	GroupStockLotPrintGs1                       *Bool      `xmlrpc:"group_stock_lot_print_gs1,omitempty"`
	GroupStockMultiLocations                    *Bool      `xmlrpc:"group_stock_multi_locations,omitempty"`
	GroupStockPackaging                         *Bool      `xmlrpc:"group_stock_packaging,omitempty"`
	GroupStockPickingWave                       *Bool      `xmlrpc:"group_stock_picking_wave,omitempty"`
	GroupStockProductionLot                     *Bool      `xmlrpc:"group_stock_production_lot,omitempty"`
	GroupStockReceptionReport                   *Bool      `xmlrpc:"group_stock_reception_report,omitempty"`
	GroupStockSignDelivery                      *Bool      `xmlrpc:"group_stock_sign_delivery,omitempty"`
	GroupStockStorageCategories                 *Bool      `xmlrpc:"group_stock_storage_categories,omitempty"`
	GroupStockTrackingLot                       *Bool      `xmlrpc:"group_stock_tracking_lot,omitempty"`
	GroupStockTrackingOwner                     *Bool      `xmlrpc:"group_stock_tracking_owner,omitempty"`
	GroupSubtaskProject                         *Bool      `xmlrpc:"group_subtask_project,omitempty"`
	GroupUom                                    *Bool      `xmlrpc:"group_uom,omitempty"`
	GroupUseLead                                *Bool      `xmlrpc:"group_use_lead,omitempty"`
	GroupUseRecurringRevenues                   *Bool      `xmlrpc:"group_use_recurring_revenues,omitempty"`
	GroupWarningAccount                         *Bool      `xmlrpc:"group_warning_account,omitempty"`
	GroupWarningSale                            *Bool      `xmlrpc:"group_warning_sale,omitempty"`
	GroupWarningStock                           *Bool      `xmlrpc:"group_warning_stock,omitempty"`
	HasAccountingEntries                        *Bool      `xmlrpc:"has_accounting_entries,omitempty"`
	HasChartOfAccounts                          *Bool      `xmlrpc:"has_chart_of_accounts,omitempty"`
	HrEmployeeSelfEdit                          *Bool      `xmlrpc:"hr_employee_self_edit,omitempty"`
	HrPresenceControlEmail                      *Bool      `xmlrpc:"hr_presence_control_email,omitempty"`
	HrPresenceControlEmailAmount                *Int       `xmlrpc:"hr_presence_control_email_amount,omitempty"`
	HrPresenceControlIp                         *Bool      `xmlrpc:"hr_presence_control_ip,omitempty"`
	HrPresenceControlIpList                     *String    `xmlrpc:"hr_presence_control_ip_list,omitempty"`
	HrPresenceControlLogin                      *Bool      `xmlrpc:"hr_presence_control_login,omitempty"`
	Id                                          *Int       `xmlrpc:"id,omitempty"`
	IncomeCurrencyExchangeAccountId             *Many2One  `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	IncotermId                                  *Many2One  `xmlrpc:"incoterm_id,omitempty"`
	InvoiceIsEmail                              *Bool      `xmlrpc:"invoice_is_email,omitempty"`
	InvoiceIsPrint                              *Bool      `xmlrpc:"invoice_is_print,omitempty"`
	InvoiceIsSnailmail                          *Bool      `xmlrpc:"invoice_is_snailmail,omitempty"`
	InvoiceMailTemplateId                       *Many2One  `xmlrpc:"invoice_mail_template_id,omitempty"`
	InvoiceTerms                                *String    `xmlrpc:"invoice_terms,omitempty"`
	InvoiceTermsHtml                            *String    `xmlrpc:"invoice_terms_html,omitempty"`
	IsDefaultPricelistDisplayed                 *Bool      `xmlrpc:"is_default_pricelist_displayed,omitempty"`
	IsMembershipMulti                           *Bool      `xmlrpc:"is_membership_multi,omitempty"`
	LanguageCount                               *Int       `xmlrpc:"language_count,omitempty"`
	LeadEnrichAuto                              *Selection `xmlrpc:"lead_enrich_auto,omitempty"`
	LeadMiningInPipeline                        *Bool      `xmlrpc:"lead_mining_in_pipeline,omitempty"`
	MassMailingMailServerId                     *Many2One  `xmlrpc:"mass_mailing_mail_server_id,omitempty"`
	MassMailingOutgoingMailServer               *Bool      `xmlrpc:"mass_mailing_outgoing_mail_server,omitempty"`
	MassMailingReports                          *Bool      `xmlrpc:"mass_mailing_reports,omitempty"`
	ModuleAccountAccountant                     *Bool      `xmlrpc:"module_account_accountant,omitempty"`
	ModuleAccountBankStatementImportCamt        *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omitempty"`
	ModuleAccountBankStatementImportCsv         *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omitempty"`
	ModuleAccountBankStatementImportOfx         *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omitempty"`
	ModuleAccountBankStatementImportQif         *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omitempty"`
	ModuleAccountBatchPayment                   *Bool      `xmlrpc:"module_account_batch_payment,omitempty"`
	ModuleAccountBudget                         *Bool      `xmlrpc:"module_account_budget,omitempty"`
	ModuleAccountCheckPrinting                  *Bool      `xmlrpc:"module_account_check_printing,omitempty"`
	ModuleAccountInterCompanyRules              *Bool      `xmlrpc:"module_account_inter_company_rules,omitempty"`
	ModuleAccountIntrastat                      *Bool      `xmlrpc:"module_account_intrastat,omitempty"`
	ModuleAccountInvoiceExtract                 *Bool      `xmlrpc:"module_account_invoice_extract,omitempty"`
	ModuleAccountPayment                        *Bool      `xmlrpc:"module_account_payment,omitempty"`
	ModuleAccountReports                        *Bool      `xmlrpc:"module_account_reports,omitempty"`
	ModuleAccountSepa                           *Bool      `xmlrpc:"module_account_sepa,omitempty"`
	ModuleAccountSepaDirectDebit                *Bool      `xmlrpc:"module_account_sepa_direct_debit,omitempty"`
	ModuleAccountTaxcloud                       *Bool      `xmlrpc:"module_account_taxcloud,omitempty"`
	ModuleAuthLdap                              *Bool      `xmlrpc:"module_auth_ldap,omitempty"`
	ModuleAuthOauth                             *Bool      `xmlrpc:"module_auth_oauth,omitempty"`
	ModuleBaseGengo                             *Bool      `xmlrpc:"module_base_gengo,omitempty"`
	ModuleBaseGeolocalize                       *Bool      `xmlrpc:"module_base_geolocalize,omitempty"`
	ModuleBaseImport                            *Bool      `xmlrpc:"module_base_import,omitempty"`
	ModuleCrmIapEnrich                          *Bool      `xmlrpc:"module_crm_iap_enrich,omitempty"`
	ModuleCrmIapMine                            *Bool      `xmlrpc:"module_crm_iap_mine,omitempty"`
	ModuleCurrencyRateLive                      *Bool      `xmlrpc:"module_currency_rate_live,omitempty"`
	ModuleDelivery                              *Bool      `xmlrpc:"module_delivery,omitempty"`
	ModuleDeliveryBpost                         *Bool      `xmlrpc:"module_delivery_bpost,omitempty"`
	ModuleDeliveryDhl                           *Bool      `xmlrpc:"module_delivery_dhl,omitempty"`
	ModuleDeliveryEasypost                      *Bool      `xmlrpc:"module_delivery_easypost,omitempty"`
	ModuleDeliveryFedex                         *Bool      `xmlrpc:"module_delivery_fedex,omitempty"`
	ModuleDeliverySendcloud                     *Bool      `xmlrpc:"module_delivery_sendcloud,omitempty"`
	ModuleDeliveryUps                           *Bool      `xmlrpc:"module_delivery_ups,omitempty"`
	ModuleDeliveryUsps                          *Bool      `xmlrpc:"module_delivery_usps,omitempty"`
	ModuleGoogleCalendar                        *Bool      `xmlrpc:"module_google_calendar,omitempty"`
	ModuleGoogleGmail                           *Bool      `xmlrpc:"module_google_gmail,omitempty"`
	ModuleGoogleRecaptcha                       *Bool      `xmlrpc:"module_google_recaptcha,omitempty"`
	ModuleHrAttendance                          *Bool      `xmlrpc:"module_hr_attendance,omitempty"`
	ModuleHrHomeworking                         *Bool      `xmlrpc:"module_hr_homeworking,omitempty"`
	ModuleHrPresence                            *Bool      `xmlrpc:"module_hr_presence,omitempty"`
	ModuleHrSkills                              *Bool      `xmlrpc:"module_hr_skills,omitempty"`
	ModuleHrTimesheet                           *Bool      `xmlrpc:"module_hr_timesheet,omitempty"`
	ModuleL10NEuOss                             *Bool      `xmlrpc:"module_l10n_eu_oss,omitempty"`
	ModuleLoyalty                               *Bool      `xmlrpc:"module_loyalty,omitempty"`
	ModuleMailPlugin                            *Bool      `xmlrpc:"module_mail_plugin,omitempty"`
	ModuleMicrosoftCalendar                     *Bool      `xmlrpc:"module_microsoft_calendar,omitempty"`
	ModuleMicrosoftOutlook                      *Bool      `xmlrpc:"module_microsoft_outlook,omitempty"`
	ModulePartnerAutocomplete                   *Bool      `xmlrpc:"module_partner_autocomplete,omitempty"`
	ModulePosAdyen                              *Bool      `xmlrpc:"module_pos_adyen,omitempty"`
	ModulePosMercury                            *Bool      `xmlrpc:"module_pos_mercury,omitempty"`
	ModulePosSix                                *Bool      `xmlrpc:"module_pos_six,omitempty"`
	ModulePosStripe                             *Bool      `xmlrpc:"module_pos_stripe,omitempty"`
	ModuleProductEmailTemplate                  *Bool      `xmlrpc:"module_product_email_template,omitempty"`
	ModuleProductExpiry                         *Bool      `xmlrpc:"module_product_expiry,omitempty"`
	ModuleProductImages                         *Bool      `xmlrpc:"module_product_images,omitempty"`
	ModuleProductMargin                         *Bool      `xmlrpc:"module_product_margin,omitempty"`
	ModuleProjectForecast                       *Bool      `xmlrpc:"module_project_forecast,omitempty"`
	ModuleQualityControl                        *Bool      `xmlrpc:"module_quality_control,omitempty"`
	ModuleQualityControlWorksheet               *Bool      `xmlrpc:"module_quality_control_worksheet,omitempty"`
	ModuleSaleAmazon                            *Bool      `xmlrpc:"module_sale_amazon,omitempty"`
	ModuleSaleLoyalty                           *Bool      `xmlrpc:"module_sale_loyalty,omitempty"`
	ModuleSaleMargin                            *Bool      `xmlrpc:"module_sale_margin,omitempty"`
	ModuleSaleProductMatrix                     *Bool      `xmlrpc:"module_sale_product_matrix,omitempty"`
	ModuleSaleQuotationBuilder                  *Bool      `xmlrpc:"module_sale_quotation_builder,omitempty"`
	ModuleSnailmailAccount                      *Bool      `xmlrpc:"module_snailmail_account,omitempty"`
	ModuleStockBarcode                          *Bool      `xmlrpc:"module_stock_barcode,omitempty"`
	ModuleStockLandedCosts                      *Bool      `xmlrpc:"module_stock_landed_costs,omitempty"`
	ModuleStockPickingBatch                     *Bool      `xmlrpc:"module_stock_picking_batch,omitempty"`
	ModuleStockSms                              *Bool      `xmlrpc:"module_stock_sms,omitempty"`
	ModuleVoip                                  *Bool      `xmlrpc:"module_voip,omitempty"`
	ModuleWebUnsplash                           *Bool      `xmlrpc:"module_web_unsplash,omitempty"`
	ModuleWebsiteCrmIapReveal                   *Bool      `xmlrpc:"module_website_crm_iap_reveal,omitempty"`
	PartnerAutocompleteInsufficientCredit       *Bool      `xmlrpc:"partner_autocomplete_insufficient_credit,omitempty"`
	PayInvoicesOnline                           *Bool      `xmlrpc:"pay_invoices_online,omitempty"`
	PointOfSaleUseTicketQrCode                  *Bool      `xmlrpc:"point_of_sale_use_ticket_qr_code,omitempty"`
	PortalAllowApiKeys                          *Bool      `xmlrpc:"portal_allow_api_keys,omitempty"`
	PortalConfirmationPay                       *Bool      `xmlrpc:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                      *Bool      `xmlrpc:"portal_confirmation_sign,omitempty"`
	PosAllowedPricelistIds                      *Relation  `xmlrpc:"pos_allowed_pricelist_ids,omitempty"`
	PosAmountAuthorizedDiff                     *Float     `xmlrpc:"pos_amount_authorized_diff,omitempty"`
	PosAvailablePricelistIds                    *Relation  `xmlrpc:"pos_available_pricelist_ids,omitempty"`
	PosCashControl                              *Bool      `xmlrpc:"pos_cash_control,omitempty"`
	PosCashRounding                             *Bool      `xmlrpc:"pos_cash_rounding,omitempty"`
	PosCompanyHasTemplate                       *Bool      `xmlrpc:"pos_company_has_template,omitempty"`
	PosConfigId                                 *Many2One  `xmlrpc:"pos_config_id,omitempty"`
	PosCrmTeamId                                *Many2One  `xmlrpc:"pos_crm_team_id,omitempty"`
	PosDefaultBillIds                           *Relation  `xmlrpc:"pos_default_bill_ids,omitempty"`
	PosDefaultFiscalPositionId                  *Many2One  `xmlrpc:"pos_default_fiscal_position_id,omitempty"`
	PosDownPaymentProductId                     *Many2One  `xmlrpc:"pos_down_payment_product_id,omitempty"`
	PosEmployeeIds                              *Relation  `xmlrpc:"pos_employee_ids,omitempty"`
	PosEpsonPrinterIp                           *String    `xmlrpc:"pos_epson_printer_ip,omitempty"`
	PosFiscalPositionIds                        *Relation  `xmlrpc:"pos_fiscal_position_ids,omitempty"`
	PosHasActiveSession                         *Bool      `xmlrpc:"pos_has_active_session,omitempty"`
	PosIfaceAvailableCategIds                   *Relation  `xmlrpc:"pos_iface_available_categ_ids,omitempty"`
	PosIfaceBigScrollbars                       *Bool      `xmlrpc:"pos_iface_big_scrollbars,omitempty"`
	PosIfaceCashdrawer                          *Bool      `xmlrpc:"pos_iface_cashdrawer,omitempty"`
	PosIfaceCustomerFacingDisplayLocal          *Bool      `xmlrpc:"pos_iface_customer_facing_display_local,omitempty"`
	PosIfaceCustomerFacingDisplayViaProxy       *Bool      `xmlrpc:"pos_iface_customer_facing_display_via_proxy,omitempty"`
	PosIfaceElectronicScale                     *Bool      `xmlrpc:"pos_iface_electronic_scale,omitempty"`
	PosIfacePrintAuto                           *Bool      `xmlrpc:"pos_iface_print_auto,omitempty"`
	PosIfacePrintSkipScreen                     *Bool      `xmlrpc:"pos_iface_print_skip_screen,omitempty"`
	PosIfacePrintViaProxy                       *Bool      `xmlrpc:"pos_iface_print_via_proxy,omitempty"`
	PosIfaceScanViaProxy                        *Bool      `xmlrpc:"pos_iface_scan_via_proxy,omitempty"`
	PosIfaceStartCategId                        *Many2One  `xmlrpc:"pos_iface_start_categ_id,omitempty"`
	PosIfaceTaxIncluded                         *Selection `xmlrpc:"pos_iface_tax_included,omitempty"`
	PosIfaceTipproduct                          *Bool      `xmlrpc:"pos_iface_tipproduct,omitempty"`
	PosInvoiceJournalId                         *Many2One  `xmlrpc:"pos_invoice_journal_id,omitempty"`
	PosIsHeaderOrFooter                         *Bool      `xmlrpc:"pos_is_header_or_footer,omitempty"`
	PosIsMarginsCostsAccessibleToEveryUser      *Bool      `xmlrpc:"pos_is_margins_costs_accessible_to_every_user,omitempty"`
	PosIsPosbox                                 *Bool      `xmlrpc:"pos_is_posbox,omitempty"`
	PosJournalId                                *Many2One  `xmlrpc:"pos_journal_id,omitempty"`
	PosLimitCategories                          *Bool      `xmlrpc:"pos_limit_categories,omitempty"`
	PosLimitedPartnersAmount                    *Int       `xmlrpc:"pos_limited_partners_amount,omitempty"`
	PosLimitedPartnersLoading                   *Bool      `xmlrpc:"pos_limited_partners_loading,omitempty"`
	PosLimitedProductsAmount                    *Int       `xmlrpc:"pos_limited_products_amount,omitempty"`
	PosLimitedProductsLoading                   *Bool      `xmlrpc:"pos_limited_products_loading,omitempty"`
	PosManualDiscount                           *Bool      `xmlrpc:"pos_manual_discount,omitempty"`
	PosModulePosDiscount                        *Bool      `xmlrpc:"pos_module_pos_discount,omitempty"`
	PosModulePosHr                              *Bool      `xmlrpc:"pos_module_pos_hr,omitempty"`
	PosModulePosRestaurant                      *Bool      `xmlrpc:"pos_module_pos_restaurant,omitempty"`
	PosOnlyRoundCashMethod                      *Bool      `xmlrpc:"pos_only_round_cash_method,omitempty"`
	PosOtherDevices                             *Bool      `xmlrpc:"pos_other_devices,omitempty"`
	PosPartnerLoadBackground                    *Bool      `xmlrpc:"pos_partner_load_background,omitempty"`
	PosPaymentMethodIds                         *Relation  `xmlrpc:"pos_payment_method_ids,omitempty"`
	PosPickingPolicy                            *Selection `xmlrpc:"pos_picking_policy,omitempty"`
	PosPickingTypeId                            *Many2One  `xmlrpc:"pos_picking_type_id,omitempty"`
	PosPricelistId                              *Many2One  `xmlrpc:"pos_pricelist_id,omitempty"`
	PosProductLoadBackground                    *Bool      `xmlrpc:"pos_product_load_background,omitempty"`
	PosProxyIp                                  *String    `xmlrpc:"pos_proxy_ip,omitempty"`
	PosReceiptFooter                            *String    `xmlrpc:"pos_receipt_footer,omitempty"`
	PosReceiptHeader                            *String    `xmlrpc:"pos_receipt_header,omitempty"`
	PosRestrictPriceControl                     *Bool      `xmlrpc:"pos_restrict_price_control,omitempty"`
	PosRoundingMethod                           *Many2One  `xmlrpc:"pos_rounding_method,omitempty"`
	PosRouteId                                  *Many2One  `xmlrpc:"pos_route_id,omitempty"`
	PosSelectableCategIds                       *Relation  `xmlrpc:"pos_selectable_categ_ids,omitempty"`
	PosSequenceId                               *Many2One  `xmlrpc:"pos_sequence_id,omitempty"`
	PosSetMaximumDifference                     *Bool      `xmlrpc:"pos_set_maximum_difference,omitempty"`
	PosShipLater                                *Bool      `xmlrpc:"pos_ship_later,omitempty"`
	PosStartCategory                            *Bool      `xmlrpc:"pos_start_category,omitempty"`
	PosTaxRegimeSelection                       *Bool      `xmlrpc:"pos_tax_regime_selection,omitempty"`
	PosTipProductId                             *Many2One  `xmlrpc:"pos_tip_product_id,omitempty"`
	PosUsePricelist                             *Bool      `xmlrpc:"pos_use_pricelist,omitempty"`
	PosWarehouseId                              *Many2One  `xmlrpc:"pos_warehouse_id,omitempty"`
	PredictiveLeadScoringFieldLabels            *String    `xmlrpc:"predictive_lead_scoring_field_labels,omitempty"`
	PredictiveLeadScoringFields                 *Relation  `xmlrpc:"predictive_lead_scoring_fields,omitempty"`
	PredictiveLeadScoringFieldsStr              *String    `xmlrpc:"predictive_lead_scoring_fields_str,omitempty"`
	PredictiveLeadScoringStartDate              *Time      `xmlrpc:"predictive_lead_scoring_start_date,omitempty"`
	PredictiveLeadScoringStartDateStr           *String    `xmlrpc:"predictive_lead_scoring_start_date_str,omitempty"`
	PreviewReady                                *Bool      `xmlrpc:"preview_ready,omitempty"`
	PrimaryColor                                *String    `xmlrpc:"primary_color,omitempty"`
	ProductPricelistSetting                     *Selection `xmlrpc:"product_pricelist_setting,omitempty"`
	ProductVolumeVolumeInCubicFeet              *Selection `xmlrpc:"product_volume_volume_in_cubic_feet,omitempty"`
	ProductWeightInLbs                          *Selection `xmlrpc:"product_weight_in_lbs,omitempty"`
	ProfilingEnabledUntil                       *Time      `xmlrpc:"profiling_enabled_until,omitempty"`
	PurchaseTaxId                               *Many2One  `xmlrpc:"purchase_tax_id,omitempty"`
	QrCode                                      *Bool      `xmlrpc:"qr_code,omitempty"`
	QuickEditMode                               *Selection `xmlrpc:"quick_edit_mode,omitempty"`
	QuotationValidityDays                       *Int       `xmlrpc:"quotation_validity_days,omitempty"`
	ReportFooter                                *String    `xmlrpc:"report_footer,omitempty"`
	ResourceCalendarId                          *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	RestrictTemplateRendering                   *Bool      `xmlrpc:"restrict_template_rendering,omitempty"`
	SaleTaxId                                   *Many2One  `xmlrpc:"sale_tax_id,omitempty"`
	SecondaryColor                              *String    `xmlrpc:"secondary_color,omitempty"`
	SecurityLead                                *Float     `xmlrpc:"security_lead,omitempty"`
	ShowBlacklistButtons                        *Bool      `xmlrpc:"show_blacklist_buttons,omitempty"`
	ShowEffect                                  *Bool      `xmlrpc:"show_effect,omitempty"`
	ShowLineSubtotalsTaxSelection               *Selection `xmlrpc:"show_line_subtotals_tax_selection,omitempty"`
	SnailmailColor                              *Bool      `xmlrpc:"snailmail_color,omitempty"`
	SnailmailCover                              *Bool      `xmlrpc:"snailmail_cover,omitempty"`
	SnailmailDuplex                             *Bool      `xmlrpc:"snailmail_duplex,omitempty"`
	StockMoveEmailValidation                    *Bool      `xmlrpc:"stock_move_email_validation,omitempty"`
	StockMoveSmsValidation                      *Bool      `xmlrpc:"stock_move_sms_validation,omitempty"`
	StockSmsConfirmationTemplateId              *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omitempty"`
	TaxCalculationRoundingMethod                *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                       *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                              *Bool      `xmlrpc:"tax_exigibility,omitempty"`
	TermsType                                   *Selection `xmlrpc:"terms_type,omitempty"`
	TransferAccountId                           *Many2One  `xmlrpc:"transfer_account_id,omitempty"`
	TwilioAccountSid                            *String    `xmlrpc:"twilio_account_sid,omitempty"`
	TwilioAccountToken                          *String    `xmlrpc:"twilio_account_token,omitempty"`
	UnsplashAccessKey                           *String    `xmlrpc:"unsplash_access_key,omitempty"`
	UnsplashAppId                               *String    `xmlrpc:"unsplash_app_id,omitempty"`
	UpdateStockQuantities                       *Selection `xmlrpc:"update_stock_quantities,omitempty"`
	UseInvoiceTerms                             *Bool      `xmlrpc:"use_invoice_terms,omitempty"`
	UseQuotationValidityDays                    *Bool      `xmlrpc:"use_quotation_validity_days,omitempty"`
	UseSecurityLead                             *Bool      `xmlrpc:"use_security_lead,omitempty"`
	UseTwilioRtcServers                         *Bool      `xmlrpc:"use_twilio_rtc_servers,omitempty"`
	UserDefaultRights                           *Bool      `xmlrpc:"user_default_rights,omitempty"`
	VatCheckVies                                *Bool      `xmlrpc:"vat_check_vies,omitempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResConfigSettingss represents array of res.config.settings model.
type ResConfigSettingss []ResConfigSettings

// ResConfigSettingsModel is the odoo model name.
const ResConfigSettingsModel = "res.config.settings"

// Many2One convert ResConfigSettings to *Many2One.
func (rcs *ResConfigSettings) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettings(rcs *ResConfigSettings) (int64, error) {
	ids, err := c.CreateResConfigSettingss([]*ResConfigSettings{rcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettingss(rcss []*ResConfigSettings) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcss {
		vv = append(vv, v)
	}
	return c.Create(ResConfigSettingsModel, vv, nil)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs, nil)
}

// DeleteResConfigSettings deletes an existing res.config.settings record.
func (c *Client) DeleteResConfigSettings(id int64) error {
	return c.DeleteResConfigSettingss([]int64{id})
}

// DeleteResConfigSettingss deletes existing res.config.settings records.
func (c *Client) DeleteResConfigSettingss(ids []int64) error {
	return c.Delete(ResConfigSettingsModel, ids)
}

// GetResConfigSettings gets res.config.settings existing record.
func (c *Client) GetResConfigSettings(id int64) (*ResConfigSettings, error) {
	rcss, err := c.GetResConfigSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// GetResConfigSettingss gets res.config.settings existing records.
func (c *Client) GetResConfigSettingss(ids []int64) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.Read(ResConfigSettingsModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettings finds res.config.settings record by querying it with criteria.
func (c *Client) FindResConfigSettings(criteria *Criteria) (*ResConfigSettings, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// FindResConfigSettingss finds res.config.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingss(criteria *Criteria, options *Options) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResConfigSettingsModel, criteria, options)
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
