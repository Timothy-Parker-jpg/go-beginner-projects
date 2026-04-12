# Static File HTTP Server

An HTTP server that serves files from a local directory — like `python -m http.server` but written in Go.

---

## What You'll Learn

- Using `http.FileServer` and `http.Dir`
- Writing custom middleware (logging, error handling)
- Using `http.StripPrefix`
- Parsing flags for port and directory
- Handling `http.Handler` vs `http.HandlerFunc`
- Graceful shutdown with `context` and `os.Signal`

---

## Project Structure

```
http-server/
├── main.go
├── middleware.go
└── server.go
```

### `main.go`
Parses two flags:
- `-port` (default `"8080"`)
- `-dir` (default `"."`)

Creates the server, registers middleware, and starts listening. Sets up a signal listener for `SIGINT`/`SIGTERM` to trigger graceful shutdown.

### `server.go`
Creates the file handler:
```go
fs := http.FileServer(http.Dir(rootDir))
handler := http.StripPrefix("/", fs)
```

`http.FileServer` handles directory listings and file serving automatically. `http.StripPrefix` strips the leading `/` from the URL path before looking up the file.

Wraps the file handler in middleware and passes it to `http.Server`:
```go
srv := &http.Server{
    Addr:    ":" + port,
    Handler: loggingMiddleware(handler),
}
```

Graceful shutdown using `srv.Shutdown(ctx)` with a 5-second timeout context.

### `middleware.go`
Implements a logging middleware using a custom `ResponseWriter` wrapper:
```go
type responseWriter struct {
    http.ResponseWriter
    statusCode int
}
```

The wrapper intercepts `WriteHeader(code int)` to capture the status code before passing it through. The middleware then logs: method, path, status code, and duration.

---

## How It Works

```
$ go run . -dir ./public -port 3000
Serving ./public on http://localhost:3000

GET  /index.html    200  1.2ms
GET  /style.css     200  0.3ms
GET  /missing.html  404  0.1ms
```

---

## Data Flow

```
HTTP Request → loggingMiddleware → http.StripPrefix → http.FileServer → reads file → response
```

---

## Suggested Extensions

- Add basic auth with a `-password` flag
- Serve a custom 404 page instead of the default Go one
- Add CORS headers for use as a local dev API mock server
- Add HTTPS support with `http.ListenAndServeTLS` and a self-signed cert
