package shared

import (
	"crypto/sha256" // Используем для вычисления хешей SHA-256
	"fmt"           // Для форматирования байтов в строки
	"strings"       // Для манипуляций со строками
)

// ToBitString конвертирует SHA256-хеш (массив из 32 байт) в строку из 256 бит ("0" и "1").
// Каждый байт преобразуется в строку из 8 символов (все ведущие нули сохраняются).
func ToBitString(hash [32]byte) string {
	var sb strings.Builder // Строим строку эффективно
	for _, b := range hash {
		sb.WriteString(fmt.Sprintf("%08b", b)) // Преобразуем байт в 8-битную строку, без обрезки
	}
	return sb.String() // Возвращаем строку длиной 256 символов
}

// CheckPoW проверяет, что SHA256(challenge + nonce) начинается с заданного количества нулевых битов.
//
// challenge — строка, выдаваемая сервером
// nonce     — подбираемое клиентом значение
// difficulty — количество ведущих нулей, которое нужно получить в битовой строке хеша
//
// Возвращает:
// - true, если хеш начинается с difficulty нулей
// - сам хеш [32]byte
// - строку битов (для логирования и отладки)
func CheckPoW(challenge, nonce string, difficulty int) (bool, [32]byte, string) {
	hash := Hash(challenge, nonce)                                      // Получаем хеш от challenge + nonce
	bits := ToBitString(hash)                                           // Конвертируем в битовую строку
	isValid := strings.HasPrefix(bits, strings.Repeat("0", difficulty)) // Проверяем по префиксу
	return isValid, hash, bits
}

// Hash возвращает SHA256 хеш от строки challenge + nonce.
// Это основной шаг PoW — создание контрольного значения.
func Hash(challenge, nonce string) [32]byte {
	return sha256.Sum256([]byte(challenge + nonce))
}
