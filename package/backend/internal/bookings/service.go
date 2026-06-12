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

func (s Service) CreateBooking(ctx context.Context, slotID, userID int) (int, error) {
	return s.repo.Create(ctx, slotID, userID)
}

func (s Service) DeleteBooking(ctx context.Context, bookingID, userID int) error {
	booking, err := s.repo.GetByID(ctx, bookingID)
	if err != nil {
		return err
	}

	if booking.UserID != userID {
		return ErrForbidden
	}
	return s.repo.Delete(ctx, bookingID)
}
