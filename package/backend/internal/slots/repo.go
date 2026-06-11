package slots

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
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

func (r *Repo) Get(ctx context.Context, id int) (Slot, error) {
	var s Slot
	err := r.pool.QueryRow(
		ctx,
		"SELECT id, provider_id, start_time, end_time, is_active FROM slots WHERE id = $1",
		id,
	).Scan(&s.ID, &s.ProviderID, &s.StartTime, &s.EndTime, &s.IsActive)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return Slot{}, ErrSlotNotFound
		default:
			return Slot{}, err
		}
	}

	return s, nil
}

func (r *Repo) GetListByProvider(ctx context.Context, providerID int) ([]Slot, error) {
	rows, err := r.pool.Query(
		ctx,
		"SELECT id, provider_id, start_time, end_time, is_active FROM slots WHERE provider_id = $1",
		providerID,
	)
	if err != nil {
		return nil, err
	}

	slots, err := pgx.CollectRows(rows, pgx.RowToStructByName[Slot])
	if err != nil {
		return nil, err
	}

	return slots, nil
}
