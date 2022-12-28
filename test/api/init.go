package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/lucasvmiguel/stock-api/cmd/api/starter"
)

const (
	START_SERVER_TRIES = 5
)

var DB *sql.DB

func init() {
	godotenv.Load("../../.env")

	s := starter.New()
	go s.Start()

	fmt.Println("Starting integration test...")

	// wait for the api to start
	tries := START_SERVER_TRIES
	for {
		if tries <= 0 {
			panic("server hasn't started")
		}

		resp, err := http.Get("http://localhost:8080/health")
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}

		fmt.Printf("Starting server (%d tries)\n", START_SERVER_TRIES-tries+1)
		tries--
		time.Sleep(1 * time.Second)
	}

	DB = s.DB
}