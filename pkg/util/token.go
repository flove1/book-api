package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
)

type Token struct {
	Plaintext string `json:"token"`
	Hash      []byte `json:"-"`
}

func GenerateToken() (*Token, error) {
	token := new(Token)

	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]
	return token, nil
}
