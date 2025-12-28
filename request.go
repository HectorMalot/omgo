package omgo

import "fmt"

// ForecastRequest represents a request to the Forecast API.
type ForecastRequest struct {
	location Location

	// Metrics to request
	hourlyMetrics     []HourlyMetric
	dailyMetrics      []DailyMetric
	currentMetrics    []CurrentMetric
	minutely15Metrics []Minutely15Metric

	// Units
	temperatureUnit   TemperatureUnit
	windSpeedUnit     WindSpeedUnit
	precipitationUnit PrecipitationUnit

	// Time options
	timezone     string
	forecastDays int
	pastDays     int
	forecastHours int
	pastHours     int

	// Date range options
	startDate string
	endDate   string
	startHour string
	endHour   string

	// Other options
	timeFormat    TimeFormat
	cellSelection CellSelection
	models        []string

	// Solar radiation options (for global_tilted_irradiance)
	tilt    *float64
	azimuth *float64
}

// NewForecastRequest creates a new ForecastRequest for the given coordinates.
func NewForecastRequest(lat, lon float64) (*ForecastRequest, error) {
	loc, err := NewLocation(lat, lon)
	if err != nil {
		return nil, err
	}
	return &ForecastRequest{
		location: loc,
	}, nil
}

// WithLocation sets the location from an existing Location struct.
func (r *ForecastRequest) WithLocation(loc Location) *ForecastRequest {
	r.location = loc
	return r
}

// WithHourly adds hourly metrics to the request.
func (r *ForecastRequest) WithHourly(metrics ...HourlyMetric) *ForecastRequest {
	r.hourlyMetrics = append(r.hourlyMetrics, metrics...)
	return r
}

// WithDaily adds daily metrics to the request.
func (r *ForecastRequest) WithDaily(metrics ...DailyMetric) *ForecastRequest {
	r.dailyMetrics = append(r.dailyMetrics, metrics...)
	return r
}

// WithCurrent adds current weather metrics to the request.
func (r *ForecastRequest) WithCurrent(metrics ...CurrentMetric) *ForecastRequest {
	r.currentMetrics = append(r.currentMetrics, metrics...)
	return r
}

// WithMinutely15 adds 15-minutely metrics to the request.
func (r *ForecastRequest) WithMinutely15(metrics ...Minutely15Metric) *ForecastRequest {
	r.minutely15Metrics = append(r.minutely15Metrics, metrics...)
	return r
}

// WithTemperatureUnit sets the temperature unit for the response.
func (r *ForecastRequest) WithTemperatureUnit(unit TemperatureUnit) *ForecastRequest {
	r.temperatureUnit = unit
	return r
}

// WithWindSpeedUnit sets the wind speed unit for the response.
func (r *ForecastRequest) WithWindSpeedUnit(unit WindSpeedUnit) *ForecastRequest {
	r.windSpeedUnit = unit
	return r
}

// WithPrecipitationUnit sets the precipitation unit for the response.
func (r *ForecastRequest) WithPrecipitationUnit(unit PrecipitationUnit) *ForecastRequest {
	r.precipitationUnit = unit
	return r
}

// WithTimezone sets the timezone for the response.
// Use "auto" to automatically detect the timezone based on coordinates.
// Any timezone name from the time zone database is supported.
func (r *ForecastRequest) WithTimezone(tz string) *ForecastRequest {
	r.timezone = tz
	return r
}

// WithForecastDays sets the number of forecast days (0-16).
func (r *ForecastRequest) WithForecastDays(days int) *ForecastRequest {
	r.forecastDays = days
	return r
}

// WithPastDays sets the number of past days to include (0-92).
func (r *ForecastRequest) WithPastDays(days int) *ForecastRequest {
	r.pastDays = days
	return r
}

// WithForecastHours sets the number of forecast hours.
func (r *ForecastRequest) WithForecastHours(hours int) *ForecastRequest {
	r.forecastHours = hours
	return r
}

// WithPastHours sets the number of past hours to include.
func (r *ForecastRequest) WithPastHours(hours int) *ForecastRequest {
	r.pastHours = hours
	return r
}

// WithDateRange sets a specific date range for the forecast.
// Dates should be in ISO8601 format (yyyy-mm-dd).
func (r *ForecastRequest) WithDateRange(startDate, endDate string) *ForecastRequest {
	r.startDate = startDate
	r.endDate = endDate
	return r
}

// WithHourRange sets a specific hour range for the forecast.
// Times should be in ISO8601 format (yyyy-mm-ddThh:mm).
func (r *ForecastRequest) WithHourRange(startHour, endHour string) *ForecastRequest {
	r.startHour = startHour
	r.endHour = endHour
	return r
}

// WithTimeFormat sets the time format for the response.
func (r *ForecastRequest) WithTimeFormat(format TimeFormat) *ForecastRequest {
	r.timeFormat = format
	return r
}

// WithCellSelection sets the grid-cell selection preference.
func (r *ForecastRequest) WithCellSelection(selection CellSelection) *ForecastRequest {
	r.cellSelection = selection
	return r
}

// WithModels sets specific weather models to use.
func (r *ForecastRequest) WithModels(models ...string) *ForecastRequest {
	r.models = append(r.models, models...)
	return r
}

// WithTilt sets the tilt angle for global_tilted_irradiance calculations (0-90 degrees).
func (r *ForecastRequest) WithTilt(degrees float64) *ForecastRequest {
	r.tilt = &degrees
	return r
}

// WithAzimuth sets the azimuth angle for global_tilted_irradiance calculations.
// 0° = south, -90° = east, 90° = west, ±180° = north.
func (r *ForecastRequest) WithAzimuth(degrees float64) *ForecastRequest {
	r.azimuth = &degrees
	return r
}

// HistoricalRequest represents a request to the Historical API.
type HistoricalRequest struct {
	location Location

	// Required date range
	startDate string
	endDate   string

	// Metrics to request
	hourlyMetrics []HourlyMetric
	dailyMetrics  []DailyMetric

	// Units
	temperatureUnit   TemperatureUnit
	windSpeedUnit     WindSpeedUnit
	precipitationUnit PrecipitationUnit

	// Other options
	timezone      string
	timeFormat    TimeFormat
	cellSelection CellSelection

	// Solar radiation options
	tilt    *float64
	azimuth *float64
}

// NewHistoricalRequest creates a new HistoricalRequest for the given coordinates and date range.
// startDate and endDate should be in ISO8601 format (yyyy-mm-dd).
func NewHistoricalRequest(lat, lon float64, startDate, endDate string) (*HistoricalRequest, error) {
	loc, err := NewLocation(lat, lon)
	if err != nil {
		return nil, err
	}
	if startDate == "" {
		return nil, fmt.Errorf("startDate is required for historical requests")
	}
	if endDate == "" {
		return nil, fmt.Errorf("endDate is required for historical requests")
	}
	return &HistoricalRequest{
		location:  loc,
		startDate: startDate,
		endDate:   endDate,
	}, nil
}

// WithLocation sets the location from an existing Location struct.
func (r *HistoricalRequest) WithLocation(loc Location) *HistoricalRequest {
	r.location = loc
	return r
}

// WithHourly adds hourly metrics to the request.
func (r *HistoricalRequest) WithHourly(metrics ...HourlyMetric) *HistoricalRequest {
	r.hourlyMetrics = append(r.hourlyMetrics, metrics...)
	return r
}

// WithDaily adds daily metrics to the request.
func (r *HistoricalRequest) WithDaily(metrics ...DailyMetric) *HistoricalRequest {
	r.dailyMetrics = append(r.dailyMetrics, metrics...)
	return r
}

// WithTemperatureUnit sets the temperature unit for the response.
func (r *HistoricalRequest) WithTemperatureUnit(unit TemperatureUnit) *HistoricalRequest {
	r.temperatureUnit = unit
	return r
}

// WithWindSpeedUnit sets the wind speed unit for the response.
func (r *HistoricalRequest) WithWindSpeedUnit(unit WindSpeedUnit) *HistoricalRequest {
	r.windSpeedUnit = unit
	return r
}

// WithPrecipitationUnit sets the precipitation unit for the response.
func (r *HistoricalRequest) WithPrecipitationUnit(unit PrecipitationUnit) *HistoricalRequest {
	r.precipitationUnit = unit
	return r
}

// WithTimezone sets the timezone for the response.
func (r *HistoricalRequest) WithTimezone(tz string) *HistoricalRequest {
	r.timezone = tz
	return r
}

// WithTimeFormat sets the time format for the response.
func (r *HistoricalRequest) WithTimeFormat(format TimeFormat) *HistoricalRequest {
	r.timeFormat = format
	return r
}

// WithCellSelection sets the grid-cell selection preference.
func (r *HistoricalRequest) WithCellSelection(selection CellSelection) *HistoricalRequest {
	r.cellSelection = selection
	return r
}

// WithTilt sets the tilt angle for global_tilted_irradiance calculations (0-90 degrees).
func (r *HistoricalRequest) WithTilt(degrees float64) *HistoricalRequest {
	r.tilt = &degrees
	return r
}

// WithAzimuth sets the azimuth angle for global_tilted_irradiance calculations.
func (r *HistoricalRequest) WithAzimuth(degrees float64) *HistoricalRequest {
	r.azimuth = &degrees
	return r
}

