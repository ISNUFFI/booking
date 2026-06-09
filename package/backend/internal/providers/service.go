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
	return s.repo.Create(ctx, name, description)
}

func (s Service) GetProvider(ctx context.Context, id int) (Provider, error) {
	return s.repo.Get(ctx, id)
}

func (s Service) GetProvidersList(ctx context.Context) ([]Provider, error) {
	return s.repo.GetList(ctx)
}
