package omgo

import (
	"encoding/json"
	"time"
)

type ForecastJSON struct {
	Latitude       float64
	Longitude      float64
	Timezone       string
	Elevation      float64
	GenerationTime float64                    `json:"generationtime_ms"`
	CurrentWeather CurrentWeather             `json:"current_weather"`
	HourlyUnits    map[string]string          `json:"hourly_units"`
	HourlyMetrics  map[string]json.RawMessage `json:"hourly"` // Parsed later, the API returns both Time and floats here
	DailyUnits     map[string]string          `json:"daily_units"`
	DailyMetrics   map[string]json.RawMessage `json:"daily"` // Parsed later, the API returns both Time and floats here

}

type Forecast struct {
	Latitude       float64
	Longitude      float64
	Elevation      float64
	GenerationTime float64
	CurrentWeather CurrentWeather
	HourlyUnits    map[string]string
	HourlyMetrics  map[string][]float64 // Parsed from ForecastJSON.HourlyMetrics
	HourlyTimes    []time.Time          // Parsed from ForecastJSON.HourlyMetrics
	DailyUnits     map[string]string
	DailyMetrics   map[string][]float64 // Parsed from ForecastJSON.DailyMetrics
	DailyTimes     []time.Time          // Parsed from ForecastJSON.DailyMetrics
}

type CurrentWeather struct {
	IsDay         float64 `json:"is_day"`
	Temperature   float64
	Time          ApiTime
	WeatherCode   float64
	WindDirection float64
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

	loc := time.UTC
	if f.Timezone != "" {
		var err error
		loc, err = time.LoadLocation(f.Timezone)
		if err != nil {
			return nil, err
		}
	}

	fc := &Forecast{
		Latitude:       f.Latitude,
		Longitude:      f.Longitude,
		Elevation:      f.Elevation,
		GenerationTime: f.GenerationTime,
		CurrentWeather: f.CurrentWeather,
		HourlyUnits:    f.HourlyUnits,
		HourlyTimes:    []time.Time{},
		HourlyMetrics:  make(map[string][]float64),
		DailyUnits:     f.DailyUnits,
		DailyTimes:     []time.Time{},
		DailyMetrics:   make(map[string][]float64),
	}

	// Fix the timezone for the current weather timestamp, if needed
	if f.Timezone != "UTC" {
		y, m, d := f.CurrentWeather.Time.Date()
		f.CurrentWeather.Time = ApiTime{
			Time: time.Date(y, m, d, f.CurrentWeather.Time.Hour(), f.CurrentWeather.Time.Minute(), f.CurrentWeather.Time.Second(), f.CurrentWeather.Time.Nanosecond(), loc)}
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
				// We can fix the timezone now: by default time.Time is UTC when unmarshalled,
				// but we now know what the location should be, based on the response
				if f.Timezone != "UTC" {
					y, m, d := at.Date()
					adjusted := time.Date(y, m, d, at.Hour(), at.Minute(), at.Second(), at.Nanosecond(), loc)
					fc.HourlyTimes = append(fc.HourlyTimes, adjusted)
				} else {
					fc.HourlyTimes = append(fc.HourlyTimes, at.Time)
				}
			}

			continue
		}
		target := []float64{}
		err := json.Unmarshal(v, &target)
		if err != nil {
			return nil, err
		}
		fc.HourlyMetrics[k] = target
	}

	for k, v := range f.DailyMetrics {
		if k == "time" {
			// We unmarshal into an ApiTime array because of the custom formatting
			// of the timestamp in the API response
			target := []ApiDate{}
			err := json.Unmarshal(v, &target)
			if err != nil {
				return nil, err
			}

			for _, at := range target {
				// We can fix the timezone now: by default time.Time is UTC when unmarshalled,
				// but we now know what the location should be, based on the response
				if f.Timezone != "UTC" {
					y, m, d := at.Date()
					adjusted := time.Date(y, m, d, at.Hour(), at.Minute(), at.Second(), at.Nanosecond(), loc)
					fc.DailyTimes = append(fc.DailyTimes, adjusted)
				} else {
					fc.DailyTimes = append(fc.DailyTimes, at.Time)
				}
			}

			continue
		}
		target := []float64{}
		err := json.Unmarshal(v, &target)
		if err != nil {
			return nil, err
		}
		fc.DailyMetrics[k] = target
	}

	return fc, nil
}
