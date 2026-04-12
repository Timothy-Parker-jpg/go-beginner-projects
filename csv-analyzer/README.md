# CSV Analyzer

A CLI tool that reads a CSV file and prints summary statistics: row count, column names, and numeric column averages.

---

## What You'll Learn

- Parsing CSV files with `encoding/csv`
- Using `strconv.ParseFloat` to detect and convert numeric columns
- Working with maps (`map[string][]float64`) to group data by column
- Computing min, max, and average from a float slice
- Printing a formatted summary report

---

## Project Structure

```
csv-analyzer/
├── main.go
├── parser.go
└── stats.go
```

### `main.go`
Accepts a file path as the first CLI argument. Opens the file, passes the reader to `ParseCSV()`, then calls `PrintSummary()` on the result.

### `parser.go`
Defines a `CSVData` struct:
```go
type CSVData struct {
    Headers []string
    Rows    [][]string
    NumericCols map[string][]float64
}
```

The `ParseCSV(r io.Reader) (CSVData, error)` function:
1. Uses `csv.NewReader(r).ReadAll()` to load all rows
2. The first row becomes `Headers`
3. For each column, attempts to parse every value as float64. If ALL values in a column parse successfully, adds them to `NumericCols`
4. Returns the populated `CSVData`

### `stats.go`
Defines `Stats` struct with Min, Max, Avg, Count fields. The `ComputeStats(values []float64) Stats` function iterates once to find min/max and accumulate sum, then divides for avg.

`PrintSummary(data CSVData)` prints: total rows, column names, and a stats table for numeric columns.

---

## How It Works

Given `sales.csv`:
```
name,region,revenue,units
Alice,North,5200.50,42
Bob,South,3100.00,28
Carol,North,7800.75,61
```

```
$ go run . sales.csv

Rows: 3
Columns: name, region, revenue, units

Numeric columns:
  revenue  →  min: 3100.00  max: 7800.75  avg: 5367.08
  units    →  min: 28.00    max: 61.00    avg: 43.67
```

---

## Data Flow

```
file path → open → csv.Reader → ParseCSV() → CSVData → ComputeStats() → PrintSummary()
```

---

## Suggested Extensions

- Filter rows by a column value: `-filter region=North`
- Export a summary report to a new CSV
- Sort rows by a numeric column
- Handle CSV files with missing/empty values gracefully
