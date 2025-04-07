package main

import (
	"bufio"
	"fmt"
	"github.com/jfixby/tcptest/client/client"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	if len(parts) != 2 {
		fmt.Println("Invalid challenge format")
		return
	}

	challenge, difficulty := parts[0], 0
	fmt.Sscanf(parts[1], "%d", &difficulty)

	nonce := client.SolvePoW(challenge, difficulty)
	fmt.Fprintf(conn, "%s\n", nonce)

	reply, _ := reader.ReadString('\n')
	fmt.Println("Server says:", strings.TrimSpace(reply))
}
