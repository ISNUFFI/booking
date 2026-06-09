package users

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ISNUFFI/booking/internal/errs"
)

type Repo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) Repo {
	return Repo{
		pool: pool,
	}
}

func (r *Repo) GetUserByID(ctx context.Context, id int) (User, error) {
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
			return User{}, errs.ErrUserNotFound
		}
		return User{}, err
	}

	return u, nil
}
