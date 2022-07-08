package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucasvmiguel/stock-api/internal/product/handler"
	"github.com/lucasvmiguel/stock-api/internal/product/repository"
	"github.com/lucasvmiguel/stock-api/pkg/cmd"
	"github.com/lucasvmiguel/stock-api/pkg/ping"

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

	// // Migrate the schema
	dbClient.AutoMigrate(&repository.Product{})

	router := chi.NewRouter()

	// middlewaress
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request
	router.Use(middleware.Timeout(60 * time.Second))

	// ping handler
	router.Handle("/ping", &ping.Handler{})

	// post handler [POST]
	postHandlerPost, err := handler.NewHandlerPost(dbClient)
	if err != nil {
		cmd.ExitWithError("post handler post had an error", err)
	}
	router.Post("/products", postHandlerPost.ServeHTTP)

	log.Printf("listening on port %d", cfg.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router))
}
