package slots

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

func (s Service) GetSlot(ctx context.Context, id int) (Slot, error) {
	return s.repo.Get(ctx, id)
}

func (s Service) GetSlotListByProvider(ctx context.Context, providerID int) ([]Slot, error) {
	return s.repo.GetListByProvider(ctx, providerID)
}
