package server

import (
	"fmt"
	"github.com/jfixby/tcptest/shared"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

func generateChallenge() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	challenge := string(b)
	log.Printf("Generated challenge: %s", challenge)
	return challenge
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func verifyPoW(challenge, nonce string, difficulty int) bool {
	log.Printf("Verifying PoW — Challenge: %s, Nonce: %s", challenge, nonce)
	result, hash, bits := shared.CheckPoW(challenge, nonce, difficulty)
	log.Printf("SHA256(%s) = %x", challenge+nonce, hash)
	log.Printf("Bits: %s...", bits[:min(len(bits), difficulty+10)])
	log.Printf("PoW valid: %t", result)
	return result
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	log.Printf("New connection from %s", addr)

	start := time.Now()

	localDifficulty := GetDifficulty()

	challenge := generateChallenge()
	log.Printf("Sending challenge to %s: %s %d", addr, challenge, localDifficulty)
	fmt.Fprintf(conn, "%s %d\n", challenge, localDifficulty)

	var nonce string
	_, err := fmt.Fscanf(conn, "%s\n", &nonce)
	if err != nil {
		log.Printf("[%s] Error reading nonce: %v", addr, err)
		return
	}
	log.Printf("[%s] Received nonce: %s", addr, nonce)

	duration := time.Since(start)

	if verifyPoW(challenge, nonce, localDifficulty) {
		quote := GetRandomQuote()
		log.Printf("[%s] PoW valid — Sending quote: %s", addr, quote)
		io.WriteString(conn, quote+"\n")
	} else {
		log.Printf("[%s] Invalid PoW — Rejecting connection", addr)
		io.WriteString(conn, "Invalid proof of work.\n")
	}

	AdjustDifficulty(duration)
}

func Start(address, wisdoms string) {
	rand.Seed(time.Now().UnixNano())

	err := LoadQuotes(wisdoms)
	if err != nil {
		log.Fatalf("Failed to load quotes: %v", err)
	}

	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server listening on port 1337")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go HandleConnection(conn)
	}
}
