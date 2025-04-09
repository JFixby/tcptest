package client

import (
	"bufio"
	"fmt"
	"github.com/jfixby/tcptest/shared"
	"log"
	"net"
	"strings"
	"time"
)

func SolvePoW(challenge string, difficulty int) string {
	log.Printf("Starting PoW solver — challenge: %s, difficulty: %d", challenge, difficulty)

	for i := 0; ; i++ {
		nonce := fmt.Sprintf("%d", i)
		valid, hash, bits := shared.CheckPoW(challenge, nonce, difficulty)
		if valid {
			log.Printf("PoW solved! Nonce: %s", nonce)
			log.Printf("Hash: %x", hash)
			log.Printf("Bits: %s...", bits[:min(len(bits), difficulty+10)])
			return nonce
		}

		if i%10000 == 0 {
			log.Printf("Hash: %x", hash)
		}
	}
}

func ConnectToServer(address string) net.Conn {
	log.Printf("Connecting to server at %s...", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	log.Println("Connected.")
	return conn
}

func ReadChallenge(reader *bufio.Reader) (string, int) {
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read challenge from server: %v", err)
	}
	log.Printf("Received challenge line: %s", strings.TrimSpace(line))

	parts := strings.Fields(line)
	if len(parts) != 2 {
		log.Fatalf("Invalid challenge format: %v", parts)
	}

	challenge := parts[0]
	var difficulty int
	fmt.Sscanf(parts[1], "%d", &difficulty)

	log.Printf("Parsed challenge: %s", challenge)
	log.Printf("Parsed difficulty: %d", difficulty)

	return challenge, difficulty
}

func SolveChallenge(challenge string, difficulty int) string {
	log.Println("Solving PoW challenge...")
	nonce := SolvePoW(challenge, difficulty)
	log.Printf("Solved PoW! Nonce: %s", nonce)
	return nonce
}

func SendNonce(conn net.Conn, nonce string) {
	fmt.Fprintf(conn, "%s\n", nonce)
	log.Println("Nonce sent to server.")
}

func ReadReply(reader *bufio.Reader) {
	reply, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read reply from server: %v", err)
	}
	log.Printf("Server reply: %s", strings.TrimSpace(reply))
	fmt.Println("Server says:", strings.TrimSpace(reply))
}

func Exchange(address string) {
	start := time.Now()

	conn := ConnectToServer(address)
	defer conn.Close()

	reader := bufio.NewReader(conn)

	challenge, difficulty := ReadChallenge(reader)
	nonce := SolveChallenge(challenge, difficulty)
	SendNonce(conn, nonce)
	ReadReply(reader)

	elapsed := time.Since(start)
	log.Printf("Exchange completed in %s\n", elapsed)
	fmt.Println()
}
