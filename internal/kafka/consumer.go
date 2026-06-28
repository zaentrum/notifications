// Package kafka wraps a sarama consumer-group for the
// `stube.notify.user.*` topics.
//
// Stub: real implementation will use github.com/IBM/sarama with a
// ConsumerGroupHandler that decodes Avro per the contract published in
// the schemas repo (stube/notify/user-push.avsc) and dispatches into
// the apns / fcm / webpush clients.
//
// Build deferred to Phase 2.
package kafka

import (
	"context"
	"errors"
)

// ErrNotImplemented marks methods whose bodies will land in Phase 2.
var ErrNotImplemented = errors.New("kafka: not implemented")

// Consumer wraps a long-lived sarama consumer-group session.
type Consumer interface {
	// Run blocks until ctx is cancelled or an unrecoverable error occurs.
	Run(ctx context.Context) error
}

// NewStub returns a placeholder Consumer that returns ErrNotImplemented.
func NewStub() Consumer { return &stub{} }

type stub struct{}

func (s *stub) Run(_ context.Context) error { return ErrNotImplemented }
