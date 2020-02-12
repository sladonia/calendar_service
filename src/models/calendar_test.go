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
	InitIndexes(db)
	os.Exit(m.Run())
}

func TestCalendar_Create(t *testing.T) {
	_ = MockDbData(db)
	//type fields struct {
	//	Base   Base
	//	Name   string
	//	UserId uuid.UUID
	//}
	//type args struct {
	//	db *gorm.DB
	//}
	//tests := []struct {
	//	name    string
	//	fields  fields
	//	args    args
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		c := &Calendar{
	//			Base:   tt.fields.Base,
	//			Name:   tt.fields.Name,
	//			UserId: tt.fields.UserId,
	//		}
	//		if err := c.Create(tt.args.db); (err != nil) != tt.wantErr {
	//			t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
	//		}
	//	})
	//}
}