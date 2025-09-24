package odoo

// PosBusMixin represents pos.bus.mixin model.
type PosBusMixin struct {
	AccessToken *String `xmlrpc:"access_token,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// PosBusMixins represents array of pos.bus.mixin model.
type PosBusMixins []PosBusMixin

// PosBusMixinModel is the odoo model name.
const PosBusMixinModel = "pos.bus.mixin"

// Many2One convert PosBusMixin to *Many2One.
func (pbm *PosBusMixin) Many2One() *Many2One {
	return NewMany2One(pbm.Id.Get(), "")
}

// CreatePosBusMixin creates a new pos.bus.mixin model and returns its id.
func (c *Client) CreatePosBusMixin(pbm *PosBusMixin) (int64, error) {
	ids, err := c.CreatePosBusMixins([]*PosBusMixin{pbm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosBusMixin creates a new pos.bus.mixin model and returns its id.
func (c *Client) CreatePosBusMixins(pbms []*PosBusMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range pbms {
		vv = append(vv, v)
	}
	return c.Create(PosBusMixinModel, vv, nil)
}

// UpdatePosBusMixin updates an existing pos.bus.mixin record.
func (c *Client) UpdatePosBusMixin(pbm *PosBusMixin) error {
	return c.UpdatePosBusMixins([]int64{pbm.Id.Get()}, pbm)
}

// UpdatePosBusMixins updates existing pos.bus.mixin records.
// All records (represented by ids) will be updated by pbm values.
func (c *Client) UpdatePosBusMixins(ids []int64, pbm *PosBusMixin) error {
	return c.Update(PosBusMixinModel, ids, pbm, nil)
}

// DeletePosBusMixin deletes an existing pos.bus.mixin record.
func (c *Client) DeletePosBusMixin(id int64) error {
	return c.DeletePosBusMixins([]int64{id})
}

// DeletePosBusMixins deletes existing pos.bus.mixin records.
func (c *Client) DeletePosBusMixins(ids []int64) error {
	return c.Delete(PosBusMixinModel, ids)
}

// GetPosBusMixin gets pos.bus.mixin existing record.
func (c *Client) GetPosBusMixin(id int64) (*PosBusMixin, error) {
	pbms, err := c.GetPosBusMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pbms)[0]), nil
}

// GetPosBusMixins gets pos.bus.mixin existing records.
func (c *Client) GetPosBusMixins(ids []int64) (*PosBusMixins, error) {
	pbms := &PosBusMixins{}
	if err := c.Read(PosBusMixinModel, ids, nil, pbms); err != nil {
		return nil, err
	}
	return pbms, nil
}

// FindPosBusMixin finds pos.bus.mixin record by querying it with criteria.
func (c *Client) FindPosBusMixin(criteria *Criteria) (*PosBusMixin, error) {
	pbms := &PosBusMixins{}
	if err := c.SearchRead(PosBusMixinModel, criteria, NewOptions().Limit(1), pbms); err != nil {
		return nil, err
	}
	return &((*pbms)[0]), nil
}

// FindPosBusMixins finds pos.bus.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosBusMixins(criteria *Criteria, options *Options) (*PosBusMixins, error) {
	pbms := &PosBusMixins{}
	if err := c.SearchRead(PosBusMixinModel, criteria, options, pbms); err != nil {
		return nil, err
	}
	return pbms, nil
}

// FindPosBusMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosBusMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosBusMixinModel, criteria, options)
}

// FindPosBusMixinId finds record id by querying it with criteria.
func (c *Client) FindPosBusMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosBusMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
