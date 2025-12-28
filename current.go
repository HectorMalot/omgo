package omgo

import "time"

// CurrentData contains current weather conditions.
type CurrentData struct {
	// Time of the current weather observation.
	Time time.Time `json:"-"` // parsed separately

	// Interval is the time interval in seconds used for aggregations.
	Interval int `json:"interval,omitempty"`

	// Weather data
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

// IsDaytime returns true if it's currently daytime at the location.
func (c *CurrentData) IsDaytime() bool {
	return c.IsDay != nil && *c.IsDay == 1
}
