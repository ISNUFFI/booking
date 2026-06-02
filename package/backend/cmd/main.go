package main

import (
	"context"
	// "fmt"
	"log"
	"net/http"
	"os"

	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"

	"github.com/ISNUFFI/booking/internal/user"
)

const serverAddr = ":8080"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("could not acquire a postgres connection pool: ", err)
	}
	defer pool.Close()

	userHandler := user.NewHandler(pool)
	userHandler.AttachHandlers(r)

	log.Println("Server listening on ", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, r))
}
