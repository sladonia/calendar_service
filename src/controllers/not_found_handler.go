package controllers

import (
	"calendar_service/src/logger"
	"fmt"
	"net/http"
)

type NotFoundHandler struct{}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debugw("incoming request", "method", r.Method, "path", r.URL.Path)
	apiErr := NewNotFoundApiError(fmt.Sprintf("resource %s %s not found", r.Method, r.URL.Path))
	RespondError(w, apiErr)
}
