package abci

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

// hashStringWithNonce hashes a string and returns the hash value.
func hashString(data string) string {
	// Concatenate data with nonce.
	input := fmt.Sprintf("%s%d", data)

	// Compute the SHA256 hash.
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashed := hasher.Sum(nil)

	return hex.EncodeToString(hashed)
}
