package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/kaisersuzaku/BE_A/handlers"

	"github.com/kaisersuzaku/BE_A/services"

	"github.com/kaisersuzaku/BE_A/repo"

	"github.com/kaisersuzaku/BE_A/utils"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cfg := utils.GetConfig(dir + "/conf.json")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cfgCors := cfg.Cors
	if cfgCors.IsEnabled {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   cfgCors.AllowedOrigins,
			AllowedMethods:   cfgCors.AllowedMethods,
			AllowedHeaders:   cfgCors.AllowedHeaders,
			AllowCredentials: cfgCors.AllowCredentials,
			MaxAge:           cfgCors.MaxAge,
		}))
	}

	opr := repo.BuildProductRepo(utils.GetDB())
	ops := services.BuildOrderProductService(opr)
	oph := handlers.BuildOrderProductHandler(ops)

	fbcs := services.FillBallContainerService{}
	fbch := handlers.BuildFBCHandler(&fbcs)
	r.Post("/ball-container-check", fbch.CheckBallContainer)
	r.Post("/order-product", oph.OrderProduct)
	r.Get("/get-product", oph.GetProductByID)

	http.ListenAndServe(":7789", r)
}
