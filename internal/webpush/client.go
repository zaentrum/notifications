// Package webpush is the W3C Web Push dispatcher for browser clients
// (chino-web, tv-web, musig-web).
//
// Stub: needs a VAPID key pair. Real implementation should use
// github.com/SherClockHolmes/webpush-go. Build deferred to Phase 2
// once a browser PWA exists with a service worker subscribed to the
// stube push endpoint.
package webpush

import (
	"context"
	"errors"
)

// ErrNotImplemented marks methods whose bodies will land in Phase 2.
var ErrNotImplemented = errors.New("webpush: not implemented (needs VAPID key pair)")

// Subscription is the W3C PushSubscription serialised form delivered
// by the browser to our subscription endpoint.
type Subscription struct {
	Endpoint string
	Auth     string
	P256dh   string
}

// Push describes a single browser push payload.
type Push struct {
	Sub   Subscription
	Title string
	Body  string
	URL   string
}

// Client is the WebPush facade.
type Client interface {
	Send(ctx context.Context, p Push) error
}

// NewStub returns a placeholder Client that returns ErrNotImplemented.
func NewStub() Client { return &stub{} }

type stub struct{}

func (s *stub) Send(_ context.Context, _ Push) error { return ErrNotImplemented }
