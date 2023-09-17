package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/scortier/gopherbank/util"
	"github.com/stretchr/testify/require"
)

// TestPasetoMaker tests the NewPasetoMaker function
// It is happy path testing
func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	// require.NotEmpty(t, payload)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

// TestExpiredJWTToken tests the NewPasetoMaker function
// It is error path testing
func TestExpiredPasetoToken(t *testing.T) {
	// Create a new NewPasetoMaker
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Create a new token
	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	// require.NotEmpty(t, payload)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, util.ErrExpiredToken.Error())
	require.Nil(t, payload)
}

// TestInvalidJWTTokenAlgNone tests the NewPasetoMaker function
func TestInvalidPasetoTokenAlgNone(t *testing.T) {
	// Create a new payload
	payload, err := NewPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	// Create a new token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)

	// Sign the token with the secret key, this string is only used for testing
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	// Create a new NewPasetoMaker
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Verify the token
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)

	// Check if the error is a validation error
	require.EqualError(t, err, util.ErrInvalidToken.Error())
	require.Nil(t, payload)
}
