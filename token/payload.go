package token

import (
	"time"

	"github.com/google/uuid"
	"github.com/scortier/gopherbank/util"
)

type Payload struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        int64(tokenID.ID()),
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// Valid checks if the token is valid or not
func (payload *Payload) Valid() error {
	// Check if the token is expired
	if time.Now().After(payload.ExpiredAt) {
		return util.ErrExpiredToken
	}

	return nil
}
