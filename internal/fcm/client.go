// Package fcm is the Firebase Cloud Messaging dispatcher.
//
// Stub: needs FCM service-account credentials JSON. Real implementation
// should use firebase.google.com/go/v4/messaging via HTTP v1 API.
// Build deferred to Phase 2 once an Android client exists.
package fcm

import (
	"context"
	"errors"
)

// ErrNotImplemented marks methods whose bodies will land in Phase 2.
var ErrNotImplemented = errors.New("fcm: not implemented (needs FCM service-account creds)")

// Push describes a single Android device push.
type Push struct {
	Token string
	Title string
	Body  string
	Data  map[string]string
}

// Client is the FCM facade.
type Client interface {
	Send(ctx context.Context, p Push) error
}

// NewStub returns a placeholder Client that returns ErrNotImplemented.
func NewStub() Client { return &stub{} }

type stub struct{}

func (s *stub) Send(_ context.Context, _ Push) error { return ErrNotImplemented }
