package starter

import (
	"database/sql"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/pkg/cmd"
	"github.com/lucasvmiguel/stock-api/pkg/http/server"
)

type Starter struct {
	DB *sql.DB

	config config
	gormDB *gorm.DB
	router *chi.Mux

	repositories repositories
	services     services
	handlers     handlers
}

func New() *Starter {
	return &Starter{}
}

func (s *Starter) Start() {
	var err error

	// loads config
	s.config, err = loadConfig()
	if err != nil {
		cmd.ExitWithError("failed to load config", err)
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
	s.gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		cmd.ExitWithError("failed to connect database", err)
	}
	spew.Dump(s.gormDB)

	// migrates the database
	s.gormDB.AutoMigrate(&entity.Product{})

	// adds sql DB to starter struct
	s.DB, err = s.gormDB.DB()
	if err != nil {
		cmd.ExitWithError("failed to return sql DB", err)
	}

	// creates repositories
	s.repositories, err = s.createRepositories()
	if err != nil {
		cmd.ExitWithError("failed to create repositories", err)
	}

	// creates services
	s.services, err = s.createServices()
	if err != nil {
		cmd.ExitWithError("failed to create services", err)
	}

	// creates handlers
	s.handlers, err = s.createHandlers()
	if err != nil {
		cmd.ExitWithError("failed to create handlers", err)
	}

	// creates routes
	s.router = chi.NewRouter()
	s.createRoutes()

	// start http server
	server.Serve(s.config.Port, s.router)
}
