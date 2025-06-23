package authrepo

import (
	"context"

	"github.com/nanagoboiler/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Check(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, user *models.User) error
	AddRefresh(ctx context.Context, token *models.Tokens, uuid string) error
}
