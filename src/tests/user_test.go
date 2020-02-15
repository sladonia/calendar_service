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

func TestUserController_Read(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		res, err := client.Get(fmt.Sprintf("%s/user/%s", testServer.URL, models.KnownUserId))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}

		var usr models.User
		err = json.Unmarshal(bodyBytes, &usr)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, models.KnownUserId, usr.ID)
	})

	t.Run("fail no such user", func(tt *testing.T) {
		res, err := client.Get(fmt.Sprintf("%s/user/%s", testServer.URL, models.UnexistingId))
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
		assert.Equal(t, "unable to get user", apiErr.GetMessage())
	})
}

func TestUserController_Create(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		res, err := client.Post(fmt.Sprintf("%s/user", testServer.URL), "application/json", strings.NewReader(
			`{"first_name": "James",
			"last_name": "Hamilgton",
			"email": "james@gmail.com"}`))

		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}

		var usr models.ResponseCreated
		err = json.Unmarshal(bodyBytes, &usr)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 201, res.StatusCode)
		assert.Equal(t, "user created", usr.Message)
	})

	t.Run("fail user exists", func(tt *testing.T) {
		res, err := client.Post(fmt.Sprintf("%s/user", testServer.URL), "application/json", strings.NewReader(
			`{"first_name": "James",
			"last_name": "Hamilgton",
			"email": "james@gmail.com"}`))

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
		assert.Equal(t, http.StatusConflict, res.StatusCode)
		assert.Equal(t, "unable to crate user", apiErr.Message)
	})

	t.Run("fail empty user name", func(tt *testing.T) {
		res, err := client.Post(fmt.Sprintf("%s/user", testServer.URL), "application/json", strings.NewReader(
			`{"first_name": "",
			"last_name": "Hamilgt",
			"email": "james@gma.com"}`))

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
		assert.Equal(t, http.StatusConflict, res.StatusCode)
		assert.Equal(t, "unable to crate user", apiErr.Message)
	})
}

func TestUserController_Delete(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/user/%s", testServer.URL, models.KnownUserId), strings.NewReader(""))
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

		var resp models.ResponseDeleted
		err = json.Unmarshal(bodyBytes, &resp)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 202, res.StatusCode)
		assert.Equal(t, models.KnownUserId, resp.DeletedId)
	})

	t.Run("fail no such user", func(tt *testing.T) {
		request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/user/%s", testServer.URL, models.UnexistingId), strings.NewReader(""))
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

		var resp controllers.ApiError
		err = json.Unmarshal(bodyBytes, &resp)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 404, res.StatusCode)
		assert.Equal(t, "unable to delete user", resp.Message)
	})
}

func TestUserController_Update(t *testing.T) {
	err := models.MockDbData(calendardb.DB)
	if err != nil {
		t.Fatal("unable to mock db")
	}
	defer models.DropAllData(calendardb.DB)

	t.Run("success", func(tt *testing.T) {
		res, err := client.Post(fmt.Sprintf("%s/user/%s", testServer.URL, models.KnownUserId), "application/json", strings.NewReader(
			`{"first_name": "Rotor",
			"last_name": "Hamilgton",
			"email": "james@gmail.com"}`))
		if err != nil {
			tt.Fatal("unable to execute request", err)
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			tt.Fatal("unable to read response body", err)
		}

		var usr models.User
		err = json.Unmarshal(bodyBytes, &usr)
		if err != nil {
			tt.Fatal("unable to unmarshal response", err)
		}
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, models.KnownUserId, usr.ID)
		assert.Equal(t, "Rotor", usr.FirstName)
	})

	t.Run("fail validation error", func(tt *testing.T) {
		res, err := client.Post(fmt.Sprintf("%s/user/%s", testServer.URL, models.KnownUserId), "application/json", strings.NewReader(
			`{"first_name": "Rotor",
			"last_name": "Hamilgton",
			"email": "jamesgmail.com"}`))
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
		assert.Equal(t, "unable to update user", apiErr.Message)
	})
}
