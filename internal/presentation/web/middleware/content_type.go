package web_middleware

import (
	"net/http"
	"strings"

	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type contentType struct{}

func NewContentType() *contentType {
	return &contentType{}
}

func (p *contentType) GetMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			ct := r.Header.Get("Content-Type")
			if ct != "" {
				mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
				if mediaType != "application/json" {
					web_response_manager.RespondUnsupportedMediaType(w)
					return
				}
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
