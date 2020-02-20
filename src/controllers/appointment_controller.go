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
	AppointmentController AppointmentControllerInterface = &appointmentController{}
)

type AppointmentControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	AddAttendees(w http.ResponseWriter, r *http.Request)
	RemoveAttendees(w http.ResponseWriter, r *http.Request)
}

type appointmentController struct{}

func (a *appointmentController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	calendarId := vars["calendar_id"]
	ok := IsValidUUID(calendarId)
	if !ok {
		logger.Logger.Infof("received invalid uuid=%s", calendarId)
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
	var appt models.Appointment
	err = json.Unmarshal(requestBody, &appt)
	if err != nil {
		errorMsg := "invalid json body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		RespondError(w, apiErr)
		return
	}

	appt.CalendarId = calendarId
	resultAppt, err := services.AppointmentService.Create(appt)
	if err != nil {
		errorMsg := "unable to crate appointment"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusConflict)
		RespondError(w, apiErr)
		return
	}
	response := models.ResponseCreated{
		Message:   "appointment created",
		CreatedId: resultAppt.ID,
	}
	RespondJSON(w, http.StatusCreated, response)
}

func (a *appointmentController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apptId := vars["appointment_id"]
	ok := IsValidUUID(apptId)
	if !ok {
		logger.Logger.Infof("received invalid uuid=%s", apptId)
		apiErr := NewBadRequestApiError("invalid uuid")
		RespondError(w, apiErr)
		return
	}

	resAppt, err := services.AppointmentService.Read(apptId)
	if err != nil {
		errorMsg := "unable to get appointment"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusNotFound)
		RespondError(w, apiErr)
		return
	}
	RespondJSON(w, http.StatusOK, resAppt)
}

func (a *appointmentController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apptId := vars["appointment_id"]
	ok := IsValidUUID(apptId)
	if !ok {
		logger.Logger.Infof("received invalid uuid=%s", apptId)
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
	var appt models.Appointment
	err = json.Unmarshal(requestBody, &appt)
	if err != nil {
		errorMsg := "invalid json body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		RespondError(w, apiErr)
		return
	}
	appt.ID = apptId
	resAppt, err := services.AppointmentService.Update(appt)
	if err != nil {
		errorMsg := "unable to update appointment"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusConflict)
		RespondError(w, apiErr)
		return
	}
	RespondJSON(w, http.StatusOK, resAppt)
}

func (a *appointmentController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apptId := vars["appointment_id"]
	ok := IsValidUUID(apptId)
	if !ok {
		logger.Logger.Infof("received invalid uuid=%s", apptId)
		apiErr := NewBadRequestApiError("invalid uuid")
		RespondError(w, apiErr)
		return
	}

	deletedId, err := services.AppointmentService.Delete(apptId)
	if err != nil {
		errorMsg := "unable to delete appointment"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusNotFound)
		RespondError(w, apiErr)
		return
	}
	response := models.ResponseDeleted{
		Message:   "appointment deleted",
		DeletedId: deletedId,
	}
	RespondJSON(w, http.StatusAccepted, response)
}

func (a *appointmentController) AddAttendees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apptId := vars["appointment_id"]
	ok := IsValidUUID(apptId)
	if !ok {
		logger.Logger.Infof("received invalid uuid=%s", apptId)
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

	var attendees []string
	err = json.Unmarshal(requestBody, &attendees)
	if err != nil {
		errorMsg := "invalid json body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		RespondError(w, apiErr)
		return
	}

	appt := models.Appointment{Base: models.Base{ID: apptId}}
	resultAppt, err := services.AppointmentService.AddAttendees(appt, attendees)
	if err != nil {
		errorMsg := "unable to add attendees to appointment"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusNotFound)
		RespondError(w, apiErr)
		return
	}
	RespondJSON(w, http.StatusOK, resultAppt)
}

func (a *appointmentController) RemoveAttendees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apptId := vars["appointment_id"]
	ok := IsValidUUID(apptId)
	if !ok {
		logger.Logger.Infof("received invalid uuid=%s", apptId)
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

	var attendees []string
	err = json.Unmarshal(requestBody, &attendees)
	if err != nil {
		errorMsg := "invalid json body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		RespondError(w, apiErr)
		return
	}

	appt := models.Appointment{Base: models.Base{ID: apptId}}
	resultAppt, err := services.AppointmentService.RemoveAttendees(appt, attendees)
	if err != nil {
		errorMsg := "unable to remove attendees from appointment"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusNotFound)
		RespondError(w, apiErr)
		return
	}
	RespondJSON(w, http.StatusOK, resultAppt)
}
