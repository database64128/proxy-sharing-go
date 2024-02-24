// Package tokenhelper provides helper functions for working with tokens.
// Tokens are assumed to be 32 bytes long.
package tokenhelper

import "crypto/rand"

// NewTokenBytes returns a new token as a byte slice.
func NewTokenBytes() ([]byte, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	return b, err
}

// NewAccessTokenAndRefreshTokenBytes returns a new access token and refresh token as byte slices.
func NewAccessTokenAndRefreshTokenBytes() (accessToken, refreshToken []byte, err error) {
	b := make([]byte, 64)
	_, err = rand.Read(b)
	return b[:32], b[32:], err
}
