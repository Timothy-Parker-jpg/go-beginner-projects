# Password Generator

A CLI tool that generates secure random passwords with configurable length and character sets.

---

## What You'll Learn

- Using `crypto/rand` for cryptographically secure random numbers (vs `math/rand`)
- Working with byte slices and rune slices
- Using the `flag` package for multiple CLI options
- String building with `strings.Builder`
- Copying text to clipboard (optional extension)

---

## Project Structure

```
password-generator/
├── main.go
└── generator.go
```

### `main.go`
Defines flags:
- `-length` (int, default 16) — number of characters
- `-upper` (bool, default true) — include uppercase letters
- `-lower` (bool, default true) — include lowercase letters
- `-digits` (bool, default true) — include numbers
- `-symbols` (bool, default false) — include special characters
- `-count` (int, default 1) — how many passwords to generate

Calls `flag.Parse()`, builds the character set, calls `Generate()`, and prints results.

### `generator.go`
Defines constant character set strings:
```go
const (
    lowercase = "abcdefghijklmnopqrstuvwxyz"
    uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digits    = "0123456789"
    symbols   = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)
```

The `Generate(length int, charset string) (string, error)` function:
1. Uses `crypto/rand.Int` to pick random indices into the charset string
2. Appends characters to a `strings.Builder`
3. Returns the completed password string

---

## How It Works

```
$ go run . -length 20 -symbols
Generated password: kR7!mX2@qP9#nL4$wB6^

$ go run . -length 12 -count 3
mXp9kL3nQr7w
bT4hN8vR2sYj
cP6mW1kD5fXn
```

---

## Why `crypto/rand` and not `math/rand`?

`math/rand` is deterministic — given the same seed it produces the same sequence. `crypto/rand` pulls from the OS entropy source (e.g., `/dev/urandom`) and is not predictable. Always use `crypto/rand` for security-sensitive randomness.

---

## Data Flow

```
flags → build charset string → Generate(length, charset) → print password(s)
```

---

## Suggested Extensions

- Add a `-exclude` flag to exclude ambiguous characters (`0`, `O`, `l`, `1`)
- Add password strength estimation (entropy bits = log2(charsetSize^length))
- Save generated passwords to a file with timestamps
- Check if the password meets common policy requirements
