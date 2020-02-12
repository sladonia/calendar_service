package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var connectString = "user=user password=password dbname=calendar_development sslmode=disable"

func InitDbConnection(user, password, dbname, sslmode string) (*gorm.DB, error) {
	connectString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		user, password, dbname, sslmode)
	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func RecreateTables(db *gorm.DB) {
	db.DropTableIfExists("users_appointments")
	db.DropTableIfExists(&Appointment{})
	db.DropTableIfExists(&Calendar{})
	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})
	db.CreateTable(&Calendar{})
	db.CreateTable(&Appointment{})
}

func InitIndexes(db *gorm.DB) {
	db.Model(&User{}).AddUniqueIndex("idx_user_first_last_name_unique", "first_name", "last_name")
	db.Model(&Calendar{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Appointment{}).AddForeignKey("calendar_id", "calendars(id)", "CASCADE", "CASCADE")
	db.Model(&Appointment{}).AddUniqueIndex("idx_calendar_id_subject_unique", "calendar_id", "subject")
}

func DropAllData(db *gorm.DB) {
	db.Where("true").Delete("users_appointments")
	db.Where("true").Delete(&Appointment{})
	db.Where("true").Delete(&Calendar{})
	db.Where("true").Delete(&User{})
}
