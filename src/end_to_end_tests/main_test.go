package end_to_end_tests

import (
	"calendar_service/src/app"
	"calendar_service/src/datasources/postgres/calendardb"
	"calendar_service/src/models"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	testServer *httptest.Server
	client     *http.Client
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		fmt.Println("unable to load test env")
		os.Exit(1)
	}
	err = app.ConfigureApp()
	if err != nil {
		fmt.Println("unable to configure app")
		os.Exit(1)
	}
	router := app.InitApp()
	testServer = httptest.NewServer(router)
	client = testServer.Client()
	models.RecreateTables(calendardb.DB)
	models.InitIndexes(calendardb.DB)
	os.Exit(m.Run())
}
