package odoo

// PosPreset represents pos.preset model.
type PosPreset struct {
	AttendanceIds      *Relation  `xmlrpc:"attendance_ids,omitempty"`
	Color              *Int       `xmlrpc:"color,omitempty"`
	CountLinkedConfig  *Int       `xmlrpc:"count_linked_config,omitempty"`
	CountLinkedOrders  *Int       `xmlrpc:"count_linked_orders,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	FiscalPositionId   *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	HasImage           *Bool      `xmlrpc:"has_image,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	Identification     *Selection `xmlrpc:"identification,omitempty"`
	Image128           *String    `xmlrpc:"image_128,omitempty"`
	Image512           *String    `xmlrpc:"image_512,omitempty"`
	IntervalTime       *Int       `xmlrpc:"interval_time,omitempty"`
	IsReturn           *Bool      `xmlrpc:"is_return,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	PricelistId        *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	ResourceCalendarId *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	SlotsPerInterval   *Int       `xmlrpc:"slots_per_interval,omitempty"`
	UseTiming          *Bool      `xmlrpc:"use_timing,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PosPresets represents array of pos.preset model.
type PosPresets []PosPreset

// PosPresetModel is the odoo model name.
const PosPresetModel = "pos.preset"

// Many2One convert PosPreset to *Many2One.
func (pp *PosPreset) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreatePosPreset creates a new pos.preset model and returns its id.
func (c *Client) CreatePosPreset(pp *PosPreset) (int64, error) {
	ids, err := c.CreatePosPresets([]*PosPreset{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPreset creates a new pos.preset model and returns its id.
func (c *Client) CreatePosPresets(pps []*PosPreset) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(PosPresetModel, vv, nil)
}

// UpdatePosPreset updates an existing pos.preset record.
func (c *Client) UpdatePosPreset(pp *PosPreset) error {
	return c.UpdatePosPresets([]int64{pp.Id.Get()}, pp)
}

// UpdatePosPresets updates existing pos.preset records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdatePosPresets(ids []int64, pp *PosPreset) error {
	return c.Update(PosPresetModel, ids, pp, nil)
}

// DeletePosPreset deletes an existing pos.preset record.
func (c *Client) DeletePosPreset(id int64) error {
	return c.DeletePosPresets([]int64{id})
}

// DeletePosPresets deletes existing pos.preset records.
func (c *Client) DeletePosPresets(ids []int64) error {
	return c.Delete(PosPresetModel, ids)
}

// GetPosPreset gets pos.preset existing record.
func (c *Client) GetPosPreset(id int64) (*PosPreset, error) {
	pps, err := c.GetPosPresets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// GetPosPresets gets pos.preset existing records.
func (c *Client) GetPosPresets(ids []int64) (*PosPresets, error) {
	pps := &PosPresets{}
	if err := c.Read(PosPresetModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPosPreset finds pos.preset record by querying it with criteria.
func (c *Client) FindPosPreset(criteria *Criteria) (*PosPreset, error) {
	pps := &PosPresets{}
	if err := c.SearchRead(PosPresetModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// FindPosPresets finds pos.preset records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPresets(criteria *Criteria, options *Options) (*PosPresets, error) {
	pps := &PosPresets{}
	if err := c.SearchRead(PosPresetModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPosPresetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPresetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPresetModel, criteria, options)
}

// FindPosPresetId finds record id by querying it with criteria.
func (c *Client) FindPosPresetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPresetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
