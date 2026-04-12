# Weather CLI

A CLI app that fetches and displays current weather for a given city using the Open-Meteo API (free, no API key required).

---

## What You'll Learn

- Making HTTP GET requests with `net/http`
- Decoding JSON responses with `encoding/json`
- Defining structs that match JSON shapes
- Chaining API calls (geocoding → weather)
- Using `url.Values` to build query strings safely
- Handling HTTP errors and non-200 status codes

---

## Project Structure

```
weather-cli/
├── main.go
├── geocode.go
├── weather.go
└── display.go
```

### `main.go`
Reads the city name from `os.Args`. Calls `Geocode()` to convert the city to coordinates, then calls `GetWeather()` with those coordinates. Passes the result to `Display()`.

### `geocode.go`
Uses the Open-Meteo Geocoding API:
```
GET https://geocoding-api.open-meteo.com/v1/search?name=London&count=1
```

Defines a struct matching the JSON response shape. Decodes with `json.NewDecoder(resp.Body).Decode(&result)`. Returns latitude, longitude, and the resolved city name.

### `weather.go`
Uses the Open-Meteo Weather API:
```
GET https://api.open-meteo.com/v1/forecast
    ?latitude=51.5&longitude=-0.1
    &current=temperature_2m,wind_speed_10m,weathercode
    &temperature_unit=celsius
```

Builds the URL with `url.Values{}` then calls `.Encode()`. Defines a `WeatherResponse` struct with nested `Current` struct matching the JSON. Returns the decoded weather data.

### `display.go`
Maps WMO weather codes (integers) to descriptions using a `map[int]string`:
- 0 → "Clear sky"
- 1, 2, 3 → "Partly cloudy"
- 61, 63, 65 → "Rain"
- etc.

Prints a clean, formatted weather card to the terminal.

---

## How It Works

```
$ go run . "New York"

📍 New York, United States
🌡  22°C
💨 14 km/h wind
⛅ Partly cloudy
```

---

## Data Flow

```
city name → Geocode() → lat/lon → GetWeather() → WeatherResponse → Display()
```

---

## API Reference

Both APIs used are from [open-meteo.com](https://open-meteo.com) — completely free with no API key needed.

---

## Suggested Extensions

- Add a `-days N` flag to show a multi-day forecast
- Cache results to a local file for 10 minutes to avoid repeat API calls
- Support `-unit imperial` to switch to Fahrenheit and mph
- Show a colored temperature (blue for cold, red for hot) using ANSI codes
