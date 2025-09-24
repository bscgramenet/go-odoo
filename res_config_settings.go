package odoo

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	AccountCashBasisBaseAccountId               *Many2One   `xmlrpc:"account_cash_basis_base_account_id,omitempty"`
	AccountDefaultCreditLimit                   *Float      `xmlrpc:"account_default_credit_limit,omitempty"`
	AccountDefaultPosReceivableAccountId        *Many2One   `xmlrpc:"account_default_pos_receivable_account_id,omitempty"`
	AccountDiscountExpenseAllocationId          *Many2One   `xmlrpc:"account_discount_expense_allocation_id,omitempty"`
	AccountDiscountIncomeAllocationId           *Many2One   `xmlrpc:"account_discount_income_allocation_id,omitempty"`
	AccountFiscalCountryId                      *Many2One   `xmlrpc:"account_fiscal_country_id,omitempty"`
	AccountIsTokenOutOfSync                     *Bool       `xmlrpc:"account_is_token_out_of_sync,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One   `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One   `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountJournalSuspenseAccountId             *Many2One   `xmlrpc:"account_journal_suspense_account_id,omitempty"`
	AccountPeppolContactEmail                   *String     `xmlrpc:"account_peppol_contact_email,omitempty"`
	AccountPeppolEas                            *Selection  `xmlrpc:"account_peppol_eas,omitempty"`
	AccountPeppolEdiIdentification              *String     `xmlrpc:"account_peppol_edi_identification,omitempty"`
	AccountPeppolEdiMode                        *Selection  `xmlrpc:"account_peppol_edi_mode,omitempty"`
	AccountPeppolEdiUser                        *Many2One   `xmlrpc:"account_peppol_edi_user,omitempty"`
	AccountPeppolEndpoint                       *String     `xmlrpc:"account_peppol_endpoint,omitempty"`
	AccountPeppolMigrationKey                   *String     `xmlrpc:"account_peppol_migration_key,omitempty"`
	AccountPeppolPhoneNumber                    *String     `xmlrpc:"account_peppol_phone_number,omitempty"`
	AccountPeppolProxyState                     *Selection  `xmlrpc:"account_peppol_proxy_state,omitempty"`
	AccountPeppolPurchaseJournalId              *Many2One   `xmlrpc:"account_peppol_purchase_journal_id,omitempty"`
	AccountPriceInclude                         *Selection  `xmlrpc:"account_price_include,omitempty"`
	AccountStorno                               *Bool       `xmlrpc:"account_storno,omitempty"`
	AccountUseCreditLimit                       *Bool       `xmlrpc:"account_use_credit_limit,omitempty"`
	ActiveProviderId                            *Many2One   `xmlrpc:"active_provider_id,omitempty"`
	ActiveUserCount                             *Int        `xmlrpc:"active_user_count,omitempty"`
	AliasDomainId                               *Many2One   `xmlrpc:"alias_domain_id,omitempty"`
	AnnualInventoryDay                          *Int        `xmlrpc:"annual_inventory_day,omitempty"`
	AnnualInventoryMonth                        *Selection  `xmlrpc:"annual_inventory_month,omitempty"`
	AuthSignupResetPassword                     *Bool       `xmlrpc:"auth_signup_reset_password,omitempty"`
	AuthSignupTemplateUserId                    *Many2One   `xmlrpc:"auth_signup_template_user_id,omitempty"`
	AuthSignupUninvited                         *Selection  `xmlrpc:"auth_signup_uninvited,omitempty"`
	AuthTotpEnforce                             *Bool       `xmlrpc:"auth_totp_enforce,omitempty"`
	AuthTotpPolicy                              *Selection  `xmlrpc:"auth_totp_policy,omitempty"`
	AutomaticInvoice                            *Bool       `xmlrpc:"automatic_invoice,omitempty"`
	AutopostBills                               *Bool       `xmlrpc:"autopost_bills,omitempty"`
	BarcodeNomenclatureId                       *Many2One   `xmlrpc:"barcode_nomenclature_id,omitempty"`
	BarcodeSeparator                            *String     `xmlrpc:"barcode_separator,omitempty"`
	ChartTemplate                               *Selection  `xmlrpc:"chart_template,omitempty"`
	CompanyCount                                *Int        `xmlrpc:"company_count,omitempty"`
	CompanyCountryCode                          *String     `xmlrpc:"company_country_code,omitempty"`
	CompanyCountryGroupCodes                    interface{} `xmlrpc:"company_country_group_codes,omitempty"`
	CompanyCurrencyId                           *Many2One   `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                                   *Many2One   `xmlrpc:"company_id,omitempty"`
	CompanyInformations                         *String     `xmlrpc:"company_informations,omitempty"`
	CompanyName                                 *String     `xmlrpc:"company_name,omitempty"`
	CompanySoTemplateId                         *Many2One   `xmlrpc:"company_so_template_id,omitempty"`
	CountryCode                                 *String     `xmlrpc:"country_code,omitempty"`
	CreateDate                                  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyExchangeJournalId                   *Many2One   `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                                  *Many2One   `xmlrpc:"currency_id,omitempty"`
	DaysToPurchase                              *Float      `xmlrpc:"days_to_purchase,omitempty"`
	DefaultInvoicePolicy                        *Selection  `xmlrpc:"default_invoice_policy,omitempty"`
	DefaultPickingPolicy                        *Selection  `xmlrpc:"default_picking_policy,omitempty"`
	DigestEmails                                *Bool       `xmlrpc:"digest_emails,omitempty"`
	DigestId                                    *Many2One   `xmlrpc:"digest_id,omitempty"`
	DisplayAccountStorno                        *Bool       `xmlrpc:"display_account_storno,omitempty"`
	DisplayInvoiceAmountTotalWords              *Bool       `xmlrpc:"display_invoice_amount_total_words,omitempty"`
	DisplayInvoiceTaxCompanyCurrency            *Bool       `xmlrpc:"display_invoice_tax_company_currency,omitempty"`
	DisplayName                                 *String     `xmlrpc:"display_name,omitempty"`
	DownpaymentAccountId                        *Many2One   `xmlrpc:"downpayment_account_id,omitempty"`
	EmailPrimaryColor                           *String     `xmlrpc:"email_primary_color,omitempty"`
	EmailSecondaryColor                         *String     `xmlrpc:"email_secondary_color,omitempty"`
	ExpenseAccountId                            *Many2One   `xmlrpc:"expense_account_id,omitempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One   `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	ExternalEmailServerDefault                  *Bool       `xmlrpc:"external_email_server_default,omitempty"`
	ExternalReportLayoutId                      *Many2One   `xmlrpc:"external_report_layout_id,omitempty"`
	FailCounter                                 *Int        `xmlrpc:"fail_counter,omitempty"`
	ForceRestrictiveAuditTrail                  *Bool       `xmlrpc:"force_restrictive_audit_trail,omitempty"`
	GoogleGmailClientIdentifier                 *String     `xmlrpc:"google_gmail_client_identifier,omitempty"`
	GoogleGmailClientSecret                     *String     `xmlrpc:"google_gmail_client_secret,omitempty"`
	GooglePlacesApiKey                          *String     `xmlrpc:"google_places_api_key,omitempty"`
	GoogleTranslateApiKey                       *String     `xmlrpc:"google_translate_api_key,omitempty"`
	GroupAnalyticAccounting                     *Bool       `xmlrpc:"group_analytic_accounting,omitempty"`
	GroupAutoDoneSetting                        *Bool       `xmlrpc:"group_auto_done_setting,omitempty"`
	GroupCashRounding                           *Bool       `xmlrpc:"group_cash_rounding,omitempty"`
	GroupDiscountPerSoLine                      *Bool       `xmlrpc:"group_discount_per_so_line,omitempty"`
	GroupLotOnDeliverySlip                      *Bool       `xmlrpc:"group_lot_on_delivery_slip,omitempty"`
	GroupLotOnInvoice                           *Bool       `xmlrpc:"group_lot_on_invoice,omitempty"`
	GroupMultiCurrency                          *Bool       `xmlrpc:"group_multi_currency,omitempty"`
	GroupPosPreset                              *Bool       `xmlrpc:"group_pos_preset,omitempty"`
	GroupProductPricelist                       *Bool       `xmlrpc:"group_product_pricelist,omitempty"`
	GroupProductVariant                         *Bool       `xmlrpc:"group_product_variant,omitempty"`
	GroupProformaSales                          *Bool       `xmlrpc:"group_proforma_sales,omitempty"`
	GroupSaleDeliveryAddress                    *Bool       `xmlrpc:"group_sale_delivery_address,omitempty"`
	GroupSaleOrderTemplate                      *Bool       `xmlrpc:"group_sale_order_template,omitempty"`
	GroupSendReminder                           *Bool       `xmlrpc:"group_send_reminder,omitempty"`
	GroupStockAdvLocation                       *Bool       `xmlrpc:"group_stock_adv_location,omitempty"`
	GroupStockLotPrintGs1                       *Bool       `xmlrpc:"group_stock_lot_print_gs1,omitempty"`
	GroupStockMultiLocations                    *Bool       `xmlrpc:"group_stock_multi_locations,omitempty"`
	GroupStockProductionLot                     *Bool       `xmlrpc:"group_stock_production_lot,omitempty"`
	GroupStockReceptionReport                   *Bool       `xmlrpc:"group_stock_reception_report,omitempty"`
	GroupStockSignDelivery                      *Bool       `xmlrpc:"group_stock_sign_delivery,omitempty"`
	GroupStockTrackingLot                       *Bool       `xmlrpc:"group_stock_tracking_lot,omitempty"`
	GroupStockTrackingOwner                     *Bool       `xmlrpc:"group_stock_tracking_owner,omitempty"`
	GroupUom                                    *Bool       `xmlrpc:"group_uom,omitempty"`
	GroupWarningPurchase                        *Bool       `xmlrpc:"group_warning_purchase,omitempty"`
	GroupWarningSale                            *Bool       `xmlrpc:"group_warning_sale,omitempty"`
	GroupWarningStock                           *Bool       `xmlrpc:"group_warning_stock,omitempty"`
	HasAccountingEntries                        *Bool       `xmlrpc:"has_accounting_entries,omitempty"`
	HasChartOfAccounts                          *Bool       `xmlrpc:"has_chart_of_accounts,omitempty"`
	HasEnabledProvider                          *Bool       `xmlrpc:"has_enabled_provider,omitempty"`
	HorizonDays                                 *Float      `xmlrpc:"horizon_days,omitempty"`
	Id                                          *Int        `xmlrpc:"id,omitempty"`
	IncomeAccountId                             *Many2One   `xmlrpc:"income_account_id,omitempty"`
	IncomeCurrencyExchangeAccountId             *Many2One   `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	IncotermId                                  *Many2One   `xmlrpc:"incoterm_id,omitempty"`
	InvoiceMailTemplateId                       *Many2One   `xmlrpc:"invoice_mail_template_id,omitempty"`
	InvoiceTerms                                *String     `xmlrpc:"invoice_terms,omitempty"`
	InvoiceTermsHtml                            *String     `xmlrpc:"invoice_terms_html,omitempty"`
	IsAccountPeppolEligible                     *Bool       `xmlrpc:"is_account_peppol_eligible,omitempty"`
	IsInstalledSale                             *Bool       `xmlrpc:"is_installed_sale,omitempty"`
	IsKioskMode                                 *Bool       `xmlrpc:"is_kiosk_mode,omitempty"`
	IsRootCompany                               *Bool       `xmlrpc:"is_root_company,omitempty"`
	L10NEsEdiVerifactuCertificateIds            *Relation   `xmlrpc:"l10n_es_edi_verifactu_certificate_ids,omitempty"`
	L10NEsEdiVerifactuRequired                  *Bool       `xmlrpc:"l10n_es_edi_verifactu_required,omitempty"`
	L10NEsEdiVerifactuSpecialVatRegime          *Selection  `xmlrpc:"l10n_es_edi_verifactu_special_vat_regime,omitempty"`
	L10NEsEdiVerifactuTestEnvironment           *Bool       `xmlrpc:"l10n_es_edi_verifactu_test_environment,omitempty"`
	L10NEsSimplifiedInvoiceLimit                *Float      `xmlrpc:"l10n_es_simplified_invoice_limit,omitempty"`
	LanguageCount                               *Int        `xmlrpc:"language_count,omitempty"`
	LinkQrCode                                  *Bool       `xmlrpc:"link_qr_code,omitempty"`
	LockConfirmedPo                             *Bool       `xmlrpc:"lock_confirmed_po,omitempty"`
	MicrosoftOutlookClientIdentifier            *String     `xmlrpc:"microsoft_outlook_client_identifier,omitempty"`
	MicrosoftOutlookClientSecret                *String     `xmlrpc:"microsoft_outlook_client_secret,omitempty"`
	ModuleAccount3WayMatch                      *Bool       `xmlrpc:"module_account_3way_match,omitempty"`
	ModuleAccountAccountant                     *Bool       `xmlrpc:"module_account_accountant,omitempty"`
	ModuleAccountBankStatementExtract           *Bool       `xmlrpc:"module_account_bank_statement_extract,omitempty"`
	ModuleAccountBankStatementImportQif         *Bool       `xmlrpc:"module_account_bank_statement_import_qif,omitempty"`
	ModuleAccountBatchPayment                   *Bool       `xmlrpc:"module_account_batch_payment,omitempty"`
	ModuleAccountBudget                         *Bool       `xmlrpc:"module_account_budget,omitempty"`
	ModuleAccountCheckPrinting                  *Bool       `xmlrpc:"module_account_check_printing,omitempty"`
	ModuleAccountExtract                        *Bool       `xmlrpc:"module_account_extract,omitempty"`
	ModuleAccountInterCompanyRules              *Bool       `xmlrpc:"module_account_inter_company_rules,omitempty"`
	ModuleAccountIntrastat                      *Bool       `xmlrpc:"module_account_intrastat,omitempty"`
	ModuleAccountInvoiceExtract                 *Bool       `xmlrpc:"module_account_invoice_extract,omitempty"`
	ModuleAccountIso20022                       *Bool       `xmlrpc:"module_account_iso20022,omitempty"`
	ModuleAccountPayment                        *Bool       `xmlrpc:"module_account_payment,omitempty"`
	ModuleAccountPeppol                         *Bool       `xmlrpc:"module_account_peppol,omitempty"`
	ModuleAccountReports                        *Bool       `xmlrpc:"module_account_reports,omitempty"`
	ModuleAccountSepaDirectDebit                *Bool       `xmlrpc:"module_account_sepa_direct_debit,omitempty"`
	ModuleAuthLdap                              *Bool       `xmlrpc:"module_auth_ldap,omitempty"`
	ModuleAuthOauth                             *Bool       `xmlrpc:"module_auth_oauth,omitempty"`
	ModuleBaseGeolocalize                       *Bool       `xmlrpc:"module_base_geolocalize,omitempty"`
	ModuleBaseImport                            *Bool       `xmlrpc:"module_base_import,omitempty"`
	ModuleCurrencyRateLive                      *Bool       `xmlrpc:"module_currency_rate_live,omitempty"`
	ModuleDelivery                              *Bool       `xmlrpc:"module_delivery,omitempty"`
	ModuleDeliveryBpost                         *Bool       `xmlrpc:"module_delivery_bpost,omitempty"`
	ModuleDeliveryDhl                           *Bool       `xmlrpc:"module_delivery_dhl,omitempty"`
	ModuleDeliveryEasypost                      *Bool       `xmlrpc:"module_delivery_easypost,omitempty"`
	ModuleDeliveryEnvia                         *Bool       `xmlrpc:"module_delivery_envia,omitempty"`
	ModuleDeliveryFedexRest                     *Bool       `xmlrpc:"module_delivery_fedex_rest,omitempty"`
	ModuleDeliverySendcloud                     *Bool       `xmlrpc:"module_delivery_sendcloud,omitempty"`
	ModuleDeliveryShiprocket                    *Bool       `xmlrpc:"module_delivery_shiprocket,omitempty"`
	ModuleDeliveryStarshipit                    *Bool       `xmlrpc:"module_delivery_starshipit,omitempty"`
	ModuleDeliveryUpsRest                       *Bool       `xmlrpc:"module_delivery_ups_rest,omitempty"`
	ModuleDeliveryUspsRest                      *Bool       `xmlrpc:"module_delivery_usps_rest,omitempty"`
	ModuleGoogleAddressAutocomplete             *Bool       `xmlrpc:"module_google_address_autocomplete,omitempty"`
	ModuleGoogleCalendar                        *Bool       `xmlrpc:"module_google_calendar,omitempty"`
	ModuleGoogleGmail                           *Bool       `xmlrpc:"module_google_gmail,omitempty"`
	ModuleGoogleRecaptcha                       *Bool       `xmlrpc:"module_google_recaptcha,omitempty"`
	ModuleLoyalty                               *Bool       `xmlrpc:"module_loyalty,omitempty"`
	ModuleMailPlugin                            *Bool       `xmlrpc:"module_mail_plugin,omitempty"`
	ModuleMicrosoftCalendar                     *Bool       `xmlrpc:"module_microsoft_calendar,omitempty"`
	ModuleMicrosoftOutlook                      *Bool       `xmlrpc:"module_microsoft_outlook,omitempty"`
	ModulePartnerAutocomplete                   *Bool       `xmlrpc:"module_partner_autocomplete,omitempty"`
	ModulePosAdyen                              *Bool       `xmlrpc:"module_pos_adyen,omitempty"`
	ModulePosMercadoPago                        *Bool       `xmlrpc:"module_pos_mercado_pago,omitempty"`
	ModulePosPineLabs                           *Bool       `xmlrpc:"module_pos_pine_labs,omitempty"`
	ModulePosPricer                             *Bool       `xmlrpc:"module_pos_pricer,omitempty"`
	ModulePosQfpay                              *Bool       `xmlrpc:"module_pos_qfpay,omitempty"`
	ModulePosRazorpay                           *Bool       `xmlrpc:"module_pos_razorpay,omitempty"`
	ModulePosStripe                             *Bool       `xmlrpc:"module_pos_stripe,omitempty"`
	ModulePosVivaCom                            *Bool       `xmlrpc:"module_pos_viva_com,omitempty"`
	ModuleProductEmailTemplate                  *Bool       `xmlrpc:"module_product_email_template,omitempty"`
	ModuleProductExpiry                         *Bool       `xmlrpc:"module_product_expiry,omitempty"`
	ModuleProductMargin                         *Bool       `xmlrpc:"module_product_margin,omitempty"`
	ModulePurchaseProductMatrix                 *Bool       `xmlrpc:"module_purchase_product_matrix,omitempty"`
	ModulePurchaseRequisition                   *Bool       `xmlrpc:"module_purchase_requisition,omitempty"`
	ModuleQualityControl                        *Bool       `xmlrpc:"module_quality_control,omitempty"`
	ModuleQualityControlWorksheet               *Bool       `xmlrpc:"module_quality_control_worksheet,omitempty"`
	ModuleSaleAmazon                            *Bool       `xmlrpc:"module_sale_amazon,omitempty"`
	ModuleSaleCommission                        *Bool       `xmlrpc:"module_sale_commission,omitempty"`
	ModuleSaleGelato                            *Bool       `xmlrpc:"module_sale_gelato,omitempty"`
	ModuleSaleLoyalty                           *Bool       `xmlrpc:"module_sale_loyalty,omitempty"`
	ModuleSaleMargin                            *Bool       `xmlrpc:"module_sale_margin,omitempty"`
	ModuleSalePdfQuoteBuilder                   *Bool       `xmlrpc:"module_sale_pdf_quote_builder,omitempty"`
	ModuleSaleProductMatrix                     *Bool       `xmlrpc:"module_sale_product_matrix,omitempty"`
	ModuleSaleShopee                            *Bool       `xmlrpc:"module_sale_shopee,omitempty"`
	ModuleSms                                   *Bool       `xmlrpc:"module_sms,omitempty"`
	ModuleSnailmailAccount                      *Bool       `xmlrpc:"module_snailmail_account,omitempty"`
	ModuleStockBarcode                          *Bool       `xmlrpc:"module_stock_barcode,omitempty"`
	ModuleStockBarcodeBarcodelookup             *Bool       `xmlrpc:"module_stock_barcode_barcodelookup,omitempty"`
	ModuleStockDropshipping                     *Bool       `xmlrpc:"module_stock_dropshipping,omitempty"`
	ModuleStockFleet                            *Bool       `xmlrpc:"module_stock_fleet,omitempty"`
	ModuleStockLandedCosts                      *Bool       `xmlrpc:"module_stock_landed_costs,omitempty"`
	ModuleStockPickingBatch                     *Bool       `xmlrpc:"module_stock_picking_batch,omitempty"`
	ModuleStockSms                              *Bool       `xmlrpc:"module_stock_sms,omitempty"`
	ModuleVoip                                  *Bool       `xmlrpc:"module_voip,omitempty"`
	ModuleWebUnsplash                           *Bool       `xmlrpc:"module_web_unsplash,omitempty"`
	ModuleWebsiteCfTurnstile                    *Bool       `xmlrpc:"module_website_cf_turnstile,omitempty"`
	OnboardingPaymentModule                     *Selection  `xmlrpc:"onboarding_payment_module,omitempty"`
	PartnerAutocompleteInsufficientCredit       *Bool       `xmlrpc:"partner_autocomplete_insufficient_credit,omitempty"`
	PayInvoicesOnline                           *Bool       `xmlrpc:"pay_invoices_online,omitempty"`
	PeppolExternalProvider                      *String     `xmlrpc:"peppol_external_provider,omitempty"`
	PeppolParentCompanyName                     *String     `xmlrpc:"peppol_parent_company_name,omitempty"`
	PeppolUseParentCompany                      *Bool       `xmlrpc:"peppol_use_parent_company,omitempty"`
	PoDoubleValidation                          *Selection  `xmlrpc:"po_double_validation,omitempty"`
	PoDoubleValidationAmount                    *Float      `xmlrpc:"po_double_validation_amount,omitempty"`
	PoLock                                      *Selection  `xmlrpc:"po_lock,omitempty"`
	PoOrderApproval                             *Bool       `xmlrpc:"po_order_approval,omitempty"`
	PointOfSaleTicketPortalUrlDisplayMode       *Selection  `xmlrpc:"point_of_sale_ticket_portal_url_display_mode,omitempty"`
	PointOfSaleTicketUniqueCode                 *Bool       `xmlrpc:"point_of_sale_ticket_unique_code,omitempty"`
	PointOfSaleUseTicketQrCode                  *Bool       `xmlrpc:"point_of_sale_use_ticket_qr_code,omitempty"`
	PortalAllowApiKeys                          *Bool       `xmlrpc:"portal_allow_api_keys,omitempty"`
	PortalConfirmationPay                       *Bool       `xmlrpc:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                      *Bool       `xmlrpc:"portal_confirmation_sign,omitempty"`
	PosAllowedPricelistIds                      *Relation   `xmlrpc:"pos_allowed_pricelist_ids,omitempty"`
	PosAmountAuthorizedDiff                     *Float      `xmlrpc:"pos_amount_authorized_diff,omitempty"`
	PosAutoValidateTerminalPayment              *Bool       `xmlrpc:"pos_auto_validate_terminal_payment,omitempty"`
	PosAvailablePresetIds                       *Relation   `xmlrpc:"pos_available_preset_ids,omitempty"`
	PosAvailablePricelistIds                    *Relation   `xmlrpc:"pos_available_pricelist_ids,omitempty"`
	PosBasicReceipt                             *Bool       `xmlrpc:"pos_basic_receipt,omitempty"`
	PosCashControl                              *Bool       `xmlrpc:"pos_cash_control,omitempty"`
	PosCashRounding                             *Bool       `xmlrpc:"pos_cash_rounding,omitempty"`
	PosCompanyHasTemplate                       *Bool       `xmlrpc:"pos_company_has_template,omitempty"`
	PosConfigId                                 *Many2One   `xmlrpc:"pos_config_id,omitempty"`
	PosCrmTeamId                                *Many2One   `xmlrpc:"pos_crm_team_id,omitempty"`
	PosCustomerDisplayBgImg                     *String     `xmlrpc:"pos_customer_display_bg_img,omitempty"`
	PosCustomerDisplayBgImgName                 *String     `xmlrpc:"pos_customer_display_bg_img_name,omitempty"`
	PosDefaultBillIds                           *Relation   `xmlrpc:"pos_default_bill_ids,omitempty"`
	PosDefaultFiscalPositionId                  *Many2One   `xmlrpc:"pos_default_fiscal_position_id,omitempty"`
	PosDefaultPresetId                          *Many2One   `xmlrpc:"pos_default_preset_id,omitempty"`
	PosDownPaymentProductId                     *Many2One   `xmlrpc:"pos_down_payment_product_id,omitempty"`
	PosEpsonPrinterIp                           *String     `xmlrpc:"pos_epson_printer_ip,omitempty"`
	PosFallbackNomenclatureId                   *Many2One   `xmlrpc:"pos_fallback_nomenclature_id,omitempty"`
	PosFastPaymentMethodIds                     *Relation   `xmlrpc:"pos_fast_payment_method_ids,omitempty"`
	PosFiscalPositionIds                        *Relation   `xmlrpc:"pos_fiscal_position_ids,omitempty"`
	PosHasActiveSession                         *Bool       `xmlrpc:"pos_has_active_session,omitempty"`
	PosIfaceAvailableCategIds                   *Relation   `xmlrpc:"pos_iface_available_categ_ids,omitempty"`
	PosIfaceBigScrollbars                       *Bool       `xmlrpc:"pos_iface_big_scrollbars,omitempty"`
	PosIfaceCashdrawer                          *Bool       `xmlrpc:"pos_iface_cashdrawer,omitempty"`
	PosIfaceElectronicScale                     *Bool       `xmlrpc:"pos_iface_electronic_scale,omitempty"`
	PosIfaceGroupByCateg                        *Bool       `xmlrpc:"pos_iface_group_by_categ,omitempty"`
	PosIfacePrintAuto                           *Bool       `xmlrpc:"pos_iface_print_auto,omitempty"`
	PosIfacePrintSkipScreen                     *Bool       `xmlrpc:"pos_iface_print_skip_screen,omitempty"`
	PosIfacePrintViaProxy                       *Bool       `xmlrpc:"pos_iface_print_via_proxy,omitempty"`
	PosIfaceScanViaProxy                        *Bool       `xmlrpc:"pos_iface_scan_via_proxy,omitempty"`
	PosIfaceTaxIncluded                         *Selection  `xmlrpc:"pos_iface_tax_included,omitempty"`
	PosIfaceTipproduct                          *Bool       `xmlrpc:"pos_iface_tipproduct,omitempty"`
	PosInvoiceJournalId                         *Many2One   `xmlrpc:"pos_invoice_journal_id,omitempty"`
	PosIsClosingEntryByProduct                  *Bool       `xmlrpc:"pos_is_closing_entry_by_product,omitempty"`
	PosIsHeaderOrFooter                         *Bool       `xmlrpc:"pos_is_header_or_footer,omitempty"`
	PosIsMarginsCostsAccessibleToEveryUser      *Bool       `xmlrpc:"pos_is_margins_costs_accessible_to_every_user,omitempty"`
	PosIsOrderPrinter                           *Bool       `xmlrpc:"pos_is_order_printer,omitempty"`
	PosIsPosbox                                 *Bool       `xmlrpc:"pos_is_posbox,omitempty"`
	PosJournalId                                *Many2One   `xmlrpc:"pos_journal_id,omitempty"`
	PosLimitCategories                          *Bool       `xmlrpc:"pos_limit_categories,omitempty"`
	PosManualDiscount                           *Bool       `xmlrpc:"pos_manual_discount,omitempty"`
	PosModulePosAppointment                     *Bool       `xmlrpc:"pos_module_pos_appointment,omitempty"`
	PosModulePosAvatax                          *Bool       `xmlrpc:"pos_module_pos_avatax,omitempty"`
	PosModulePosDiscount                        *Bool       `xmlrpc:"pos_module_pos_discount,omitempty"`
	PosModulePosHr                              *Bool       `xmlrpc:"pos_module_pos_hr,omitempty"`
	PosModulePosRestaurant                      *Bool       `xmlrpc:"pos_module_pos_restaurant,omitempty"`
	PosModulePosSms                             *Bool       `xmlrpc:"pos_module_pos_sms,omitempty"`
	PosNoteIds                                  *Relation   `xmlrpc:"pos_note_ids,omitempty"`
	PosOnlyRoundCashMethod                      *Bool       `xmlrpc:"pos_only_round_cash_method,omitempty"`
	PosOrderEditTracking                        *Bool       `xmlrpc:"pos_order_edit_tracking,omitempty"`
	PosOtherDevices                             *Bool       `xmlrpc:"pos_other_devices,omitempty"`
	PosPaymentMethodIds                         *Relation   `xmlrpc:"pos_payment_method_ids,omitempty"`
	PosPickingPolicy                            *Selection  `xmlrpc:"pos_picking_policy,omitempty"`
	PosPickingTypeId                            *Many2One   `xmlrpc:"pos_picking_type_id,omitempty"`
	PosPricelistId                              *Many2One   `xmlrpc:"pos_pricelist_id,omitempty"`
	PosPrinterIds                               *Relation   `xmlrpc:"pos_printer_ids,omitempty"`
	PosProxyIp                                  *String     `xmlrpc:"pos_proxy_ip,omitempty"`
	PosReceiptFooter                            *String     `xmlrpc:"pos_receipt_footer,omitempty"`
	PosReceiptHeader                            *String     `xmlrpc:"pos_receipt_header,omitempty"`
	PosRestrictPriceControl                     *Bool       `xmlrpc:"pos_restrict_price_control,omitempty"`
	PosRoundingMethod                           *Many2One   `xmlrpc:"pos_rounding_method,omitempty"`
	PosRouteId                                  *Many2One   `xmlrpc:"pos_route_id,omitempty"`
	PosSelectableCategIds                       *Relation   `xmlrpc:"pos_selectable_categ_ids,omitempty"`
	PosSetMaximumDifference                     *Bool       `xmlrpc:"pos_set_maximum_difference,omitempty"`
	PosShipLater                                *Bool       `xmlrpc:"pos_ship_later,omitempty"`
	PosShowCategoryImages                       *Bool       `xmlrpc:"pos_show_category_images,omitempty"`
	PosShowProductImages                        *Bool       `xmlrpc:"pos_show_product_images,omitempty"`
	PosTaxRegimeSelection                       *Bool       `xmlrpc:"pos_tax_regime_selection,omitempty"`
	PosTipProductId                             *Many2One   `xmlrpc:"pos_tip_product_id,omitempty"`
	PosTrustedConfigIds                         *Relation   `xmlrpc:"pos_trusted_config_ids,omitempty"`
	PosUseFastPayment                           *Bool       `xmlrpc:"pos_use_fast_payment,omitempty"`
	PosUsePresets                               *Bool       `xmlrpc:"pos_use_presets,omitempty"`
	PosUsePricelist                             *Bool       `xmlrpc:"pos_use_pricelist,omitempty"`
	PosWarehouseId                              *Many2One   `xmlrpc:"pos_warehouse_id,omitempty"`
	PrepaymentPercent                           *Float      `xmlrpc:"prepayment_percent,omitempty"`
	PreviewReady                                *Bool       `xmlrpc:"preview_ready,omitempty"`
	ProductVolumeVolumeInCubicFeet              *Selection  `xmlrpc:"product_volume_volume_in_cubic_feet,omitempty"`
	ProductWeightInLbs                          *Selection  `xmlrpc:"product_weight_in_lbs,omitempty"`
	ProfilingEnabledUntil                       *Time       `xmlrpc:"profiling_enabled_until,omitempty"`
	PurchaseTaxId                               *Many2One   `xmlrpc:"purchase_tax_id,omitempty"`
	QrCode                                      *Bool       `xmlrpc:"qr_code,omitempty"`
	QuickEditMode                               *Selection  `xmlrpc:"quick_edit_mode,omitempty"`
	QuotationValidityDays                       *Int        `xmlrpc:"quotation_validity_days,omitempty"`
	ReplenishOnOrder                            *Bool       `xmlrpc:"replenish_on_order,omitempty"`
	ReportFooter                                *String     `xmlrpc:"report_footer,omitempty"`
	RestrictTemplateRendering                   *Bool       `xmlrpc:"restrict_template_rendering,omitempty"`
	RestrictiveAuditTrail                       *Bool       `xmlrpc:"restrictive_audit_trail,omitempty"`
	SaleTaxId                                   *Many2One   `xmlrpc:"sale_tax_id,omitempty"`
	SecurityLead                                *Float      `xmlrpc:"security_lead,omitempty"`
	SfuServerKey                                *String     `xmlrpc:"sfu_server_key,omitempty"`
	SfuServerUrl                                *String     `xmlrpc:"sfu_server_url,omitempty"`
	ShowEffect                                  *Bool       `xmlrpc:"show_effect,omitempty"`
	ShowSaleReceipts                            *Bool       `xmlrpc:"show_sale_receipts,omitempty"`
	SnailmailColor                              *Bool       `xmlrpc:"snailmail_color,omitempty"`
	SnailmailCover                              *Bool       `xmlrpc:"snailmail_cover,omitempty"`
	SnailmailCoverReadonly                      *Bool       `xmlrpc:"snailmail_cover_readonly,omitempty"`
	SnailmailDuplex                             *Bool       `xmlrpc:"snailmail_duplex,omitempty"`
	StockConfirmationType                       *Selection  `xmlrpc:"stock_confirmation_type,omitempty"`
	StockMoveEmailValidation                    *Bool       `xmlrpc:"stock_move_email_validation,omitempty"`
	StockSmsConfirmationTemplateId              *Many2One   `xmlrpc:"stock_sms_confirmation_template_id,omitempty"`
	StockTextConfirmation                       *Bool       `xmlrpc:"stock_text_confirmation,omitempty"`
	TaxCalculationRoundingMethod                *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                       *Many2One   `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                              *Bool       `xmlrpc:"tax_exigibility,omitempty"`
	TenorApiKey                                 *String     `xmlrpc:"tenor_api_key,omitempty"`
	TermsType                                   *Selection  `xmlrpc:"terms_type,omitempty"`
	TransferAccountId                           *Many2One   `xmlrpc:"transfer_account_id,omitempty"`
	TwilioAccountSid                            *String     `xmlrpc:"twilio_account_sid,omitempty"`
	TwilioAccountToken                          *String     `xmlrpc:"twilio_account_token,omitempty"`
	UnsplashAccessKey                           *String     `xmlrpc:"unsplash_access_key,omitempty"`
	UnsplashAppId                               *String     `xmlrpc:"unsplash_app_id,omitempty"`
	UpdateStockQuantities                       *Selection  `xmlrpc:"update_stock_quantities,omitempty"`
	UseInvoiceTerms                             *Bool       `xmlrpc:"use_invoice_terms,omitempty"`
	UseSecurityLead                             *Bool       `xmlrpc:"use_security_lead,omitempty"`
	UseSfuServer                                *Bool       `xmlrpc:"use_sfu_server,omitempty"`
	UseTwilioRtcServers                         *Bool       `xmlrpc:"use_twilio_rtc_servers,omitempty"`
	VatCheckVies                                *Bool       `xmlrpc:"vat_check_vies,omitempty"`
	WebAppName                                  *String     `xmlrpc:"web_app_name,omitempty"`
	WriteDate                                   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One   `xmlrpc:"write_uid,omitempty"`
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
