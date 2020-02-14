package models

import (
	"time"
)

type Base struct {
	ID        string    `sql:"type:uuid;primary_key;default:uuid_generate_v1()" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (b *Base) EmptyID() bool {
	if b.ID == "" {
		return true
	}
	return false
}

func IdIsEmpty(id string) bool {
	if id == "" {
		return true
	}
	return false
}
