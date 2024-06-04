package infra_adapters

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
)

type chiRouteAdapter struct {
	chi chi.Router
}

func NewChiRouteAdapter() *chiRouteAdapter {
	return &chiRouteAdapter{
		chi: chi.NewRouter(),
	}
}

func (a *chiRouteAdapter) Use(middlewares ...func(http.Handler) http.Handler) {
	a.chi.Use(middlewares...)
}

func (a *chiRouteAdapter) With(middlewares ...func(http.Handler) http.Handler) protocol_application.Router {
	return &chiRouteAdapter{chi: a.chi.With(middlewares...)}
}

func (a *chiRouteAdapter) Route(pattern string, fn func(r protocol_application.Router)) protocol_application.Router {
	chiFn := func(r chi.Router) {
		newAdapter := &chiRouteAdapter{
			chi: r,
		}
		fn(newAdapter)
	}

	a.chi.Route(pattern, chiFn)
	return a
}

func (a *chiRouteAdapter) Group(fn func(r protocol_application.Router)) protocol_application.Router {
	chiFn := func(r chi.Router) {
		newAdapter := &chiRouteAdapter{
			chi: r,
		}
		fn(newAdapter)
	}

	a.chi.Group(chiFn)
	return a
}

// HTTP-method routing along `pattern`
func (a *chiRouteAdapter) Connect(pattern string, h http.HandlerFunc) {
	a.chi.Connect(pattern, h)
}

func (a *chiRouteAdapter) Delete(pattern string, h http.HandlerFunc) {
	a.chi.Delete(pattern, h)
}

func (a *chiRouteAdapter) Get(pattern string, h http.HandlerFunc) {
	a.chi.Get(pattern, h)
}

func (a *chiRouteAdapter) Head(pattern string, h http.HandlerFunc) {
	a.chi.Head(pattern, h)
}

func (a *chiRouteAdapter) Options(pattern string, h http.HandlerFunc) {
	a.chi.Get(pattern, h)
}

func (a *chiRouteAdapter) Patch(pattern string, h http.HandlerFunc) {
	a.chi.Patch(pattern, h)
}

func (a *chiRouteAdapter) Post(pattern string, h http.HandlerFunc) {
	a.chi.Post(pattern, h)
}

func (a *chiRouteAdapter) Put(pattern string, h http.HandlerFunc) {
	a.chi.Put(pattern, h)
}

func (a *chiRouteAdapter) Trace(pattern string, h http.HandlerFunc) {
	a.chi.Trace(pattern, h)
}

func (a *chiRouteAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.chi.ServeHTTP(w, r)
}

func (a *chiRouteAdapter) GetPathParamFromURL(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func (a *chiRouteAdapter) PrintRoutes() {
	a.printSubRoutes(a.chi.Routes(), "")
}

func (a *chiRouteAdapter) printSubRoutes(routes []chi.Route, parentRoute string) {
	for _, route := range routes {
		fmt.Println("route:", parentRoute, route.Pattern)
		if route.SubRoutes != nil {
			a.printSubRoutes(route.SubRoutes.Routes(), route.Pattern)
		}
	}
}
