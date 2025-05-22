package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)
	mux.Get("/GetContexParams", app.GetContexParams)
	mux.Get("/CheckSiteURL/{sitename}", app.CheckSiteURL)
	mux.Get("/GetAuditScripts", app.GetAuditScripts)
	mux.Post("/ExecuteAudit", app.ExecuteAudit)

	/*
		mux.Get("/Home", app.Home)


		mux.Post("/authenticate", app.authenticate)
		mux.Get("/refresh", app.refreshToken)
		mux.Get("/logout", app.logout)
		mux.Get("/genres", app.AllGenres)

		mux.Get("/movies", app.AllMovies)
		mux.Get("/movies/{id}", app.GetMovie)

		mux.Post("/graph", app.movieGraphQL)

		mux.Get("/movies/genres/{id}", app.AllMoviesByGenre)
	*/
	mux.Route("/admin", func(mux chi.Router) {
		/*
			mux.Use(app.authRequired)


				mux.Get("/movies", app.movieCatalog)
				mux.Get("/movies/{id}", app.GetMovieForEdit)

				mux.Put("/movies/0", app.InsertMovie)
				mux.Patch("/movies/{id}", app.UpdateMovie)
				mux.Delete("/movies/{id}", app.DeleteMovie)
		*/

	})
	return mux
}
