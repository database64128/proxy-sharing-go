// Package tokens provides utilities for working with tokens.
// Tokens are assumed to be 32 bytes long.
package tokens

import "crypto/rand"

// NewTokenBytes returns a new token as a byte slice.
func NewTokenBytes() []byte {
	b := make([]byte, 32)
	rand.Read(b)
	return b
}

// NewAccessTokenAndRefreshTokenBytes returns a new access token and refresh token as byte slices.
func NewAccessTokenAndRefreshTokenBytes() (accessToken, refreshToken []byte) {
	b := make([]byte, 64)
	rand.Read(b)
	return b[:32], b[32:]
}
