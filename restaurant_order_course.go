package odoo

// RestaurantOrderCourse represents restaurant.order.course model.
type RestaurantOrderCourse struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Fired       *Bool     `xmlrpc:"fired,omitempty"`
	FiredDate   *Time     `xmlrpc:"fired_date,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Index       *Int      `xmlrpc:"index,omitempty"`
	LineIds     *Relation `xmlrpc:"line_ids,omitempty"`
	OrderId     *Many2One `xmlrpc:"order_id,omitempty"`
	Uuid        *String   `xmlrpc:"uuid,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// RestaurantOrderCourses represents array of restaurant.order.course model.
type RestaurantOrderCourses []RestaurantOrderCourse

// RestaurantOrderCourseModel is the odoo model name.
const RestaurantOrderCourseModel = "restaurant.order.course"

// Many2One convert RestaurantOrderCourse to *Many2One.
func (roc *RestaurantOrderCourse) Many2One() *Many2One {
	return NewMany2One(roc.Id.Get(), "")
}

// CreateRestaurantOrderCourse creates a new restaurant.order.course model and returns its id.
func (c *Client) CreateRestaurantOrderCourse(roc *RestaurantOrderCourse) (int64, error) {
	ids, err := c.CreateRestaurantOrderCourses([]*RestaurantOrderCourse{roc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRestaurantOrderCourse creates a new restaurant.order.course model and returns its id.
func (c *Client) CreateRestaurantOrderCourses(rocs []*RestaurantOrderCourse) ([]int64, error) {
	var vv []interface{}
	for _, v := range rocs {
		vv = append(vv, v)
	}
	return c.Create(RestaurantOrderCourseModel, vv, nil)
}

// UpdateRestaurantOrderCourse updates an existing restaurant.order.course record.
func (c *Client) UpdateRestaurantOrderCourse(roc *RestaurantOrderCourse) error {
	return c.UpdateRestaurantOrderCourses([]int64{roc.Id.Get()}, roc)
}

// UpdateRestaurantOrderCourses updates existing restaurant.order.course records.
// All records (represented by ids) will be updated by roc values.
func (c *Client) UpdateRestaurantOrderCourses(ids []int64, roc *RestaurantOrderCourse) error {
	return c.Update(RestaurantOrderCourseModel, ids, roc, nil)
}

// DeleteRestaurantOrderCourse deletes an existing restaurant.order.course record.
func (c *Client) DeleteRestaurantOrderCourse(id int64) error {
	return c.DeleteRestaurantOrderCourses([]int64{id})
}

// DeleteRestaurantOrderCourses deletes existing restaurant.order.course records.
func (c *Client) DeleteRestaurantOrderCourses(ids []int64) error {
	return c.Delete(RestaurantOrderCourseModel, ids)
}

// GetRestaurantOrderCourse gets restaurant.order.course existing record.
func (c *Client) GetRestaurantOrderCourse(id int64) (*RestaurantOrderCourse, error) {
	rocs, err := c.GetRestaurantOrderCourses([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rocs)[0]), nil
}

// GetRestaurantOrderCourses gets restaurant.order.course existing records.
func (c *Client) GetRestaurantOrderCourses(ids []int64) (*RestaurantOrderCourses, error) {
	rocs := &RestaurantOrderCourses{}
	if err := c.Read(RestaurantOrderCourseModel, ids, nil, rocs); err != nil {
		return nil, err
	}
	return rocs, nil
}

// FindRestaurantOrderCourse finds restaurant.order.course record by querying it with criteria.
func (c *Client) FindRestaurantOrderCourse(criteria *Criteria) (*RestaurantOrderCourse, error) {
	rocs := &RestaurantOrderCourses{}
	if err := c.SearchRead(RestaurantOrderCourseModel, criteria, NewOptions().Limit(1), rocs); err != nil {
		return nil, err
	}
	return &((*rocs)[0]), nil
}

// FindRestaurantOrderCourses finds restaurant.order.course records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRestaurantOrderCourses(criteria *Criteria, options *Options) (*RestaurantOrderCourses, error) {
	rocs := &RestaurantOrderCourses{}
	if err := c.SearchRead(RestaurantOrderCourseModel, criteria, options, rocs); err != nil {
		return nil, err
	}
	return rocs, nil
}

// FindRestaurantOrderCourseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRestaurantOrderCourseIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(RestaurantOrderCourseModel, criteria, options)
}

// FindRestaurantOrderCourseId finds record id by querying it with criteria.
func (c *Client) FindRestaurantOrderCourseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RestaurantOrderCourseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
