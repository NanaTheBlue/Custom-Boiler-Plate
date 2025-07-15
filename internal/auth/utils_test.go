package auth

import (
	"testing"

	"github.com/google/uuid"

	"github.com/nanagoboiler/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	type testCase struct {
		//input params
		password []byte

		//expected values
		err error
	}
	t.Run("Valid PasswordHash", func(t *testing.T) {
		tests := []testCase{
			{password: []byte("BingusBongus123")},
		}

		for _, test := range tests {
			actual, err := HashPassword(test.password)
			assert.NoError(t, err)
			assert.NotEmpty(t, actual)

			err = bcrypt.CompareHashAndPassword([]byte(actual), test.password)
			assert.NoError(t, err, "hashed password should validate against original")
		}
	})
}

func TestValidateJWT(t *testing.T) {
	type testCase struct {
		//input params
		user models.User

		//expected values
		err error
	}
	tests := []testCase{
		{
			user: models.User{
				ID:       uuid.New().String(),
				Username: "Bingus",
				Email:    "Bingus@bongus.com",
			},
		},
	}
	for _, test := range tests {
		// need to mock
		token, jti, err := s.generateTokens(&test.user)
		assert.NoError(t, err)
		assert.NotEmpty(t, token, jti)
		user, err := validateJWT(token.Auth_token)
		assert.NoError(t, err)

		assert.Equal(t, test.user.Username, user.Username, "Token Username doesn't match")

	}

}
