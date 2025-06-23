package auth

import (
	"encoding/json"
	"net/http"

	"github.com/nanagoboiler/models"
)

func Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid Request Json", http.StatusBadRequest)
		}

		err = validateRegistration(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}

}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}
