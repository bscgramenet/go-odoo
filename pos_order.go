package odoo

// PosOrder represents pos.order model.
type PosOrder struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken         *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl           *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning       *String    `xmlrpc:"access_warning,omitempty"`
	AccountMove         *Many2One  `xmlrpc:"account_move,omitempty"`
	AmountPaid          *Float     `xmlrpc:"amount_paid,omitempty"`
	AmountReturn        *Float     `xmlrpc:"amount_return,omitempty"`
	AmountTax           *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTotal         *Float     `xmlrpc:"amount_total,omitempty"`
	Cashier             *String    `xmlrpc:"cashier,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	ConfigId            *Many2One  `xmlrpc:"config_id,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmTeamId           *Many2One  `xmlrpc:"crm_team_id,omitempty"`
	CurrencyId          *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrencyRate        *Float     `xmlrpc:"currency_rate,omitempty"`
	DateOrder           *Time      `xmlrpc:"date_order,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId          *Many2One  `xmlrpc:"employee_id,omitempty"`
	FailedPickings      *Bool      `xmlrpc:"failed_pickings,omitempty"`
	FiscalPositionId    *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	HasRefundableLines  *Bool      `xmlrpc:"has_refundable_lines,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	IsInvoiced          *Bool      `xmlrpc:"is_invoiced,omitempty"`
	IsRefunded          *Bool      `xmlrpc:"is_refunded,omitempty"`
	IsTipped            *Bool      `xmlrpc:"is_tipped,omitempty"`
	IsTotalCostComputed *Bool      `xmlrpc:"is_total_cost_computed,omitempty"`
	Lines               *Relation  `xmlrpc:"lines,omitempty"`
	Margin              *Float     `xmlrpc:"margin,omitempty"`
	MarginPercent       *Float     `xmlrpc:"margin_percent,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	NbPrint             *Int       `xmlrpc:"nb_print,omitempty"`
	Note                *String    `xmlrpc:"note,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty"`
	PaymentIds          *Relation  `xmlrpc:"payment_ids,omitempty"`
	PickingCount        *Int       `xmlrpc:"picking_count,omitempty"`
	PickingIds          *Relation  `xmlrpc:"picking_ids,omitempty"`
	PickingTypeId       *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	PosReference        *String    `xmlrpc:"pos_reference,omitempty"`
	PricelistId         *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	ProcurementGroupId  *Many2One  `xmlrpc:"procurement_group_id,omitempty"`
	RefundOrdersCount   *Int       `xmlrpc:"refund_orders_count,omitempty"`
	RefundedOrderIds    *Relation  `xmlrpc:"refunded_order_ids,omitempty"`
	RefundedOrdersCount *Int       `xmlrpc:"refunded_orders_count,omitempty"`
	SaleJournal         *Many2One  `xmlrpc:"sale_journal,omitempty"`
	SaleOrderCount      *Int       `xmlrpc:"sale_order_count,omitempty"`
	SequenceNumber      *Int       `xmlrpc:"sequence_number,omitempty"`
	SessionId           *Many2One  `xmlrpc:"session_id,omitempty"`
	SessionMoveId       *Many2One  `xmlrpc:"session_move_id,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty"`
	TipAmount           *Float     `xmlrpc:"tip_amount,omitempty"`
	ToInvoice           *Bool      `xmlrpc:"to_invoice,omitempty"`
	ToShip              *Bool      `xmlrpc:"to_ship,omitempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PosOrders represents array of pos.order model.
type PosOrders []PosOrder

// PosOrderModel is the odoo model name.
const PosOrderModel = "pos.order"

// Many2One convert PosOrder to *Many2One.
func (po *PosOrder) Many2One() *Many2One {
	return NewMany2One(po.Id.Get(), "")
}

// CreatePosOrder creates a new pos.order model and returns its id.
func (c *Client) CreatePosOrder(po *PosOrder) (int64, error) {
	ids, err := c.CreatePosOrders([]*PosOrder{po})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosOrder creates a new pos.order model and returns its id.
func (c *Client) CreatePosOrders(pos []*PosOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range pos {
		vv = append(vv, v)
	}
	return c.Create(PosOrderModel, vv, nil)
}

// UpdatePosOrder updates an existing pos.order record.
func (c *Client) UpdatePosOrder(po *PosOrder) error {
	return c.UpdatePosOrders([]int64{po.Id.Get()}, po)
}

// UpdatePosOrders updates existing pos.order records.
// All records (represented by ids) will be updated by po values.
func (c *Client) UpdatePosOrders(ids []int64, po *PosOrder) error {
	return c.Update(PosOrderModel, ids, po, nil)
}

// DeletePosOrder deletes an existing pos.order record.
func (c *Client) DeletePosOrder(id int64) error {
	return c.DeletePosOrders([]int64{id})
}

// DeletePosOrders deletes existing pos.order records.
func (c *Client) DeletePosOrders(ids []int64) error {
	return c.Delete(PosOrderModel, ids)
}

// GetPosOrder gets pos.order existing record.
func (c *Client) GetPosOrder(id int64) (*PosOrder, error) {
	pos, err := c.GetPosOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pos)[0]), nil
}

// GetPosOrders gets pos.order existing records.
func (c *Client) GetPosOrders(ids []int64) (*PosOrders, error) {
	pos := &PosOrders{}
	if err := c.Read(PosOrderModel, ids, nil, pos); err != nil {
		return nil, err
	}
	return pos, nil
}

// FindPosOrder finds pos.order record by querying it with criteria.
func (c *Client) FindPosOrder(criteria *Criteria) (*PosOrder, error) {
	pos := &PosOrders{}
	if err := c.SearchRead(PosOrderModel, criteria, NewOptions().Limit(1), pos); err != nil {
		return nil, err
	}
	return &((*pos)[0]), nil
}

// FindPosOrders finds pos.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosOrders(criteria *Criteria, options *Options) (*PosOrders, error) {
	pos := &PosOrders{}
	if err := c.SearchRead(PosOrderModel, criteria, options, pos); err != nil {
		return nil, err
	}
	return pos, nil
}

// FindPosOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosOrderModel, criteria, options)
}

// FindPosOrderId finds record id by querying it with criteria.
func (c *Client) FindPosOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
