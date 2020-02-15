package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type Calendar struct {
	Base
	Name         string        `gorm:"unique_index;not null" json:"name"`
	UserId       string        `gorm:"type:uuid;not null;" json:"-"`
	Appointments []Appointment `json:"appointments"`
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
	return db.Debug().Create(c).Error
}

func (c *Calendar) Delete(db *gorm.DB) error {
	if c.EmptyID() {
		return EmptyIdError
	}
	dbState := db.Delete(c)
	if dbState.Error != nil {
		return dbState.Error
	}
	if dbState.RowsAffected == 0 {
		return NewModeError(fmt.Sprintf("calendar with id=%s not present in the db", c.ID))
	}
	return nil
}

func (c *Calendar) Update(db *gorm.DB) error {
	if err := c.Validate(); err != nil {
		return err
	}
	if c.EmptyID() {
		return EmptyIdError
	}
	dbState := db.Model(&Calendar{}).Updates(c)
	if dbState.Error != nil {
		return dbState.Error
	}
	if dbState.RowsAffected == 0 {
		return NewModeError(fmt.Sprintf("calendar with id=%s not present in the db", c.ID))
	}
	return nil
}

func (c *Calendar) Read(db *gorm.DB) error {
	if c.EmptyID() {
		return EmptyIdError
	}
	dbState := db.Preload("Appointments").Find(c, "id = ?", c.ID)
	if dbState.Error != nil {
		return dbState.Error
	}
	if dbState.RowsAffected == 0 {
		return NewModeError(fmt.Sprintf("calendar with id=%s not present in the db", c.ID))
	}
	if c.Appointments == nil {
		c.Appointments = []Appointment{}
	}
	return nil
}
