package main

import (
	"github.com/jfixby/tcptest/server/server"
)

func main() {
	server.Start(":1337", "wisdoms.json")
}
