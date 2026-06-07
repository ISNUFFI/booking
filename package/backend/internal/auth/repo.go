package auth

import (
	"context"
	"errors"

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

func (r *Repo) CreateUser(ctx context.Context, email, hash string) error {
	_, err := r.pool.Exec(
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

func (r *Repo) GetUserByID(ctx context.Context, id int) (*User, error) {
	var u User

	err := r.pool.QueryRow(
		ctx,
		"SELECT id, email, role, password_hash FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Role,
		&u.passwordHash,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var u User

	err := r.pool.QueryRow(
		ctx,
		"SELECT id, email, role, password_hash FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Role,
		&u.passwordHash,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
