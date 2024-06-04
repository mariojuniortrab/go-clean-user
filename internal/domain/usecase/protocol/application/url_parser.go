package protocol_application

import "net/http"

type URLParser interface {
	GetPathParamFromURL(r *http.Request, key string) string
	GetQueryParamFromURL(r *http.Request, key string) string
}
