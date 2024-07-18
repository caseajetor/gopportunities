package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Initialize() {
	router := chi.NewRouter()

	initializeRoutes(router)

	http.ListenAndServe(":3000", router)
}
