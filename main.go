package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/i-zaitsev/weather/internal/openmeteo"
)

const defaultCity = "Hamburg"

func main() {
	day := flag.String("d", "today", "day to report: today or tomorrow")
	more := flag.Bool("more", false, "print hourly lines")
	version := flag.Bool("version", false, "print version and attribution")
	flag.Parse()

	if *version {
		fmt.Println("weather - Weather data by Open-Meteo.com")
		return
	}

	if *day != "today" && *day != "tomorrow" {
		fmt.Fprintf(os.Stderr, "invalid -d %q: want today or tomorrow\n", *day)
		os.Exit(2)
	}

	city := defaultCity
	if args := flag.Args(); len(args) > 0 {
		city = strings.Join(args, " ")
	}

	if err := run(city, *day, *more); err != nil {
		fmt.Fprintln(os.Stderr, "weather:", err)
		os.Exit(1)
	}
}

func run(city, day string, more bool) error {
	c := openmeteo.NewClient()

	loc, err := c.Geocode(city)
	if err != nil {
		return err
	}
	f, err := c.GetForecast(loc)
	if err != nil {
		return err
	}

	idx := 0
	if day == "tomorrow" {
		idx = 1
	}
	if idx >= len(f.Daily.Time) {
		return fmt.Errorf("no forecast for %s", day)
	}

	fmt.Println(summary(f, day, idx))

	if more {
		for _, line := range hourlyLines(f, f.Daily.Time[idx]) {
			fmt.Println(line)
		}
	}
	return nil
}

func summary(f openmeteo.Forecast, day string, idx int) string {
	var temp float64
	var code, prob int

	if day == "today" {
		temp = f.Current.Temperature
		code = f.Current.WeatherCode
	} else {
		temp = f.Daily.TemperatureMax[idx]
		code = f.Daily.WeatherCode[idx]
	}
	if idx < len(f.Daily.PrecipitationProbabilityMax) {
		prob = f.Daily.PrecipitationProbabilityMax[idx]
	}

	return fmt.Sprintf("%s %s %d%%", formatTemp(temp), openmeteo.Condition(code), prob)
}

func hourlyLines(f openmeteo.Forecast, date string) []string {
	var lines []string
	for i, t := range f.Hourly.Time {
		if !strings.HasPrefix(t, date) {
			continue
		}
		lines = append(lines, fmt.Sprintf("%s %s %s %d%%",
			clock(t),
			formatTemp(f.Hourly.Temperature[i]),
			openmeteo.Condition(f.Hourly.WeatherCode[i]),
			f.Hourly.PrecipitationProbability[i],
		))
	}
	return lines
}

func formatTemp(t float64) string {
	return fmt.Sprintf("%+d°C", int(math.Round(t)))
}

func clock(t string) string {
	if _, hm, ok := strings.Cut(t, "T"); ok {
		return hm
	}
	return t
}
