package main

import (
	"net/http"

	authapi "github.com/nanagoboiler/internal/api/auth"
	"github.com/nanagoboiler/internal/auth"
	"github.com/nanagoboiler/internal/bootstrap"
	authrepo "github.com/nanagoboiler/internal/repository/auth"

	"context"
)

func main() {
	router := http.NewServeMux()
	ctx := context.Background()
	pool, err := bootstrap.NewPostgresPool(ctx)
	if err != nil {
		panic(err)
	}
	authRepo := authrepo.NewUserRepository(pool)
	tokenRepo := authrepo.NewTokensRepository(pool)
	authService := auth.NewAuthService(authRepo, tokenRepo)
	authRegister := authapi.Register(authService)
	authLogin := authapi.Login(authService)
	bingus := authapi.Bingus()

	router.HandleFunc("POST /register/", authRegister)
	router.HandleFunc("POST /login/", authLogin)
	router.HandleFunc("POST /test/", auth.AuthMiddleware(bingus))

	println("Server Listening on Port 8085")
	http.ListenAndServe(":8085", router)
}
