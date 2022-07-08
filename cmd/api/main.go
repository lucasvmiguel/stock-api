package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/internal/product/handler"
	"github.com/lucasvmiguel/stock-api/internal/product/repository"
	"github.com/lucasvmiguel/stock-api/pkg/cmd"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type config struct {
	Port int `env:"PORT" envDefault:"8080"`
}

func main() {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		cmd.ExitWithError("failed to read config", err)
	}

	dbClient, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		cmd.ExitWithError("failed to connect database", err)
	}

	// Migrate the schema (TODO: fix schema?)
	dbClient.AutoMigrate(&entity.Product{})

	productRepository, err := repository.NewRepository(dbClient)
	if err != nil {
		cmd.ExitWithError("failed to create product repository", err)
	}

	router := chi.NewRouter()

	// middlewaress
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// product handler
	productHandler, err := handler.NewHandler(productRepository)
	if err != nil {
		cmd.ExitWithError("product handler had an error", err)
	}

	// product routes
	router.Get("/products", productHandler.HandleGetAll)
	router.Post("/products", productHandler.HandleCreate)
	router.Get("/products/{id}", productHandler.HandleGetByID)
	router.Delete("/products/{id}", productHandler.HandleDeleteByID)
	router.Put("/products/{id}", productHandler.HandleUpdate)
	router.Patch("/products/{id}", productHandler.HandleUpdate)

	log.Printf("listening on port %d", cfg.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router))
}
