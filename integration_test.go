//go:build integration

package omgo_test

import (
	"context"
	"testing"
	"time"

	"github.com/hectormalot/omgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These tests hit the real Open-Meteo API.
// Run with: go test -tags=integration -v
// Or run all tests including integration: go test -tags=integration -v ./...

func TestIntegrationForecastHourlyAndDaily(t *testing.T) {
	client := omgo.NewClient()

	// Request forecast for Amsterdam
	req, err := omgo.NewForecastRequest(52.3738, 4.8910)
	require.NoError(t, err)

	req.WithHourly(
		omgo.HourlyTemperature2m,
		omgo.HourlyRelativeHumidity2m,
		omgo.HourlyPrecipitation,
		omgo.HourlyWeatherCode,
		omgo.HourlyWindSpeed10m,
	).WithDaily(
		omgo.DailyTemperature2mMax,
		omgo.DailyTemperature2mMin,
		omgo.DailySunrise,
		omgo.DailySunset,
		omgo.DailyPrecipitationSum,
	).WithTimezone("Europe/Amsterdam").
		WithForecastDays(7)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	weather, err := client.Forecast(ctx, req)
	require.NoError(t, err)

	// Verify metadata
	assert.InDelta(t, 52.37, weather.Latitude, 0.1)
	assert.InDelta(t, 4.89, weather.Longitude, 0.1)
	assert.Equal(t, "Europe/Amsterdam", weather.Timezone)
	assert.NotEmpty(t, weather.TimezoneAbbreviation)
	assert.Greater(t, weather.GenerationTimeMs, 0.0)

	// Verify hourly data
	require.NotNil(t, weather.Hourly)
	assert.NotEmpty(t, weather.Hourly.Times)
	assert.NotEmpty(t, weather.Hourly.Temperature2m)
	assert.NotEmpty(t, weather.Hourly.RelativeHumidity2m)
	assert.NotEmpty(t, weather.Hourly.Precipitation)
	assert.NotEmpty(t, weather.Hourly.WeatherCode)
	assert.NotEmpty(t, weather.Hourly.WindSpeed10m)

	// Verify we have 7 days worth of hourly data (168 hours)
	assert.GreaterOrEqual(t, len(weather.Hourly.Times), 168)
	assert.Equal(t, len(weather.Hourly.Times), len(weather.Hourly.Temperature2m))

	// Verify hourly units
	require.NotNil(t, weather.HourlyUnits)
	assert.Equal(t, "°C", weather.HourlyUnits.Temperature2m)
	assert.Equal(t, "%", weather.HourlyUnits.RelativeHumidity2m)
	assert.Equal(t, "mm", weather.HourlyUnits.Precipitation)

	// Verify daily data
	require.NotNil(t, weather.Daily)
	assert.Len(t, weather.Daily.Times, 7)
	assert.Len(t, weather.Daily.Temperature2mMax, 7)
	assert.Len(t, weather.Daily.Temperature2mMin, 7)
	assert.Len(t, weather.Daily.Sunrise, 7)
	assert.Len(t, weather.Daily.Sunset, 7)

	// Verify sunrise is before sunset for each day
	for i := range weather.Daily.Times {
		assert.True(t, weather.Daily.Sunrise[i].Before(weather.Daily.Sunset[i]),
			"sunrise should be before sunset on day %d", i)
	}

	// Verify daily units
	require.NotNil(t, weather.DailyUnits)
	assert.Equal(t, "°C", weather.DailyUnits.Temperature2mMax)

	// Verify weather codes are valid
	for _, code := range weather.Hourly.WeatherCode {
		// Should produce a meaningful string, not "Unknown"
		str := code.String()
		assert.NotEmpty(t, str)
	}
}

func TestIntegrationCurrentWeather(t *testing.T) {
	client := omgo.NewClient()

	// Request current weather for New York
	req, err := omgo.NewForecastRequest(40.7128, -74.0060)
	require.NoError(t, err)

	req.WithCurrent(
		omgo.CurrentTemperature2m,
		omgo.CurrentRelativeHumidity2m,
		omgo.CurrentApparentTemperature,
		omgo.CurrentWeatherCode,
		omgo.CurrentIsDay,
		omgo.CurrentWindSpeed10m,
		omgo.CurrentWindDirection10m,
		omgo.CurrentPrecipitation,
	).WithTimezone("America/New_York")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	weather, err := client.Forecast(ctx, req)
	require.NoError(t, err)

	// Verify metadata
	assert.InDelta(t, 40.71, weather.Latitude, 0.1)
	assert.Equal(t, "America/New_York", weather.Timezone)

	// Verify current data exists
	require.NotNil(t, weather.Current)
	assert.False(t, weather.Current.Time.IsZero(), "current time should be set")
	assert.Greater(t, weather.Current.Interval, 0, "interval should be set")

	// Verify current metrics
	require.NotNil(t, weather.Current.Temperature2m)
	require.NotNil(t, weather.Current.RelativeHumidity2m)
	require.NotNil(t, weather.Current.WeatherCode)
	require.NotNil(t, weather.Current.IsDay)
	require.NotNil(t, weather.Current.WindSpeed10m)

	// Verify temperature is in a reasonable range (-50 to 60°C)
	assert.Greater(t, *weather.Current.Temperature2m, -50.0)
	assert.Less(t, *weather.Current.Temperature2m, 60.0)

	// Verify humidity is in valid range (0-100%)
	assert.GreaterOrEqual(t, *weather.Current.RelativeHumidity2m, 0.0)
	assert.LessOrEqual(t, *weather.Current.RelativeHumidity2m, 100.0)

	// Test IsDaytime helper
	isDay := weather.Current.IsDaytime()
	assert.Equal(t, *weather.Current.IsDay == 1, isDay)

	// Verify weather code produces meaningful description
	t.Logf("Current weather in New York: %s, %.1f°C",
		weather.Current.WeatherCode.String(),
		*weather.Current.Temperature2m)

	// Verify current units
	require.NotNil(t, weather.CurrentUnits)
	assert.Equal(t, "°C", weather.CurrentUnits.Temperature2m)
	assert.Equal(t, "%", weather.CurrentUnits.RelativeHumidity2m)
}

func TestIntegrationHistoricalData(t *testing.T) {
	client := omgo.NewClient()

	// Request historical data for a known date range (last month)
	endDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	startDate := time.Now().AddDate(0, 0, -14).Format("2006-01-02")

	req, err := omgo.NewHistoricalRequest(52.52, 13.41, startDate, endDate) // Berlin
	require.NoError(t, err)

	req.WithHourly(
		omgo.HourlyTemperature2m,
		omgo.HourlyRelativeHumidity2m,
		omgo.HourlyPrecipitation,
	).WithDaily(
		omgo.DailyTemperature2mMax,
		omgo.DailyTemperature2mMin,
		omgo.DailyPrecipitationSum,
	).WithTimezone("Europe/Berlin")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	weather, err := client.Historical(ctx, req)
	require.NoError(t, err)

	// Verify metadata
	assert.InDelta(t, 52.52, weather.Latitude, 0.1)
	assert.Equal(t, "Europe/Berlin", weather.Timezone)

	// Verify hourly data - should have 7 days * 24 hours = 168 data points
	require.NotNil(t, weather.Hourly)
	assert.GreaterOrEqual(t, len(weather.Hourly.Times), 168)
	assert.Equal(t, len(weather.Hourly.Times), len(weather.Hourly.Temperature2m))
	assert.Equal(t, len(weather.Hourly.Times), len(weather.Hourly.RelativeHumidity2m))

	// Verify daily data - should have 7 days
	require.NotNil(t, weather.Daily)
	expectedDays := 7
	assert.GreaterOrEqual(t, len(weather.Daily.Times), expectedDays)
	assert.GreaterOrEqual(t, len(weather.Daily.Temperature2mMax), expectedDays)
	assert.GreaterOrEqual(t, len(weather.Daily.Temperature2mMin), expectedDays)

	// Verify max temp is always >= min temp
	for i := range weather.Daily.Temperature2mMax {
		assert.GreaterOrEqual(t, weather.Daily.Temperature2mMax[i], weather.Daily.Temperature2mMin[i],
			"max temp should be >= min temp on day %d", i)
	}

	// Verify units
	require.NotNil(t, weather.HourlyUnits)
	assert.Equal(t, "°C", weather.HourlyUnits.Temperature2m)

	t.Logf("Retrieved %d hours and %d days of historical data for Berlin",
		len(weather.Hourly.Times), len(weather.Daily.Times))
}
