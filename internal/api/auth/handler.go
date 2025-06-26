package authapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nanagoboiler/internal/auth"
	"github.com/nanagoboiler/models"
)

func Register(s auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RegisterRequest
		csrf, err := uuid.NewRandom()
		if err != nil {
			return
		}
		now := time.Now()

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid Request Json", http.StatusBadRequest)
			return
		}

		err = validateRegistration(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tokens, err := s.RegisterUser(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// i should maybe rewrite this into a helper function

		//Auth Cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    tokens.Auth_token,
			Expires:  now.Add(10 * time.Minute),
			SameSite: http.SameSiteNoneMode,
			HttpOnly: true,
			Secure:   true,
		},
		)
		//Refresh Cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    tokens.Refresh_token,
			Expires:  now.Add(24 * 30 * time.Hour),
			SameSite: http.SameSiteNoneMode,
			HttpOnly: true,
			Secure:   true,
		},
		)
		//CSRF Cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    csrf.String(),
			Expires:  now.Add(10 * time.Minute),
			SameSite: http.SameSiteNoneMode,
			HttpOnly: false,
			Secure:   true,
		})

		w.WriteHeader(http.StatusCreated)

	}

}

func Login(s auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.LoginRequest
		csrf, err := uuid.NewRandom()
		if err != nil {
			http.Error(w, "Failed to generate CSRF token", http.StatusInternalServerError)
			return
		}
		now := time.Now()

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid Request JSON", http.StatusBadRequest)
			return
		}

		tokens, err := s.LoginUser(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Auth Cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    tokens.Auth_token,
			Expires:  now.Add(10 * time.Minute),
			SameSite: http.SameSiteNoneMode,
			HttpOnly: true,
			Secure:   true,
		},
		)
		//Refresh Cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    tokens.Refresh_token,
			Expires:  now.Add(24 * 30 * time.Hour),
			SameSite: http.SameSiteNoneMode,
			HttpOnly: true,
			Secure:   true,
		},
		)
		//CSRF Cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    csrf.String(),
			Expires:  now.Add(10 * time.Minute),
			SameSite: http.SameSiteNoneMode,
			HttpOnly: false,
			Secure:   true,
		})

		w.WriteHeader(http.StatusCreated)

	}

}
