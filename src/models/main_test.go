package models

import (
	"calendar_service/src/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

var (
	db *gorm.DB
)

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load("../../.env.test")
	if err != nil {
		fmt.Println("unable to load test env")
		os.Exit(1)
	}
	if err := config.Load(); err != nil {
		fmt.Println("unable to load config", err)
		os.Exit(1)
	}
	db, err = InitDbConnection(
		config.Config.CalendarDb.User,
		config.Config.CalendarDb.Password,
		config.Config.CalendarDb.DbName,
		config.Config.CalendarDb.SslMode,
		config.Config.CalendarDb.MaxOpenConnections,
		config.Config.CalendarDb.MaxIdleConnections,
		config.Config.CalendarDb.ConnectionMaxLifetime)
	if err != nil {
		fmt.Println("unable to connect to db", err)
		os.Exit(1)
	}
	RecreateTables(db)
	InitIndexes(db)
	os.Exit(m.Run())
}
