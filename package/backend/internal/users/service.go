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
	return s.repo.GetUserByID(ctx, userID)
}
