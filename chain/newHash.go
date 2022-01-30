package chain

import (
	"crypto/sha256"
	"encoding/hex"
)

func NewHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
