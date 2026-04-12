# URL Shortener

An HTTP server that creates short URLs (e.g., `/abc123`) and redirects visitors to the original long URL.

---

## What You'll Learn

- Setting up an HTTP server with `net/http`
- Handling routes with `http.HandleFunc`
- Reading from and writing to a `map` (with mutex for thread safety)
- Generating random short codes
- HTTP redirects with `http.Redirect` and status 301/302
- Parsing request bodies with `io.ReadAll`

---

## Project Structure

```
url-shortener/
├── main.go
├── store.go
└── handlers.go
```

### `main.go`
Creates a `Store`, registers two routes, and starts the server on `:8080`:
- `POST /shorten` — accepts a long URL, returns a short code
- `GET /{code}` — looks up the code and redirects

### `store.go`
Defines the `Store` struct:
```go
type Store struct {
    mu   sync.RWMutex
    urls map[string]string // short code → original URL
}
```

Methods:
- `Save(code, url string)` — acquires write lock, stores the mapping
- `Get(code string) (string, bool)` — acquires read lock, returns the URL
- `GenerateCode() string` — generates a random 6-character alphanumeric code using `crypto/rand`

The `sync.RWMutex` allows multiple concurrent readers but only one writer — important because HTTP handlers run in separate goroutines.

### `handlers.go`
`ShortenHandler(store *Store)` returns an `http.HandlerFunc`:
1. Only accepts POST requests
2. Reads the URL from the request body with `io.ReadAll`
3. Validates it starts with `http://` or `https://`
4. Calls `store.GenerateCode()` and `store.Save()`
5. Writes back the full short URL as plain text

`RedirectHandler(store *Store)` returns an `http.HandlerFunc`:
1. Extracts the code from the URL path
2. Calls `store.Get(code)`
3. If found: `http.Redirect(w, r, originalURL, http.StatusFound)`
4. If not found: `http.Error(w, "Not found", http.StatusNotFound)`

---

## How It Works

```
$ go run .
Server running on :8080

# In another terminal:
$ curl -X POST http://localhost:8080/shorten -d "https://www.google.com/search?q=golang"
http://localhost:8080/xK9mP2

$ curl -L http://localhost:8080/xK9mP2
# → Redirects to https://www.google.com/search?q=golang
```

---

## Data Flow

```
POST /shorten → read body → validate → generate code → store.Save() → return short URL
GET /{code}   → extract code → store.Get() → http.Redirect() to original URL
```

---

## Suggested Extensions

- Persist the map to a JSON file on shutdown (using `os.Signal` / `signal.Notify`)
- Add a simple HTML form at `GET /` to shorten URLs in a browser
- Add a hit counter to each short URL
- Add expiration times to URLs
