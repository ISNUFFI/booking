package providers

import (
	"context"

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

func (r *Repo) CreateProvider(ctx context.Context, name, description string) (int, error) {
	var id int

	err := r.pool.QueryRow(
		ctx,
		"INSERT INTO providers(name, description) VALUES ($1, $2) RETURNING id",
		name, description,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}
