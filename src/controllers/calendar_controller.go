package controllers

import (
	"calendar_service/src/logger"
	"calendar_service/src/models"
	"calendar_service/src/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var (
	CalendarController CalendarControllerInterface = &calendarController{}
)

type CalendarControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type calendarController struct{}

func (c *calendarController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	ok := IsValidUUID(userId)
	if !ok {
		logger.Logger.Infof("received invalid uuid=%s", userId)
		apiErr := NewBadRequestApiError("invalid uuid")
		RespondError(w, apiErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMsg := "invalid request body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()
	var calendar models.Calendar
	err = json.Unmarshal(requestBody, &calendar)
	if err != nil {
		errorMsg := "invalid json body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		RespondError(w, apiErr)
		return
	}

	calendar.UserId = userId
	resultCalendar, err := services.CalendarService.Create(calendar)
	if err != nil {
		errorMsg := "unable to crate calendar"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusConflict)
		RespondError(w, apiErr)
		return
	}
	response := models.ResponseCreated{
		Message:   "calendar created",
		CreatedId: resultCalendar.ID,
	}
	RespondJSON(w, http.StatusCreated, response)
}

func (c *calendarController) Read(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c *calendarController) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c *calendarController) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
