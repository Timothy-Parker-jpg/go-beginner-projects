# Directory Tree Printer

A CLI tool that prints a visual tree of a directory's contents — like the Unix `tree` command.

---

## What You'll Learn

- Walking a directory with `os.ReadDir`
- Writing recursive functions in Go
- Passing and threading state (prefix strings) through recursion
- Sorting directory entries
- Using `filepath.Join` for cross-platform paths

---

## Project Structure

```
dir-tree/
├── main.go
└── tree.go
```

### `main.go`
Reads the target path from `os.Args` (defaults to `.` if none given). Accepts optional flags:
- `-depth N` — max depth to traverse (default: unlimited)
- `-hidden` — include hidden files/directories (those starting with `.`)

Calls `PrintTree()` and prints a summary of total dirs and files.

### `tree.go`
The core recursive function:
```go
func PrintTree(path string, prefix string, depth int, maxDepth int, showHidden bool, counts *Counts)
```

Logic:
1. Call `os.ReadDir(path)` to get directory entries (already sorted alphabetically)
2. Filter out hidden entries if `showHidden` is false
3. For each entry, determine if it's the last item in the list
4. If last: use `└──` as the connector; otherwise use `├──`
5. Print `prefix + connector + name`
6. If the entry is a directory and depth < maxDepth: recurse with updated prefix
   - If last item: append `"    "` to prefix (4 spaces)
   - Otherwise: append `"│   "` to prefix

The `Counts` struct tracks total files and directories found.

---

## How It Works

```
$ go run . ./myproject

myproject
├── main.go
├── go.mod
└── internal
    ├── config.go
    └── handlers
        ├── auth.go
        └── user.go

3 directories, 5 files
```

---

## The Prefix Logic (Key Concept)

The visual indentation is built by threading a `prefix` string through recursion:

```
"" + "├── " → first item
"│   " + "├── " → nested first item (parent was NOT last)
"    " + "└── " → nested last item (parent WAS last)
```

This is the core algorithm to understand in this project.

---

## Suggested Extensions

- Add a `-pattern` flag to filter by file extension (e.g., `*.go`)
- Add file sizes next to file names
- Add color output (dirs in blue, files in white) using ANSI codes
- Count total disk usage
