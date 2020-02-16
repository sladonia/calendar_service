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

func TestAppointmentCreate(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		res, err := client.Post(
			fmt.Sprintf("%s/calendar/%s/appointment", testServer.URL, models.KnownCalendarId),
			"application/json", strings.NewReader(`
				{"subject": "first_appt",
				"whole_day": true,
				"start": "2018-09-22T12:42:31Z"}`))
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
		assert.Equal(t, "appointment created", responseCreated.Message)
	})

	t.Run("fail already exists", func(tt *testing.T) {
		res, err := client.Post(
			fmt.Sprintf("%s/calendar/%s/appointment", testServer.URL, models.KnownCalendarId),
			"application/json", strings.NewReader(`
				{"subject": "first_appt",
				"whole_day": true,
				"start": "2018-09-22T12:42:31Z"}`))
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
		assert.Equal(t, "unable to crate appointment", apiErr.Message)
	})

	t.Run("fail calendar does not exist", func(tt *testing.T) {
		res, err := client.Post(
			fmt.Sprintf("%s/calendar/%s/appointment", testServer.URL, models.UnexistingId),
			"application/json", strings.NewReader(`
				{"subject": "first_appt",
				"whole_day": true,
				"start": "2018-09-22T12:42:31Z"}`))
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
		assert.Equal(t, "unable to crate appointment", apiErr.Message)
	})

	t.Run("fail validation error", func(tt *testing.T) {
		res, err := client.Post(
			fmt.Sprintf("%s/calendar/%s/appointment", testServer.URL, models.KnownCalendarId),
			"application/json", strings.NewReader(`
				{"subject": "first_appt",
				"start": "2018-09-22T12:42:31Z"}`))
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
		assert.Equal(t, "unable to crate appointment", apiErr.Message)
	})
}

func TestAppointmentRead(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		res, err := client.Get(fmt.Sprintf("%s/appointment/%s", testServer.URL, models.AppointmentWholeDayId))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var resAppt models.Appointment
		err = json.Unmarshal(bodyBytes, &resAppt)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, models.AppointmentWholeDayId, resAppt.ID)
	})

	t.Run("fail no such appointment", func(tt *testing.T) {
		res, err := client.Get(fmt.Sprintf("%s/appointment/%s", testServer.URL, models.UnexistingId))
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
		assert.Equal(t, "unable to get appointment", apiErr.Message)
	})
}

func TestAppointmentUpdate(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		requestBody := fmt.Sprintf(
			`{"calendar_id": "%s",
			"subject": "first_appt",
			"start": "2018-09-22T12:42:31Z",
			"end": "2018-10-22T12:42:31Z",
			"whole_day": false}`, models.KnownCalendarId)
		res, err := client.Post(
			fmt.Sprintf("%s/appointment/%s", testServer.URL, models.AppointmentWholeDayId),
			"application/json", strings.NewReader(requestBody))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var resAppt models.Appointment
		err = json.Unmarshal(bodyBytes, &resAppt)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, models.AppointmentWholeDayId, resAppt.ID)
	})

	t.Run("fail no such calendar", func(tt *testing.T) {
		requestBody := fmt.Sprintf(
			`{"calendar_id": "%s",
			"subject": "first_appt",
			"start": "2018-09-22T12:42:31Z",
			"end": "2018-10-22T12:42:31Z",
			"whole_day": false}`, models.UnexistingId)
		res, err := client.Post(
			fmt.Sprintf("%s/appointment/%s", testServer.URL, models.AppointmentWholeDayId),
			"application/json", strings.NewReader(requestBody))
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
		assert.Equal(t, "unable to update appointment", apiErr.Message)
	})

	t.Run("fail no such appointment", func(tt *testing.T) {
		requestBody := fmt.Sprintf(
			`{"calendar_id": "%s",
			"subject": "first_appt",
			"start": "2018-09-22T12:42:31Z",
			"end": "2018-10-22T12:42:31Z",
			"whole_day": false}`, models.KnownCalendarId)
		res, err := client.Post(
			fmt.Sprintf("%s/appointment/%s", testServer.URL, models.UnexistingId),
			"application/json", strings.NewReader(requestBody))
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
		assert.Equal(t, "unable to update appointment", apiErr.Message)
	})
}

func TestAppointmentDelete(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		request, err := http.NewRequest(
			"DELETE",
			fmt.Sprintf("%s/appointment/%s", testServer.URL, models.AppointmentWholeDayId),
			strings.NewReader(""))
		if err != nil {
			tt.Fatal("unable to create request", err)
		}
		res, err := client.Do(request)
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}
		var respDeleted models.ResponseDeleted
		err = json.Unmarshal(bodyBytes, &respDeleted)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 202, res.StatusCode)
		assert.Equal(t, "appointment deleted", respDeleted.Message)
		assert.Equal(t, models.AppointmentWholeDayId, respDeleted.DeletedId)
	})

	t.Run("fail no such appointment", func(tt *testing.T) {
		request, err := http.NewRequest(
			"DELETE",
			fmt.Sprintf("%s/appointment/%s", testServer.URL, models.AppointmentWholeDayId),
			strings.NewReader(""))
		if err != nil {
			tt.Fatal("unable to create request", err)
		}
		res, err := client.Do(request)
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
		assert.Equal(t, "unable to delete appointment", apiErr.Message)
	})
}
