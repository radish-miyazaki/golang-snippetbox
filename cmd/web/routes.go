package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(secureHeaders)

	mux.Group(func(r chi.Router) {
		r.Use(app.session.LoadAndSave)

		r.Get("/", app.home)
		r.Get("/snippet/create", app.createSnippetForm)
		r.Post("/snippet/create", app.createSnippet)
		r.Get("/snippet/{id}", app.showSnippet)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
