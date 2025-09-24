package odoo

// PosLoadMixin represents pos.load.mixin model.
type PosLoadMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// PosLoadMixins represents array of pos.load.mixin model.
type PosLoadMixins []PosLoadMixin

// PosLoadMixinModel is the odoo model name.
const PosLoadMixinModel = "pos.load.mixin"

// Many2One convert PosLoadMixin to *Many2One.
func (plm *PosLoadMixin) Many2One() *Many2One {
	return NewMany2One(plm.Id.Get(), "")
}

// CreatePosLoadMixin creates a new pos.load.mixin model and returns its id.
func (c *Client) CreatePosLoadMixin(plm *PosLoadMixin) (int64, error) {
	ids, err := c.CreatePosLoadMixins([]*PosLoadMixin{plm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosLoadMixin creates a new pos.load.mixin model and returns its id.
func (c *Client) CreatePosLoadMixins(plms []*PosLoadMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range plms {
		vv = append(vv, v)
	}
	return c.Create(PosLoadMixinModel, vv, nil)
}

// UpdatePosLoadMixin updates an existing pos.load.mixin record.
func (c *Client) UpdatePosLoadMixin(plm *PosLoadMixin) error {
	return c.UpdatePosLoadMixins([]int64{plm.Id.Get()}, plm)
}

// UpdatePosLoadMixins updates existing pos.load.mixin records.
// All records (represented by ids) will be updated by plm values.
func (c *Client) UpdatePosLoadMixins(ids []int64, plm *PosLoadMixin) error {
	return c.Update(PosLoadMixinModel, ids, plm, nil)
}

// DeletePosLoadMixin deletes an existing pos.load.mixin record.
func (c *Client) DeletePosLoadMixin(id int64) error {
	return c.DeletePosLoadMixins([]int64{id})
}

// DeletePosLoadMixins deletes existing pos.load.mixin records.
func (c *Client) DeletePosLoadMixins(ids []int64) error {
	return c.Delete(PosLoadMixinModel, ids)
}

// GetPosLoadMixin gets pos.load.mixin existing record.
func (c *Client) GetPosLoadMixin(id int64) (*PosLoadMixin, error) {
	plms, err := c.GetPosLoadMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*plms)[0]), nil
}

// GetPosLoadMixins gets pos.load.mixin existing records.
func (c *Client) GetPosLoadMixins(ids []int64) (*PosLoadMixins, error) {
	plms := &PosLoadMixins{}
	if err := c.Read(PosLoadMixinModel, ids, nil, plms); err != nil {
		return nil, err
	}
	return plms, nil
}

// FindPosLoadMixin finds pos.load.mixin record by querying it with criteria.
func (c *Client) FindPosLoadMixin(criteria *Criteria) (*PosLoadMixin, error) {
	plms := &PosLoadMixins{}
	if err := c.SearchRead(PosLoadMixinModel, criteria, NewOptions().Limit(1), plms); err != nil {
		return nil, err
	}
	return &((*plms)[0]), nil
}

// FindPosLoadMixins finds pos.load.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosLoadMixins(criteria *Criteria, options *Options) (*PosLoadMixins, error) {
	plms := &PosLoadMixins{}
	if err := c.SearchRead(PosLoadMixinModel, criteria, options, plms); err != nil {
		return nil, err
	}
	return plms, nil
}

// FindPosLoadMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosLoadMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosLoadMixinModel, criteria, options)
}

// FindPosLoadMixinId finds record id by querying it with criteria.
func (c *Client) FindPosLoadMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosLoadMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
