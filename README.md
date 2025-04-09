````markdown
# 🧠 tcptest — минималистичный TCP-сервер с Proof-of-Work

`tcptest` — это лёгкий и безопасный TCP-сервер, который перед выдачей случайной цитаты из _StarCraft: Brood War_ требует от клиента решения задачи **Proof-of-Work**. Это предотвращает спам и избыточные подключения.

---

## 🚀 Возможности

- 💬 Отправка случайных цитат из `wisdoms.json` (юниты Протоссов)
- 🔐 PoW-защита: вычислительная задача перед получением цитаты
- 📈 Адаптивная сложность в зависимости от времени решения
- 🧪 Интеграционный тест с реальным TCP-соединением
- 🐳 Полноценная Docker-поддержка и запуск через `docker-compose`
- ✅ CI через GitHub Actions: автоматическая проверка работоспособности

---

## ⚙️ Быстрый старт

### 🧠 Сервер (локально)

```bash
go run server/main.go -address=:1337 -wisdoms=server/wisdoms.json
```

### 🤖 Клиент (локально)

```bash
go run client/main.go -address=127.0.0.1:1337 -count=3
```

### 🐳 Запуск через Docker Compose

```bash
docker compose up --build
```

---

## 🧪 Интеграционный тест

```bash
go test -v integration_test.go
```

- Запускается реальный сервер
- Подключаются клиенты, решают PoW
- Получают цитаты
- Тестирует стабильность и корректность взаимодействия

---

## 📄 Структура проекта

```
.
├── client/          # TCP-клиент
│   ├── main.go
│   ├── Dockerfile
│   └── ...
├── server/          # TCP-сервер
│   ├── main.go
│   ├── wisdoms.json
│   └── Dockerfile
├── shared/          # Общая PoW-логика
│   ├── pow.go
│   └── pow_test.go
├── integration_test.go
├── docker-compose.yml
└── .github/workflows/docker-compose-test.yml
```

---

## 🧠 Пример wisdoms.json

```json
[
  {
    "unit": "Zealot",
    "quote": "My life for Aiur!"
  },
  {
    "unit": "High Templar",
    "quote": "The merging is complete."
  }
]
```

Можно расширить цитаты, добавить фракции Zerg и Terran.

---

## 🔒 Как работает PoW

- Сервер генерирует случайный challenge
- Клиент подбирает nonce, чтобы SHA-256(challenge + nonce) начинался с N нулевых битов
- Сложность (N) растёт, если клиент справляется слишком быстро

Общая логика PoW реализована в [shared/pow.go], и используется как клиентом, так и сервером.

---

## ✅ CI / GitHub Actions

Автоматическая проверка запуска `docker-compose`, падает, если клиент или сервер крашатся. См. [.github/workflows/docker-compose-test.yml]

---

## 💡 Идеи на будущее

- [ ] TLS-поддержка
- [ ] HTTP UI для отображения сложности и логов
- [ ] Расширенная база цитат
- [ ] Подключение через WebSocket
````

