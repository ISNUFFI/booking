package providers

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

func (r *Repo) Create(ctx context.Context, name, description string) (int, error) {
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

func (r *Repo) Get(ctx context.Context, id int) (Provider, error) {
	var p Provider
	err := r.pool.QueryRow(
		ctx,
		"SELECT name, description FROM providers WHERE id = $1",
		id,
	).Scan(&p.Name, &p.Description)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Provider{}, ErrProviderNotFound
		}
		return Provider{}, err
	}

	return p, nil
}

func (r *Repo) GetList(ctx context.Context) ([]Provider, error) {
	rows, err := r.pool.Query(
		ctx,
		"SELECT name, description FROM providers",
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	providers, err := pgx.CollectRows(rows, pgx.RowToStructByName[Provider])
	if err != nil {
		return nil, err
	}

	return providers, nil
}
