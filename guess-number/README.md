# Guess the Number

A simple CLI game where the computer picks a random number and the player tries to guess it with hints.

---

## What You'll Learn

- Seeding and using `math/rand`
- Reading user input with `bufio.Scanner` on `os.Stdin`
- Converting input strings to integers with `strconv.Atoi`
- Game loop design with `for` and `break`
- Tracking state across loop iterations (attempts counter, bounds)

---

## Project Structure

```
guess-number/
├── main.go
└── game.go
```

### `main.go`
Parses optional flags:
- `-max N` — upper bound of the range (default: `100`)
- `-attempts N` — max allowed guesses, 0 = unlimited (default: `0`)

Creates a `Game` and calls `game.Run()`.

### `game.go`
Defines the `Game` struct:
```go
type Game struct {
    Secret   int
    Max      int
    MaxTries int
    Attempts int
}
```

`NewGame(max, maxTries int) *Game`:
1. Seeds `rand` with the current time using `rand.New(rand.NewSource(time.Now().UnixNano()))`
2. Generates the secret: `rand.Intn(max) + 1`
3. Returns a new `Game`

`Run()` method:
1. Print the game intro: "I'm thinking of a number between 1 and N"
2. Create a `bufio.Scanner` on `os.Stdin`
3. Loop:
   a. Print prompt (and remaining attempts if limited)
   b. Read a line with `scanner.Scan()`
   c. Parse it with `strconv.Atoi` — print error and continue if invalid
   d. Increment `Attempts`
   e. Compare guess to `Secret`:
      - Too low → "Higher!"
      - Too high → "Lower!"
      - Correct → print win message with attempt count, `break`
   f. If `MaxTries > 0` and `Attempts >= MaxTries` → print loss message, `break`

---

## How It Works

```
$ go run . -max 50 -attempts 5

I'm thinking of a number between 1 and 50.
You have 5 attempts.

Guess: 25
Too high!

Guess: 12
Too low!

Guess: 18
Too high!

Guess: 15
🎉 Correct! You got it in 4 attempts.
```

---

## Data Flow

```
flags → NewGame() → Run() loop → read stdin → parse → compare → hint or win/lose
```

---

## Suggested Extensions

- Track a high score (fewest attempts) and save it to a file
- Add difficulty levels (Easy: 1-50, Medium: 1-100, Hard: 1-500)
- Let the roles reverse: the player thinks of a number and the computer guesses using binary search
- Add a timer to race against the clock
