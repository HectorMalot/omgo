package omgo

import (
	"encoding/json"
	"time"
)

// rawResponse represents the raw JSON response from the API.
// This is used as an intermediate step for parsing.
type rawResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	Elevation            float64 `json:"elevation"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	UTCOffsetSeconds     int     `json:"utc_offset_seconds"`
	GenerationTimeMs     float64 `json:"generationtime_ms"`

	Current      json.RawMessage `json:"current,omitempty"`
	CurrentUnits *CurrentUnits   `json:"current_units,omitempty"`

	Hourly      json.RawMessage `json:"hourly,omitempty"`
	HourlyUnits *HourlyUnits    `json:"hourly_units,omitempty"`

	Minutely15      json.RawMessage  `json:"minutely_15,omitempty"`
	Minutely15Units *Minutely15Units `json:"minutely_15_units,omitempty"`

	Daily      json.RawMessage `json:"daily,omitempty"`
	DailyUnits *DailyUnits     `json:"daily_units,omitempty"`
}

// rawCurrent represents the raw current weather data.
type rawCurrent struct {
	Time                string       `json:"time"`
	Interval            int          `json:"interval"`
	Temperature2m       *float64     `json:"temperature_2m,omitempty"`
	RelativeHumidity2m  *float64     `json:"relative_humidity_2m,omitempty"`
	ApparentTemperature *float64     `json:"apparent_temperature,omitempty"`
	IsDay               *int         `json:"is_day,omitempty"`
	Precipitation       *float64     `json:"precipitation,omitempty"`
	Rain                *float64     `json:"rain,omitempty"`
	Showers             *float64     `json:"showers,omitempty"`
	Snowfall            *float64     `json:"snowfall,omitempty"`
	WeatherCode         *WeatherCode `json:"weather_code,omitempty"`
	CloudCover          *float64     `json:"cloud_cover,omitempty"`
	PressureMSL         *float64     `json:"pressure_msl,omitempty"`
	SurfacePressure     *float64     `json:"surface_pressure,omitempty"`
	WindSpeed10m        *float64     `json:"wind_speed_10m,omitempty"`
	WindDirection10m    *float64     `json:"wind_direction_10m,omitempty"`
	WindGusts10m        *float64     `json:"wind_gusts_10m,omitempty"`
}

// rawHourly represents the raw hourly data with time as strings.
type rawHourly struct {
	Time []string `json:"time"`
	// All other fields will be unmarshaled directly into HourlyData
}

// rawDaily represents the raw daily data with time and sun times as strings.
type rawDaily struct {
	Time    []string `json:"time"`
	Sunrise []string `json:"sunrise,omitempty"`
	Sunset  []string `json:"sunset,omitempty"`
	// All other fields will be unmarshaled directly into DailyData
}

// rawMinutely15 represents the raw 15-minutely data with time as strings.
type rawMinutely15 struct {
	Time []string `json:"time"`
	// All other fields will be unmarshaled directly into Minutely15Data
}

// parseWeatherResponse parses the API response into a Weather struct.
func parseWeatherResponse(body []byte) (*Weather, error) {
	var raw rawResponse
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	// Load timezone for proper time parsing
	var loc *time.Location
	if raw.Timezone != "" {
		var err error
		loc, err = time.LoadLocation(raw.Timezone)
		if err != nil {
			// Fall back to UTC if timezone is invalid
			loc = time.UTC
		}
	}

	weather := &Weather{
		Latitude:             raw.Latitude,
		Longitude:            raw.Longitude,
		Elevation:            raw.Elevation,
		Timezone:             raw.Timezone,
		TimezoneAbbreviation: raw.TimezoneAbbreviation,
		UTCOffsetSeconds:     raw.UTCOffsetSeconds,
		GenerationTimeMs:     raw.GenerationTimeMs,
		CurrentUnits:         raw.CurrentUnits,
		HourlyUnits:          raw.HourlyUnits,
		Minutely15Units:      raw.Minutely15Units,
		DailyUnits:           raw.DailyUnits,
	}

	// Parse current weather
	if len(raw.Current) > 0 {
		current, err := parseCurrent(raw.Current, loc)
		if err != nil {
			return nil, err
		}
		weather.Current = current
	}

	// Parse hourly data
	if len(raw.Hourly) > 0 {
		hourly, err := parseHourly(raw.Hourly, loc)
		if err != nil {
			return nil, err
		}
		weather.Hourly = hourly
	}

	// Parse 15-minutely data
	if len(raw.Minutely15) > 0 {
		minutely15, err := parseMinutely15(raw.Minutely15, loc)
		if err != nil {
			return nil, err
		}
		weather.Minutely15 = minutely15
	}

	// Parse daily data
	if len(raw.Daily) > 0 {
		daily, err := parseDaily(raw.Daily, loc)
		if err != nil {
			return nil, err
		}
		weather.Daily = daily
	}

	return weather, nil
}

// parseCurrent parses current weather data.
func parseCurrent(data json.RawMessage, loc *time.Location) (*CurrentData, error) {
	var raw rawCurrent
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	t, err := parseDateTime(raw.Time, loc)
	if err != nil {
		return nil, err
	}

	return &CurrentData{
		Time:                t,
		Interval:            raw.Interval,
		Temperature2m:       raw.Temperature2m,
		RelativeHumidity2m:  raw.RelativeHumidity2m,
		ApparentTemperature: raw.ApparentTemperature,
		IsDay:               raw.IsDay,
		Precipitation:       raw.Precipitation,
		Rain:                raw.Rain,
		Showers:             raw.Showers,
		Snowfall:            raw.Snowfall,
		WeatherCode:         raw.WeatherCode,
		CloudCover:          raw.CloudCover,
		PressureMSL:         raw.PressureMSL,
		SurfacePressure:     raw.SurfacePressure,
		WindSpeed10m:        raw.WindSpeed10m,
		WindDirection10m:    raw.WindDirection10m,
		WindGusts10m:        raw.WindGusts10m,
	}, nil
}

// parseHourly parses hourly weather data.
func parseHourly(data json.RawMessage, loc *time.Location) (*HourlyData, error) {
	// First, parse just the time array
	var rawTime rawHourly
	if err := json.Unmarshal(data, &rawTime); err != nil {
		return nil, err
	}

	// Parse times
	times, err := parseDateTimeArray(rawTime.Time, loc)
	if err != nil {
		return nil, err
	}

	// Parse all other fields into HourlyData
	hourly := &HourlyData{}
	if err := json.Unmarshal(data, hourly); err != nil {
		return nil, err
	}

	hourly.Times = times
	return hourly, nil
}

// parseMinutely15 parses 15-minutely weather data.
func parseMinutely15(data json.RawMessage, loc *time.Location) (*Minutely15Data, error) {
	// First, parse just the time array
	var rawTime rawMinutely15
	if err := json.Unmarshal(data, &rawTime); err != nil {
		return nil, err
	}

	// Parse times
	times, err := parseDateTimeArray(rawTime.Time, loc)
	if err != nil {
		return nil, err
	}

	// Parse all other fields into Minutely15Data
	minutely15 := &Minutely15Data{}
	if err := json.Unmarshal(data, minutely15); err != nil {
		return nil, err
	}

	minutely15.Times = times
	return minutely15, nil
}

// parseDaily parses daily weather data.
func parseDaily(data json.RawMessage, loc *time.Location) (*DailyData, error) {
	// First, parse time and sun times
	var rawTime rawDaily
	if err := json.Unmarshal(data, &rawTime); err != nil {
		return nil, err
	}

	// Parse dates
	times, err := parseDateArray(rawTime.Time, loc)
	if err != nil {
		return nil, err
	}

	// Parse all other fields into DailyData
	daily := &DailyData{}
	if err := json.Unmarshal(data, daily); err != nil {
		return nil, err
	}

	daily.Times = times

	// Parse sunrise/sunset times
	if len(rawTime.Sunrise) > 0 {
		sunrise, err := parseDateTimeArray(rawTime.Sunrise, loc)
		if err != nil {
			return nil, err
		}
		daily.Sunrise = sunrise
	}
	if len(rawTime.Sunset) > 0 {
		sunset, err := parseDateTimeArray(rawTime.Sunset, loc)
		if err != nil {
			return nil, err
		}
		daily.Sunset = sunset
	}

	return daily, nil
}
