package hasher

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword receives password string, return hashed string by sha256 or error
func HashPassword(password string) (string, error) {
	var hash = sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]), nil
}

//CheckPasswordHash receives user password and hashed string, after compare return true or false depend on result
func CheckPasswordHash(password, hashed string) bool {
	hash, _ := hex.DecodeString(hashed)
	var pass = sha256.Sum256([]byte(password))
	return bytes.Equal(hash, pass[:])

}
