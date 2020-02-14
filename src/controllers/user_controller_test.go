package controllers

import (
	"bytes"
	"calendar_service/src/config"
	"calendar_service/src/datasources/postgres/calendardb"
	"calendar_service/src/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	if err := config.Load(); err != nil {
		fmt.Println("unable to load config", err)
		os.Exit(1)
	}
	//var err error
	db, err := models.InitDbConnection(
		config.Config.TestCalendarDb.User,
		config.Config.TestCalendarDb.Password,
		config.Config.TestCalendarDb.DbName,
		config.Config.TestCalendarDb.SslMode,
		config.Config.TestCalendarDb.MaxOpenConnections,
		config.Config.TestCalendarDb.MaxIdleConnections,
		config.Config.TestCalendarDb.ConnectionMaxLifetime)
	if err != nil {
		fmt.Println("unable to connect to db", err)
		os.Exit(1)
	}
	calendardb.DB = db
	models.RecreateTables(db)
	models.InitIndexes(db)
	os.Exit(m.Run())
}

func TestUserController_Create(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/user", strings.NewReader(
			`{"first_name": "James",
		"last_name": "Hamilgton",
		"email": "james@gmail.com"}`))
		UserController.Create(response, request)
		var result models.ResponseCreated
		err = json.Unmarshal(response.Body.Bytes(), &result)
		assert.Nil(tt, err)
		assert.Equal(t, 201, response.Code)
		assert.Equal(tt, "user created", result.Message)
	})

	t.Run("fail user exists", func(tt *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/user", strings.NewReader(
			`{"first_name": "James",
			"last_name": "Hamilgton",
			"email": "james@gmail.com"}`))
		UserController.Create(response, request)
		var result apiError
		err = json.Unmarshal(response.Body.Bytes(), &result)
		assert.Nil(tt, err)
		assert.Equal(t, 409, response.Code)
		assert.Equal(tt, "unable to crate user", result.Message)
	})

	t.Run("fail validation error", func(tt *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/user", strings.NewReader(
			`{"first_name": "",
			"last_name": "Hamilgton",
			"email": "james@gmail.com"}`))
		UserController.Create(response, request)
		var result apiError
		err = json.Unmarshal(response.Body.Bytes(), &result)
		assert.Nil(tt, err)
		assert.Equal(t, 409, response.Code)
		assert.Equal(tt, "unable to crate user", result.Message)
	})
}

func TestUserController_Read(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/user/user_id", bytes.NewReader([]byte{}))
		request = mux.SetURLVars(request, map[string]string{"id": models.KnownUserId})
		UserController.Read(response, request)
		var result models.User
		err = json.Unmarshal(response.Body.Bytes(), &result)
		assert.Nil(t, err)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, models.KnownUserId, result.ID)
	})

	t.Run("fail no such user", func(tt *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/user/user_id", bytes.NewReader([]byte{}))
		request = mux.SetURLVars(request, map[string]string{"id": models.UnexistingId})
		UserController.Read(response, request)
		var result apiError
		err = json.Unmarshal(response.Body.Bytes(), &result)
		assert.Nil(t, err)
		assert.Equal(t, 404, response.Code)
		assert.Equal(t, "unable to get user", result.Message)
	})
}

func TestUserController_Delete(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("DELETE", "/user/user_id", bytes.NewReader([]byte{}))
		request = mux.SetURLVars(request, map[string]string{"id": models.KnownUserId})
		UserController.Delete(response, request)
		var result models.ResponseDeleted
		err = json.Unmarshal(response.Body.Bytes(), &result)
		assert.Nil(t, err)
		assert.Equal(t, 202, response.Code)
		assert.Equal(t, models.KnownUserId, result.DeletedId)
	})

	t.Run("fail no such user", func(tt *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("DELETE", "/user/user_id", bytes.NewReader([]byte{}))
		request = mux.SetURLVars(request, map[string]string{"id": models.KnownUserId})
		UserController.Delete(response, request)
		var result apiError
		err = json.Unmarshal(response.Body.Bytes(), &result)
		assert.Nil(t, err)
		assert.Equal(t, 404, response.Code)
		assert.Equal(t, "unable to delete user", result.Message)
	})
}

func TestUserController_Update(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/user", strings.NewReader(
			`{"first_name": "James",
		"last_name": "Hamilgton",
		"email": "james@gmail.com"}`))
		request = mux.SetURLVars(request, map[string]string{"id": models.KnownUserId})
		UserController.Update(response, request)
		var result models.User
		err = json.Unmarshal(response.Body.Bytes(), &result)
		assert.Nil(tt, err)
		assert.Equal(tt, 200, response.Code)
		assert.Equal(tt, models.KnownUserId, result.ID)
		assert.Equal(tt, "Hamilgton", result.LastName)
	})

	t.Run("fail validation error", func(tt *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/user", strings.NewReader(
			`{"first_name": "James",
		"last_name": "Hamilgton",
		"email": "jamesil.com"}`))
		request = mux.SetURLVars(request, map[string]string{"id": models.KnownUserId})
		UserController.Update(response, request)
		var result apiError
		err = json.Unmarshal(response.Body.Bytes(), &result)
		assert.Nil(tt, err)
		assert.Equal(tt, http.StatusConflict, response.Code)
		assert.Equal(tt, "unable to update user", result.Message)
	})
}
