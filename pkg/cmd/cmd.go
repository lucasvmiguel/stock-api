package cmd

import (
	"log"
	"os"
)

// helper to exit the application
func ExitWithError(message string, err error) {
	log.Printf("%s: %s", message, err.Error())
	os.Exit(1)
}
