package sockjs

import (
	"net/http"

	"github.com/i9m/sockjs-go/v3/sockjs"
)

type Session = sockjs.Session
type Options = sockjs.Options
type Handler = sockjs.Handler
type SessionState = sockjs.SessionState

const (
	SessionOpening = sockjs.SessionOpening
	SessionActive  = sockjs.SessionActive
	SessionClosing = sockjs.SessionClosing
	SessionClosed  = sockjs.SessionClosed
)

var DefaultOptions = sockjs.DefaultOptions
var ErrErrSessionNotOpen = sockjs.ErrSessionNotOpen

func NewHandler(prefix string, opts Options, handlerFunc func(Session)) *Handler {
	return sockjs.NewHandler(prefix, opts, handlerFunc)
}

func DefaultJSessionID(rw http.ResponseWriter, req *http.Request) {
	sockjs.DefaultJSessionID(rw, req)
}
