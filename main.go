package main

import (
	"net/http"

	"github.com/kaisersuzaku/BE_A/handlers"

	"github.com/kaisersuzaku/BE_A/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fbcs := services.FillBallContainerService{}
	fbch := handlers.BuildFBCHandler(&fbcs)
	r.Post("/ball-container-check", fbch.CheckBallContainer)

	http.ListenAndServe(":7789", r)
}
