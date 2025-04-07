package main

import (
	"bufio"
	"fmt"
	"github.com/jfixby/tcptest/client/client"
	"github.com/jfixby/tcptest/server/server"
	"net"
	"strings"
	"testing"
	"time"
)

func TestExchangeWithServer(t *testing.T) {
	test_server := ":1337"

	server.Start(test_server, "./server/wisdoms.json")

	const attempts = 10

	for i := 0; i < attempts; i++ {
		t.Run(fmt.Sprintf("Exchange #%d", i+1), func(t *testing.T) {
			exchangeOnce(t, test_server)
		})
	}
}

func exchangeOnce(t *testing.T, test_server string) {
	start := time.Now()

	t.Log("Connecting to server at " + test_server + "...")
	conn, err := net.Dial("tcp", test_server)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	t.Log("Connected.")

	reader := bufio.NewReader(conn)

	// Step 1: Read challenge
	line, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read challenge from server: %v", err)
	}
	t.Logf("Received challenge line: %s", strings.TrimSpace(line))

	parts := strings.Fields(line)
	if len(parts) != 2 {
		t.Fatalf("Invalid challenge format: %v", parts)
	}

	challenge, difficulty := parts[0], 0
	fmt.Sscanf(parts[1], "%d", &difficulty)
	t.Logf("Parsed challenge: %s", challenge)
	t.Logf("Parsed difficulty: %d", difficulty)

	// Step 2: Solve PoW
	t.Log("Solving PoW challenge...")
	nonce := client.SolvePoW(challenge, difficulty)
	t.Logf("Solved PoW! Nonce: %s", nonce)

	// Step 3: Send nonce
	fmt.Fprintf(conn, "%s\n", nonce)
	t.Log("Nonce sent to server.")

	// Step 4: Read server reply
	reply, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read reply from server: %v", err)
	}
	reply = strings.TrimSpace(reply)
	t.Logf("Server reply: %s", reply)

	if strings.HasPrefix(reply, "Invalid") {
		t.Errorf("Server rejected valid PoW: %s", reply)
	}

	// Print duration
	elapsed := time.Since(start)
	t.Logf("Exchange completed in %s", elapsed)
}
