# Web Scraper

A CLI tool that fetches a webpage and extracts all links, headings, or plain text — configurable by flag.

---

## What You'll Learn

- Making HTTP GET requests and reading response bodies
- Parsing HTML with the `golang.org/x/net/html` package
- Walking an HTML node tree recursively
- Filtering nodes by type and attribute
- Using `go mod init` and `go get` to add external dependencies

---

## Project Structure

```
web-scraper/
├── main.go
├── fetcher.go
└── scraper.go
```

### `main.go`
Accepts flags:
- `-url` — the page to scrape (required)
- `-mode` — what to extract: `links`, `headings`, or `text` (default: `links`)

Calls `Fetch()` then passes the result to the appropriate scraper function.

### `fetcher.go`
The `Fetch(url string) (io.ReadCloser, error)` function:
1. Creates an `http.Client` with a 10-second timeout: `&http.Client{Timeout: 10 * time.Second}`
2. Makes a GET request with a realistic `User-Agent` header (some sites block Go's default)
3. Checks `resp.StatusCode` — returns an error if not 200
4. Returns `resp.Body` (the caller is responsible for closing it)

### `scraper.go`
All functions take an `io.Reader` and use `html.Parse(r)` to get the root `*html.Node`.

**`ExtractLinks(r io.Reader) []string`**
Recursively walks nodes. When a node is type `html.ElementNode` with tag `"a"`, look through its `Attr` slice for `Key == "href"`. Collect and return all found href values.

**`ExtractHeadings(r io.Reader) []string`**
Same walk pattern. Match tags `"h1"` through `"h6"`. For each heading node, collect the text content from its child `html.TextNode` children.

**`ExtractText(r io.Reader) string`**
Walk all nodes. When a node is type `html.TextNode` and its parent is not a `<script>` or `<style>` tag, append its `Data` to a `strings.Builder`. Return the full text.

---

## How It Works

```
$ go run . -url https://example.com -mode links

https://www.iana.org/domains/reserved

$ go run . -url https://example.com -mode headings

h1: Example Domain
```

---

## Data Flow

```
URL → Fetch() → io.Reader → html.Parse() → *html.Node → walk tree → extract → print
```

---

## The Node Walk Pattern

```go
func walk(n *html.Node, fn func(*html.Node)) {
    fn(n)
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        walk(c, fn)
    }
}
```

This pattern comes up constantly in HTML scraping — learn it well.

---

## Suggested Extensions

- Follow links recursively up to a `-depth N` limit (basic crawler)
- Save output to a file with a `-output` flag
- Extract all images (`<img src="...">`)
- Respect `robots.txt` before crawling
