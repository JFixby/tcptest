package main

import (
	"github.com/jfixby/tcptest/server/server"
)

func main() {
	s := server.NewServer()

	err := s.Start(":1337", "/wisdoms.json")
	if err != nil {
		panic(err)
	}
}
