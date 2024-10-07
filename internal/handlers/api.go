package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/pythonakoto/goapi/internal/middleware"
)

// functiona that beigin with capital letters tell the compiler that it can be used GLOBALLY
func Handler(r *chi.Mux) {
	// global middleware
	r.Use(chimiddle.StripSlashes) // gets rid of trailing slashes from a /route/ -> /route

	// set up route that we will use
	r.Route("/account", func(router chi.Router) {

		// middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
