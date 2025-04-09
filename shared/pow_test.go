package shared

import (
	"strings"
	"testing"
)

// TestToBitString проверяет корректность перевода хеша в битовую строку
func TestToBitString(t *testing.T) {
	// Хеш из нулей
	var hash [32]byte
	bitStr := ToBitString(hash)

	if len(bitStr) != 256 {
		t.Errorf("Ожидалась длина 256 бит, получено %d", len(bitStr))
	}

	if strings.Trim(bitStr, "0") != "" {
		t.Errorf("Все биты должны быть нулями, получено: %s", bitStr)
	}
}

// TestHash проверяет, что хеш одинаков при одинаковых входных
func TestHash(t *testing.T) {
	h1 := Hash("test", "123")
	h2 := Hash("test", "123")

	if h1 != h2 {
		t.Errorf("Ожидался одинаковый хеш, но получены разные:\n%v\n%v", h1, h2)
	}

	h3 := Hash("test", "456")
	if h1 == h3 {
		t.Errorf("Ожидался разный хеш для разных входных данных")
	}
}

// TestCheckPoW проверяет валидацию PoW при простом примере
func TestCheckPoW(t *testing.T) {
	challenge := "test"
	nonce := "0000"

	valid, hash, _ := CheckPoW(challenge, nonce, 0)
	if !valid {
		t.Errorf("difficulty=0 должно всегда быть валидным, но вернуло false")
	}

	valid, _, _ = CheckPoW(challenge, nonce, 256)
	if valid {
		t.Errorf("difficulty=256 не должен быть валидным почти никогда (слишком сложно)")
	}

	// Для наглядности: проверим, что хеш совпадает с ручным
	expectedHash := Hash(challenge, nonce)
	if hash != expectedHash {
		t.Errorf("Hash внутри CheckPoW не совпадает с Hash(): %v != %v", hash, expectedHash)
	}
}
