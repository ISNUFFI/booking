package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Register(ctx context.Context, email, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	if err = s.repo.CreateUser(ctx, email, string(bytes)); err != nil {
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return ErrEmailAlreadyExists
		default:
			return err
		}
	}

	return nil
}

func (s Service) Login(password string) (string, error) {
	return "", nil
}
