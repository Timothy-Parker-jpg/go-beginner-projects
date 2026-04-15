# Unit Converter — Code Review

This review follows the instructions in your learner_prompt: lead with the most important issue, evaluate across four dimensions (Correctness, Code quality, Professional standards, Performance & design), show concrete before/after fixes, and end with a single takeaway.

**Most important issue**
- Problem: CLI output and control flow produce confusing results and non-zero exits on success.
- Why it matters: Users and automation (CI, scripts) expect clear output and proper exit codes. Printing struct values instead of unit names and always exiting 1 cause misinterpretation and failing automation pipelines.

**Correctness**
- **Symptom**: Several code paths print `UnitDef` structs (e.g., `%v` with structs) instead of canonical unit strings.
- **Edge cases**: Unit lookup is case-sensitive and whitespace-sensitive (e.g., `Gal` or ` gal` may fail). Temperature conversions handle supported units but return a generic error message for unsupported units.
- **Silent failures**: `main.go` previously exits with code `1` even on successful conversions which looks like an error to callers.

**Code quality**
- **Names**: Types `UnitType`, `UnitDef` are clear. Some helper names (like `normalizeUnit`) are fine but should document expectations (input casing, trimming).
- **Duplication**: Conversion print logic repeats similar patterns across Length/Weight/Volume — extract a small helper to format/print conversions if desired.
- **Readability**: Use `.Unit` when printing; prefer explicit format strings and avoid printing whole structs.

**Professional standards**
- **Error handling**: Errors from `convert` should be returned to `main` and cause non-zero exit; success should exit `0`.
- **CLI UX**: Improve `flag.Usage` and error messages. Example: when `-from` or `-to` are missing, show usage rather than a terse error.
- **Tests**: No unit tests present. Add table-driven tests for conversions (length/weight/volume) and for `normalizeUnit` behavior.

**Performance & design**
- **Complexity**: All conversions are O(1) map lookups — perfectly adequate.
- **Extensibility**: Current alias→canonical→factor mapping is a simple registry. For larger sets, consider a registration API or JSON/YAML-backed registry.

---

**Concrete fixes (before / after)**

1) Fix `Volume` UnitType typo

Before (`units.go`):
```go
const (
    Length UnitType = "length"
    Weight UnitType = "weight"
    Temp   UnitType = "temp"
    Volume UnitType = "volumne"
)
```

After:
```go
const (
    Length UnitType = "length"
    Weight UnitType = "weight"
    Temp   UnitType = "temp"
    Volume UnitType = "volume"
)
```

Why: Typos in type labels can cause logic branches to miss matches and are a source of subtle bugs.

2) Make unit normalization robust (trim + lowercase)

Before (`convert.go` / `normalizeUnit`):
```go
func normalizeUnit(unit string, aliasMap map[string]UnitDef) (UnitDef, bool) {
    normalizedUnit, ok := aliasMap[unit]
    return normalizedUnit, ok
}
```

After:
```go
import "strings"

func normalizeUnit(unit string, aliasMap map[string]UnitDef) (UnitDef, bool) {
    key := strings.ToLower(strings.TrimSpace(unit))
    normalizedUnit, ok := aliasMap[key]
    return normalizedUnit, ok
}
```

Why: Makes CLI forgiving of casing/whitespace, improving UX and reducing surprises.

3) Print canonical unit strings rather than structs

Before (example in conversion):
```go
format := "%.2f %s = %.2f %s\n"
convertedValue := *value * lengthToMeter[normFrom.Unit] / lengthToMeter[normTo.Unit]
fmt.Printf(format, *value, normFrom, convertedValue, normTo)
```

After:
```go
format := "%.2f %s = %.2f %s\n"
convertedValue := *value * lengthToMeter[normFrom.Unit] / lengthToMeter[normTo.Unit]
fmt.Printf(format, *value, normFrom.Unit, convertedValue, normTo.Unit)
```

Why: Printing the `.Unit` field shows the canonical unit name (e.g., "pint"), not the struct representation.

4) Correct CLI exit codes and messaging

Before (`main.go`):
```go
if *value <= 0 {
    fmt.Println("Error: -value must use a value greater than 1.")
    os.Exit(1)
}
// ...
err := convert(from, to, value)
fmt.Println("Conversion Complete")
if err != nil {
    fmt.Println(fmt.Errorf("ERROR: Convert() : %w", err))
}
os.Exit(1)
```

After:
```go
if *value <= 0 {
    fmt.Println("Error: -value must be greater than 0.")
    os.Exit(1)
}
// ...
err := convert(from, to, value)
if err != nil {
    fmt.Println(fmt.Errorf("ERROR: Convert() : %w", err))
    os.Exit(1)
}
fmt.Println("Conversion Complete")
os.Exit(0)
```

Why: Exit `0` on success; return non-zero when errors occur. Also clarify validation message so users know permitted values.

---

**Suggested tests (high priority)**
- Table-driven tests for `ConvertLength` and `ConvertWeight`:
  - Cases: mm→m, km→m, ft→m, lb→kg, oz→g.
  - Check numeric tolerance (e.g., within 1e-6).
- Tests for `TempConvert` with known values (0°C → 32°F; 273.15K → 0°C).
- Tests for `normalizeUnit` (cases: "L", " l ", "Gallon", "GAL").
- Integration test: build binary, run common example and assert stdout/exit code.

---

**Small refactor suggestions (optional)**
- Extract a `formatConversion(value, fromUnit, toUnit string, converted float64)` helper to reduce repetition.
- Centralize alias normalization: provide a small helper `cleanUnit(input string) string` used by CLI parsing and tests.
- Consider exposing a programmatic API (package functions) separate from CLI so unit tests can call conversion logic directly without spawning subprocesses.

---

**Next steps (recommendation)
- Apply the concrete fixes above in the code (you indicated you'll edit files yourself).
- Add the tests described and run `go test ./...`.
- Improve `flag.Usage` to show examples when users pass wrong flags.

**Single takeaway**
- Make unit normalization and output explicit: always normalize input deterministically (trim + lowercase), and print canonical unit names — this fixes the majority of UX and correctness problems in this tool.

---

If you'd like, I can also prepare a small set of table-driven Go tests and/or patch files that you can apply directly; say whether you want `tests only` or `patches + tests` and I will prepare them for you.
