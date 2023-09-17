package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"github.com/scortier/gopherbank/util"
)

// PasetoMaker is a struct that contains the key to sign the token
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmKey string) (Maker, error) {

	// Check if the key is valid, it must be 32 characters
	if (len(symmKey)) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be %d characters", chacha20poly1305.KeySize)
	}

	// Create a new PasetoMaker, which contains the key to sign the token
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmKey),
	}

	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(email string, duration time.Duration) (string, error) {

	// Create a new payload
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}

	// Create a new token
	// payload is the data we want to store in the token
	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	if err != nil {
		return "", err
	}

	// Return the token
	return token, nil
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {

	// Create a new emty payload object to store the decrypt data.
	payload := &Payload{}

	// Verify the token
	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, util.ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	// Return the payload
	return payload, nil
}
