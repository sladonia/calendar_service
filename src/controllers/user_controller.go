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
	UserController UserControllerInterface = &userController{}
)

type UserControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type userController struct{}

func (u *userController) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMsg := "invalid request body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()
	var usr models.User
	err = json.Unmarshal(requestBody, &usr)
	if err != nil {
		errorMsg := "invalid json body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		RespondError(w, apiErr)
		return
	}
	resultUsr, err := services.UserService.Create(usr)
	if err != nil {
		errorMsg := "unable to crate user"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusConflict)
		RespondError(w, apiErr)
		return
	}
	response := models.ResponseCreated{
		Message:   "user created",
		CreatedId: resultUsr.ID,
	}
	RespondJSON(w, http.StatusCreated, response)
}

func (u *userController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	ok := IsValidUUID(userId)
	if !ok {
		logger.Logger.Infof("received invalid uuid=%s", userId)
		apiErr := NewBadRequestApiError("invalid uuid")
		RespondError(w, apiErr)
		return
	}
	response, err := services.UserService.Read(userId)
	if err != nil {
		errorMsg := "unable to get user"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusNotFound)
		RespondError(w, apiErr)
		return
	}

	RespondJSON(w, http.StatusOK, response)
}

func (u *userController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
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
	var usr models.User
	err = json.Unmarshal(requestBody, &usr)
	if err != nil {
		errorMsg := "invalid json body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		RespondError(w, apiErr)
		return
	}
	usr.ID = userId
	result, err := services.UserService.Update(usr)
	if err != nil {
		errorMsg := "unable to update user"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusConflict)
		RespondError(w, apiErr)
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func (u *userController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	ok := IsValidUUID(userId)
	if !ok {
		logger.Logger.Infof("received invalid uuid=%s", userId)
		apiErr := NewBadRequestApiError("invalid uuid")
		RespondError(w, apiErr)
		return
	}

	deletedId, err := services.UserService.Delete(userId)
	if err != nil {
		errorMsg := "unable to delete user"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusNotFound)
		RespondError(w, apiErr)
		return
	}
	response := models.ResponseDeleted{
		Message:   "deleted",
		DeletedId: deletedId,
	}
	RespondJSON(w, http.StatusAccepted, response)
}
