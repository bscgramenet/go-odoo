package odoo

// RestaurantFloor represents restaurant.floor model.
type RestaurantFloor struct {
	Active               *Bool     `xmlrpc:"active,omitempty"`
	BackgroundColor      *String   `xmlrpc:"background_color,omitempty"`
	BackgroundImage      *String   `xmlrpc:"background_image,omitempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	FloorBackgroundImage *String   `xmlrpc:"floor_background_image,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	Name                 *String   `xmlrpc:"name,omitempty"`
	PosConfigIds         *Relation `xmlrpc:"pos_config_ids,omitempty"`
	Sequence             *Int      `xmlrpc:"sequence,omitempty"`
	TableIds             *Relation `xmlrpc:"table_ids,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// RestaurantFloors represents array of restaurant.floor model.
type RestaurantFloors []RestaurantFloor

// RestaurantFloorModel is the odoo model name.
const RestaurantFloorModel = "restaurant.floor"

// Many2One convert RestaurantFloor to *Many2One.
func (rf *RestaurantFloor) Many2One() *Many2One {
	return NewMany2One(rf.Id.Get(), "")
}

// CreateRestaurantFloor creates a new restaurant.floor model and returns its id.
func (c *Client) CreateRestaurantFloor(rf *RestaurantFloor) (int64, error) {
	ids, err := c.CreateRestaurantFloors([]*RestaurantFloor{rf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRestaurantFloor creates a new restaurant.floor model and returns its id.
func (c *Client) CreateRestaurantFloors(rfs []*RestaurantFloor) ([]int64, error) {
	var vv []interface{}
	for _, v := range rfs {
		vv = append(vv, v)
	}
	return c.Create(RestaurantFloorModel, vv, nil)
}

// UpdateRestaurantFloor updates an existing restaurant.floor record.
func (c *Client) UpdateRestaurantFloor(rf *RestaurantFloor) error {
	return c.UpdateRestaurantFloors([]int64{rf.Id.Get()}, rf)
}

// UpdateRestaurantFloors updates existing restaurant.floor records.
// All records (represented by ids) will be updated by rf values.
func (c *Client) UpdateRestaurantFloors(ids []int64, rf *RestaurantFloor) error {
	return c.Update(RestaurantFloorModel, ids, rf, nil)
}

// DeleteRestaurantFloor deletes an existing restaurant.floor record.
func (c *Client) DeleteRestaurantFloor(id int64) error {
	return c.DeleteRestaurantFloors([]int64{id})
}

// DeleteRestaurantFloors deletes existing restaurant.floor records.
func (c *Client) DeleteRestaurantFloors(ids []int64) error {
	return c.Delete(RestaurantFloorModel, ids)
}

// GetRestaurantFloor gets restaurant.floor existing record.
func (c *Client) GetRestaurantFloor(id int64) (*RestaurantFloor, error) {
	rfs, err := c.GetRestaurantFloors([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rfs)[0]), nil
}

// GetRestaurantFloors gets restaurant.floor existing records.
func (c *Client) GetRestaurantFloors(ids []int64) (*RestaurantFloors, error) {
	rfs := &RestaurantFloors{}
	if err := c.Read(RestaurantFloorModel, ids, nil, rfs); err != nil {
		return nil, err
	}
	return rfs, nil
}

// FindRestaurantFloor finds restaurant.floor record by querying it with criteria.
func (c *Client) FindRestaurantFloor(criteria *Criteria) (*RestaurantFloor, error) {
	rfs := &RestaurantFloors{}
	if err := c.SearchRead(RestaurantFloorModel, criteria, NewOptions().Limit(1), rfs); err != nil {
		return nil, err
	}
	return &((*rfs)[0]), nil
}

// FindRestaurantFloors finds restaurant.floor records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRestaurantFloors(criteria *Criteria, options *Options) (*RestaurantFloors, error) {
	rfs := &RestaurantFloors{}
	if err := c.SearchRead(RestaurantFloorModel, criteria, options, rfs); err != nil {
		return nil, err
	}
	return rfs, nil
}

// FindRestaurantFloorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRestaurantFloorIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(RestaurantFloorModel, criteria, options)
}

// FindRestaurantFloorId finds record id by querying it with criteria.
func (c *Client) FindRestaurantFloorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RestaurantFloorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
