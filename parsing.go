package omgo

import (
	"encoding/json"
	"time"
)

type ForecastJSON struct {
	Latitude       float64
	Longitude      float64
	Elevation      float64
	GenerationTime float64                    `json:"generationtime_ms"`
	CurrentWeather CurrentWeather             `json:"current_weather"`
	HourlyUnits    map[string]string          `json:"hourly_units"`
	HourlyMetrics  map[string]json.RawMessage `json:"hourly"` // Parsed later, the API returns both Time and floats here
}

type Forecast struct {
	Latitude       float64
	Longitude      float64
	Elevation      float64
	GenerationTime float64
	CurrentWeather CurrentWeather
	HourlyUnits    map[string]string
	Metrics        map[string][]float64 // Parsed from ForecastJSON.HourlyMetrics
	Hours          []time.Time          // Parsed from ForecastJSON.HourlyMetrics
}

type CurrentWeather struct {
	Temperature   float64
	Time          ApiTime
	WeatherCode   int
	WindDirection int
	WindSpeed     float64
}

// ParseBody converts the API response body into a Forecast struct
// Rationale: The API returns a map with both times as well as floats, this function
// unmarshalls in 2 steps in order to not return a map[string][]interface{}
func ParseBody(body []byte) (*Forecast, error) {
	f := &ForecastJSON{}
	err := json.Unmarshal(body, f)
	if err != nil {
		return nil, err
	}

	fc := &Forecast{
		Latitude:       f.Latitude,
		Longitude:      f.Longitude,
		Elevation:      f.Elevation,
		GenerationTime: f.GenerationTime,
		CurrentWeather: f.CurrentWeather,
		HourlyUnits:    f.HourlyUnits,
		Hours:          []time.Time{},
		Metrics:        make(map[string][]float64),
	}

	for k, v := range f.HourlyMetrics {
		if k == "time" {
			// We unmarshal into an ApiTime array because of the custom formatting
			// of the timestamp in the API response
			target := []ApiTime{}
			err := json.Unmarshal(v, &target)
			if err != nil {
				return nil, err
			}

			for _, at := range target {
				fc.Hours = append(fc.Hours, at.Time)
			}

			continue
		}
		target := []float64{}
		err := json.Unmarshal(v, &target)
		if err != nil {
			return nil, err
		}
		fc.Metrics[k] = target
	}

	return fc, nil
}
