package main

import (
	"flag"
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
