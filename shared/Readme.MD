# 🔐 shared — Proof-of-Work Utilities (PoW)

Этот пакет предоставляет простую реализацию **Proof-of-Work (PoW)** на основе SHA-256, используемую для защиты от спама и злоупотреблений при подключении к TCP-серверу.

---

## ⚙️ Как работает PoW

1. **Сервер** генерирует `challenge` (строку) и `difficulty` (целое число).
2. **Клиент** подбирает строку `nonce`, такую что:
   ```
   sha256(challenge + nonce)
   ```
   имеет **не менее N нулей в начале битовой строки хеша**.
3. Сервер проверяет корректность этого `nonce` при подключении клиента.

---

## 📦 Функции

### `Hash(challenge, nonce string) [32]byte`

Возвращает SHA-256 от строки `challenge + nonce`.

```go
hash := shared.Hash("abc", "123")
```

---

### `ToBitString(hash [32]byte) string`

Конвертирует хеш (массив из 32 байт) в строку из 256 бит.

```go
bits := shared.ToBitString(hash)
// Пример: "0000010110011010..."
```

---

### `CheckPoW(challenge, nonce string, difficulty int) (bool, [32]byte, string)`

Проверяет, что `sha256(challenge + nonce)` начинается с нужного количества `difficulty` нулевых битов.

```go
ok, hash, bits := shared.CheckPoW("hello", "42", 20)
if ok {
    fmt.Println("✔ Valid PoW!")
} else {
    fmt.Println("❌ Invalid nonce")
}
```

---

## 🧪 Пример: ручной перебор nonce

```go
challenge := "abc123"
difficulty := 20
nonce := 0

for {
    candidate := fmt.Sprintf("%d", nonce)
    ok, hash, bits := shared.CheckPoW(challenge, candidate, difficulty)
    if ok {
        fmt.Printf("✅ Found nonce: %s\nHash: %x\nBits: %s\n", candidate, hash, bits[:32])
        break
    }
    nonce++
}
```

---

## 💡 Применение

- Простая защита TCP-сервера от спама / DDoS
- Демонстрация принципов майнинга и распределённого вычисления
- Обратимая и быстрая проверка

---

## 📁 Структура

- `Hash` — вычисляет SHA-256(challenge + nonce)
- `ToBitString` — представление хеша в виде битов
- `CheckPoW` — основная проверка корректности PoW

---

## ⚠️ Замечания

- Это **упрощённый алгоритм**, не защищённый от GPU/ASIC.
- Используется только для **обучающих или анти-спам задач**.

---

## 📚 См. также

- [Bitcoin whitepaper](https://bitcoin.org/bitcoin.pdf)
- [Hashcash](https://en.wikipedia.org/wiki/Hashcash)
