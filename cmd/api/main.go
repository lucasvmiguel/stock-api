package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/internal/product/handler"
	"github.com/lucasvmiguel/stock-api/internal/product/repository"
	"github.com/lucasvmiguel/stock-api/pkg/cmd"
	"gorm.io/driver/postgres"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

type Config struct {
	DBPort     string
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	Port       string
}

func main() {
	config := loadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	dbClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		cmd.ExitWithError("failed to connect database", err)
	}

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

	log.Printf("listening on port %s", config.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router))
}

func loadConfig() Config {
	return Config{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}
