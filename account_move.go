package odoo

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// AccountMove represents account.move model.
type AccountMove struct {
	LastUpdate                            *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                           *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                             *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning                         *String    `xmlrpc:"access_warning,omitempty"`
	ActivityCalendarEventId               *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                  *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration           *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                 *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                           *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                         *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                       *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                      *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                        *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                        *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AlwaysTaxExigible                     *Bool      `xmlrpc:"always_tax_exigible,omitempty"`
	AmountPaid                            *Float     `xmlrpc:"amount_paid,omitempty"`
	AmountResidual                        *Float     `xmlrpc:"amount_residual,omitempty"`
	AmountResidualSigned                  *Float     `xmlrpc:"amount_residual_signed,omitempty"`
	AmountTax                             *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTaxSigned                       *Float     `xmlrpc:"amount_tax_signed,omitempty"`
	AmountTotal                           *Float     `xmlrpc:"amount_total,omitempty"`
	AmountTotalInCurrencySigned           *Float     `xmlrpc:"amount_total_in_currency_signed,omitempty"`
	AmountTotalSigned                     *Float     `xmlrpc:"amount_total_signed,omitempty"`
	AmountUntaxed                         *Float     `xmlrpc:"amount_untaxed,omitempty"`
	AmountUntaxedSigned                   *Float     `xmlrpc:"amount_untaxed_signed,omitempty"`
	AttachmentIds                         *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorizedTransactionIds              *Relation  `xmlrpc:"authorized_transaction_ids,omitempty"`
	AutoPost                              *Selection `xmlrpc:"auto_post,omitempty"`
	AutoPostOriginId                      *Many2One  `xmlrpc:"auto_post_origin_id,omitempty"`
	AutoPostUntil                         *Time      `xmlrpc:"auto_post_until,omitempty"`
	BankPartnerId                         *Many2One  `xmlrpc:"bank_partner_id,omitempty"`
	CampaignId                            *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CommercialPartnerId                   *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyCurrencyId                     *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryCode                           *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                                  *Time      `xmlrpc:"date,omitempty"`
	DirectionSign                         *Int       `xmlrpc:"direction_sign,omitempty"`
	DisplayInactiveCurrencyWarning        *Bool      `xmlrpc:"display_inactive_currency_warning,omitempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omitempty"`
	DisplayQrCode                         *Bool      `xmlrpc:"display_qr_code,omitempty"`
	DuplicatedRefIds                      *Relation  `xmlrpc:"duplicated_ref_ids,omitempty"`
	EdiBlockingLevel                      *Selection `xmlrpc:"edi_blocking_level,omitempty"`
	EdiDocumentIds                        *Relation  `xmlrpc:"edi_document_ids,omitempty"`
	EdiErrorCount                         *Int       `xmlrpc:"edi_error_count,omitempty"`
	EdiErrorMessage                       *String    `xmlrpc:"edi_error_message,omitempty"`
	EdiShowAbandonCancelButton            *Bool      `xmlrpc:"edi_show_abandon_cancel_button,omitempty"`
	EdiShowCancelButton                   *Bool      `xmlrpc:"edi_show_cancel_button,omitempty"`
	EdiState                              *Selection `xmlrpc:"edi_state,omitempty"`
	EdiWebServicesToProcess               *String    `xmlrpc:"edi_web_services_to_process,omitempty"`
	FiscalPositionId                      *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	HasMessage                            *Bool      `xmlrpc:"has_message,omitempty"`
	HasReconciledEntries                  *Bool      `xmlrpc:"has_reconciled_entries,omitempty"`
	HidePostButton                        *Bool      `xmlrpc:"hide_post_button,omitempty"`
	HighestName                           *String    `xmlrpc:"highest_name,omitempty"`
	Id                                    *Int       `xmlrpc:"id,omitempty"`
	InalterableHash                       *String    `xmlrpc:"inalterable_hash,omitempty"`
	InvoiceCashRoundingId                 *Many2One  `xmlrpc:"invoice_cash_rounding_id,omitempty"`
	InvoiceDate                           *Time      `xmlrpc:"invoice_date,omitempty"`
	InvoiceDateDue                        *Time      `xmlrpc:"invoice_date_due,omitempty"`
	InvoiceFilterTypeDomain               *String    `xmlrpc:"invoice_filter_type_domain,omitempty"`
	InvoiceHasOutstanding                 *Bool      `xmlrpc:"invoice_has_outstanding,omitempty"`
	InvoiceIncotermId                     *Many2One  `xmlrpc:"invoice_incoterm_id,omitempty"`
	InvoiceLineIds                        *Relation  `xmlrpc:"invoice_line_ids,omitempty"`
	InvoiceOrigin                         *String    `xmlrpc:"invoice_origin,omitempty"`
	InvoiceOutstandingCreditsDebitsWidget *String    `xmlrpc:"invoice_outstanding_credits_debits_widget,omitempty"`
	InvoicePartnerDisplayName             *String    `xmlrpc:"invoice_partner_display_name,omitempty"`
	InvoicePaymentTermId                  *Many2One  `xmlrpc:"invoice_payment_term_id,omitempty"`
	InvoicePaymentsWidget                 *String    `xmlrpc:"invoice_payments_widget,omitempty"`
	InvoiceSourceEmail                    *String    `xmlrpc:"invoice_source_email,omitempty"`
	InvoiceUserId                         *Many2One  `xmlrpc:"invoice_user_id,omitempty"`
	InvoiceVendorBillId                   *Many2One  `xmlrpc:"invoice_vendor_bill_id,omitempty"`
	IsMoveSent                            *Bool      `xmlrpc:"is_move_sent,omitempty"`
	IsStorno                              *Bool      `xmlrpc:"is_storno,omitempty"`
	JournalId                             *Many2One  `xmlrpc:"journal_id,omitempty"`
	LineIds                               *Relation  `xmlrpc:"line_ids,omitempty"`
	MadeSequenceHole                      *Bool      `xmlrpc:"made_sequence_hole,omitempty"`
	MediumId                              *Many2One  `xmlrpc:"medium_id,omitempty"`
	MessageAttachmentCount                *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                    *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                       *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                    *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                            *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                     *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId               *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction                     *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter              *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                     *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MoveType                              *Selection `xmlrpc:"move_type,omitempty"`
	MyActivityDateDeadline                *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                  *String    `xmlrpc:"name,omitempty"`
	Narration                             *String    `xmlrpc:"narration,omitempty"`
	NeededTerms                           *String    `xmlrpc:"needed_terms,omitempty"`
	NeededTermsDirty                      *Bool      `xmlrpc:"needed_terms_dirty,omitempty"`
	PartnerBankId                         *Many2One  `xmlrpc:"partner_bank_id,omitempty"`
	PartnerCreditWarning                  *String    `xmlrpc:"partner_credit_warning,omitempty"`
	PartnerId                             *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerShippingId                     *Many2One  `xmlrpc:"partner_shipping_id,omitempty"`
	PaymentId                             *Many2One  `xmlrpc:"payment_id,omitempty"`
	PaymentIds                            *Relation  `xmlrpc:"payment_ids,omitempty"`
	PaymentReference                      *String    `xmlrpc:"payment_reference,omitempty"`
	PaymentState                          *Selection `xmlrpc:"payment_state,omitempty"`
	PaymentTermDetails                    *String    `xmlrpc:"payment_term_details,omitempty"`
	PosOrderIds                           *Relation  `xmlrpc:"pos_order_ids,omitempty"`
	PosPaymentIds                         *Relation  `xmlrpc:"pos_payment_ids,omitempty"`
	PostedBefore                          *Bool      `xmlrpc:"posted_before,omitempty"`
	QrCodeMethod                          *Selection `xmlrpc:"qr_code_method,omitempty"`
	QuickEditMode                         *Bool      `xmlrpc:"quick_edit_mode,omitempty"`
	QuickEditTotalAmount                  *Float     `xmlrpc:"quick_edit_total_amount,omitempty"`
	QuickEncodingVals                     *String    `xmlrpc:"quick_encoding_vals,omitempty"`
	Ref                                   *String    `xmlrpc:"ref,omitempty"`
	RestrictModeHashTable                 *Bool      `xmlrpc:"restrict_mode_hash_table,omitempty"`
	ReversalMoveId                        *Relation  `xmlrpc:"reversal_move_id,omitempty"`
	ReversedEntryId                       *Many2One  `xmlrpc:"reversed_entry_id,omitempty"`
	SaleOrderCount                        *Int       `xmlrpc:"sale_order_count,omitempty"`
	SecureSequenceNumber                  *Int       `xmlrpc:"secure_sequence_number,omitempty"`
	SequenceNumber                        *Int       `xmlrpc:"sequence_number,omitempty"`
	SequencePrefix                        *String    `xmlrpc:"sequence_prefix,omitempty"`
	ShowDiscountDetails                   *Bool      `xmlrpc:"show_discount_details,omitempty"`
	ShowNameWarning                       *Bool      `xmlrpc:"show_name_warning,omitempty"`
	ShowPaymentTermDetails                *Bool      `xmlrpc:"show_payment_term_details,omitempty"`
	ShowResetToDraftButton                *Bool      `xmlrpc:"show_reset_to_draft_button,omitempty"`
	SourceId                              *Many2One  `xmlrpc:"source_id,omitempty"`
	State                                 *Selection `xmlrpc:"state,omitempty"`
	StatementLineId                       *Many2One  `xmlrpc:"statement_line_id,omitempty"`
	StatementLineIds                      *Relation  `xmlrpc:"statement_line_ids,omitempty"`
	StockMoveId                           *Many2One  `xmlrpc:"stock_move_id,omitempty"`
	StockValuationLayerIds                *Relation  `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	StringToHash                          *String    `xmlrpc:"string_to_hash,omitempty"`
	SuitableJournalIds                    *Relation  `xmlrpc:"suitable_journal_ids,omitempty"`
	TaxCashBasisCreatedMoveIds            *Relation  `xmlrpc:"tax_cash_basis_created_move_ids,omitempty"`
	TaxCashBasisOriginMoveId              *Many2One  `xmlrpc:"tax_cash_basis_origin_move_id,omitempty"`
	TaxCashBasisRecId                     *Many2One  `xmlrpc:"tax_cash_basis_rec_id,omitempty"`
	TaxCountryCode                        *String    `xmlrpc:"tax_country_code,omitempty"`
	TaxCountryId                          *Many2One  `xmlrpc:"tax_country_id,omitempty"`
	TaxLockDateMessage                    *String    `xmlrpc:"tax_lock_date_message,omitempty"`
	TaxTotals                             *String    `xmlrpc:"tax_totals,omitempty"`
	TeamId                                *Many2One  `xmlrpc:"team_id,omitempty"`
	ToCheck                               *Bool      `xmlrpc:"to_check,omitempty"`
	TransactionIds                        *Relation  `xmlrpc:"transaction_ids,omitempty"`
	TypeName                              *String    `xmlrpc:"type_name,omitempty"`
	UserId                                *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds                     *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountMoves represents array of account.move model.
type AccountMoves []AccountMove

// AccountMoveModel is the odoo model name.
const AccountMoveModel = "account.move"

// Many2One convert AccountMove to *Many2One.
func (am *AccountMove) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAccountMove creates a new account.move model and returns its id.
func (c *Client) CreateAccountMove(am *AccountMove) (int64, error) {
	ids, err := c.CreateAccountMoves([]*AccountMove{am})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMove creates a new account.move model and returns its id.
func (c *Client) CreateAccountMoves(ams []*AccountMove) ([]int64, error) {
	var vv []interface{}
	for _, v := range ams {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveModel, vv, nil)
}

// UpdateAccountMove updates an existing account.move record.
func (c *Client) UpdateAccountMove(am *AccountMove) error {
	return c.UpdateAccountMoves([]int64{am.Id.Get()}, am)
}

// UpdateAccountMoves updates existing account.move records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAccountMoves(ids []int64, am *AccountMove) error {
	return c.Update(AccountMoveModel, ids, am, nil)
}

// DeleteAccountMove deletes an existing account.move record.
func (c *Client) DeleteAccountMove(id int64) error {
	return c.DeleteAccountMoves([]int64{id})
}

// DeleteAccountMoves deletes existing account.move records.
func (c *Client) DeleteAccountMoves(ids []int64) error {
	return c.Delete(AccountMoveModel, ids)
}

// GetAccountMove gets account.move existing record.
func (c *Client) GetAccountMove(id int64) (*AccountMove, error) {
	ams, err := c.GetAccountMoves([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// GetAccountMoves gets account.move existing records.
func (c *Client) GetAccountMoves(ids []int64) (*AccountMoves, error) {
	ams := &AccountMoves{}
	if err := c.Read(AccountMoveModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountMove finds account.move record by querying it with criteria.
func (c *Client) FindAccountMove(criteria *Criteria) (*AccountMove, error) {
	ams := &AccountMoves{}
	if err := c.SearchRead(AccountMoveModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// FindAccountMoves finds account.move records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoves(criteria *Criteria, options *Options) (*AccountMoves, error) {
	ams := &AccountMoves{}
	if err := c.SearchRead(AccountMoveModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountMoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveModel, criteria, options)
}

// FindAccountMoveId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}

// ReportDownloadInvoicePDF 通过 Odoo 16+ 登录态下载发票 PDF（自动登录获取 session cookie，自动获取 csrf_token）
func (c *Client) ReportDownloadInvoicePDF(moveId int64, username, password string) ([]byte, error) {
	if c.cfg == nil || c.cfg.URL == "" {
		return nil, errors.New("Odoo URL 未配置")
	}
	loginUrl := fmt.Sprintf("%s/web/login", c.cfg.URL)
	pdfUrl := fmt.Sprintf("%s/report/pdf/account.account_invoices/%d", c.cfg.URL, moveId)
	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}

	// 1. 先GET登录页，提取csrf_token
	resp0, err := client.Get(loginUrl)
	if err != nil {
		return nil, fmt.Errorf("Odoo 登录页获取失败: %v", err)
	}
	defer resp0.Body.Close()
	body0, err := io.ReadAll(resp0.Body)
	if err != nil {
		return nil, err
	}
	csrfToken := ""
	// 简单正则提取 <input type="hidden" name="csrf_token" value="..."/>
	re := regexp.MustCompile(`name=["']csrf_token["']\s+value=["']([^"']+)["']`)
	match := re.FindSubmatch(body0)
	if len(match) == 2 {
		csrfToken = string(match[1])
	}
	if csrfToken == "" {
		return nil, errors.New("Odoo 登录页未获取到 csrf_token")
	}

	// 2. 登录获取 session_id，带上首次GET登录页返回的所有cookies
	loginData := url.Values{}
	loginData.Set("login", username)
	loginData.Set("password", password)
	//loginData.Set("db", c.cfg.Database)
	loginData.Set("csrf_token", csrfToken)

	// 构造POST请求，带上cookie
	loginReq, err := http.NewRequest("POST", loginUrl, strings.NewReader(loginData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("Odoo 登录请求构造失败: %v", err)
	}
	loginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 收集首次GET登录页的所有cookie
	var cookies []string
	for _, c := range resp0.Cookies() {
		cookies = append(cookies, c.Name+"="+c.Value)
	}
	if len(cookies) > 0 {
		loginReq.Header.Set("Cookie", strings.Join(cookies, "; "))
	}
	loginResp, err := client.Do(loginReq)
	if err != nil {
		return nil, fmt.Errorf("Odoo 登录失败: %v", err)
	}
	defer loginResp.Body.Close()
	var sessionId string
	for _, cookie := range loginResp.Cookies() {
		if cookie.Name == "session_id" {
			sessionId = cookie.Value
			break
		}
	}
	if sessionId == "" {
		return nil, errors.New("Odoo 登录未获取到 session_id，账号或密码错误")
	}

	// 3. 用所有cookie访问 PDF（合并首次GET和登录后所有cookie，去重）
	cookieMap := make(map[string]string)
	for _, c := range resp0.Cookies() {
		cookieMap[c.Name] = c.Value
	}
	for _, c := range loginResp.Cookies() {
		cookieMap[c.Name] = c.Value
	}
	var pdfCookies []string
	for k, v := range cookieMap {
		pdfCookies = append(pdfCookies, k+"="+v)
	}

	// 3.1 检查 session 有效性（用POST+JSON）
	infoUrl := fmt.Sprintf("%s/web/session/get_session_info", c.cfg.URL)
	infoReq, err := http.NewRequest("POST", infoUrl, strings.NewReader("{}"))
	if err == nil {
		if len(pdfCookies) > 0 {
			infoReq.Header.Set("Cookie", strings.Join(pdfCookies, "; "))
		}
		infoReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		infoReq.Header.Set("Content-Type", "application/json")
		infoResp, err := client.Do(infoReq)
		if err == nil {
			defer infoResp.Body.Close()
			infoBody, _ := io.ReadAll(infoResp.Body)
			os.WriteFile("odoo_session_info.json", infoBody, 0644)
		}
	}

	// 4. 用所有cookie访问 PDF，支持 303 跳转自动跟进
	maxRedirect := 3
	curUrl := pdfUrl
	var lastResp *http.Response
	for i := 0; i < maxRedirect; i++ {
		req, err := http.NewRequest("GET", curUrl, nil)
		if err != nil {
			return nil, err
		}
		if len(pdfCookies) > 0 {
			req.Header.Set("Cookie", strings.Join(pdfCookies, "; "))
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		req.Header.Set("Referer", fmt.Sprintf("%s/web", c.cfg.URL))
		resp2, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp2.Body.Close()
		lastResp = resp2
		if resp2.StatusCode == 200 {
			contentType := resp2.Header.Get("Content-Type")
			body, err := io.ReadAll(resp2.Body)
			if err != nil {
				return nil, err
			}
			if contentType != "application/pdf" {
				errFile := fmt.Sprintf("odoo_invoice_%d_error.html", moveId)
				_ = os.WriteFile(errFile, body, 0644)
				preview := string(body)
				if len(preview) > 300 {
					preview = preview[:300] + "..."
				}
				return nil, fmt.Errorf("Odoo 返回非 PDF，Content-Type: %s，内容预览: %s，已保存为 %s", contentType, preview, errFile)
			}
			return body, nil
		}
		if resp2.StatusCode == 303 || resp2.StatusCode == 302 {
			loc := resp2.Header.Get("Location")
			headersFile := fmt.Sprintf("odoo_invoice_%d_303_headers.txt", moveId)
			bodyFile := fmt.Sprintf("odoo_invoice_%d_303.html", moveId)
			_ = os.WriteFile(headersFile, []byte(resp2.Status+"\n"+loc+"\n"+fmt.Sprint(resp2.Header)), 0644)
			b, _ := io.ReadAll(resp2.Body)
			_ = os.WriteFile(bodyFile, b, 0644)
			if strings.Contains(loc, "/web/login") {
				return nil, fmt.Errorf("Odoo 跳转到登录页，session 可能无效或权限不足，Location: %s", loc)
			}
			if strings.HasPrefix(loc, "/") && !strings.HasPrefix(loc, "http") {
				curUrl = c.cfg.URL + loc
			} else {
				curUrl = loc
			}
			continue
		}
		// 其它状态码
		b, _ := io.ReadAll(resp2.Body)
		errFile := fmt.Sprintf("odoo_invoice_%d_error.html", moveId)
		_ = os.WriteFile(errFile, b, 0644)
		return nil, fmt.Errorf("Odoo PDF 下载失败: %s，内容已保存为 %s", resp2.Status, errFile)
	}
	// 超过最大跳转次数
	if lastResp != nil {
		b, _ := io.ReadAll(lastResp.Body)
		errFile := fmt.Sprintf("odoo_invoice_%d_error.html", moveId)
		_ = os.WriteFile(errFile, b, 0644)
	}
	return nil, fmt.Errorf("Odoo PDF 下载失败，超过最大跳转次数")
}

// ReportDownloadPosTicketPDF 通过 Odoo 16+ 登录态下载 POS 小票 PDF（自动登录获取 session cookie，自动获取 csrf_token）
func (c *Client) ReportDownloadPosTicketPDF(posOrderId int64, username, password string) ([]byte, error) {
	if c.cfg == nil || c.cfg.URL == "" {
		return nil, errors.New("Odoo URL 未配置")
	}
	loginUrl := fmt.Sprintf("%s/web/login", c.cfg.URL)
	pdfUrl := fmt.Sprintf("%s/report/pdf/point_of_sale.pos_ticket/%d", c.cfg.URL, posOrderId)
	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}

	// 1. 先GET登录页，提取csrf_token
	resp0, err := client.Get(loginUrl)
	if err != nil {
		return nil, fmt.Errorf("Odoo 登录页获取失败: %v", err)
	}
	defer resp0.Body.Close()
	body0, err := io.ReadAll(resp0.Body)
	if err != nil {
		return nil, err
	}
	csrfToken := ""
	re := regexp.MustCompile(`name=["']csrf_token["']\s+value=["']([^"']+)["']`)
	match := re.FindSubmatch(body0)
	if len(match) == 2 {
		csrfToken = string(match[1])
	}
	if csrfToken == "" {
		return nil, errors.New("Odoo 登录页未获取到 csrf_token")
	}

	// 2. 登录获取 session_id，带上首次GET登录页返回的所有cookies
	loginData := url.Values{}
	loginData.Set("login", username)
	loginData.Set("password", password)
	loginData.Set("csrf_token", csrfToken)
	loginReq, err := http.NewRequest("POST", loginUrl, strings.NewReader(loginData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("Odoo 登录请求构造失败: %v", err)
	}
	loginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var cookies []string
	for _, c := range resp0.Cookies() {
		cookies = append(cookies, c.Name+"="+c.Value)
	}
	if len(cookies) > 0 {
		loginReq.Header.Set("Cookie", strings.Join(cookies, "; "))
	}
	loginResp, err := client.Do(loginReq)
	if err != nil {
		return nil, fmt.Errorf("Odoo 登录失败: %v", err)
	}
	defer loginResp.Body.Close()
	var sessionId string
	for _, cookie := range loginResp.Cookies() {
		if cookie.Name == "session_id" {
			sessionId = cookie.Value
			break
		}
	}
	if sessionId == "" {
		return nil, errors.New("Odoo 登录未获取到 session_id，账号或密码错误")
	}

	// 3. 合并所有cookie
	cookieMap := make(map[string]string)
	for _, c := range resp0.Cookies() {
		cookieMap[c.Name] = c.Value
	}
	for _, c := range loginResp.Cookies() {
		cookieMap[c.Name] = c.Value
	}
	var pdfCookies []string
	for k, v := range cookieMap {
		pdfCookies = append(pdfCookies, k+"="+v)
	}

	// 4. 检查 session 有效性
	infoUrl := fmt.Sprintf("%s/web/session/get_session_info", c.cfg.URL)
	infoReq, err := http.NewRequest("POST", infoUrl, strings.NewReader("{}"))
	if err == nil {
		if len(pdfCookies) > 0 {
			infoReq.Header.Set("Cookie", strings.Join(pdfCookies, "; "))
		}
		infoReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		infoReq.Header.Set("Content-Type", "application/json")
		infoResp, err := client.Do(infoReq)
		if err == nil {
			defer infoResp.Body.Close()
			infoBody, _ := io.ReadAll(infoResp.Body)
			os.WriteFile("odoo_session_info.json", infoBody, 0644)
		}
	}

	// 5. 用所有cookie访问 PDF，支持 303 跳转自动跟进
	maxRedirect := 3
	curUrl := pdfUrl
	var lastResp *http.Response
	for i := 0; i < maxRedirect; i++ {
		req, err := http.NewRequest("GET", curUrl, nil)
		if err != nil {
			return nil, err
		}
		if len(pdfCookies) > 0 {
			req.Header.Set("Cookie", strings.Join(pdfCookies, "; "))
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		req.Header.Set("Referer", fmt.Sprintf("%s/web", c.cfg.URL))
		resp2, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp2.Body.Close()
		lastResp = resp2
		if resp2.StatusCode == 200 {
			contentType := resp2.Header.Get("Content-Type")
			body, err := io.ReadAll(resp2.Body)
			if err != nil {
				return nil, err
			}
			if contentType != "application/pdf" {
				errFile := fmt.Sprintf("odoo_pos_ticket_%d_error.html", posOrderId)
				_ = os.WriteFile(errFile, body, 0644)
				preview := string(body)
				if len(preview) > 300 {
					preview = preview[:300] + "..."
				}
				return nil, fmt.Errorf("Odoo 返回非 PDF，Content-Type: %s，内容预览: %s，已保存为 %s", contentType, preview, errFile)
			}
			return body, nil
		}
		if resp2.StatusCode == 303 || resp2.StatusCode == 302 {
			loc := resp2.Header.Get("Location")
			headersFile := fmt.Sprintf("odoo_pos_ticket_%d_303_headers.txt", posOrderId)
			bodyFile := fmt.Sprintf("odoo_pos_ticket_%d_303.html", posOrderId)
			_ = os.WriteFile(headersFile, []byte(resp2.Status+"\n"+loc+"\n"+fmt.Sprint(resp2.Header)), 0644)
			b, _ := io.ReadAll(resp2.Body)
			_ = os.WriteFile(bodyFile, b, 0644)
			if strings.Contains(loc, "/web/login") {
				return nil, fmt.Errorf("Odoo 跳转到登录页，session 可能无效或权限不足，Location: %s", loc)
			}
			if strings.HasPrefix(loc, "/") && !strings.HasPrefix(loc, "http") {
				curUrl = c.cfg.URL + loc
			} else {
				curUrl = loc
			}
			continue
		}
		// 其它状态码
		b, _ := io.ReadAll(resp2.Body)
		errFile := fmt.Sprintf("odoo_pos_ticket_%d_error.html", posOrderId)
		_ = os.WriteFile(errFile, b, 0644)
		return nil, fmt.Errorf("Odoo POS 小票 PDF 下载失败: %s，内容已保存为 %s", resp2.Status, errFile)
	}
	// 超过最大跳转次数
	if lastResp != nil {
		b, _ := io.ReadAll(lastResp.Body)
		errFile := fmt.Sprintf("odoo_pos_ticket_%d_error.html", posOrderId)
		_ = os.WriteFile(errFile, b, 0644)
	}
	return nil, fmt.Errorf("Odoo POS 小票 PDF 下载失败，超过最大跳转次数")
}

// 移除本文件多余的 call 方法定义，避免与 odoo.go 冲突
