package auth

import (
	"testing"

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
