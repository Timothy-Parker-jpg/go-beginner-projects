# To-Do List CLI

A command-line to-do list app that saves tasks to a local JSON file so they persist between runs.

---

## What You'll Learn

- Defining structs and working with struct slices
- Reading and writing JSON files with `encoding/json`
- Using `os.ReadFile` and `os.WriteFile`
- Routing sub-commands (`add`, `list`, `done`, `delete`) via `os.Args`
- Formatting console output with `fmt.Printf` and padding

---

## Project Structure

```
todo-cli/
├── main.go
├── todo.go
└── storage.go
```

### `main.go`
Reads the first argument as a sub-command and dispatches:
- `add "Buy milk"` → calls `AddTodo()`
- `list` → calls `ListTodos()`
- `done 2` → calls `MarkDone(id)`
- `delete 2` → calls `DeleteTodo(id)`

Prints usage instructions if no sub-command is given.

### `todo.go`
Defines the `Todo` struct:
```go
type Todo struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
    CreatedAt string `json:"created_at"`
}
```
Functions: `AddTodo()`, `ListTodos()`, `MarkDone()`, `DeleteTodo()`. Each function loads todos from disk, modifies the slice, and saves back to disk.

### `storage.go`
Two functions:
- `LoadTodos() ([]Todo, error)` — reads `todos.json`, unmarshals JSON into `[]Todo`. Returns empty slice if file doesn't exist yet.
- `SaveTodos(todos []Todo) error` — marshals to indented JSON and writes to `todos.json`.

---

## How It Works

```
$ go run . add "Buy groceries"
✓ Added: "Buy groceries" (ID: 1)

$ go run . list
[ ] 1 - Buy groceries
[ ] 2 - Walk the dog

$ go run . done 1
✓ Marked as done: "Buy groceries"

$ go run . list
[✓] 1 - Buy groceries
[ ] 2 - Walk the dog
```

---

## Data Flow

```
os.Args → route command → load todos.json → modify []Todo → save todos.json
```

---

## Suggested Extensions

- Add a `clear` command to delete all completed tasks
- Add due dates to the `Todo` struct
- Sort tasks by creation date or completion status
- Color completed tasks differently using ANSI escape codes
