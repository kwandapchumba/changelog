package token

import "time"

type Payload struct {
	ID       string    `json:"id"`
	UserID   string    `json:"account_id"`
	Email    string    `json:"email"`
	IssuedAt time.Time `json:"issued_at"`
	Expiry   time.Time `json:"expiry"`
	IsAdmin  bool      `json:"is_admin"`
}
