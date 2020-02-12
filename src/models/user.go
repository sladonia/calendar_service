package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"regexp"
	"strings"
)

type UserInterface interface {
	Validate() error
	Create(db *gorm.DB) error
	Delete(db *gorm.DB) error
	Update(db *gorm.DB) error
	Read(db *gorm.DB) error
}

type User struct {
	Base
	FirstName    string        `sql:"not null"`
	LastName     string        `sql:"not null"`
	Email        string        `sql:"unique_index; not null"`
	Appointments []Appointment `gorm:"many2many:users_appointments;"`
}

func (u *User) Validate() error {
	const emailValidatePattern = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	u.FirstName = strings.TrimSpace(u.FirstName)
	if u.FirstName == "" {
		return NewModeError("first name can not be empty")
	}
	u.LastName = strings.TrimSpace(u.LastName)
	if u.LastName == "" {
		return NewModeError("last name can not be empty")
	}
	u.Email = strings.TrimSpace(u.Email)
	if u.Email == "" {
		return NewModeError("email can not be empty")
	}
	re := regexp.MustCompile(emailValidatePattern)
	if !re.MatchString(u.Email) {
		return NewModeError(fmt.Sprintf("%s is not a valid email", u.Email))
	}
	return nil
}

func (u *User) Create(db *gorm.DB) error {
	if err := u.Validate(); err != nil {
		return err
	}
	err := db.Create(u).Error
	return err
}

func (u *User) Delete(db *gorm.DB) error {
	if u.EmptyID() {
		return EmptyIdError
	}
	dbState := db.Delete(u)
	if dbState.Error != nil {
		return nil
	}
	if dbState.RowsAffected == 0 {
		return NewModeError(fmt.Sprintf("user with id=%s not present in the db", u.ID))
	}
	return nil
}

func (u *User) Update(db *gorm.DB) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if u.EmptyID() {
		return EmptyIdError
	}
	dbState := db.Model(&User{}).Updates(u)
	if dbState.Error != nil {
		return dbState.Error
	}
	if dbState.RowsAffected == 0 {
		return NewModeError(fmt.Sprintf("user with id=%s not present in the db", u.ID))
	}
	return nil
}

func (u *User) Read(db *gorm.DB) error {
	if u.EmptyID() {
		return EmptyIdError
	}
	dbState := db.Find(u, "id = ?", u.ID)
	if dbState.Error != nil {
		return dbState.Error
	}
	if dbState.RowsAffected == 0 {
		return NewModeError(fmt.Sprintf("user with id=%s not present in the db", u.ID))
	}
	return nil
}
