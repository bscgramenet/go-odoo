package odoo

// RestaurantTable represents restaurant.table model.
type RestaurantTable struct {
	Active      *Bool      `xmlrpc:"active,omitempty"`
	Color       *String    `xmlrpc:"color,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	FloorId     *Many2One  `xmlrpc:"floor_id,omitempty"`
	Height      *Float     `xmlrpc:"height,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Identifier  *String    `xmlrpc:"identifier,omitempty"`
	ParentId    *Many2One  `xmlrpc:"parent_id,omitempty"`
	PositionH   *Float     `xmlrpc:"position_h,omitempty"`
	PositionV   *Float     `xmlrpc:"position_v,omitempty"`
	Seats       *Int       `xmlrpc:"seats,omitempty"`
	Shape       *Selection `xmlrpc:"shape,omitempty"`
	TableNumber *Int       `xmlrpc:"table_number,omitempty"`
	Width       *Float     `xmlrpc:"width,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// RestaurantTables represents array of restaurant.table model.
type RestaurantTables []RestaurantTable

// RestaurantTableModel is the odoo model name.
const RestaurantTableModel = "restaurant.table"

// Many2One convert RestaurantTable to *Many2One.
func (rt *RestaurantTable) Many2One() *Many2One {
	return NewMany2One(rt.Id.Get(), "")
}

// CreateRestaurantTable creates a new restaurant.table model and returns its id.
func (c *Client) CreateRestaurantTable(rt *RestaurantTable) (int64, error) {
	ids, err := c.CreateRestaurantTables([]*RestaurantTable{rt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRestaurantTable creates a new restaurant.table model and returns its id.
func (c *Client) CreateRestaurantTables(rts []*RestaurantTable) ([]int64, error) {
	var vv []interface{}
	for _, v := range rts {
		vv = append(vv, v)
	}
	return c.Create(RestaurantTableModel, vv, nil)
}

// UpdateRestaurantTable updates an existing restaurant.table record.
func (c *Client) UpdateRestaurantTable(rt *RestaurantTable) error {
	return c.UpdateRestaurantTables([]int64{rt.Id.Get()}, rt)
}

// UpdateRestaurantTables updates existing restaurant.table records.
// All records (represented by ids) will be updated by rt values.
func (c *Client) UpdateRestaurantTables(ids []int64, rt *RestaurantTable) error {
	return c.Update(RestaurantTableModel, ids, rt, nil)
}

// DeleteRestaurantTable deletes an existing restaurant.table record.
func (c *Client) DeleteRestaurantTable(id int64) error {
	return c.DeleteRestaurantTables([]int64{id})
}

// DeleteRestaurantTables deletes existing restaurant.table records.
func (c *Client) DeleteRestaurantTables(ids []int64) error {
	return c.Delete(RestaurantTableModel, ids)
}

// GetRestaurantTable gets restaurant.table existing record.
func (c *Client) GetRestaurantTable(id int64) (*RestaurantTable, error) {
	rts, err := c.GetRestaurantTables([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rts)[0]), nil
}

// GetRestaurantTables gets restaurant.table existing records.
func (c *Client) GetRestaurantTables(ids []int64) (*RestaurantTables, error) {
	rts := &RestaurantTables{}
	if err := c.Read(RestaurantTableModel, ids, nil, rts); err != nil {
		return nil, err
	}
	return rts, nil
}

// FindRestaurantTable finds restaurant.table record by querying it with criteria.
func (c *Client) FindRestaurantTable(criteria *Criteria) (*RestaurantTable, error) {
	rts := &RestaurantTables{}
	if err := c.SearchRead(RestaurantTableModel, criteria, NewOptions().Limit(1), rts); err != nil {
		return nil, err
	}
	return &((*rts)[0]), nil
}

// FindRestaurantTables finds restaurant.table records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRestaurantTables(criteria *Criteria, options *Options) (*RestaurantTables, error) {
	rts := &RestaurantTables{}
	if err := c.SearchRead(RestaurantTableModel, criteria, options, rts); err != nil {
		return nil, err
	}
	return rts, nil
}

// FindRestaurantTableIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRestaurantTableIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(RestaurantTableModel, criteria, options)
}

// FindRestaurantTableId finds record id by querying it with criteria.
func (c *Client) FindRestaurantTableId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RestaurantTableModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
