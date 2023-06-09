package token

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/kwandapchumba/changelog/utils"
)

func CreateToken(userID, email, companyID string, issuedAt time.Time, duration time.Duration, isAdmin bool) (string, *Payload, error) {
	id := utils.RandomString()

	expiry := time.Now().UTC().Add(duration)

	payload := NewPayload(id, userID, email, companyID, issuedAt, expiry, isAdmin)

	token := paseto.NewToken()

	if err := token.Set("payload", payload); err != nil {
		return "", nil, err
	}

	config, err := utils.LoadConfig(".")
	if err != nil {
		return "", nil, err
	}

	key, err := paseto.V4SymmetricKeyFromHex(config.Hex)
	if err != nil {
		return "", nil, err
	}

	encrypted := token.V4Encrypt(key, []byte(config.Secret))

	return encrypted, payload, nil
}
