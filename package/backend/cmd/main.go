package main

import (
	// "context"
	"fmt"
	"log"
	"net/http"
	// "os"

	"github.com/gorilla/mux"
	// "github.com/jackc/pgx/v5"
)

const serverAddr = ":8080"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello_world", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello world!")
	}).Methods("GET")

	log.Println("Server listening on ", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, r))
}
