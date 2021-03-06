package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type Appointment struct {
	Base
	Subject     string    `gorm:"index;not null" json:"subject"`
	Description string    `json:"description"`
	WholeDay    bool      `json:"whole_day"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Attendees   []*User   `gorm:"many2many:users_appointments;" json:"attendees"`
	CalendarId  string    `gorm:"type:uuid;not null;" json:"calendar_id"`
}

func (a *Appointment) AfterFind() (err error) {
	if a.Attendees == nil {
		a.Attendees = []*User{}
	}
	return
}

func (a *Appointment) AfterUpdate() (err error) {
	if a.Attendees == nil {
		a.Attendees = []*User{}
	}
	return
}

func (a *Appointment) validateTime() error {
	defaultTime := time.Time{}
	if a.Start == defaultTime {
		return NewModeError("start time can not be empty")
	}
	if !a.WholeDay && a.End == defaultTime {
		return NewModeError("end time can not be empty")
	}
	if !a.WholeDay && a.End.Before(a.Start) {
		return NewModeError("appointment start time should be before end time")
	}
	if a.WholeDay && a.End != defaultTime {
		return NewModeError("both whole_day=true and end time provided")
	}
	return nil
}

func (a *Appointment) Validate() error {
	a.Subject = strings.TrimSpace(a.Subject)
	if a.Subject == "" {
		return NewModeError("appointment subject can not be empty")
	}
	if IdIsEmpty(a.CalendarId) {
		return NewModeError("appointment calendar_id can not be empty")
	}
	return a.validateTime()
}

func (a *Appointment) Create(db *gorm.DB) error {
	if err := a.Validate(); err != nil {
		return err
	}
	return db.Create(a).Error
}

func (a *Appointment) Delete(db *gorm.DB) error {
	if a.EmptyID() {
		return EmptyIdError
	}
	dbState := db.Delete(a)
	if dbState.Error != nil {
		return dbState.Error
	}
	if dbState.RowsAffected == 0 {
		return NewModeError(fmt.Sprintf("appointment with id=%s not present in the db", a.ID))
	}
	return nil
}

func (a *Appointment) Update(db *gorm.DB) error {
	if err := a.Validate(); err != nil {
		return err
	}
	if a.EmptyID() {
		return EmptyIdError
	}
	dbState := db.Model(&Appointment{}).Updates(a)
	if dbState.Error != nil {
		return dbState.Error
	}
	if dbState.RowsAffected == 0 {
		return NewModeError(fmt.Sprintf("appointment with id=%s not present in the db", a.ID))
	}
	if a.Attendees == nil {
		a.Attendees = []*User{}
	}
	return nil
}

func (a *Appointment) Read(db *gorm.DB) error {
	if a.EmptyID() {
		return EmptyIdError
	}
	dbState := db.Preload("Attendees").Find(a, "id = ?", a.ID)
	if dbState.Error != nil {
		return dbState.Error
	}
	if dbState.RowsAffected == 0 {
		return NewModeError(fmt.Sprintf("appointment with with id=%s not present in the db", a.ID))
	}
	return nil
}

func (a *Appointment) AddAttendees(userIds []string, db *gorm.DB) error {
	usrs := make([]*User, 0, 2)
	for _, userId := range userIds {
		_, err := uuid.Parse(userId)
		if err != nil {
			return err
		}
		usrs = append(usrs, &User{Base: Base{ID: userId}})
	}
	return db.Model(a).Association("Attendees").Append(usrs).Error
}

func (a *Appointment) RemoveAttendees(userIds []string, db *gorm.DB) error {
	usrs := make([]*User, 0, 2)
	for _, userId := range userIds {
		_, err := uuid.Parse(userId)
		if err != nil {
			return err
		}
		usrs = append(usrs, &User{Base: Base{ID: userId}})
	}
	res := db.Model(a).Association("Attendees").Delete(usrs)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
