# Hangman

A classic CLI hangman game where the player guesses letters to reveal a hidden word.

---

## What You'll Learn

- Storing a word list in a slice and picking randomly
- Using a `map[rune]bool` to track guessed letters
- Manipulating `[]rune` (Go's representation of characters)
- Checking win/loss conditions in a game loop
- Printing multi-line ASCII art

---

## Project Structure

```
hangman/
├── main.go
├── game.go
├── display.go
└── words.go
```

### `main.go`
Seeds the random number generator. Creates a new game with a random word, then runs the game loop.

### `words.go`
Contains a `var wordList = []string{...}` of 50–100 common words. The `RandomWord() string` function picks one using `rand.Intn(len(wordList))`.

### `game.go`
Defines the `Game` struct:
```go
type Game struct {
    Word    []rune
    Guessed map[rune]bool
    Lives   int
}
```

`NewGame(word string) *Game` initializes with `Lives = 6`.

Key methods:
- `Guess(letter rune) bool` — adds letter to `Guessed` map, returns true if it's in the word
- `IsWon() bool` — returns true if every rune in `Word` exists in `Guessed`
- `IsLost() bool` — returns true if `Lives == 0`
- `DisplayWord() string` — iterates over `Word`; for each rune, if it's in `Guessed` append it, otherwise append `"_"`; join with spaces

`Run()` runs the main loop:
1. Print the gallows, word display, and wrong guesses
2. Prompt for a letter
3. Read one character from stdin
4. Call `Guess()` — if miss, decrement `Lives`
5. Check `IsWon()` / `IsLost()` and break accordingly

### `display.go`
Stores the 7 stages of the hangman (empty gallows through full figure) as a `[7]string` array of multi-line ASCII art strings. The current stage shown is `6 - lives`.

---

## How It Works

```
$ go run .

  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
Word: _ _ _ _ _
Wrong guesses: x, q
Lives: 4

Guess a letter: 
```

---

## Data Flow

```
RandomWord() → NewGame() → Run() loop → Guess() → update Lives → DisplayWord() → win/lose
```

---

## Suggested Extensions

- Load words from a text file instead of the hardcoded list
- Add word categories (animals, countries, food) with a `-category` flag
- Show a hint (first letter or category name) if the player is stuck
- Add a two-player mode where one player types the word and the other guesses
