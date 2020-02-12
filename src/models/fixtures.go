package models

import "github.com/jinzhu/gorm"

const (
	knownUserId     = "123e4567-e89b-12d3-a456-426655440000"
	knownCalendarId = "b09b7b26-4d83-11ea-b1e0-c83a35cc61f1"

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

	return nil
}
