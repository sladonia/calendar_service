package services

import (
	"calendar_service/src/datasources/postgres/calendardb"
	"calendar_service/src/models"
)

var (
	CalendarService CalendarServiceInterface = &calendarService{}
)

type CalendarServiceInterface interface {
	Create(cal models.Calendar) (*models.Calendar, error)
	Read(calendarId string) (*models.Calendar, error)
	Update(cal models.Calendar) (*models.Calendar, error)
	Delete(calendarId string) (string, error)
}

type calendarService struct{}

func (c *calendarService) Create(cal models.Calendar) (*models.Calendar, error) {
	err := cal.Validate()
	if err != nil {
		return nil, err
	}
	err = cal.Create(calendardb.DB)
	if err != nil {
		return nil, err
	}
	return &cal, nil
}

func (c *calendarService) Read(calendarId string) (*models.Calendar, error) {
	cal := models.Calendar{Base: models.Base{ID: calendarId}}
	err := cal.Read(calendardb.DB)
	return &cal, err
}

func (c *calendarService) Update(cal models.Calendar) (*models.Calendar, error) {
	err := cal.Update(calendardb.DB)
	return &cal, err
}

func (c *calendarService) Delete(calendarId string) (string, error) {
	cal := models.Calendar{Base: models.Base{ID: calendarId}}
	err := cal.Delete(calendardb.DB)
	return cal.ID, err
}
