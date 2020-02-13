package controllers

import "net/http"

var (
	RootController RootControllerInterface = &rootController{}
)

type RootControllerInterface interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type rootController struct{}

func (c *rootController) Get(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: "welcome to calendar api"})
}
