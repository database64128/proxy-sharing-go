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
