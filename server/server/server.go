package server

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
)

var quotes []string

const difficulty = 5 // bits of leading zero

func LoadQuotes(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(bufio.NewReader(file)).Decode(&quotes)
}

func generateChallenge() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func verifyPoW(challenge, nonce string) bool {
	hash := sha256.Sum256([]byte(challenge + nonce))
	bits := fmt.Sprintf("%08b", hash)
	return strings.HasPrefix(bits, strings.Repeat("0", difficulty))
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	challenge := generateChallenge()
	fmt.Fprintf(conn, "%s %d\n", challenge, difficulty)

	var nonce string
	_, err := fmt.Fscanf(conn, "%s\n", &nonce)
	if err != nil {
		log.Println("Error reading nonce:", err)
		return
	}

	if verifyPoW(challenge, nonce) {
		quote := quotes[rand.Intn(len(quotes))]
		io.WriteString(conn, quote+"\n")
	} else {
		io.WriteString(conn, "Invalid proof of work.\n")
	}
}
