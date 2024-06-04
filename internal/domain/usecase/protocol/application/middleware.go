package protocol_application

import "net/http"

type MiddlewareGetter interface {
	Middleware(next http.Handler) http.Handler
}

type Middleware func(next http.Handler) http.Handler
