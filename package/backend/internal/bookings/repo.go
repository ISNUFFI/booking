package bookings

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) Repo {
	return Repo{
		pool: pool,
	}
}

func (r *Repo) GetByID(ctx context.Context, bookingID int) (Booking, error) {
	var b Booking
	err := r.pool.QueryRow(
		ctx,
		"SELECT id, user_id, slot_id, created_at FROM bookings WHERE id = $1",
		bookingID,
	).Scan(&b.ID, &b.UserID, &b.SlotID, &b.CreatedAt)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return Booking{}, ErrBookingNotFound
		default:
			return Booking{}, err
		}
	}

	return b, nil
}

func (r *Repo) GetListByUserID(ctx context.Context, userID int) ([]Booking, error) {
	rows, err := r.pool.Query(
		ctx,
		"SELECT id, slot_id, created_at FROM bookings WHERE user_id = $1",
		userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bookings, err := pgx.CollectRows(rows, pgx.RowToStructByName[Booking])
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *Repo) Create(ctx context.Context, slotID, userID int) (int, error) {
	var id int

	err := r.pool.QueryRow(
		ctx,
		"INSERT INTO bookings(slot_id, user_id) VALUES ($1, $2) RETURNING id",
		slotID, userID,
	).Scan(&id)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return 0, ErrSlotAlreadyTaken
			case pgerrcode.ForeignKeyViolation:
				return 0, ErrSlotDoesNotExist
			}
		}

		return 0, err
	}

	return id, nil
}

func (r *Repo) Delete(ctx context.Context, bookingID int) error {
	res, err := r.pool.Exec(
		ctx,
		"DELETE FROM bookings WHERE id = $1",
		bookingID,
	)

	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return ErrBookingNotFound
	}

	return nil
}
