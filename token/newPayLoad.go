package token

import "time"

func NewPayload(id, userID, email string, issueAt, expiry time.Time, isAdmin bool) *Payload {
	return &Payload{
		ID:       id,
		UserID:   userID,
		Email:    email,
		IssuedAt: issueAt,
		Expiry:   expiry,
		IsAdmin:  isAdmin,
	}
}
