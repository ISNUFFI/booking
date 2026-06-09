package users

import (
	"context"
)

type Service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Me(ctx context.Context, userID int) (User, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
