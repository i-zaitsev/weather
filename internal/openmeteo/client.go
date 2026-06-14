package openmeteo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	geocodeURL  = "https://geocoding-api.open-meteo.com/v1/search"
	forecastURL = "https://api.open-meteo.com/v1/forecast"
	timezone    = "Europe/Berlin"
)

// Client talks to the Open-Meteo REST APIs.
type Client struct {
	http *http.Client
}

// NewClient returns a Client with a 5s request timeout.
func NewClient() *Client {
	return &Client{http: &http.Client{Timeout: 5 * time.Second}}
}

// Location is a geocoded place.
type Location struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Forecast is the decoded forecast response.
type Forecast struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Current   Current `json:"current"`
	Hourly    Hourly  `json:"hourly"`
	Daily     Daily   `json:"daily"`
}

// Current holds the current conditions.
type Current struct {
	Time          string  `json:"time"`
	Temperature   float64 `json:"temperature_2m"`
	WeatherCode   int     `json:"weather_code"`
	Precipitation float64 `json:"precipitation"`
}

// Hourly holds the hourly series.
type Hourly struct {
	Time                     []string  `json:"time"`
	Temperature              []float64 `json:"temperature_2m"`
	PrecipitationProbability []int     `json:"precipitation_probability"`
	WeatherCode              []int     `json:"weather_code"`
}

// Daily holds the daily series.
type Daily struct {
	Time                        []string  `json:"time"`
	TemperatureMax              []float64 `json:"temperature_2m_max"`
	TemperatureMin              []float64 `json:"temperature_2m_min"`
	PrecipitationProbabilityMax []int     `json:"precipitation_probability_max"`
	WeatherCode                 []int     `json:"weather_code"`
}

// Geocode resolves a city name to a single Location.
func (c *Client) Geocode(city string) (Location, error) {
	q := url.Values{}
	q.Set("name", city)
	q.Set("count", "1")

	var body struct {
		Results []Location `json:"results"`
	}
	if err := c.getJSON(geocodeURL+"?"+q.Encode(), &body); err != nil {
		return Location{}, fmt.Errorf("geocode %q: %w", city, err)
	}
	if len(body.Results) == 0 {
		return Location{}, fmt.Errorf("geocode %q: no results", city)
	}
	return body.Results[0], nil
}

// GetForecast fetches the forecast for a location.
func (c *Client) GetForecast(loc Location) (Forecast, error) {
	q := url.Values{}
	q.Set("latitude", fmt.Sprintf("%g", loc.Latitude))
	q.Set("longitude", fmt.Sprintf("%g", loc.Longitude))
	q.Set("current", "temperature_2m,weather_code,precipitation")
	q.Set("hourly", "temperature_2m,precipitation_probability,weather_code")
	q.Set("daily", "temperature_2m_max,temperature_2m_min,precipitation_probability_max,weather_code")
	q.Set("timezone", timezone)

	var f Forecast
	if err := c.getJSON(forecastURL+"?"+q.Encode(), &f); err != nil {
		return Forecast{}, fmt.Errorf("forecast: %w", err)
	}
	return f, nil
}

func (c *Client) getJSON(u string, dst any) error {
	resp, err := c.http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
		return fmt.Errorf("decode: %w", err)
	}
	return nil
}
