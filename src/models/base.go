package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v1()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
