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

func (r *Repo) Create(ctx context.Context, name, description string, ownerID int) (int, error) {
	var id int

	err := r.pool.QueryRow(
		ctx,
		"INSERT INTO providers(name, description, owner) VALUES ($1, $2, $3) RETURNING id",
		name, description, ownerID,
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
		"SELECT name, description, owner FROM providers WHERE id = $1",
		id,
	).Scan(&p.Name, &p.Description, &p.Owner)

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
		"SELECT name, description, owner FROM providers",
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

func (r *Repo) Delete(ctx context.Context, id int) error {
	res, err := r.pool.Exec(
		ctx,
		"DELETE FROM providers WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return ErrProviderNotFound
	}

	return nil
}
