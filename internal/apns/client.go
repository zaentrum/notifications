// Package apns is the Apple Push Notification service dispatcher.
//
// Stub: needs an APNs auth key (.p8), Team ID, and Key ID. The real
// implementation should use github.com/sideshow/apns2 with HTTP/2 token
// authentication. Build deferred to Phase 2 once an iOS client exists.
package apns

import (
	"context"
	"errors"
)

// ErrNotImplemented marks methods whose bodies will land in Phase 2.
var ErrNotImplemented = errors.New("apns: not implemented (needs APNs key + Team ID + Key ID)")

// Push describes a single device push.
type Push struct {
	DeviceToken string
	Title       string
	Body        string
	Category    string
	Payload     map[string]any
}

// Client is the APNs facade.
type Client interface {
	Send(ctx context.Context, p Push) error
}

// NewStub returns a placeholder Client that returns ErrNotImplemented.
func NewStub() Client { return &stub{} }

type stub struct{}

func (s *stub) Send(_ context.Context, _ Push) error { return ErrNotImplemented }
