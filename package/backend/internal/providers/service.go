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

func (s Service) CreateProvider(ctx context.Context, name, description string, userID int) (int, error) {
	return s.repo.Create(ctx, name, description, userID)
}

func (s Service) GetProvider(ctx context.Context, id int) (Provider, error) {
	return s.repo.Get(ctx, id)
}

func (s Service) GetProvidersList(ctx context.Context) ([]Provider, error) {
	return s.repo.GetList(ctx)
}

func (s Service) DeleteProvider(ctx context.Context, id, userID int) error {
	provider, err := s.repo.Get(ctx, id)
	if err != nil {
		return err
	}

	if provider.Owner != userID {
		return ErrProviderOwnerMismatch
	}

	return s.repo.Delete(ctx, id)
}
