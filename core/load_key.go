package core

import (
	"crypto/ecdsa"
	"fmt"
	"io"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// LoadPublicKey reads an ECDSA public key from a PEM-encoded file.
//
// Parameters:
//   - PUBLIC_KEY_PATH: the file path to the PEM-encoded ECDSA public key.
//
// Purpose:
//
//	This function is typically used to load the public key required to verify JWT tokens
//	that were signed using the corresponding ECDSA private key.
//
// Returns:
//   - *ecdsa.PublicKey: the successfully parsed public key.
//   - error: an error if the file could not be opened, read, or the key could not be parsed.
//
// Example usage:
//
//	publicKey, err := LoadPublicKey("/path/to/public.pem")
//	if err != nil {
//	    log.Fatal(err)
//	}
func LoadPublicKey(PUBLIC_KEY_PATH string) (*ecdsa.PublicKey, error) {
	file, err := os.Open(PUBLIC_KEY_PATH)
	if err != nil {
		return nil, fmt.Errorf("unable to open public key file %s: %w", PUBLIC_KEY_PATH, err)
	}
	defer file.Close()

	keyData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read public key file %s: %w", PUBLIC_KEY_PATH, err)
	}

	publicKey, err := jwt.ParseECPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA public key from %s: %w", PUBLIC_KEY_PATH, err)
	}

	return publicKey, nil
}
