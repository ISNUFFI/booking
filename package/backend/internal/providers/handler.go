package providers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	service Service
}

func NewHandler(pool *pgxpool.Pool) Handler {
	return Handler{
		service: NewService(NewRepo(pool)),
	}
}

func (h Handler) AttachHandlers(router chi.Router) {
	router.Post("/providers", h.CreateProviderHandler)
	router.Get("/providers", h.GetProviderHandler)
}

type CreateProviderRequest struct {
	Name        string
	Description string
}

func (h Handler) CreateProviderHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateProviderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateProvider(r.Context(), req.Name, req.Description)
	if err != nil {
		log.Println("could not create a provider: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	type Response struct {
		ID int `json:"id"`
	}

	response, err := json.Marshal(Response{ID: id})
	if err != nil {
		log.Println("could not marshal the response: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, string(response))
}

func (h Handler) GetProviderHandler(w http.ResponseWriter, r *http.Request) {
}
