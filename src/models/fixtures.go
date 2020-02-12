package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	knownUserId            = "123e4567-e89b-12d3-a456-426655440000"
	knownCalendarId        = "b09b7b26-4d83-11ea-b1e0-c83a35cc61f1"
	appointmentFixedTimeId = "4e7ca4b6-4da5-11ea-b1e0-c83a35cc61f1"
	appointmentWholeDayId  = "33148505-7595-4c2a-9a45-bc885d0910a6"
	// 415f3e7a-4320-4a3b-be10-ccbbebd21bcf
	// ec819614-fb93-498e-9360-3e3d41301599

	unexistingId = "12345678-1234-5678-1234-567812345678"
)

func MockDbData(db *gorm.DB) error {
	userJhon := &User{Base: Base{ID: knownUserId}, FirstName: "John", LastName: "Carmack", Email: "jhon@gmail.com"}
	userKotlin := &User{FirstName: "Kotlin", LastName: "Jackson", Email: "kotlinjackson@gmail.com"}
	db.Create(userJhon)
	db.Create(userKotlin)

	johnsCalendar := &Calendar{Base: Base{ID: knownCalendarId}, Name: "John's personal calendar", UserId: userJhon.ID}
	kotlinsCalendar := &Calendar{Name: "Kotlin's meetings", UserId: userKotlin.ID}

	db.Create(johnsCalendar)
	db.Create(kotlinsCalendar)

	jonsFixedTimeAppointment := &Appointment{
		Base:        Base{ID: appointmentFixedTimeId},
		Subject:     "Meet friends",
		Description: "just have fun",
		CalendarId:  knownCalendarId,
		Start:       time.Date(2020, 1, 17, 20, 0, 0, 0, time.UTC),
		End:         time.Date(2020, 1, 17, 22, 30, 0, 0, time.UTC),
		WholeDay:    false,
	}

	jonsWholeDayAppointment := &Appointment{
		Base:        Base{ID: appointmentWholeDayId},
		Subject:     "take a rest",
		Description: "have fun twice a dau",
		CalendarId:  knownCalendarId,
		Start:       time.Date(2020, 1, 18, 11, 0, 0, 0, time.UTC),
		WholeDay:    true,
	}

	db.Create(jonsFixedTimeAppointment)
	db.Create(jonsWholeDayAppointment)

	return nil
}
