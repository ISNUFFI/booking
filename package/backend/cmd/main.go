package main

import (
	"context"
	"log"
	"net/http"

	// "github.com/jackc/pgx/v5"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ISNUFFI/booking/internal/auth"
	"github.com/ISNUFFI/booking/internal/config"
	"github.com/ISNUFFI/booking/internal/providers"
	"github.com/ISNUFFI/booking/internal/slots"
	"github.com/ISNUFFI/booking/internal/users"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	config := config.Load()

	pool, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		log.Fatal("could not acquire a postgres connection pool: ", err)
	}
	defer pool.Close()

	authHandler := auth.NewHandler(pool, config)
	usersHandler := users.NewHandler(pool)
	providersHandler := providers.NewHandler(pool)
	slotsHandler := slots.NewHandler(pool)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.RegisterHandler)
		r.Post("/login", authHandler.LoginHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.JWTMiddleware([]byte(config.JWTSecret)))

		r.Get("/me", usersHandler.MeHandler)

		r.Route("/providers", func(r chi.Router) {
			r.Get("/", providersHandler.GetListHandler)
			r.Post("/", providersHandler.CreateHandler)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", providersHandler.GetHandler)
				r.Delete("/", providersHandler.DeleteHandler)

				r.Get("/slots", slotsHandler.GetProviderSlotsHandler)
				r.Post("/slots/bulk", slotsHandler.CreateBulkHandler)
			})
		})

		r.Route("/slots", func(r chi.Router) {
			r.Get("/{id}", slotsHandler.GetHandler)
		})
	})

	log.Println("Server listening on ", config.AppAddress)
	log.Fatal(http.ListenAndServe(config.AppAddress, r))
}
