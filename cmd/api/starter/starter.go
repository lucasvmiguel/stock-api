package starter

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/internal/product/handler"
	"github.com/lucasvmiguel/stock-api/internal/product/repository"
	"github.com/lucasvmiguel/stock-api/internal/product/service"
	"github.com/lucasvmiguel/stock-api/pkg/cmd"
	"github.com/lucasvmiguel/stock-api/pkg/env"
	"github.com/lucasvmiguel/stock-api/pkg/http/server"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	DBPort                 string
	DBHost                 string
	DBName                 string
	DBUser                 string
	DBPassword             string
	Port                   string
	ENV                    env.Environment
	PaginationDefaultLimit int
}

type Starter struct {
	DB *sql.DB
}

func New() *Starter {
	return &Starter{}
}

func (s *Starter) Start() {
	config, err := loadConfig()
	if err != nil {
		cmd.ExitWithError("failed to load config", err)
	}

	var gormDB *gorm.DB

	// starts connection with database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		cmd.ExitWithError("failed to connect database", err)
	}

	// migrates the database
	gormDB.AutoMigrate(&entity.Product{})

	// creates product repository
	productRepository, err := repository.NewRepository(gormDB)
	if err != nil {
		cmd.ExitWithError("failed to create product repository", err)
	}

	s.DB, err = gormDB.DB()
	if err != nil {
		cmd.ExitWithError("failed to return sql DB", err)
	}

	router := chi.NewRouter()

	// http middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// product service
	productService, err := service.NewService(productRepository)
	if err != nil {
		cmd.ExitWithError("product service had an error", err)
	}

	// product http handler
	productHandler, err := handler.NewHandler(handler.NewHandlerArgs{
		Service:                productService,
		PaginationDefaultLimit: config.PaginationDefaultLimit,
	})
	if err != nil {
		cmd.ExitWithError("product handler had an error", err)
	}

	router.Route("/api/v1", func(r chi.Router) {
		// product http routes
		r.Get("/products", productHandler.HandleGetPaginated)
		r.Get("/products/all", productHandler.HandleGetAll)
		r.Post("/products", productHandler.HandleCreate)
		r.Get(fmt.Sprintf("/products/{%s}", handler.FieldID), productHandler.HandleGetByID)
		r.Delete(fmt.Sprintf("/products/{%s}", handler.FieldID), productHandler.HandleDeleteByID)
		r.Put(fmt.Sprintf("/products/{%s}", handler.FieldID), productHandler.HandleUpdate)
		r.Patch(fmt.Sprintf("/products/{%s}", handler.FieldID), productHandler.HandleUpdate)
	})

	// health http route
	router.Get("/health", func(w http.ResponseWriter, req *http.Request) { w.Write([]byte("Up and running")) })

	// start http server
	server.Serve(config.Port, router)
}

func loadConfig() (config, error) {
	paginationDefaultLimitStr := os.Getenv("PAGINATION_DEFAULT_LIMIT")
	paginationDefaultLimit, err := strconv.Atoi(paginationDefaultLimitStr)
	if err != nil {
		return config{}, errors.Wrap(err, "failed to convert PAGINATION_DEFAULT_LIMIT env var to int")
	}

	return config{
		Port:                   os.Getenv("PORT"),
		DBHost:                 os.Getenv("DB_HOST"),
		DBName:                 os.Getenv("DB_NAME"),
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBPort:                 os.Getenv("DB_PORT"),
		ENV:                    env.Environment(os.Getenv("ENV")),
		PaginationDefaultLimit: paginationDefaultLimit,
	}, nil
}
