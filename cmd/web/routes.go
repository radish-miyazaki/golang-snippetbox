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

	mux.Get("/ping", ping)

	mux.Group(func(mux chi.Router) {
		mux.Use(app.session.LoadAndSave)
		mux.Use(noSurf)
		mux.Use(app.authenticate)

		mux.Get("/", app.home)
		mux.Get("/about", app.about)

		mux.Group(func(mux chi.Router) {
			mux.Use(app.requireAuthentication)

			mux.Get("/snippet/create", app.createSnippetForm)
			mux.Post("/snippet/create", app.createSnippet)
			mux.Post("/user/logout", app.logoutUser)
		})

		mux.Get("/snippet/{id}", app.showSnippet)

		mux.Get("/user/signup", app.signupUserForm)
		mux.Post("/user/signup", app.signupUser)
		mux.Get("/user/login", app.loginUserForm)
		mux.Post("/user/login", app.loginUser)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
