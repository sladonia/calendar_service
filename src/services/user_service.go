package services

import (
	"calendar_service/src/datasources/postgres/calendardb"
	"calendar_service/src/models"
)

var (
	UserService UserServiceInterface = &userService{}
)

type UserServiceInterface interface {
	Create(usr models.User) (*models.User, error)
	Read(userId string) (*models.User, error)
	Delete(userId string) (string, error)
	Update(usr models.User) (*models.User, error)
}

type userService struct{}

func (s *userService) Create(usr models.User) (*models.User, error) {
	err := usr.Create(calendardb.DB)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func (s *userService) Read(userId string) (*models.User, error) {
	usr := models.User{Base: models.Base{ID: userId}}
	err := usr.Read(calendardb.DB)
	return &usr, err
}

func (s *userService) Delete(userId string) (string, error) {
	usr := models.User{Base: models.Base{ID: userId}}
	err := usr.Delete(calendardb.DB)
	if err != nil {
		return "", err
	}
	return usr.ID, nil
}

func (s *userService) Update(usr models.User) (*models.User, error) {
	usr.Appointments = nil // we do not update appointments using this api
	err := usr.Update(calendardb.DB)
	return &usr, err
}
