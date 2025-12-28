package omgo

import "time"

// DailyData contains daily aggregated weather data.
type DailyData struct {
	// Time contains timestamps for each day (at 00:00).
	Times []time.Time `json:"-"` // parsed separately from "time" field

	// Weather code (most severe of the day)
	WeatherCode []WeatherCode `json:"weather_code,omitempty"`

	// Temperature
	Temperature2mMax  []float64 `json:"temperature_2m_max,omitempty"`
	Temperature2mMin  []float64 `json:"temperature_2m_min,omitempty"`
	Temperature2mMean []float64 `json:"temperature_2m_mean,omitempty"`

	// Apparent temperature
	ApparentTemperatureMax  []float64 `json:"apparent_temperature_max,omitempty"`
	ApparentTemperatureMin  []float64 `json:"apparent_temperature_min,omitempty"`
	ApparentTemperatureMean []float64 `json:"apparent_temperature_mean,omitempty"`

	// Sun times (as time.Time)
	Sunrise []time.Time `json:"-"` // parsed separately
	Sunset  []time.Time `json:"-"` // parsed separately

	// Sunshine and daylight duration (seconds)
	SunshineDuration []float64 `json:"sunshine_duration,omitempty"`
	DaylightDuration []float64 `json:"daylight_duration,omitempty"`

	// Precipitation
	PrecipitationSum   []float64 `json:"precipitation_sum,omitempty"`
	RainSum            []float64 `json:"rain_sum,omitempty"`
	ShowersSum         []float64 `json:"showers_sum,omitempty"`
	SnowfallSum        []float64 `json:"snowfall_sum,omitempty"`
	PrecipitationHours []float64 `json:"precipitation_hours,omitempty"`

	// Precipitation probability
	PrecipitationProbabilityMax  []float64 `json:"precipitation_probability_max,omitempty"`
	PrecipitationProbabilityMin  []float64 `json:"precipitation_probability_min,omitempty"`
	PrecipitationProbabilityMean []float64 `json:"precipitation_probability_mean,omitempty"`

	// Wind
	WindSpeed10mMax          []float64 `json:"wind_speed_10m_max,omitempty"`
	WindGusts10mMax          []float64 `json:"wind_gusts_10m_max,omitempty"`
	WindDirection10mDominant []float64 `json:"wind_direction_10m_dominant,omitempty"`

	// Radiation (MJ/m²)
	ShortwaveRadiationSum []float64 `json:"shortwave_radiation_sum,omitempty"`

	// Evapotranspiration
	ET0FAOEvapotranspiration []float64 `json:"et0_fao_evapotranspiration,omitempty"`

	// UV Index
	UVIndexMax         []float64 `json:"uv_index_max,omitempty"`
	UVIndexClearSkyMax []float64 `json:"uv_index_clear_sky_max,omitempty"`
}
