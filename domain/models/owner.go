package models

type Owner struct {
	id             string
	name           string
	documentNumber string
}

func NewOwner(name, documentNumber string) *Owner {
	return &Owner{
		name:           name,
		documentNumber: documentNumber,
	}
}

func (o Owner) GetId() string {
	return o.id
}

func (o Owner) GetName() string {
	return o.name
}

func (o Owner) GetDocumentNumber() string {
	return o.documentNumber
}
