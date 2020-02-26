package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bramsedana/law-cots-1/entity"
	"github.com/bramsedana/law-cots-1/usecase"
	"github.com/julienschmidt/httprouter"
)

// Handler for now is intended for checking application health testing purposes only
type Handler struct {
}

// NewHandler returns a pointer of Handler instance
func NewHandler() *Handler {
	return &Handler{}
}

// Healthz is used to control the flow of GET /healthz endpoint
func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	OK(w, nil, "ok")
}

func (h *Handler) Arithmetic(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	equation := entity.Equation{}

	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&equation)
	defer r.Body.Close()
	result := usecase.Arithmetic(equation)
	OK(w, result, "")
}
