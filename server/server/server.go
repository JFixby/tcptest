package server

import (
	"fmt"
	"github.com/jfixby/tcptest/shared"
	"io"
	"log"
	"math/rand"
	"net"
)

const difficulty = 15 // bits of leading zero

func generateChallenge() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	ch := string(b)
	log.Printf("Generated challenge: %s", ch)
	return ch
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func verifyPoW(challenge, nonce string) bool {
	hash := shared.Hash(challenge, nonce)
	bits := shared.ToBitString(hash)

	log.Printf("Verifying PoW — Challenge: %s, Nonce: %s", challenge, nonce)
	log.Printf("SHA256(%s) = %x", challenge+nonce, hash)
	log.Printf("Bits: %s...", bits[:min(len(bits), difficulty+10)])

	result := shared.CheckPoW(challenge, nonce, difficulty)
	log.Printf("PoW valid: %t", result)
	return result
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	log.Printf("New connection from %s", addr)

	challenge := generateChallenge()
	log.Printf("Sending challenge to %s: %s %d", addr, challenge, difficulty)
	fmt.Fprintf(conn, "%s %d\n", challenge, difficulty)

	var nonce string
	_, err := fmt.Fscanf(conn, "%s\n", &nonce)
	if err != nil {
		log.Printf("[%s] Error reading nonce: %v", addr, err)
		return
	}
	log.Printf("[%s] Received nonce: %s", addr, nonce)

	if verifyPoW(challenge, nonce) {
		quote := GetRandomQuote()
		log.Printf("[%s] PoW valid — Sending quote: %s", addr, quote)
		io.WriteString(conn, quote+"\n")
	} else {
		log.Printf("[%s] Invalid PoW — Rejecting connection", addr)
		io.WriteString(conn, "Invalid proof of work.\n")
	}
}
