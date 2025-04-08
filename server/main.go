package main

import (
	"fmt"
	"github.com/jfixby/tcptest/server/server"
	"os"
)

func main() {
	// Проверяем наличие аргументов: адрес и путь к файлу с мудростями
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <address> <wisdoms_file>")
		os.Exit(1)
	}

	address := os.Args[1]     // Пример: ":1337"
	wisdomsFile := os.Args[2] // Пример: "server/wisdoms.json"

	s := server.NewServer()
	if err := s.Start(address, wisdomsFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
