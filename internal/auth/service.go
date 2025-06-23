package auth

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	authrepo "github.com/nanagoboiler/internal/repository/auth"
	"github.com/nanagoboiler/models"
)

type authService struct {
	Repo   authrepo.UserRepository
	secret string
}

func NewAuthService(repo authrepo.UserRepository) Service {
	return &authService{Repo: repo, secret: os.Getenv("JWT_SECRET")}
}

func (s *authService) RegisterUser(ctx context.Context, req *models.RegisterRequest) (models.Tokens, error) {
	var id = uuid.New().String()
	passwordHash, err := HashPassword([]byte(req.Password))
	if err != nil {
		return models.Tokens{}, err
	}

	user := models.User{
		ID:           id,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: passwordHash,
	}

	err = s.Repo.Create(ctx, &user)
	if err != nil {
		return models.Tokens{}, err
	}
	token, jti, err := s.generateTokens(&user)
	if err != nil {
		return models.Tokens{}, err
	}
	err = s.Repo.AddRefresh(ctx, &token, jti)
	if err != nil {
		return models.Tokens{}, err
	}

	return token, nil
}

func (s *authService) generateTokens(user *models.User) (token models.Tokens, jti string, err error) {
	jti = uuid.NewString()
	now := time.Now()
	access_claims := jwt.MapClaims{
		"userName": user.Username,
		"userId":   user.ID,
		"exp":      now.Add(10 * time.Minute).Unix(),
		"iat":      time.Now().Unix(),
	}
	refresh_claims := jwt.MapClaims{
		"userName": user.Username,
		"userId":   user.ID,
		"exp":      now.Add(30 * 24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
		"jti":      jti,
	}

	authTok := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)

	authToken, err := authTok.SignedString([]byte(s.secret))
	if err != nil {
		return token, "", err
	}
	refreshTok := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)

	refreshToken, err := refreshTok.SignedString([]byte(s.secret))
	if err != nil {
		return token, "", err
	}

	token = models.Tokens{
		Auth_token:    authToken,
		Refresh_token: refreshToken,
	}

	return token, jti, nil

}
