# Temperature Converter

A CLI tool to convert temperatures between Celsius, Fahrenheit, and Kelvin.

---

## What You'll Learn

- Reading command-line arguments with `os.Args`
- Parsing strings to floats with `strconv.ParseFloat`
- Writing functions with multiple return values
- Basic error handling with `fmt.Errorf`
- Printing formatted output with `fmt.Printf`

---

## Project Structure

```
temp-converter/
├── main.go
└── converter.go
```

### `main.go`
Entry point. Reads `os.Args` to get the temperature value and unit (e.g., `100 C`). Validates that exactly two arguments are passed, then calls the conversion functions and prints all three results.

### `converter.go`
Contains three pure conversion functions:
- `CelsiusToFahrenheit(c float64) float64`
- `CelsiusToKelvin(c float64) float64`
- A dispatcher function `Convert(value float64, unit string) (float64, float64, float64, error)` that routes to the correct conversion path depending on the input unit.

---

## How It Works

1. User runs: `go run . 100 C`
2. `main.go` parses `"100"` → `float64`, and reads `"C"` as the source unit
3. `Convert()` is called — it converts to all three units
4. Results are printed:

```
100.00°C = 212.00°F = 373.15K
```

---

## Data Flow

```
os.Args → parse string → float64 → Convert() → print results
```

---

## Suggested Extensions

- Accept lowercase unit names (`c`, `f`, `k`)
- Add a `--from` and `--to` flag using the `flag` package
- Add a loop so the user can convert multiple values interactively
- Write unit tests in `converter_test.go`
