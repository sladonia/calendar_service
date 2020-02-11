package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"testing"
)

var (
	db *gorm.DB
)

func TestMain(m *testing.M) {
	var err error
	db, err = InitDbConnection("user", "password", "calendar_test", "disable")
	if err != nil {
		fmt.Println("unable to connect to db", err)
		os.Exit(1)
	}
	RecreateTables(db)
	os.Exit(m.Run())
}

//func TestUser_Validate(t *testing.T) {
//	validUser := User{
//		FirstName: "Jorge",
//		LastName:  "TheGreat",
//		Email:     "gorge@gmail.com",
//	}
//	err := validUser.Validate()
//	assert.Nil(t, err)
//}
