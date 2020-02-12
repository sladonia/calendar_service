package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Base struct {
	//ID        string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ID        uuid.UUID `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Base) EmptyID() bool {
	if b.ID == [16]byte{} {
		return true
	}
	return false
}

func IdIsEmpty(id uuid.UUID) bool {
	if id == [16]byte{} {
		return true
	}
	return false
}
