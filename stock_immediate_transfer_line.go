package odoo

// StockImmediateTransferLine represents stock.immediate.transfer.line model.
type StockImmediateTransferLine struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	ImmediateTransferId *Many2One `xmlrpc:"immediate_transfer_id,omitempty"`
	PickingId           *Many2One `xmlrpc:"picking_id,omitempty"`
	ToImmediate         *Bool     `xmlrpc:"to_immediate,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockImmediateTransferLines represents array of stock.immediate.transfer.line model.
type StockImmediateTransferLines []StockImmediateTransferLine

// StockImmediateTransferLineModel is the odoo model name.
const StockImmediateTransferLineModel = "stock.immediate.transfer.line"

// Many2One convert StockImmediateTransferLine to *Many2One.
func (sitl *StockImmediateTransferLine) Many2One() *Many2One {
	return NewMany2One(sitl.Id.Get(), "")
}

// CreateStockImmediateTransferLine creates a new stock.immediate.transfer.line model and returns its id.
func (c *Client) CreateStockImmediateTransferLine(sitl *StockImmediateTransferLine) (int64, error) {
	ids, err := c.CreateStockImmediateTransferLines([]*StockImmediateTransferLine{sitl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockImmediateTransferLine creates a new stock.immediate.transfer.line model and returns its id.
func (c *Client) CreateStockImmediateTransferLines(sitls []*StockImmediateTransferLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range sitls {
		vv = append(vv, v)
	}
	return c.Create(StockImmediateTransferLineModel, vv, nil)
}

// UpdateStockImmediateTransferLine updates an existing stock.immediate.transfer.line record.
func (c *Client) UpdateStockImmediateTransferLine(sitl *StockImmediateTransferLine) error {
	return c.UpdateStockImmediateTransferLines([]int64{sitl.Id.Get()}, sitl)
}

// UpdateStockImmediateTransferLines updates existing stock.immediate.transfer.line records.
// All records (represented by ids) will be updated by sitl values.
func (c *Client) UpdateStockImmediateTransferLines(ids []int64, sitl *StockImmediateTransferLine) error {
	return c.Update(StockImmediateTransferLineModel, ids, sitl, nil)
}

// DeleteStockImmediateTransferLine deletes an existing stock.immediate.transfer.line record.
func (c *Client) DeleteStockImmediateTransferLine(id int64) error {
	return c.DeleteStockImmediateTransferLines([]int64{id})
}

// DeleteStockImmediateTransferLines deletes existing stock.immediate.transfer.line records.
func (c *Client) DeleteStockImmediateTransferLines(ids []int64) error {
	return c.Delete(StockImmediateTransferLineModel, ids)
}

// GetStockImmediateTransferLine gets stock.immediate.transfer.line existing record.
func (c *Client) GetStockImmediateTransferLine(id int64) (*StockImmediateTransferLine, error) {
	sitls, err := c.GetStockImmediateTransferLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sitls)[0]), nil
}

// GetStockImmediateTransferLines gets stock.immediate.transfer.line existing records.
func (c *Client) GetStockImmediateTransferLines(ids []int64) (*StockImmediateTransferLines, error) {
	sitls := &StockImmediateTransferLines{}
	if err := c.Read(StockImmediateTransferLineModel, ids, nil, sitls); err != nil {
		return nil, err
	}
	return sitls, nil
}

// FindStockImmediateTransferLine finds stock.immediate.transfer.line record by querying it with criteria.
func (c *Client) FindStockImmediateTransferLine(criteria *Criteria) (*StockImmediateTransferLine, error) {
	sitls := &StockImmediateTransferLines{}
	if err := c.SearchRead(StockImmediateTransferLineModel, criteria, NewOptions().Limit(1), sitls); err != nil {
		return nil, err
	}
	return &((*sitls)[0]), nil
}

// FindStockImmediateTransferLines finds stock.immediate.transfer.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockImmediateTransferLines(criteria *Criteria, options *Options) (*StockImmediateTransferLines, error) {
	sitls := &StockImmediateTransferLines{}
	if err := c.SearchRead(StockImmediateTransferLineModel, criteria, options, sitls); err != nil {
		return nil, err
	}
	return sitls, nil
}

// FindStockImmediateTransferLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockImmediateTransferLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockImmediateTransferLineModel, criteria, options)
}

// FindStockImmediateTransferLineId finds record id by querying it with criteria.
func (c *Client) FindStockImmediateTransferLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockImmediateTransferLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
