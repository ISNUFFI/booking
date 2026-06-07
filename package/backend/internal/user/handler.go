package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ISNUFFI/booking/internal/config"
)

type Handler struct {
	service Service
	config *config.Config
}

func NewHandler(pool *pgxpool.Pool, config *config.Config) Handler {
	return Handler{
		service: NewService(NewRepo(pool), config),
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterValidationError struct {
	message string
}

func (e RegisterValidationError) Error() string {
	return e.message
}

func (req RegisterRequest) Validate() error {
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return RegisterValidationError{
			message: "invalid email",
		}
	}

	if len(req.Password) < 10 {
		return RegisterValidationError{
			message: "password is too short",
		}
	}

	return nil
}

func (h Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		log.Println("could not register: ", err)
		switch {
		case errors.Is(err, ErrEmailAlreadyExists):
			http.Error(w, "email already exists", http.StatusConflict)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "success")
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	jwt, err := h.service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		log.Println("could not login: ", err)
		http.Error(w, "invalid login", http.StatusUnauthorized)
		return
	}

	type Response struct {
		Token string `json:"token"`
	}

	resp := Response{Token: jwt}

	response, err := json.Marshal(resp)
	if err != nil {
		log.Println("")
		http.Error(w, "internal error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, string(response))
}

func (h Handler) AttachHandlers(router chi.Router) {
	router.Post("/register", h.RegisterHandler)
	router.Post("/login", h.LoginHandler)
}
