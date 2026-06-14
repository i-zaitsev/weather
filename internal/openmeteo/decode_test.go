package openmeteo

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestForecastDecode(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("testdata", "forecast.json"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	var f Forecast
	if err := json.Unmarshal(data, &f); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if f.Timezone != "Europe/Berlin" {
		t.Errorf("Timezone = %q, want Europe/Berlin", f.Timezone)
	}
	if f.Current.Temperature != 26.7 {
		t.Errorf("Current.Temperature = %v, want 26.7", f.Current.Temperature)
	}
	if f.Current.WeatherCode != 61 {
		t.Errorf("Current.WeatherCode = %d, want 61", f.Current.WeatherCode)
	}
	if got := len(f.Hourly.Time); got != 3 {
		t.Fatalf("len(Hourly.Time) = %d, want 3", got)
	}
	if f.Hourly.PrecipitationProbability[0] != 20 {
		t.Errorf("Hourly.PrecipitationProbability[0] = %d, want 20", f.Hourly.PrecipitationProbability[0])
	}
	if got := len(f.Daily.Time); got != 2 {
		t.Fatalf("len(Daily.Time) = %d, want 2", got)
	}
	if f.Daily.TemperatureMax[1] != 22.1 {
		t.Errorf("Daily.TemperatureMax[1] = %v, want 22.1", f.Daily.TemperatureMax[1])
	}
	if f.Daily.PrecipitationProbabilityMax[1] != 60 {
		t.Errorf("Daily.PrecipitationProbabilityMax[1] = %d, want 60", f.Daily.PrecipitationProbabilityMax[1])
	}
}
