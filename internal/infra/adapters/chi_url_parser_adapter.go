package infra_adapters

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiUrlParserAdapter struct{}

func NewChiUrlParserAdapter() *chiUrlParserAdapter {
	return &chiUrlParserAdapter{}
}

func (a *chiUrlParserAdapter) GetPathParamFromURL(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func (a *chiUrlParserAdapter) GetQueryParamFromURL(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
