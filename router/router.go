package router

import (
	"net/http"

	// go-chi/chi
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	// go-chi/render
	"github.com/go-chi/render"
)

type Router interface {
	Mux() *http.ServeMux
	Route()
}

func new() router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))
}
