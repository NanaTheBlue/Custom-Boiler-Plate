package auth

import (
	"log"
	"net/http"
)

func AuthMiddleware(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := r.Cookie("auth_token")
		if sessionToken != nil {
			log.Println(err)
			http.Error(w, "cookie not found", http.StatusBadRequest)
			return
		}

		//next(w, r)
	}
}
