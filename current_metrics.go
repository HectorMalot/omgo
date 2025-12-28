package omgo

// CurrentMetric represents a metric that can be requested for current conditions.
type CurrentMetric string

// Current weather metrics available from the Open-Meteo API.
const (
	CurrentTemperature2m       CurrentMetric = "temperature_2m"
	CurrentRelativeHumidity2m  CurrentMetric = "relative_humidity_2m"
	CurrentApparentTemperature CurrentMetric = "apparent_temperature"
	CurrentIsDay               CurrentMetric = "is_day"
	CurrentPrecipitation       CurrentMetric = "precipitation"
	CurrentRain                CurrentMetric = "rain"
	CurrentShowers             CurrentMetric = "showers"
	CurrentSnowfall            CurrentMetric = "snowfall"
	CurrentWeatherCode         CurrentMetric = "weather_code"
	CurrentCloudCover          CurrentMetric = "cloud_cover"
	CurrentPressureMSL         CurrentMetric = "pressure_msl"
	CurrentSurfacePressure     CurrentMetric = "surface_pressure"
	CurrentWindSpeed10m        CurrentMetric = "wind_speed_10m"
	CurrentWindDirection10m    CurrentMetric = "wind_direction_10m"
	CurrentWindGusts10m        CurrentMetric = "wind_gusts_10m"
)

// String returns the API parameter string for the metric.
func (m CurrentMetric) String() string {
	return string(m)
}
