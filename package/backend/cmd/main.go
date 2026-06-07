package main

import (
	"context"
	// "fmt"
	"log"
	"net/http"

	// "github.com/jackc/pgx/v5"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ISNUFFI/booking/internal/user"
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

	userHandler := user.NewHandler(pool, config)
	userHandler.AttachHandlers(r)

	log.Println("Server listening on ", config.AppAddress)
	log.Fatal(http.ListenAndServe(config.AppAddress, r))
}
