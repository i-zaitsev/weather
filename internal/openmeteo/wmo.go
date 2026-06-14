package openmeteo

var wmoConditions = map[int]string{
	0:  "clear",
	1:  "mostly clear",
	2:  "partly cloudy",
	3:  "cloudy",
	45: "fog",
	48: "rime fog",
	51: "drizzle",
	53: "drizzle",
	55: "drizzle",
	56: "freezing drizzle",
	57: "freezing drizzle",
	61: "rain",
	63: "rain",
	65: "heavy rain",
	66: "freezing rain",
	67: "freezing rain",
	71: "snow",
	73: "snow",
	75: "heavy snow",
	77: "snow grains",
	80: "rain showers",
	81: "rain showers",
	82: "violent rain showers",
	85: "snow showers",
	86: "snow showers",
	95: "thunderstorm",
	96: "thunderstorm",
	99: "thunderstorm",
}

// Condition returns a short human-readable string for a WMO weather code.
func Condition(code int) string {
	if s, ok := wmoConditions[code]; ok {
		return s
	}
	return "unknown"
}
