package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jfixby/tcptest/client/client"
	"log"
	"time"
)

func main() {
	// CLI флаги
	address := flag.String("address", "localhost:1337", "Адрес TCP-сервера (host:port)")
	count := flag.Int("count", 4, "Сколько раз выполнить обмен с сервером")
	flag.Parse()

	// Повторяем обмен указанное количество раз
	for i := 0; i < *count; i++ {
		Exchange(*address)
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
