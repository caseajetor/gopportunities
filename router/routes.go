package router

import (
	"github.com/caseajetor/gopportunities/handler"
	"github.com/go-chi/chi/v5"
)

func initializeRoutes(router *chi.Mux) {
	router.Route("/api/v1", func(router chi.Router) {
		router.Get("/opening", handler.GetOpening)
		router.Post("/opening", handler.CreateOpening)
		router.Put("/opening", handler.UpdateOpening)
		router.Delete("/opening", handler.DeleteOpening)
		router.Get("/openings", handler.ListOpening)
	})
}
