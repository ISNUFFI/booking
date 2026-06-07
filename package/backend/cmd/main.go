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
	authHandler.AttachHandlers(r)

	r.Group(func(pr chi.Router) {
		pr.Use(auth.JWTMiddleware([]byte(config.JWTSecret)))

		// private endpoints
		pr.Get("/me", authHandler.MeHandler)
	})

	log.Println("Server listening on ", config.AppAddress)
	log.Fatal(http.ListenAndServe(config.AppAddress, r))
}
