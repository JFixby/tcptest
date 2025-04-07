package main

import (
	"bufio"
	"fmt"
	"github.com/jfixby/tcptest/client/client"
	"log"
	"net"
	"strings"
)

func main() {
	log.Println("Connecting to server at localhost:1337...")
	conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	log.Println("Connected.")

	reader := bufio.NewReader(conn)

	// Step 1: Read challenge
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read challenge from server: %v", err)
	}
	log.Printf("Received challenge line: %s", strings.TrimSpace(line))

	parts := strings.Fields(line)
	if len(parts) != 2 {
		log.Fatalf("Invalid challenge format: %v", parts)
	}

	challenge, difficulty := parts[0], 0
	fmt.Sscanf(parts[1], "%d", &difficulty)
	log.Printf("Parsed challenge: %s", challenge)
	log.Printf("Parsed difficulty: %d", difficulty)

	// Step 2: Solve PoW
	log.Println("Solving PoW challenge...")
	nonce := client.SolvePoW(challenge, difficulty)
	log.Printf("Solved PoW! Nonce: %s", nonce)

	// Step 3: Send nonce
	fmt.Fprintf(conn, "%s\n", nonce)
	log.Println("Nonce sent to server.")

	// Step 4: Read server reply
	reply, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read reply from server: %v", err)
	}
	log.Printf("Server reply: %s", strings.TrimSpace(reply))
	fmt.Println("Server says:", strings.TrimSpace(reply))
}
