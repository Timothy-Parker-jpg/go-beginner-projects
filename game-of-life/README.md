# Conway's Game of Life

A terminal-based simulation of Conway's Game of Life — cells live, die, and reproduce based on simple rules.

---

## What You'll Learn

- Using a 2D slice (`[][]bool`) as a grid
- Double-buffering (writing to a new grid while reading from the current one)
- Neighbor-counting with bounded or wrapping grid edges
- Clearing and redrawing the terminal with ANSI escape codes
- Controlling simulation speed with `time.Sleep`
- Randomly initializing a grid with `math/rand`

---

## Project Structure

```
game-of-life/
├── main.go
├── grid.go
└── rules.go
```

### `main.go`
Parses flags:
- `-width` (default: 40)
- `-height` (default: 20)
- `-fps` (default: 10) — frames per second
- `-density` (default: 0.3) — fraction of cells alive at start (0.0–1.0)

Creates a `Grid`, seeds it randomly, then runs the simulation loop: step → draw → sleep.

### `grid.go`
Defines the `Grid` struct:
```go
type Grid struct {
    Width  int
    Height int
    Cells  [][]bool
}
```

`NewGrid(w, h int) *Grid` allocates a `w × h` 2D slice.

`Seed(density float64)` iterates all cells and sets each to `true` with probability `density` using `rand.Float64() < density`.

`Draw()` prints the grid to the terminal:
1. Print ANSI escape code `\033[H` to move the cursor to the top-left (avoids screen flicker vs clearing)
2. For each row, build a string: `"█"` for alive cells, `" "` for dead cells
3. Print each row with `fmt.Println`

### `rules.go`
`CountNeighbors(g *Grid, x, y int) int` counts how many of the 8 surrounding cells are alive. Use modulo wrapping so the grid is toroidal: `(x-1+g.Width) % g.Width`.

`Step(g *Grid) *Grid`:
1. Create a new `Grid` of the same size
2. For every cell `(x, y)`:
   - Count its neighbors
   - Apply Conway's rules:
     - Alive + 2 or 3 neighbors → stays alive
     - Alive + other count → dies
     - Dead + exactly 3 neighbors → becomes alive
     - Dead + other count → stays dead
3. Return the new grid

---

## How It Works

```
$ go run . -width 60 -height 30 -fps 15

█  ██  █   ██     █   ██  █ ...
 █  █ ██  █  █  ██  █  █  ...
...
(updates in place at ~15 fps)
```

---

## Data Flow

```
NewGrid() → Seed() → loop: Step() → Draw() → time.Sleep(1/fps)
```

---

## The Double-Buffer Pattern (Key Concept)

You CANNOT update cells in place — a cell's new state must be computed from its neighbors' CURRENT state. If you update cell A, then compute cell B using A's new value, the result is wrong. The solution: always compute the next generation into a brand-new grid, then swap it in.

---

## Suggested Extensions

- Add pause/resume with keyboard input using raw terminal mode
- Load known patterns from a file (Gosper Glider Gun, gliders, etc.)
- Add a generation counter and live cell counter
- Color cells based on their age (how many consecutive generations they've been alive)
