package services

import (
	"calendar_service/src/datasources/postgres/calendardb"
	"calendar_service/src/models"
)

var (
	AppointmentService AppointmentServiceInterface = &appointmentService{}
)

type AppointmentServiceInterface interface {
	Create(appt models.Appointment) (*models.Appointment, error)
	Read(apptId string) (*models.Appointment, error)
	Update(appt models.Appointment) (*models.Appointment, error)
	Delete(apptId string) (string, error)
}

type appointmentService struct{}

func (a *appointmentService) Create(appt models.Appointment) (*models.Appointment, error) {
	err := appt.Create(calendardb.DB)
	return &appt, err
}

func (a *appointmentService) Read(apptId string) (*models.Appointment, error) {
	appt := models.Appointment{Base: models.Base{ID: apptId}}
	err := appt.Read(calendardb.DB)
	return &appt, err
}

func (a *appointmentService) Update(appt models.Appointment) (*models.Appointment, error) {
	err := appt.Update(calendardb.DB)
	return &appt, err
}

func (a *appointmentService) Delete(apptId string) (string, error) {
	appt := models.Appointment{Base: models.Base{ID: apptId}}
	err := appt.Delete(calendardb.DB)
	return appt.ID, err
}
