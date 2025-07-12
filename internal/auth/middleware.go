package auth

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := r.Cookie("auth_token")
		csrfToken, err := r.Cookie("csrf_token")
		csrfTokenHeader := r.Header.Get("X-CSRF-TOKEN")
		type contextKey string

		if err != nil {
			log.Println(err)
			http.Error(w, "cookie not found", http.StatusBadRequest)
			return
		}

		err = validateCSRF(csrfToken.Value, csrfTokenHeader)
		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		user, err := validateJWT(sessionToken.Value)
		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		//const userKey contextKey = string(user.ID)

		log.Println(user.Username)

		next(w, r)
	}
}
