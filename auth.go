package auth

import "encoding/base64"

func NewSessionID() (string, error) {
	bytes, err := generateRandomBytes(32)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
