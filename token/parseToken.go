package token

import (
	"errors"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/kwandapchumba/changelog/utils"
)

func ParseToken(encrypted string) (*Payload, error) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		return nil, err
	}

	parser := paseto.NewParserWithoutExpiryCheck()

	key, err := paseto.V4SymmetricKeyFromHex(config.Hex)
	if err != nil {
		return nil, err
	}

	token, err := parser.ParseV4Local(key, encrypted, []byte(config.Secret))
	if err != nil {
		return nil, err
	}

	var payload Payload

	if err := token.Get("payload", &payload); err != nil {
		return nil, err
	}

	if time.Now().UTC().After(payload.Expiry) {
		return nil, errors.New("expired token")
	}

	return &payload, nil
}
