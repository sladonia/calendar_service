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
	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})
}
