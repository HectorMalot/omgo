package omgo

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	forecastBaseURL   = "https://api.open-meteo.com/v1/forecast"
	historicalBaseURL = "https://archive-api.open-meteo.com/v1/archive"
)

// buildURL builds the URL for a forecast request.
func (r *ForecastRequest) buildURL(baseURL, apiKey string) string {
	params := url.Values{}

	// Location
	params.Set("latitude", formatFloat(r.location.Latitude))
	params.Set("longitude", formatFloat(r.location.Longitude))
	if r.location.Elevation != nil {
		params.Set("elevation", formatFloat(*r.location.Elevation))
	}

	// Metrics
	if len(r.hourlyMetrics) > 0 {
		params.Set("hourly", joinMetrics(r.hourlyMetrics))
	}
	if len(r.dailyMetrics) > 0 {
		params.Set("daily", joinMetrics(r.dailyMetrics))
	}
	if len(r.currentMetrics) > 0 {
		params.Set("current", joinMetrics(r.currentMetrics))
	}
	if len(r.minutely15Metrics) > 0 {
		params.Set("minutely_15", joinMetrics(r.minutely15Metrics))
	}

	// Units
	if r.temperatureUnit != "" {
		params.Set("temperature_unit", string(r.temperatureUnit))
	}
	if r.windSpeedUnit != "" {
		params.Set("wind_speed_unit", string(r.windSpeedUnit))
	}
	if r.precipitationUnit != "" {
		params.Set("precipitation_unit", string(r.precipitationUnit))
	}

	// Time options
	if r.timezone != "" {
		params.Set("timezone", r.timezone)
	}
	if r.forecastDays > 0 {
		params.Set("forecast_days", strconv.Itoa(r.forecastDays))
	}
	if r.pastDays > 0 {
		params.Set("past_days", strconv.Itoa(r.pastDays))
	}
	if r.forecastHours > 0 {
		params.Set("forecast_hours", strconv.Itoa(r.forecastHours))
	}
	if r.pastHours > 0 {
		params.Set("past_hours", strconv.Itoa(r.pastHours))
	}

	// Date/time range
	if r.startDate != "" {
		params.Set("start_date", r.startDate)
	}
	if r.endDate != "" {
		params.Set("end_date", r.endDate)
	}
	if r.startHour != "" {
		params.Set("start_hour", r.startHour)
	}
	if r.endHour != "" {
		params.Set("end_hour", r.endHour)
	}

	// Other options
	if r.timeFormat != "" {
		params.Set("timeformat", string(r.timeFormat))
	}
	if r.cellSelection != "" {
		params.Set("cell_selection", string(r.cellSelection))
	}
	if len(r.models) > 0 {
		params.Set("models", strings.Join(r.models, ","))
	}

	// Solar options
	if r.tilt != nil {
		params.Set("tilt", formatFloat(*r.tilt))
	}
	if r.azimuth != nil {
		params.Set("azimuth", formatFloat(*r.azimuth))
	}

	// API key for commercial access
	if apiKey != "" {
		params.Set("apikey", apiKey)
	}

	return baseURL + "?" + params.Encode()
}

// buildURL builds the URL for a historical request.
func (r *HistoricalRequest) buildURL(baseURL, apiKey string) string {
	params := url.Values{}

	// Location
	params.Set("latitude", formatFloat(r.location.Latitude))
	params.Set("longitude", formatFloat(r.location.Longitude))
	if r.location.Elevation != nil {
		params.Set("elevation", formatFloat(*r.location.Elevation))
	}

	// Required date range
	params.Set("start_date", r.startDate)
	params.Set("end_date", r.endDate)

	// Metrics
	if len(r.hourlyMetrics) > 0 {
		params.Set("hourly", joinMetrics(r.hourlyMetrics))
	}
	if len(r.dailyMetrics) > 0 {
		params.Set("daily", joinMetrics(r.dailyMetrics))
	}

	// Units
	if r.temperatureUnit != "" {
		params.Set("temperature_unit", string(r.temperatureUnit))
	}
	if r.windSpeedUnit != "" {
		params.Set("wind_speed_unit", string(r.windSpeedUnit))
	}
	if r.precipitationUnit != "" {
		params.Set("precipitation_unit", string(r.precipitationUnit))
	}

	// Other options
	if r.timezone != "" {
		params.Set("timezone", r.timezone)
	}
	if r.timeFormat != "" {
		params.Set("timeformat", string(r.timeFormat))
	}
	if r.cellSelection != "" {
		params.Set("cell_selection", string(r.cellSelection))
	}

	// Solar options
	if r.tilt != nil {
		params.Set("tilt", formatFloat(*r.tilt))
	}
	if r.azimuth != nil {
		params.Set("azimuth", formatFloat(*r.azimuth))
	}

	// API key for commercial access
	if apiKey != "" {
		params.Set("apikey", apiKey)
	}

	return baseURL + "?" + params.Encode()
}

// Helper functions

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// joinMetrics converts a slice of metric constants to a comma-separated string.
// Works with any metric type (HourlyMetric, DailyMetric, CurrentMetric, Minutely15Metric).
func joinMetrics[T ~string](metrics []T) string {
	strs := make([]string, len(metrics))
	for i, m := range metrics {
		strs[i] = string(m)
	}
	return strings.Join(strs, ",")
}
