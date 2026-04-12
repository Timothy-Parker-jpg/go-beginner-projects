# Word / Line Counter

A CLI tool that counts lines, words, and characters in a file — similar to the Unix `wc` command.

---

## What You'll Learn

- Reading files with `os.Open` and `bufio.Scanner`
- Scanning line by line vs word by word
- Using `strings.Fields` to split on whitespace
- Accepting multiple file paths from `os.Args`
- Printing aligned tabular output with `fmt.Printf` and `%d` width specifiers

---

## Project Structure

```
word-counter/
├── main.go
└── counter.go
```

### `main.go`
Reads file paths from `os.Args[1:]`. If no files are given, reads from `os.Stdin`. For each file, calls `CountFile()` and prints the results. If multiple files are given, prints a total row at the bottom.

### `counter.go`
Defines a `Counts` struct:
```go
type Counts struct {
    Lines int
    Words int
    Chars int
}
```

The `CountFile(r io.Reader) Counts` function:
1. Creates a `bufio.Scanner` on the reader
2. Sets scanner mode to `ScanLines`
3. For each line: increments `Lines`, adds `len(strings.Fields(line))` to `Words`, adds `len(line)` to `Chars`
4. Returns the `Counts`

---

## How It Works

```
$ go run . main.go
  42   187  1203 main.go

$ go run . *.go
  42   187  1203 main.go
  28   110   876 counter.go
  70   297  2079 total
```

---

## Data Flow

```
os.Args → open file → io.Reader → CountFile() → Counts → print
```

The use of `io.Reader` (not `*os.File` directly) is intentional — it makes `CountFile()` testable with any reader, including `strings.NewReader("...")` in tests.

---

## Suggested Extensions

- Add a `-l` flag for lines only, `-w` for words only, `-c` for chars only
- Count unique words and print the most frequent ones
- Add support for reading from a URL (using `net/http`)
- Write a benchmark test using `testing.B`
