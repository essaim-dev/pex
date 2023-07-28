// Package pex implements an event subscription multiplexer.
//
// It is designed to be used jointly with the go-pex plugin for protoc.
package pex

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

// SubMux is an event subscriber multiplexer.
type SubMux struct {
	handlers map[string]HandlerFunc
	errCh    chan error
}

// HandlerFunc defines the spec of a SubMux event handler.
type HandlerFunc func(*message.Message) error

// NewSubMux allocates and returns a new SubMux.
func NewSubMux() *SubMux {
	return &SubMux{
		handlers: make(map[string]HandlerFunc),
		errCh:    make(chan error),
	}
}

// Handle registers the given handler for the given event topic.
//
// If a handler is aleady defined for the given topic, the handler will be overwritten.
func (s *SubMux) Handle(topic string, handler HandlerFunc) {
	s.handlers[topic] = handler
}

// Run starts the event multiplexer with the given subscriber.
//
// It is a blocking method which subscribes to its configured event topics
// and processes incoming events with the given handlers.
func (s *SubMux) Run(ctx context.Context, sub message.Subscriber) error {
	for topic, handler := range s.handlers {
		go s.runHandler(ctx, sub, topic, handler)
	}

	select {
	case err := <-s.errCh:
		return err
	case <-ctx.Done():
		return nil
	}
}

func (s *SubMux) runHandler(ctx context.Context, sub message.Subscriber, topic string, handler HandlerFunc) {
	messages, err := sub.Subscribe(ctx, topic)
	if err != nil {
		s.errCh <- err
		return
	}

	for {
		select {

		case message := <-messages:
			if err := handler(message); err != nil {
				message.Nack()
				continue
			}
			message.Ack()

		case <-ctx.Done():
			return

		}
	}
}
