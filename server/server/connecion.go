package server

import (
	"bufio"
	"fmt"
	"github.com/jfixby/tcptest/shared"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	log.Printf("New connection from %s", addr)

	start := time.Now()
	localDifficulty := GetDifficulty()

	challenge := generateChallenge()
	log.Printf("Sending challenge to %s: %s %d", addr, challenge, localDifficulty)
	fmt.Fprintf(conn, "%s %d\n", challenge, localDifficulty)

	reader := bufio.NewReader(conn)
	nonceLine, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("[%s] Error reading nonce: %v", addr, err)
		return
	}
	nonce := strings.TrimSpace(nonceLine)
	log.Printf("[%s] Received nonce: %s", addr, nonce)

	duration := time.Since(start)

	result, hash, bits := shared.CheckPoW(challenge, nonce, localDifficulty)
	log.Printf("Verifying PoW — Challenge: %s, Nonce: %s", challenge, nonce)
	log.Printf("SHA256(%s) = %x", challenge+nonce, hash)
	log.Printf("Bits: %s...", bits[:min(len(bits), localDifficulty+10)])
	log.Printf("PoW valid: %t", result)

	if result {
		quote := GetRandomQuote()
		log.Printf("[%s] PoW valid — Sending quote: %s", addr, quote)
		io.WriteString(conn, quote+"\n")
	} else {
		log.Printf("[%s] Invalid PoW — Rejecting connection", addr)
		io.WriteString(conn, "Invalid proof of work.\n")
	}

	AdjustDifficulty(duration)
}

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
