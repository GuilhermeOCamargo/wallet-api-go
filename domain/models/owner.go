package models

import "time"

type Owner struct {
	id             uint
	name           string
	documentNumber string
	createdAt      time.Time
	updatedAt      time.Time
}

func NewOwner(name, documentNumber string) *Owner {
	return &Owner{
		name:           name,
		documentNumber: documentNumber,
	}
}

func (o Owner) Id() uint {
	return o.id
}

func (o Owner) Name() string {
	return o.name
}

func (o Owner) DocumentNumber() string {
	return o.documentNumber
}

func (o *Owner) SetId(id uint) {
	o.id = id
}

func (o *Owner) SetCreatedAt(createdAt time.Time) {
	o.createdAt = createdAt
}

func (o *Owner) SetUpdatedAt(updatedAt time.Time) {
	o.updatedAt = updatedAt
}
