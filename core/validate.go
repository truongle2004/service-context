package core

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// ValidateToken verifies a JWT token using the provided ECDSA public key.
//
// Parameters:
//   - tokenString: the JWT token string to validate.
//   - publicKey: the ECDSA public key used to verify the token's signature.
//
// Returns:
//   - *jwt.MapClaims: a pointer to the token claims if the token is valid.
//   - error: an error if the token is invalid, malformed, or the signature is not valid.
//
// Behavior:
//   - Ensures the token uses the expected ECDSA signing method (e.g., ES256).
//   - Parses and validates the token signature and expiration (if set).
//   - Returns the claims only if the token is valid.
//
// Example usage:
//   claims, err := ValidateToken(tokenString, publicKey)
//   if err != nil {
//       log.Println("Invalid token:", err)
//   } else {
//       fmt.Println("Token claims:", claims)
//   }

func ValidateToken(tokenString string, publicKey *ecdsa.PublicKey) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure token uses ES256
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Validate token and return claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
