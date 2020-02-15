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
