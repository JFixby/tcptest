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

// TestExchangeWithServer starts a live TCP server and runs multiple PoW exchanges against it.
// This is an integration-style test, verifying full client-server interaction via TCP.
func TestExchangeWithServer(t *testing.T) {
	testAddress := ":1337"
	s := server.NewServer()

	// Start the TCP server in a background goroutine
	go func() {
		err := s.Start(testAddress, "./server/wisdoms.json")
		if err != nil {
			t.Logf("Server exited with error: %v", err)
		}
	}()
	defer s.Stop()

	// Wait for the TCP server to be ready to accept connections
	waitForServer(t, testAddress)

	const calls = 4
	for i := 0; i < calls; i++ {
		t.Run(fmt.Sprintf("Exchange #%d", i+1), func(t *testing.T) {
			exchangeOnce(t, testAddress)
		})
	}
}

// exchangeOnce performs a full client-server proof-of-work interaction over TCP.
func exchangeOnce(t *testing.T, address string) {
	start := time.Now()

	t.Logf("📡 Connecting to TCP server at %s...", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		t.Fatalf("❌ Failed to connect to server: %v", err)
	}
	defer conn.Close()
	t.Log("✅ Connection established.")

	reader := bufio.NewReader(conn)

	// Step 1: Receive challenge line from server
	line, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("❌ Failed to read challenge from server: %v", err)
	}
	line = strings.TrimSpace(line)
	t.Logf("📨 Received challenge line: \"%s\"", line)

	// Parse challenge and difficulty
	parts := strings.Fields(line)
	if len(parts) != 2 {
		t.Fatalf("❌ Invalid challenge format: %v", parts)
	}
	challenge, difficulty := parts[0], 0
	fmt.Sscanf(parts[1], "%d", &difficulty)
	t.Logf("🧩 Parsed challenge: %s", challenge)
	t.Logf("🧠 Parsed difficulty: %d", difficulty)

	// Step 2: Solve proof of work
	t.Log("⚙️  Solving PoW...")
	nonce := client.SolvePoW(challenge, difficulty)
	t.Logf("✅ PoW solved — Nonce: %s", nonce)

	// Step 3: Send nonce to server
	_, err = fmt.Fprintf(conn, "%s\n", nonce)
	if err != nil {
		t.Fatalf("❌ Failed to send nonce: %v", err)
	}
	t.Log("📤 Nonce sent to server.")

	// Step 4: Read server response (expected: a Protoss quote or error)
	reply, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("❌ Failed to read reply from server: %v", err)
	}
	reply = strings.TrimSpace(reply)
	t.Logf("📬 Server reply: \"%s\"", reply)

	// Validate reply
	if strings.HasPrefix(reply, "Invalid") {
		t.Errorf("❌ Server rejected PoW: %s", reply)
	} else {
		t.Logf("💬 Received wisdom: %s", reply)
	}

	// Log timing
	elapsed := time.Since(start)
	t.Logf("⏱ Exchange completed in %s", elapsed)
}

// waitForServer tries to connect until the TCP server becomes available or times out.
// Prevents race conditions between starting the server and attempting to connect.
func waitForServer(t *testing.T, address string) {
	const maxWait = 2 * time.Second
	deadline := time.Now().Add(maxWait)

	for {
		conn, err := net.Dial("tcp", address)
		if err == nil {
			conn.Close()
			t.Log("✅ TCP server is ready.")
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("❌ Server did not start within %s", maxWait)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
