package bookings

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

func (s Service) GetBookingByID(ctx context.Context, bookingID int) (Booking, error) {
	return s.repo.GetByID(ctx, bookingID)
}

func (s Service) GetBookingsByUserID(ctx context.Context, userID int) ([]Booking, error) {
	return s.repo.GetListByUserID(ctx, userID)
}
