package auth

import (
	"crypto/rand"
	"encoding/base64"
)

func genRandBytes(len int) ([]byte, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func newToken() (string, error) {
	bytes, err := genRandBytes(TokenLen)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
