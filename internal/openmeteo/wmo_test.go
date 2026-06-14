package openmeteo

import "testing"

func TestCondition(t *testing.T) {
	tests := []struct {
		name string
		code int
		want string
	}{
		{"clear sky", 0, "clear"},
		{"partly cloudy", 2, "partly cloudy"},
		{"fog", 45, "fog"},
		{"drizzle", 53, "drizzle"},
		{"rain", 61, "rain"},
		{"heavy rain", 65, "heavy rain"},
		{"snow", 73, "snow"},
		{"rain showers", 80, "rain showers"},
		{"thunderstorm", 95, "thunderstorm"},
		{"unmapped code", 7, "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Condition(tt.code); got != tt.want {
				t.Errorf("Condition(%d) = %q, want %q", tt.code, got, tt.want)
			}
		})
	}
}
