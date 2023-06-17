// starter package is responsible for starting the application
package starter

import (
	"database/sql"
	"fmt"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/pkg/http/server"
	"github.com/lucasvmiguel/stock-api/pkg/logger"
)

// Starter is the struct that holds all dependencies
type Starter struct {
	// DB is the sql database
	// the database is exported to facilitate integration testing
	DB *sql.DB

	// config is the application config
	// it is used everywhere in this package
	// therefore, it is not passed as an argument
	// but as a struct field of Starter
	config config
}

// New creates a new Starter
func New() *Starter {
	return &Starter{}
}

// Start starts the application
func (s *Starter) Start() {
	var err error

	// loads config
	s.config, err = loadConfig()
	if err != nil {
		logger.Fatal("failed to load config", err)
	}

	// creates dsn string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		s.config.DBHost,
		s.config.DBUser,
		s.config.DBPassword,
		s.config.DBName,
		s.config.DBPort,
	)

	// starts connection with database
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database", err)
	}

	// migrates the database
	gormDB.AutoMigrate(&entity.Product{})

	// adds sql DB to starter struct
	s.DB, err = gormDB.DB()
	if err != nil {
		logger.Fatal("failed to return sql DB", err)
	}

	// creates repositories
	repositories, err := s.createRepositories(createRepositoriesArgs{gormDB: gormDB})
	if err != nil {
		logger.Fatal("failed to create repositories", err)
	}

	// creates services
	services, err := s.createServices(createServicesArgs{repositories: repositories})
	if err != nil {
		logger.Fatal("failed to create services", err)
	}

	// creates handlers
	handlers, err := s.createHandlers(createHandlersArgs{services: services})
	if err != nil {
		logger.Fatal("failed to create handlers", err)
	}

	// creates routes
	router := chi.NewRouter()
	s.createRoutes(createRoutesArgs{router: router, handlers: handlers})

	// start http server
	server.Serve(s.config.Port, router)
}
