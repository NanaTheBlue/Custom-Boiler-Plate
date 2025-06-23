package auth

import (
	"context"

	"github.com/nanagoboiler/models"
)

type Service interface {
	RegisterUser(ctx context.Context, req *models.RegisterRequest) (models.Tokens, error)
}
