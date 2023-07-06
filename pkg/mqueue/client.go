package mqueue

import (
	"fmt"

	"github.com/hibiken/asynq"
)

const (
	// DefaultQueue is the default queue name
	DefaultQueue = "default"
)

// Client accesses the message queue
type Client struct {
	client    *asynq.Client
	inspector *asynq.Inspector
}

// NewClientArgs are the arguments to create a new Client
type NewClientArgs struct {
	Host string
	Port string
}

// NewClient creates a new Client
func NewClient(args NewClientArgs) *Client {
	address := queueAddress(args.Host, args.Port)

	return &Client{
		client:    asynq.NewClient(asynq.RedisClientOpt{Addr: address}),
		inspector: asynq.NewInspector(asynq.RedisClientOpt{Addr: address}),
	}
}

// Enqueue enqueues a task
func (c *Client) Enqueue(task *asynq.Task) error {
	_, err := c.client.Enqueue(task, asynq.Queue(DefaultQueue))
	return err
}

// Close closes the connection to the message queue
func (c *Client) Close() error {
	return c.client.Close()
}

// Inspector returns the inspector
func (c *Client) Inspector() *asynq.Inspector {
	return c.inspector
}

func queueAddress(host, port string) string {
	return fmt.Sprint(host, ":", port)
}
