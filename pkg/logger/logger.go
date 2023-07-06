// logger is a package responsible for logging
package logger

import (
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Fatal logs a message and exits with status 1
func Fatal(message string, err error) {
	log.Fatal().Err(err).Msg(message)
}

// Error logs an errror message
func Error(message string, err error) {
	log.Error().Err(err).Msg(message)
}

// Infof logs a message with arguments
func Infof(message string, args ...interface{}) {
	log.Info().Msgf(message, args...)
}

// Info logs a message
func Info(message string) {
	log.Info().Msg(message)
}

// HTTPMiddleware returns a middleware for logging HTTP requests
func HTTPMiddleware(name string) func(next http.Handler) http.Handler {
	logger := httplog.NewLogger(name)

	return httplog.RequestLogger(logger)
}

// HTTPLogEntry returns a logger for HTTP requests
func HTTPLogEntry(req *http.Request) zerolog.Logger {
	return httplog.LogEntry(req.Context())
}
