package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() string {
	slice0 := make([]byte, 32)
	rand.Read(slice0)
	return hex.EncodeToString(slice0)
}
