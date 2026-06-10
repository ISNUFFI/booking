package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ISNUFFI/booking/internal/model"
)

type Handler struct {
	service Service
}

func NewHandler(pool *pgxpool.Pool) Handler {
	return Handler{
		service: NewService(NewRepo(pool)),
	}
}

type CreateProviderRequest struct {
	Name        string
	Description string
}

func (h Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateProviderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	userID, ok := ctx.Value(model.UserIDKey).(int)
	if !ok {
		log.Println("User ID not found in the context")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	id, err := h.service.CreateProvider(ctx, req.Name, req.Description, userID)
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

func (h Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	p, err := h.service.GetProvider(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrProviderNotFound) {
			http.Error(w, "provider not found", http.StatusNotFound)
		} else {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(p)
	if err != nil {
		log.Println("could not marshal the response: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(response))
}

func (h Handler) GetListHandler(w http.ResponseWriter, r *http.Request) {
	p, err := h.service.GetProvidersList(r.Context())
	if err != nil {
		log.Println("GetProvidersList error: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(p)
	if err != nil {
		log.Println("could not marshal the response: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(response))
}

func (h Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	userID, ok := ctx.Value(model.UserIDKey).(int)
	if !ok {
		log.Println("User ID not found in the context")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteProvider(ctx, id, userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrProviderOwnerMismatch):
			http.Error(w, "forbidden", http.StatusForbidden)
		case errors.Is(err, ErrProviderNotFound):
			http.Error(w, "provider not found", http.StatusNotFound)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
