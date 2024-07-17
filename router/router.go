package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Initialize() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/welcome", func(w http.ResponseWriter, router *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", router)

}
