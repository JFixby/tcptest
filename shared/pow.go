package shared

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// toBitString converts a SHA256 hash ([32]byte) to a binary string.
func ToBitString(hash [32]byte) string {
	var sb strings.Builder
	for _, b := range hash {
		sb.WriteString(FormatByteToBits(b))
	}
	return sb.String()
}

func FormatByteToBits(b byte) string {
	return strings.TrimLeft(fmt.Sprintf("%08b", b), "")
}

// CheckPoW verifies that the hash of (challenge + nonce)
// starts with the required number of leading zero bits.
func CheckPoW(challenge, nonce string, difficulty int) bool {
	hash := Hash(challenge, nonce)
	bits := ToBitString(hash)
	return strings.HasPrefix(bits, strings.Repeat("0", difficulty))
}

func Hash(challenge, nonce string) [32]byte {
	return sha256.Sum256([]byte(challenge + nonce))
}
