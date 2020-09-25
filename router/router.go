package router

import (
	"net/http"

	// go-chi/chi
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	// go-chi/render
	"github.com/go-chi/render"
)

// // Router will have Mux() methods to return each multiplexer and 
// type Router interface {
// 	Mux() *http.ServeMux
// 	Route()
// }

// New : Initializes a chi.NewRouter instance and return its multiplexier
func New() *chi.Mux {
	// Initialize a new router
	router := chi.NewRouter()

	// Get everything set... 
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	
	// Tasks for GET request
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		// TODO : Write something

	})

	router.Route("/something", func(router chi.Router) {
		// TODO: May write something
	})

	return router
}
