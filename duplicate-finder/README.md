# Duplicate File Finder

A CLI tool that scans a directory and finds files with identical content by comparing their checksums.

---

## What You'll Learn

- Walking directory trees with `filepath.WalkDir`
- Hashing file content with `crypto/md5` or `crypto/sha256`
- Grouping values by key with `map[string][]string`
- Working with `io.Copy` to stream file content
- Formatting file sizes in human-readable units

---

## Project Structure

```
duplicate-finder/
├── main.go
├── walker.go
└── hasher.go
```

### `main.go`
Accepts a root directory path and optional flags:
- `-min-size N` — ignore files smaller than N bytes (default: 1)
- `-delete` — delete duplicate files (keep the first found), prompt for confirmation

Calls `FindDuplicates()` and prints results grouped by hash.

### `walker.go`
The `WalkFiles(root string, minSize int64) ([]string, error)` function:
1. Uses `filepath.WalkDir` with a callback function
2. The callback receives a `fs.DirEntry` — skip directories and files smaller than `minSize`
3. Appends the full file path to a results slice
4. Returns all discovered file paths

### `hasher.go`
The `HashFile(path string) (string, error)` function:
1. Opens the file with `os.Open`
2. Creates a `sha256.New()` hasher
3. Uses `io.Copy(hasher, file)` to stream the file content through the hasher — this avoids loading the whole file into memory
4. Returns `fmt.Sprintf("%x", hasher.Sum(nil))` as the hex hash string

Back in `main.go`, build a `map[string][]string` where key = hash, value = list of paths. Any key with 2+ values is a duplicate group.

---

## How It Works

```
$ go run . ~/Downloads

Found 3 groups of duplicate files:

[a3f9c2...] 2 files, 4.2 MB each
  /Downloads/report.pdf
  /Downloads/report_copy.pdf

[7b81d4...] 3 files, 128 KB each
  /Downloads/photo.jpg
  /Downloads/photos/photo.jpg
  /Downloads/backup/photo.jpg
```

---

## Data Flow

```
root path → WalkFiles() → []string paths → HashFile() each → map[hash][]paths → print duplicates
```

---

## Why Stream with io.Copy?

Loading a 2GB file with `os.ReadFile` would require 2GB of RAM. `io.Copy(hasher, file)` streams the file in chunks — the hasher processes each chunk and discards it, keeping memory usage constant regardless of file size.

---

## Suggested Extensions

- First compare file sizes before hashing (files with different sizes can't be duplicates — skip hashing them)
- Add a `-dry-run` flag to preview deletions without executing them
- Show total disk space recoverable by deleting duplicates
- Sort duplicate groups by wasted space (largest first)
