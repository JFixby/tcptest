package main

import (
	"bufio"
	"fmt"
	"github.com/jfixby/tcptest/client/client"
	"log"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		Exchange("localhost:1337")
	}
}

func Exchange(address string) {
	start := time.Now()

	conn := client.ConnectToServer(address)
	defer conn.Close()

	reader := bufio.NewReader(conn)

	challenge, difficulty := client.ReadChallenge(reader)
	nonce := client.SolveChallenge(challenge, difficulty)
	client.SendNonce(conn, nonce)
	client.ReadReply(reader)

	elapsed := time.Since(start)
	log.Printf("Exchange completed in %s\n", elapsed)
	fmt.Println()
}
