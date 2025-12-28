package omgo

// Minutely15Metric represents a metric that can be requested for 15-minutely data.
type Minutely15Metric string

// 15-minutely weather metrics available from the Open-Meteo API.
// Note: 15-minutely data is based on NOAA HRRR for North America and
// DWD ICON-D2 / Météo-France AROME for Central Europe.
const (
	// Temperature
	Minutely15Temperature2m       Minutely15Metric = "temperature_2m"
	Minutely15RelativeHumidity2m  Minutely15Metric = "relative_humidity_2m"
	Minutely15DewPoint2m          Minutely15Metric = "dew_point_2m"
	Minutely15ApparentTemperature Minutely15Metric = "apparent_temperature"

	// Radiation
	Minutely15ShortwaveRadiation          Minutely15Metric = "shortwave_radiation"
	Minutely15DirectRadiation             Minutely15Metric = "direct_radiation"
	Minutely15DirectNormalIrradiance      Minutely15Metric = "direct_normal_irradiance"
	Minutely15DiffuseRadiation            Minutely15Metric = "diffuse_radiation"
	Minutely15GlobalTiltedIrradiance      Minutely15Metric = "global_tilted_irradiance"
	Minutely15GlobalTiltedIrradianceInstant Minutely15Metric = "global_tilted_irradiance_instant"
	Minutely15SunshineDuration            Minutely15Metric = "sunshine_duration"

	// Lightning
	Minutely15LightningPotential Minutely15Metric = "lightning_potential"

	// Precipitation
	Minutely15Precipitation Minutely15Metric = "precipitation"
	Minutely15Snowfall      Minutely15Metric = "snowfall"
	Minutely15SnowfallHeight Minutely15Metric = "snowfall_height"
	Minutely15Rain          Minutely15Metric = "rain"
	Minutely15Showers       Minutely15Metric = "showers"
	Minutely15FreezingLevelHeight Minutely15Metric = "freezing_level_height"

	// Weather
	Minutely15Cape        Minutely15Metric = "cape"
	Minutely15WeatherCode Minutely15Metric = "weather_code"
	Minutely15Visibility  Minutely15Metric = "visibility"

	// Wind
	Minutely15WindSpeed10m     Minutely15Metric = "wind_speed_10m"
	Minutely15WindSpeed80m     Minutely15Metric = "wind_speed_80m"
	Minutely15WindDirection10m Minutely15Metric = "wind_direction_10m"
	Minutely15WindDirection80m Minutely15Metric = "wind_direction_80m"
	Minutely15WindGusts10m     Minutely15Metric = "wind_gusts_10m"
)

// String returns the API parameter string for the metric.
func (m Minutely15Metric) String() string {
	return string(m)
}

