package protocol_application

import "net/http"

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
