# Unit Converter

A CLI tool to convert units of weight and length (e.g., kg → lbs, miles → km).

---

## What You'll Learn

- Using `map[string]float64` to store conversion factors
- Structuring logic with `switch` statements
- Accepting and validating CLI flags with the `flag` package
- Organizing code across multiple files
- Returning custom error types

---

## Project Structure

```
unit-converter/
├── main.go
├── length.go
└── weight.go
```

### `main.go`
Defines three flags using the `flag` package:
- `-value` — the number to convert (float64)
- `-from` — source unit (string, e.g. `"km"`)
- `-to` — target unit (string, e.g. `"miles"`)

Calls `flag.Parse()`, then routes to the correct conversion module.

### `length.go`
Defines a `map[string]float64` of all length units expressed in meters (the base unit). For example: `"km": 1000.0`, `"miles": 1609.344`, `"feet": 0.3048`. Conversion is done by converting the input to meters first, then to the output unit.

### `weight.go`
Same pattern as `length.go` but with weight units expressed in grams as the base unit. For example: `"kg": 1000.0`, `"lbs": 453.592`, `"oz": 28.3495`.

---

## How It Works

1. User runs: `go run . -value 10 -from km -to miles`
2. `main.go` parses the flags
3. Both `from` and `to` units are looked up in the same map
4. Value is converted: `result = value * (fromFactor / toFactor)`
5. Result is printed:

```
10 km = 6.214 miles
```

---

## Conversion Formula

```
result = inputValue * (factorOfFromUnit / factorOfToUnit)
```

For example: `10 km → miles`
```
10 * (1000 / 1609.344) = 6.214 miles
```

---

## Suggested Extensions

- Add temperature units (requires offset math, not just multiplication)
- Add volume units in a `volume.go` file
- Print a list of all available units when `-list` flag is passed
- Allow unit aliases (e.g., `"kilometer"` maps to `"km"`)
