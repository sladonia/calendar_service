package tests

import (
	"calendar_service/src/controllers"
	"calendar_service/src/datasources/postgres/calendardb"
	"calendar_service/src/models"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestCalendarCreate(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		res, err := client.Post(fmt.Sprintf("%s/user/%s/calendar", testServer.URL, models.KnownUserId),
			"application/json", strings.NewReader(`{"name": "First callendar"}`))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var responseCreated models.ResponseCreated
		err = json.Unmarshal(bodyBytes, &responseCreated)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 201, res.StatusCode)
		assert.Equal(t, "calendar created", responseCreated.Message)
	})

	t.Run("fail calendar exists", func(tt *testing.T) {
		res, err := client.Post(fmt.Sprintf("%s/user/%s/calendar", testServer.URL, models.KnownUserId),
			"application/json", strings.NewReader(`{"name": "First callendar"}`))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var apiErr controllers.ApiError
		err = json.Unmarshal(bodyBytes, &apiErr)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 409, res.StatusCode)
		assert.Equal(t, "unable to crate calendar", apiErr.Message)
	})

	t.Run("fail no such user", func(tt *testing.T) {
		res, err := client.Post(fmt.Sprintf("%s/user/%s/calendar", testServer.URL, models.UnexistingId),
			"application/json", strings.NewReader(`{"name": "First callendar"}`))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var apiErr controllers.ApiError
		err = json.Unmarshal(bodyBytes, &apiErr)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 409, res.StatusCode)
		assert.Equal(t, "unable to crate calendar", apiErr.Message)
	})
}

func TestCalendarRead(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		res, err := client.Get(fmt.Sprintf("%s/calendar/%s", testServer.URL, models.KnownCalendarId))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var cal models.Calendar
		err = json.Unmarshal(bodyBytes, &cal)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, models.KnownCalendarId, cal.ID)
	})

	t.Run("fail no such calendar", func(tt *testing.T) {
		res, err := client.Get(fmt.Sprintf("%s/calendar/%s", testServer.URL, models.UnexistingId))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var apiErr controllers.ApiError
		err = json.Unmarshal(bodyBytes, &apiErr)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 404, res.StatusCode)
		assert.Equal(t, "unable to get calendar", apiErr.Message)
	})
}

func TestCalendarUpdate(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		newCalendarName := "Third calendar"
		res, err := client.Post(fmt.Sprintf("%s/calendar/%s", testServer.URL, models.KnownCalendarId),
			"application/json",
			strings.NewReader(
				fmt.Sprintf("{\"user_id\": \"%s\", \"name\": \"%s\"}", models.KnownUserId, newCalendarName)))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var cal models.Calendar
		err = json.Unmarshal(bodyBytes, &cal)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, models.KnownCalendarId, cal.ID)
	})

	t.Run("fail unknown calendar_id", func(tt *testing.T) {
		newCalendarName := "Third calendar"
		res, err := client.Post(fmt.Sprintf("%s/calendar/%s", testServer.URL, models.UnexistingId),
			"application/json",
			strings.NewReader(
				fmt.Sprintf("{\"user_id\": \"%s\", \"name\": \"%s\"}", models.KnownUserId, newCalendarName)))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var apiErr controllers.ApiError
		err = json.Unmarshal(bodyBytes, &apiErr)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 404, res.StatusCode)
		assert.Equal(t, "unable to update calendar", apiErr.Message)
	})

	t.Run("fail unknown user_id", func(tt *testing.T) {
		newCalendarName := "Third calendar"
		res, err := client.Post(fmt.Sprintf("%s/calendar/%s", testServer.URL, models.KnownCalendarId),
			"application/json",
			strings.NewReader(
				fmt.Sprintf("{\"user_id\": \"%s\", \"name\": \"%s\"}", models.UnexistingId, newCalendarName)))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var apiErr controllers.ApiError
		err = json.Unmarshal(bodyBytes, &apiErr)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 404, res.StatusCode)
		assert.Equal(t, "unable to update calendar", apiErr.Message)
	})
}

func TestCalendarDelete(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		req, err := http.NewRequest(
			"DELETE",
			fmt.Sprintf("%s/calendar/%s", testServer.URL, models.KnownCalendarId),
			strings.NewReader(""))
		if err != nil {
			tt.Fatal("unable to create request", err)
		}

		res, err := client.Do(req)
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var resp models.ResponseDeleted
		err = json.Unmarshal(bodyBytes, &resp)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 202, res.StatusCode)
		assert.Equal(t, "calendar deleted", resp.Message)
	})

	t.Run("fail no such calendar", func(tt *testing.T) {
		req, err := http.NewRequest(
			"DELETE",
			fmt.Sprintf("%s/calendar/%s", testServer.URL, models.KnownCalendarId),
			strings.NewReader(""))
		if err != nil {
			tt.Fatal("unable to create request", err)
		}

		res, err := client.Do(req)
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var apiErr controllers.ApiError
		err = json.Unmarshal(bodyBytes, &apiErr)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 404, res.StatusCode)
		assert.Equal(t, "unable to delete calendar", apiErr.Message)
	})
}
