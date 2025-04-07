package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "server:1337")
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

	nonce := solvePoW(challenge, difficulty)
	fmt.Fprintf(conn, "%s\n", nonce)

	reply, _ := reader.ReadString('\n')
	fmt.Println("Server says:", strings.TrimSpace(reply))
}
