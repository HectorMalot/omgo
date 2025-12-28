package omgo

// DailyMetric represents a metric that can be requested for daily aggregated data.
type DailyMetric string

// Daily weather metrics available from the Open-Meteo API.
const (
	// Weather code
	DailyWeatherCode DailyMetric = "weather_code"

	// Temperature
	DailyTemperature2mMax  DailyMetric = "temperature_2m_max"
	DailyTemperature2mMin  DailyMetric = "temperature_2m_min"
	DailyTemperature2mMean DailyMetric = "temperature_2m_mean"

	// Apparent temperature
	DailyApparentTemperatureMax  DailyMetric = "apparent_temperature_max"
	DailyApparentTemperatureMin  DailyMetric = "apparent_temperature_min"
	DailyApparentTemperatureMean DailyMetric = "apparent_temperature_mean"

	// Sun
	DailySunrise          DailyMetric = "sunrise"
	DailySunset           DailyMetric = "sunset"
	DailySunshineDuration DailyMetric = "sunshine_duration"
	DailyDaylightDuration DailyMetric = "daylight_duration"

	// Precipitation
	DailyPrecipitationSum   DailyMetric = "precipitation_sum"
	DailyRainSum            DailyMetric = "rain_sum"
	DailyShowersSum         DailyMetric = "showers_sum"
	DailySnowfallSum        DailyMetric = "snowfall_sum"
	DailyPrecipitationHours DailyMetric = "precipitation_hours"

	// Precipitation probability
	DailyPrecipitationProbabilityMax  DailyMetric = "precipitation_probability_max"
	DailyPrecipitationProbabilityMin  DailyMetric = "precipitation_probability_min"
	DailyPrecipitationProbabilityMean DailyMetric = "precipitation_probability_mean"

	// Wind
	DailyWindSpeed10mMax          DailyMetric = "wind_speed_10m_max"
	DailyWindGusts10mMax          DailyMetric = "wind_gusts_10m_max"
	DailyWindDirection10mDominant DailyMetric = "wind_direction_10m_dominant"

	// Radiation
	DailyShortwaveRadiationSum DailyMetric = "shortwave_radiation_sum"

	// Evapotranspiration
	DailyET0FAOEvapotranspiration DailyMetric = "et0_fao_evapotranspiration"

	// UV Index
	DailyUVIndexMax         DailyMetric = "uv_index_max"
	DailyUVIndexClearSkyMax DailyMetric = "uv_index_clear_sky_max"
)

// String returns the API parameter string for the metric.
func (m DailyMetric) String() string {
	return string(m)
}
