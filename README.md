# 🧠 tcptest — Proof-of-Work TCP Wisdom Server

A fun and secure TCP server that challenges clients to solve a **Proof-of-Work** puzzle before sharing a random Protoss quote from _StarCraft: Brood War_.

---

## 🚀 Features

- 💬 Returns random quotes from Protoss units (`wisdoms.json`)
- 🔐 Defends against abuse with Proof-of-Work (PoW) challenges
- 🎯 Dynamically adjusts PoW difficulty based on solve time
- 🧪 Includes full integration test (real client/server exchange)
- ♻️ Graceful server shutdown on `Ctrl+C`

---

## 📦 Project Structure

```
tcptest/
├── client/              # PoW client
│   ├── main.go          # CLI client
│   └── client/client.go # SolvePoW logic
├── server/              # TCP server
│   ├── main.go          # Entry point
│   ├── wisdoms.json     # Quotes source
│   └── server/          # Server logic
│       ├── connection.go
│       ├── difficulty.go
│       ├── quotes.go
│       └── server.go
├── shared/              # Shared PoW logic
│   └── pow.go
├── integration_test.go  # Full end-to-end test
└── README.md
```

---

## 🛠 Usage

### 🧠 Run the Server

```bash
go run server/main.go
```

> Listens on TCP port `:1337`  
> Sends quotes only after verifying a valid PoW nonce

### 🤖 Run the Client

```bash
go run client/main.go
```

> Connects to the server, solves PoW, and prints the wisdom  
> Can be run multiple times to see dynamic difficulty in action

### 🧪 Run Integration Tests

```bash
go test -v
```

> Launches a full server in the background, connects via real TCP, solves PoW, and verifies responses.

---

## 📄 wisdoms.json Example

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

Customize this with your own Protoss wisdoms, or add Zerg and Terran!

---

## 🔒 Proof-of-Work Details

- Challenge: Random 16-character string
- PoW: SHA-256 hash of (challenge + nonce) must start with N leading zero **bits**
- Difficulty auto-tunes based on client solve time

---

## 📦 TODO / Ideas

- [ ] Docker support for client/server
- [ ] Add TLS support for encrypted quotes
- [ ] HTTP dashboard for live difficulty
- [ ] Extend `wisdoms.json` to support factions
