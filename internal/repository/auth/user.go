package authrepo

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nanagoboiler/models"
)

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepo{pool: pool}
}

func (r *userRepo) Create(ctx context.Context, user *models.User) error {
	_, err := r.pool.Exec(ctx, "INSERT into users (email,username,hashed_password) VALUES ($1, $2, $3);", user.Email, user.Username, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Delete(ctx context.Context, user *models.User) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM users WHERE id = $1", user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Check(ctx context.Context, user *models.User) error {
	var id string
	err := r.pool.QueryRow(ctx, "SELECT Id from users WHERE Id = $1", user.ID).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) AddRefresh(ctx context.Context, token *models.Tokens, uuid string) error {
	refreshExpiry := time.Now().UTC().Add(30 * 24 * time.Hour)
	_, err := r.pool.Exec(ctx, "UPDATE users SET refresh_token = $1, expires_at = $2 WHERE id = $3;", token.Refresh_token, refreshExpiry, uuid)
	if err != nil {
		return err
	}

	return nil

}
