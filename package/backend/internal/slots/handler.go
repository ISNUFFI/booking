package slots

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func (h Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	slot, err := h.service.GetSlot(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrSlotNotFound) {
			http.Error(w, "slot not found", http.StatusNotFound)
		} else {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(slot)
	if err != nil {
		log.Println("could not marshal the response: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(response))
}

func (h Handler) GetProviderSlotsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	providerID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid provider id", http.StatusBadRequest)
		return
	}

	slot, err := h.service.GetSlotListByProvider(r.Context(), providerID)
	if err != nil {
		if errors.Is(err, ErrSlotNotFound) {
			http.Error(w, "slot not found", http.StatusNotFound)
		} else {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(slot)
	if err != nil {
		log.Println("could not marshal the response: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(response))
}
