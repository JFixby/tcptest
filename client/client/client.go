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

		if shared.CheckPoW(challenge, nonce, difficulty) {
			hash := shared.Hash(challenge, nonce)
			bits := shared.ToBitString(hash)

			log.Printf("PoW solved! Nonce: %s", nonce)
			log.Printf("Hash: %x", hash)
			log.Printf("Bits: %s...", bits[:min(len(bits), difficulty+10)])

			return nonce
		}

		if i%100 == 0 {
			hash := shared.Hash(challenge, nonce)
			log.Printf("Hash: %x", hash)
		}
	}
}
