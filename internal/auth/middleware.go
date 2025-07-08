package auth

import (
	"log"
	"net/http"
)

func AuthMiddleware(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := r.Cookie("auth_token")
		csrfToken, err := r.Cookie("csrf_token")
		csrfTokenHeader := r.Header.Get("bingus")

		if err != nil {
			log.Println(err)
			http.Error(w, "cookie not found", http.StatusBadRequest)
			return
		}

		err = validateCSRF(csrfToken.Value, csrfTokenHeader)
		if err != nil {
			return
		}

		user, err := validateJWT(sessionToken.Value)
		if err != nil {
			return
		}

		log.Println(user.Username)

		//next(w, r)
	}
}
