package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	err := loadQuotes("wisdoms.json")
	if err != nil {
		log.Fatalf("Failed to load quotes: %v", err)
	}

	ln, err := net.Listen("tcp", ":1337")
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
		go handleConnection(conn)
	}
}
