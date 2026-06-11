package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ISNUFFI/booking/internal/errs"
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

func (h Handler) MeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := ctx.Value(model.UserIDKey).(int)
	if !ok {
		log.Println("User ID not found in the context")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	user, err := h.service.Me(ctx, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		if errors.Is(err, errs.ErrUserNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
		} else {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(response))
}
