package bookings

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

func (h Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	bookingID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	booking, err := h.service.GetBookingByID(r.Context(), bookingID)
	if err != nil {
		if errors.Is(err, ErrBookingNotFound) {
			http.Error(w, "booking not found", http.StatusNotFound)
		} else {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(booking)
	if err != nil {
		log.Println("could not marshal the response: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(response))
}

func (h Handler) GetMeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := ctx.Value(model.UserIDKey).(int)
	if !ok {
		log.Println("User ID not found in the context")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	bookings, err := h.service.GetBookingsByUserID(ctx, userID)
	if err != nil {
		log.Println("GetBookingsByUserID error: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(bookings)
	if err != nil {
		log.Println("could not marshal the response: ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(response))
}
