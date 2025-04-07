package client

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func SolvePoW(challenge string, difficulty int) string {
	prefix := strings.Repeat("0", difficulty)
	for i := 0; ; i++ {
		nonce := fmt.Sprintf("%d", i)
		hash := sha256.Sum256([]byte(challenge + nonce))
		bits := fmt.Sprintf("%08b", hash)
		if strings.HasPrefix(bits, prefix) {
			return nonce
		}
	}
}
