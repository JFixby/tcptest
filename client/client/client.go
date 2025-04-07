package client

import (
	"fmt"
	"github.com/jfixby/tcptest/shared"
	"log"
)

func SolvePoW(challenge string, difficulty int) string {
	log.Printf("Starting PoW solver — challenge: %s, difficulty: %d", challenge, difficulty)

	for i := 0; ; i++ {
		nonce := fmt.Sprintf("%d", i)
		valid, hash, bits := shared.CheckPoW(challenge, nonce, difficulty)
		if valid {
			log.Printf("PoW solved! Nonce: %s", nonce)
			log.Printf("Hash: %x", hash)
			log.Printf("Bits: %s...", bits[:min(len(bits), difficulty+10)])
			return nonce
		}

		if i%1000 == 0 {
			log.Printf("Hash: %x", hash)
		}
	}
}
