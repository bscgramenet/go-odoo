package odoo

// L10NEsEdiVerifactuDocument represents l10n_es_edi_verifactu.document model.
type L10NEsEdiVerifactuDocument struct {
	ChainIndex             *Int       `xmlrpc:"chain_index,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	DocumentType           *Selection `xmlrpc:"document_type,omitempty"`
	Errors                 *String    `xmlrpc:"errors,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	JsonAttachmentBase64   *String    `xmlrpc:"json_attachment_base64,omitempty"`
	JsonAttachmentFilename *String    `xmlrpc:"json_attachment_filename,omitempty"`
	JsonAttachmentId       *Many2One  `xmlrpc:"json_attachment_id,omitempty"`
	MoveId                 *Many2One  `xmlrpc:"move_id,omitempty"`
	PosOrderId             *Many2One  `xmlrpc:"pos_order_id,omitempty"`
	ResponseCsv            *String    `xmlrpc:"response_csv,omitempty"`
	State                  *Selection `xmlrpc:"state,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// L10NEsEdiVerifactuDocuments represents array of l10n_es_edi_verifactu.document model.
type L10NEsEdiVerifactuDocuments []L10NEsEdiVerifactuDocument

// L10NEsEdiVerifactuDocumentModel is the odoo model name.
const L10NEsEdiVerifactuDocumentModel = "l10n_es_edi_verifactu.document"

// Many2One convert L10NEsEdiVerifactuDocument to *Many2One.
func (ld *L10NEsEdiVerifactuDocument) Many2One() *Many2One {
	return NewMany2One(ld.Id.Get(), "")
}

// CreateL10NEsEdiVerifactuDocument creates a new l10n_es_edi_verifactu.document model and returns its id.
func (c *Client) CreateL10NEsEdiVerifactuDocument(ld *L10NEsEdiVerifactuDocument) (int64, error) {
	ids, err := c.CreateL10NEsEdiVerifactuDocuments([]*L10NEsEdiVerifactuDocument{ld})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NEsEdiVerifactuDocument creates a new l10n_es_edi_verifactu.document model and returns its id.
func (c *Client) CreateL10NEsEdiVerifactuDocuments(lds []*L10NEsEdiVerifactuDocument) ([]int64, error) {
	var vv []interface{}
	for _, v := range lds {
		vv = append(vv, v)
	}
	return c.Create(L10NEsEdiVerifactuDocumentModel, vv, nil)
}

// UpdateL10NEsEdiVerifactuDocument updates an existing l10n_es_edi_verifactu.document record.
func (c *Client) UpdateL10NEsEdiVerifactuDocument(ld *L10NEsEdiVerifactuDocument) error {
	return c.UpdateL10NEsEdiVerifactuDocuments([]int64{ld.Id.Get()}, ld)
}

// UpdateL10NEsEdiVerifactuDocuments updates existing l10n_es_edi_verifactu.document records.
// All records (represented by ids) will be updated by ld values.
func (c *Client) UpdateL10NEsEdiVerifactuDocuments(ids []int64, ld *L10NEsEdiVerifactuDocument) error {
	return c.Update(L10NEsEdiVerifactuDocumentModel, ids, ld, nil)
}

// DeleteL10NEsEdiVerifactuDocument deletes an existing l10n_es_edi_verifactu.document record.
func (c *Client) DeleteL10NEsEdiVerifactuDocument(id int64) error {
	return c.DeleteL10NEsEdiVerifactuDocuments([]int64{id})
}

// DeleteL10NEsEdiVerifactuDocuments deletes existing l10n_es_edi_verifactu.document records.
func (c *Client) DeleteL10NEsEdiVerifactuDocuments(ids []int64) error {
	return c.Delete(L10NEsEdiVerifactuDocumentModel, ids)
}

// GetL10NEsEdiVerifactuDocument gets l10n_es_edi_verifactu.document existing record.
func (c *Client) GetL10NEsEdiVerifactuDocument(id int64) (*L10NEsEdiVerifactuDocument, error) {
	lds, err := c.GetL10NEsEdiVerifactuDocuments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lds)[0]), nil
}

// GetL10NEsEdiVerifactuDocuments gets l10n_es_edi_verifactu.document existing records.
func (c *Client) GetL10NEsEdiVerifactuDocuments(ids []int64) (*L10NEsEdiVerifactuDocuments, error) {
	lds := &L10NEsEdiVerifactuDocuments{}
	if err := c.Read(L10NEsEdiVerifactuDocumentModel, ids, nil, lds); err != nil {
		return nil, err
	}
	return lds, nil
}

// FindL10NEsEdiVerifactuDocument finds l10n_es_edi_verifactu.document record by querying it with criteria.
func (c *Client) FindL10NEsEdiVerifactuDocument(criteria *Criteria) (*L10NEsEdiVerifactuDocument, error) {
	lds := &L10NEsEdiVerifactuDocuments{}
	if err := c.SearchRead(L10NEsEdiVerifactuDocumentModel, criteria, NewOptions().Limit(1), lds); err != nil {
		return nil, err
	}
	return &((*lds)[0]), nil
}

// FindL10NEsEdiVerifactuDocuments finds l10n_es_edi_verifactu.document records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NEsEdiVerifactuDocuments(criteria *Criteria, options *Options) (*L10NEsEdiVerifactuDocuments, error) {
	lds := &L10NEsEdiVerifactuDocuments{}
	if err := c.SearchRead(L10NEsEdiVerifactuDocumentModel, criteria, options, lds); err != nil {
		return nil, err
	}
	return lds, nil
}

// FindL10NEsEdiVerifactuDocumentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NEsEdiVerifactuDocumentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NEsEdiVerifactuDocumentModel, criteria, options)
}

// FindL10NEsEdiVerifactuDocumentId finds record id by querying it with criteria.
func (c *Client) FindL10NEsEdiVerifactuDocumentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NEsEdiVerifactuDocumentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
