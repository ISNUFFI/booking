package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/ISNUFFI/booking/internal/config"
	"github.com/ISNUFFI/booking/internal/errs"
)

type Service struct {
	repo   Repo
	config *config.Config
}

func NewService(repo Repo, config *config.Config) Service {
	return Service{
		repo:   repo,
		config: config,
	}
}

func (s Service) Register(ctx context.Context, email, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	if err = s.repo.CreateUser(ctx, email, string(bytes)); err != nil {
		switch {
		case errors.Is(err, errs.ErrDuplicateKey):
			return errs.ErrEmailAlreadyExists
		default:
			return err
		}
	}

	return nil
}

func (s Service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	hash := user.passwordHash
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return "", errs.ErrInvalidPassword
	}

	claims := JWTClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.config.JWTSecret))
}
