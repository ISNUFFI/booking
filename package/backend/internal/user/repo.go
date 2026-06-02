package user

import (
	"context"
	"errors"

	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) Repo {
	return Repo {
		pool: pool,
	}
}

func (r *Repo) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return conn.Exec(ctx, query, args...)
}

func (r *Repo) CreateUser(ctx context.Context, email, hash string) error {
	_, err := r.Exec(
		ctx,
		"INSERT INTO users(email, password_hash) VALUES ($1, $2)",
		email, hash,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return ErrDuplicateKey
		}
	}

	return err
}
