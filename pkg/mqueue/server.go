package mqueue

import (
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/pkg/logger"
)

// Handlers is a map of message queue handlers
type Handlers map[string]asynq.HandlerFunc

// Server is the message queue server
type Server struct {
	address  string
	handlers Handlers
}

// NewServerArgs are the arguments to create a new Server
type NewServerArgs struct {
	Address  string
	Handlers Handlers
}

// NewServer creates a new Server
func NewServer(args NewServerArgs) *Server {
	return &Server{
		handlers: args.Handlers,
		address:  args.Address,
	}
}

// Consumes the message queue
func (s *Server) Consume() error {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: s.address},
		asynq.Config{},
	)

	mux := asynq.NewServeMux()

	for name, handler := range s.handlers {
		mux.HandleFunc(name, handler)
	}

	logger.Info("message queue consumer starting")

	err := srv.Run(mux)
	if err != nil {
		logger.Error("failed to start the message queue consumer", err)
		return errors.Wrap(err, "failed to start the message queue consumer")
	}

	logger.Info("message queue consumer stopped")

	return nil
}
