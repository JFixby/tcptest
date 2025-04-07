package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// Quote represents a unit quote from wisdoms.json
type Quote struct {
	Unit  string `json:"unit"`
	Quote string `json:"quote"`
}

var quotes []Quote

// LoadQuotes loads quotes from a JSON file into memory
func LoadQuotes(filename string) error {
	log.Printf("Loading quotes from file: %s", filename)

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(bufio.NewReader(file)).Decode(&quotes)
	if err != nil {
		log.Printf("Failed to decode quotes: %v", err)
	} else {
		log.Printf("Loaded %d quotes", len(quotes))
		for i, q := range quotes {
			log.Printf("[%d] %s: \"%s\"", i+1, q.Unit, q.Quote)
		}
	}

	return err
}

// GetRandomQuote returns a random quote from the loaded list
func GetRandomQuote() string {
	if len(quotes) == 0 {
		return "No Protoss wisdom found."
	}

	q := quotes[rand.Intn(len(quotes))]
	return fmt.Sprintf("[%s] %s", q.Unit, q.Quote)
}
