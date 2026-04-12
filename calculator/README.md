# Calculator

A CLI calculator that evaluates simple math expressions like `3 + 4 * 2`.

---

## What You'll Learn

- Tokenizing an input string (splitting into numbers and operators)
- Using a stack data structure (implemented as a slice)
- Operator precedence logic
- Converting strings to numbers with `strconv`
- Reading from stdin with `bufio.Scanner`

---

## Project Structure

```
calculator/
├── main.go
├── tokenizer.go
└── evaluator.go
```

### `main.go`
Creates a `bufio.Scanner` reading from `os.Stdin`. Runs a loop: reads a line, passes it to the tokenizer, then evaluates the result. Prints the answer or an error. Type `exit` to quit.

### `tokenizer.go`
Defines a `Token` struct with a `Type` (NUMBER or OPERATOR) and `Value` (string). The `Tokenize(input string)` function iterates over the input character by character, grouping digits into number tokens and treating `+`, `-`, `*`, `/` as operator tokens. Returns `[]Token`.

### `evaluator.go`
Implements a simple two-pass evaluator:
1. **Pass 1:** Process `*` and `/` first (higher precedence) — scan through tokens and collapse any `number * number` or `number / number` into a single result token.
2. **Pass 2:** Process remaining `+` and `-` left-to-right.

This avoids needing a full expression tree and is beginner-friendly.

---

## How It Works

```
Input:  "3 + 4 * 2"
Tokens: [3] [+] [4] [*] [2]

Pass 1 (handle * /):   [3] [+] [8]
Pass 2 (handle + -):   [11]

Output: 11
```

---

## Data Flow

```
stdin → Tokenize() → []Token → Evaluate() → float64 → print
```

---

## Suggested Extensions

- Support parentheses (requires a recursive descent parser)
- Support `^` for exponentiation using `math.Pow`
- Add a history of previous calculations (slice of strings)
- Handle division by zero gracefully
