package models

import "github.com/jinzhu/gorm"

func MockDbData(db *gorm.DB) error {
	userJhon := &User{FirstName: "John", LastName: "Carmack", Email: "jhon@gmail.com"}
	userKotlin := &User{FirstName: "Kotlin", LastName: "Jackson", Email: "kotlinjackson@gmail.com"}

	db.Create(userJhon)
	db.Create(userKotlin)

	johnsCalendar := &Calendar{Name: "John's personal calendar", UserId: userJhon.ID}
	kotlinsCalendar := &Calendar{Name: "Kotlin's meetings", UserId: userKotlin.ID}

	db.Create(johnsCalendar)
	db.Create(kotlinsCalendar)

	return nil
}
