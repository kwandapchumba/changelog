package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/kwandapchumba/changelog/api"
	"github.com/kwandapchumba/changelog/db"
	myMiddleware "github.com/kwandapchumba/changelog/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.AllowContentEncoding("application/json", "application/x-www-form-urlencoded"))
	r.Use(middleware.CleanPath)
	r.Use(middleware.RedirectSlashes)

	h := api.NewBaseHandler(db.ConnectDB())

	r.Route("/public", func(r chi.Router) {
		r.Post("/addNewUser", h.AddNewUser)
		r.Post("/getOtp", h.GetOtp)
		r.Post("/verifyOtp", h.VerifyOtp)
	})

	r.Route("/private", func(r chi.Router) {
		r.Use(myMiddleware.Authenticator())
		r.Post("/addCompany", h.AddCompany)
	})

	return r
}
