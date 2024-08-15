package handler

import (
	"encoding/json"
	"net/http"

	handler_service "github.com/caseajetor/gopportunities/handler/service"
)

func GetOpening(w http.ResponseWriter, r *http.Request) {
	/* ctx := r.Context() */

	request := struct {
		role string
	}{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
func CreateOpening(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateOpeningRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	if err = handler_service.CreateOpening(ctx, req); err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
func UpdateOpening(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
func DeleteOpening(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
func ListOpening(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
