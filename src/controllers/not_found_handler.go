package controllers

import (
	"fmt"
	"net/http"
)

type NotFoundHandler struct{}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiErr := NewNotFoundApiError(fmt.Sprintf("resource %s %s not found", r.Method, r.URL.Path))
	RespondError(w, apiErr)
}
