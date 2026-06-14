# weather

A minimal, standard-library-only CLI that prints the weather for a city using
the [Open-Meteo](https://open-meteo.com) APIs.

## Build

```bash
go build ./...
```

## Usage

```bash
weather [city] [-d today|tomorrow] [-more]
```

- Default city is `Hamburg`.
- The default output is a one-line summary, e.g. `+27°C rain 20%`.
- `-d tomorrow` reports tomorrow instead of today.
- `-more` additionally prints one line per hour for the selected day.
- `-version` prints version and data attribution.

```bash
weather hamburg
weather hamburg -d tomorrow
weather -more
```

## Test

```bash
go test ./...
```

Tests run fully offline against a committed JSON fixture.

Weather data by Open-Meteo.com (CC BY 4.0).
