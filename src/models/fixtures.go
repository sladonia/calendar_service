package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	KnownUserId            = "123e4567-e89b-12d3-a456-426655440000"
	SecondKnownUserId      = "415f3e7a-4320-4a3b-be10-ccbbebd21bcf"
	ThirdKnownUserId       = "ec819614-fb93-498e-9360-3e3d41301599"
	KnownCalendarId        = "b09b7b26-4d83-11ea-b1e0-c83a35cc61f1"
	AppointmentFixedTimeId = "4e7ca4b6-4da5-11ea-b1e0-c83a35cc61f1"
	AppointmentWholeDayId  = "33148505-7595-4c2a-9a45-bc885d0910a6"

	UnexistingId = "12345678-1234-5678-1234-567812345678"
)

func MockDbData(db *gorm.DB) error {
	userJhon := &User{Base: Base{ID: KnownUserId}, FirstName: "John", LastName: "Carmack", Email: "jhon@gmail.com"}
	userKotlin := &User{Base: Base{ID: SecondKnownUserId}, FirstName: "Kotlin", LastName: "Jackson", Email: "kotlinjackson@gmail.com"}
	thirdUser := &User{Base: Base{ID: ThirdKnownUserId}, FirstName: "Patric", LastName: "Kolakovski", Email: "pkolakovski@gmail.com"}
	db.Create(userJhon)
	db.Create(userKotlin)
	db.Create(thirdUser)

	johnsCalendar := &Calendar{Base: Base{ID: KnownCalendarId}, Name: "John's personal calendar", UserId: userJhon.ID}
	kotlinsCalendar := &Calendar{Name: "Kotlin's meetings", UserId: userKotlin.ID}

	db.Create(johnsCalendar)
	db.Create(kotlinsCalendar)

	jonsFixedTimeAppointment := &Appointment{
		Base:        Base{ID: AppointmentFixedTimeId},
		Subject:     "Meet friends",
		Description: "just have fun",
		CalendarId:  KnownCalendarId,
		Start:       time.Date(2020, 1, 17, 20, 0, 0, 0, time.UTC),
		End:         time.Date(2020, 1, 17, 22, 30, 0, 0, time.UTC),
		WholeDay:    false,
	}

	jonsWholeDayAppointment := &Appointment{
		Base:        Base{ID: AppointmentWholeDayId},
		Subject:     "take a rest",
		Description: "have fun twice a dau",
		CalendarId:  KnownCalendarId,
		Start:       time.Date(2020, 1, 18, 11, 0, 0, 0, time.UTC),
		WholeDay:    true,
		Attendees:   []*User{thirdUser},
	}

	db.Create(jonsFixedTimeAppointment)
	db.Create(jonsWholeDayAppointment)

	return nil
}
