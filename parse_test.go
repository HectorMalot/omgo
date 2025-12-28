package omgo

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseHourlyForecast(t *testing.T) {
	data, err := os.ReadFile("testdata/forecast_hourly.json")
	require.NoError(t, err)

	weather, err := parseWeatherResponse(data)
	require.NoError(t, err)

	// Check metadata
	assert.InDelta(t, 52.52, weather.Latitude, 0.01)
	assert.InDelta(t, 13.42, weather.Longitude, 0.01)
	assert.Equal(t, "Europe/Berlin", weather.Timezone)
	assert.Equal(t, "CET", weather.TimezoneAbbreviation)
	assert.Equal(t, 3600, weather.UTCOffsetSeconds)

	// Check hourly data
	require.NotNil(t, weather.Hourly)
	assert.Len(t, weather.Hourly.Times, 3)
	assert.Len(t, weather.Hourly.Temperature2m, 3)
	assert.Len(t, weather.Hourly.RelativeHumidity2m, 3)
	assert.Len(t, weather.Hourly.Precipitation, 3)
	assert.Len(t, weather.Hourly.WeatherCode, 3)
	assert.Len(t, weather.Hourly.WindSpeed10m, 3)

	// Check values
	assert.Equal(t, 2.5, weather.Hourly.Temperature2m[0])
	assert.Equal(t, WeatherCode(3), weather.Hourly.WeatherCode[0])
	assert.Equal(t, WeatherCode(61), weather.Hourly.WeatherCode[1])

	// Check time parsing with timezone
	loc, _ := time.LoadLocation("Europe/Berlin")
	expectedTime := time.Date(2024, 1, 15, 0, 0, 0, 0, loc)
	assert.Equal(t, expectedTime, weather.Hourly.Times[0])

	// Check units
	require.NotNil(t, weather.HourlyUnits)
	assert.Equal(t, "°C", weather.HourlyUnits.Temperature2m)
	assert.Equal(t, "mm", weather.HourlyUnits.Precipitation)
	assert.Equal(t, "km/h", weather.HourlyUnits.WindSpeed10m)
}

func TestParseDailyForecast(t *testing.T) {
	data, err := os.ReadFile("testdata/forecast_daily.json")
	require.NoError(t, err)

	weather, err := parseWeatherResponse(data)
	require.NoError(t, err)

	// Check daily data
	require.NotNil(t, weather.Daily)
	assert.Len(t, weather.Daily.Times, 3)
	assert.Len(t, weather.Daily.Temperature2mMax, 3)
	assert.Len(t, weather.Daily.Temperature2mMin, 3)
	assert.Len(t, weather.Daily.Sunrise, 3)
	assert.Len(t, weather.Daily.Sunset, 3)

	// Check values
	assert.Equal(t, 5.2, weather.Daily.Temperature2mMax[0])
	assert.Equal(t, -1.2, weather.Daily.Temperature2mMin[0])

	// Check time parsing
	loc, _ := time.LoadLocation("Europe/Berlin")
	expectedDate := time.Date(2024, 1, 15, 0, 0, 0, 0, loc)
	assert.Equal(t, expectedDate, weather.Daily.Times[0])

	// Check sunrise/sunset parsing
	expectedSunrise := time.Date(2024, 1, 15, 8, 15, 0, 0, loc)
	assert.Equal(t, expectedSunrise, weather.Daily.Sunrise[0])

	expectedSunset := time.Date(2024, 1, 15, 16, 30, 0, 0, loc)
	assert.Equal(t, expectedSunset, weather.Daily.Sunset[0])

	// Check units
	require.NotNil(t, weather.DailyUnits)
	assert.Equal(t, "°C", weather.DailyUnits.Temperature2mMax)
}

func TestParseCurrentWeather(t *testing.T) {
	data, err := os.ReadFile("testdata/forecast_current.json")
	require.NoError(t, err)

	weather, err := parseWeatherResponse(data)
	require.NoError(t, err)

	// Check current data
	require.NotNil(t, weather.Current)
	assert.Equal(t, 900, weather.Current.Interval)

	require.NotNil(t, weather.Current.Temperature2m)
	assert.Equal(t, 3.5, *weather.Current.Temperature2m)

	require.NotNil(t, weather.Current.IsDay)
	assert.Equal(t, 1, *weather.Current.IsDay)
	assert.True(t, weather.Current.IsDaytime())

	require.NotNil(t, weather.Current.WeatherCode)
	assert.Equal(t, PartlyCloudy, *weather.Current.WeatherCode)

	// Check time
	loc, _ := time.LoadLocation("Europe/Berlin")
	expectedTime := time.Date(2024, 1, 15, 14, 0, 0, 0, loc)
	assert.Equal(t, expectedTime, weather.Current.Time)

	// Check units
	require.NotNil(t, weather.CurrentUnits)
	assert.Equal(t, "°C", weather.CurrentUnits.Temperature2m)
}

func TestParseHistorical(t *testing.T) {
	data, err := os.ReadFile("testdata/historical.json")
	require.NoError(t, err)

	weather, err := parseWeatherResponse(data)
	require.NoError(t, err)

	require.NotNil(t, weather.Hourly)
	assert.Len(t, weather.Hourly.Times, 3)
	assert.Equal(t, 18.5, weather.Hourly.Temperature2m[0])
}

func TestWeatherCodeString(t *testing.T) {
	tests := []struct {
		code     WeatherCode
		expected string
	}{
		{ClearSky, "Clear sky"},
		{PartlyCloudy, "Partly cloudy"},
		{Fog, "Fog"},
		{RainSlight, "Slight rain"},
		{ThunderstormSlight, "Thunderstorm"},
		{WeatherCode(999), "Unknown (999)"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.code.String())
		})
	}
}
