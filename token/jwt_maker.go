package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/scortier/gopherbank/util"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

// NewJWTMaker returns a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(email string, duration time.Duration) (string, error) {
	// Create a new payload
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}

	// Create a new token
	// payload is the data we want to store in the token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Sign the token with the secret key
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	// Create a key function, which will return the secret key
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return []byte(maker.secretKey), nil
	}

	// Parse the token
	// payload is the data we want to store in the token
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		// Check if the error is a validation error
		validationErr, ok := err.(*jwt.ValidationError)

		// Check if the error is an expired token error
		if ok && errors.Is(validationErr.Inner, util.ErrExpiredToken) {
			// Token is expired
			return nil, util.ErrExpiredToken
		}

		return nil, util.ErrInvalidToken
	}

	// Check if the token is valid, and get the payload
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, fmt.Errorf("invalid token claims: %v", jwtToken.Claims)
	}

	// Check if the token is valid
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil

}
