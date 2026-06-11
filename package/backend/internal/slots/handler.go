package slots

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
			log.Println("could not get slot info: ", err)
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
			log.Println("could not provider slots: ", err)
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

type CreateBulkRequest struct {
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	DurationMinutes uint      `json:"duration_minutes"`
}

func (req CreateBulkRequest) Validate() error {
	if req.Start.After(req.End) {
		return ErrStartAfterEnd
	}

	if req.DurationMinutes <= 0 {
		return ErrDurationNotPositive
	}

	return nil
}

func (h Handler) CreateBulkHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	providerID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req CreateBulkRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("could not parse json body: ", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	duration := time.Duration(req.DurationMinutes) * time.Minute
	err = h.service.CreateBulk(r.Context(), providerID, req.Start, req.End, duration)
	if err != nil {
		log.Println("could not bulk create slots: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
