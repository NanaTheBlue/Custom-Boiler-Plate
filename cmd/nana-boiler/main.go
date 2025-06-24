package main

import (
	"net/http"

	"github.com/nanagoboiler/internal/auth"
	authrepo "github.com/nanagoboiler/internal/repository/auth"
	db "github.com/nanagoboiler/internal/repository/postgres"

	"context"
)

func main() {
	router := http.NewServeMux()
	ctx := context.Background()
	pool, err := db.NewPostgresPool(ctx)
	if err != nil {
		panic(err)
	}
	authRepo := authrepo.NewUserRepository(pool)
	tokenRepo := authrepo.NewTokensRepository(pool)
	authService := auth.NewAuthService(authRepo, tokenRepo)
	authRegister := auth.Register(authService)

	router.HandleFunc("POST /register/", authRegister)

	println("Server Listening on Port 8085")
	http.ListenAndServe(":8085", router)
}
