package providers

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

func (s Service) CreateProvider(ctx context.Context, name, description string) (int, error) {
	return s.repo.CreateProvider(ctx, name, description)
}
