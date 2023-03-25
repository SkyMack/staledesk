package controller

import (
	"crypto/rand"
	"encoding/hex"
)

type ErrorResp struct {
	Description string         `json:"description"`
	Errors      []ErrorDetails `json:"errors"`
}

type ErrorDetails struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// randHexString generates a random hex value of length `n`
func randHexString(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
