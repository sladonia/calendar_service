package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"strings"
)

type Calendar struct {
	Base
	Name   string    `gorm:"unique_index"`
	UserId uuid.UUID `gorm:"not null"`
}

func (c *Calendar) Validate() error {
	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		return NewModeError("calendar name can not be empty")
	}
	if IdIsEmpty(c.UserId) {
		return NewModeError("calendar user_id can not be empty")
	}
	return nil
}

func (c *Calendar) Create(db *gorm.DB) error {
	if err := c.Validate(); err != nil {
		return err
	}
	return db.Create(c).Error
}
