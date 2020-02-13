package models

import (
	"time"
)

type Base struct {
	ID        string `sql:"type:uuid;primary_key;default:uuid_generate_v1()"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
