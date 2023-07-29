package server

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lucasvmiguel/stock-api/pkg/logger"
)

const (
	timeout = 10 * time.Second
)

type ServeOpts struct {
	TLSCertificatePath string
	TLSPrivateKeyPath  string
}

// Serve an HTTP server (with graceful shutdown)
// Reference: https://medium.com/honestbee-tw-engineer/gracefully-shutdown-in-go-http-server-5f5e6b83da5a
func Serve(port string, router *chi.Mux, opts ServeOpts) {
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		Handler:      router,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		var err error

		if opts.TLSCertificatePath != "" && opts.TLSPrivateKeyPath != "" {
			err = srv.ListenAndServeTLS(opts.TLSCertificatePath, opts.TLSPrivateKeyPath)
		} else {
			err = srv.ListenAndServe()
		}

		if err != nil && err != http.ErrServerClosed {
			logger.Error("failed to start the http server", err)
		}
	}()

	logger.Infof("HTTP server on %s", fmt.Sprintf("http://%s", srv.Addr))

	<-done

	logger.Info("HTTP server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer func() { cancel() }()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("HTTP server shutdown failed", err)
	}
	logger.Info("HTTP server exited properly")
}
