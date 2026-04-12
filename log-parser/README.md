# Log File Parser

A CLI tool that reads a server log file and summarizes it: request counts, error rates, top endpoints, and status code breakdown.

---

## What You'll Learn

- Line-by-line file reading with `bufio.Scanner`
- Parsing structured strings with `strings.Split` and `regexp`
- Using maps to count occurrences (`map[string]int`)
- Sorting map keys by value using `sort.Slice`
- Time parsing with `time.Parse`

---

## Project Structure

```
log-parser/
├── main.go
├── parser.go
└── report.go
```

### `main.go`
Accepts a log file path and optional flags:
- `-top N` — show top N endpoints (default 5)
- `-errors-only` — only show lines with 4xx/5xx status codes

Opens the file and passes it to `ParseLog()`, then calls `PrintReport()`.

### `parser.go`
Defines a `LogEntry` struct:
```go
type LogEntry struct {
    IP        string
    Timestamp time.Time
    Method    string
    Path      string
    Status    int
    Bytes     int
}
```

The `ParseLog(r io.Reader) []LogEntry` function uses a `bufio.Scanner` and a compiled `regexp.MustCompile` pattern to extract fields from each line of the Common Log Format:
```
127.0.0.1 - - [10/Oct/2024:13:55:36 -0700] "GET /index.html HTTP/1.1" 200 1234
```

Lines that don't match are skipped.

### `report.go`
The `PrintReport(entries []LogEntry, topN int)` function:
1. Counts total requests and groups by status code
2. Counts hits per path using `map[string]int`
3. Sorts paths by hit count using `sort.Slice`
4. Prints a formatted summary table

---

## How It Works

```
$ go run . access.log -top 3

Total requests: 1842
Errors (4xx/5xx): 37 (2.0%)

Status codes:
  200: 1723
  404:   28
  500:    9

Top endpoints:
  /index.html      →  843
  /api/users       →  412
  /static/main.js  →  301
```

---

## Data Flow

```
log file → bufio.Scanner → regexp.FindStringSubmatch → []LogEntry → maps → PrintReport()
```

---

## Suggested Extensions

- Group requests by hour to see traffic patterns
- Filter by IP address with a `-ip` flag
- Detect repeated 401/403 errors from the same IP (brute force detection)
- Write matching entries to a new filtered log file
