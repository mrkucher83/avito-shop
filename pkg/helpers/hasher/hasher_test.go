package hasher

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

var tests = map[string]string{
	"string":       "securepassword",
	"empty string": "",
	"space":        " ",
	"cyrillic":     "Привет",
	"chinese":      "世界",
	"emoji":        "🔥🔥🔥",
}

func TestHashPassword(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			hash, err := HashPassword(tc)

			require.NoError(t, err, "Password hashing error")
			require.NotEmpty(t, hash, "Password hash must not be empty")
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			hash, err := HashPassword(tc)

			err = CheckPasswordHash(tc, hash)
			require.NoError(t, err, "Password does not match the hash")
		})
	}
}

func TestCheckPasswordHash_Fail(t *testing.T) {
	password := "securepassword"
	wrongPassword := "wrongpassword"

	hash, _ := HashPassword(password)
	err := CheckPasswordHash(wrongPassword, hash)
	require.Error(t, err)
	require.ErrorIs(t, err, bcrypt.ErrMismatchedHashAndPassword)
}
