# AnBubble Chat

A small, self-hosted chat app for your local network. Written in Go with no external dependencies: one binary, one in-memory message store, and a plain vanilla JS/HTML/CSS frontend.

Everyone on the same LAN opens the same URL and starts chatting. No accounts, no database, no setup beyond running the server.

## Features

- Simple bubble-style chat UI, mobile-friendly
- Only one person needs to run the app. Everyone else just opens a link in their browser, nothing to install
- Privacy-friendly by design: messages live only in memory and are gone for good once the server stops, nothing is written to disk
- Thread-safe message store, safe for many concurrent users
- Auto-discovers your LAN IP so others can join instantly
- Zero external dependencies, pure Go standard library and vanilla JS

## Requirements

- [Go](https://go.dev/dl/) 1.26.4 or later

## Project Structure

```
AnBubble/
├── main.go          # entry point, HTTP server & routing
├── store.go         # thread-safe in-memory message store
├── message.go        # Message struct
├── handlers.go       # HTTP handlers (/send, /messages)
├── go.mod
├── LICENSE
└── web/               # static frontend, served at "/"
    ├── index.html
    ├── style.css
    └── app.js
```

Note: the server serves static files from a `web/` folder, so make sure `index.html`, `style.css`, and `app.js` live inside a `web/` directory next to `main.go`.

## Setup

1. Clone the repository

   ```bash
   git clone https://github.com/AnduanBerisha/AnBubble.git
   cd AnBubble
   ```

2. Run the server

   ```bash
   go run .
   ```

   You should see something like:

   ```
   Server running on LAN: http://192.168.1.42:8080
   ```

3. Open the app

   Visit that address in your browser. Anyone else on the same Wi-Fi/LAN can open the same URL from their own device to join the chat.

### Build a binary (optional)

```bash
go build -o anbubble .
./anbubble
```

## How It Works

- On first visit, the browser asks for a nickname and stores it locally.
- New messages are sent via `POST /send` and appended to an in-memory list.
- The frontend polls `GET /messages` every 2 seconds and renders any new bubbles.

| Endpoint    | Method | Description                          |
|-------------|--------|---------------------------------------|
| `/`         | GET    | Serves the frontend (`web/` folder)   |
| `/send`     | POST   | Sends a message (`{sender, text}`)    |
| `/messages` | GET    | Returns all messages as JSON          |

## Good to Know

- Chat history does not persist across restarts, by design. If you need long-term history, this is not the right tool.
- Designed for trusted local networks. There is no authentication or encryption.

## License

Licensed under the [MIT License](./LICENSE).