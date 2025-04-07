package server

import (
	"log"
	"sync"
	"time"
)

var (
	difficulty     = 15 // initial difficulty
	difficultyLock sync.Mutex
)

// GetDifficulty returns the current difficulty safely.
func GetDifficulty() int {
	difficultyLock.Lock()
	defer difficultyLock.Unlock()
	return difficulty
}

// AdjustDifficulty modifies difficulty based on the time taken.
func AdjustDifficulty(duration time.Duration) {
	difficultyLock.Lock()
	defer difficultyLock.Unlock()

	switch {
	case duration > 2*time.Second && difficulty > 1:
		difficulty--
		log.Printf("💡 Exchange took %s — lowering difficulty to %d", duration, difficulty)
	case duration < 2*time.Second && difficulty < 30:
		difficulty++
		log.Printf("⚡ Exchange took %s — increasing difficulty to %d", duration, difficulty)
	default:
		log.Printf("⏱ Exchange took %s — difficulty remains at %d", duration, difficulty)
	}
}
