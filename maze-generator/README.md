# Maze Generator

A CLI tool that generates and prints a random maze using recursive backtracking (depth-first search).

---

## What You'll Learn

- Representing a grid as a 2D slice of structs
- Implementing recursive backtracking (a form of DFS)
- Working with directions as coordinate offsets
- Shuffling slices with `rand.Shuffle`
- Printing a 2D grid with box-drawing characters

---

## Project Structure

```
maze-generator/
├── main.go
├── maze.go
└── display.go
```

### `main.go`
Parses flags:
- `-width` (default: 21) — must be odd for the algorithm to work correctly
- `-height` (default: 21) — must be odd
- `-seed` (default: current Unix time) — for reproducible mazes

Creates a `Maze`, calls `Generate()`, then `Print()`.

### `maze.go`
Defines a `Cell` struct:
```go
type Cell struct {
    Visited bool
    Walls   [4]bool // North, East, South, West — all true initially
}
```

Defines the `Maze` struct:
```go
type Maze struct {
    Width  int
    Height int
    Grid   [][]Cell
    Rand   *rand.Rand
}
```

`NewMaze(w, h int, seed int64) *Maze` allocates the grid and initializes all walls to `true`.

`Generate(x, y int)` implements recursive backtracking:
1. Mark cell `(x, y)` as visited
2. Create a slice of the 4 directions (N, E, S, W) as `[2]int` offsets: `{0,-2}, {2,0}, {0,2}, {-2,0}`
3. Shuffle the directions with `m.Rand.Shuffle`
4. For each direction `(dx, dy)`:
   - Compute neighbor: `(nx, ny) = (x+dx, y+dy)`
   - If `(nx, ny)` is in bounds and not visited:
     - Remove the wall between `(x,y)` and `(nx,ny)` — this is the cell at `(x+dx/2, y+dy/2)`
     - Recurse: `Generate(nx, ny)`

The grid uses odd coordinates for cells and even coordinates for walls. Cells sit at odd `(x,y)`, and the passage between two cells is at the midpoint.

### `display.go`
`Print(m *Maze)` renders the maze:
- Iterate over every position in the grid
- If it's a wall cell (`Walls` are intact or it's a border): print `"█"`
- If it's a passage (wall was removed): print `" "`
- Add entrance at top-left and exit at bottom-right

---

## How It Works

```
$ go run . -width 21 -height 11

█████████████████████
█ █     █         █ █
█ █ ███ █ █████ ███ █
█     █   █   █     █
█████ █████ █ █████ █
█         █ █   █   █
█ ███████ █ ███ █ ███
█ █     █ █     █   █
█ █ ███ █ █████ █ ███
█   █           █   █
█████████████████████
```

---

## Data Flow

```
NewMaze() → Generate(1,1) [recursive DFS] → Print()
```

---

## The Recursive Backtracking Algorithm

This is a classic algorithm. The key insight is moving 2 cells at a time — each "step" in the maze corresponds to moving 2 positions in the grid, with the wall between them at position 1.

```
Start at (1,1)
Mark visited
Shuffle directions: [East, North, South, West]
Try East → (3,1) not visited → remove wall at (2,1) → recurse into (3,1)
  Try North → out of bounds → skip
  Try East → (5,1) not visited → recurse...
  ...
  All neighbors visited → backtrack
```

---

## Suggested Extensions

- Add a solver that finds the shortest path using BFS and marks it with `.`
- Save the maze to a PNG image using the `image` package
- Add a `-shape` flag to carve non-rectangular mazes (circular, etc.)
- Animate the generation process in the terminal using `time.Sleep`
