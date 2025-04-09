package main

import (
	"flag"
	"fmt"
	"github.com/jfixby/tcptest/server/server"
	"os"
)

func main() {
	// Флаги командной строки
	address := flag.String("address", ":1337", "Адрес TCP-сервера (например, :1337)")
	wisdoms := flag.String("wisdoms", "server/wisdoms.json", "Путь к файлу с мудростями (JSON)")
	flag.Parse()

	// Запуск сервера
	s := server.NewServer()
	if err := s.Start(*address, *wisdoms); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
		os.Exit(1)
	}
}
