package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestGeneratePassword(t *testing.T) {
	randomPassword := RandomHashedPassword()

	genPass, err := HashedPassword(randomPassword)
	require.NoError(t, err, "Error generating password")
	require.NotEmpty(t, genPass, "Generated password is empty")

	err = CheckPassword(randomPassword, genPass)
	require.NoError(t, err, "Error checking password")

	wrongPass := RandomHashedPassword()
	err = CheckPassword(wrongPass, genPass)
	require.Error(t, err, bcrypt.ErrMismatchedHashAndPassword, "Error checking password")

	err = CheckPassword("", genPass)
	require.Error(t, err, "Error checking password")

	err = CheckPassword(randomPassword, "")
	require.Error(t, err, "Error checking password")

	// check if two hashed passwords are different
	genPass2, err := HashedPassword(randomPassword)
	require.NoError(t, err, "Error generating password")
	require.NotEmpty(t, genPass2, "Generated password is empty")
	require.NotEqual(t, genPass, genPass2, "Generated passwords are the same")
}
