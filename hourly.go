package omgo

import "time"

// BaseMetrics contains fields shared between HourlyData and Minutely15Data.
// These are embedded into both structs.
type BaseMetrics struct {
	// Time contains timestamps for each data point.
	Times []time.Time `json:"-"` // parsed separately from "time" field

	// Basic weather
	Temperature2m       []float64 `json:"temperature_2m,omitempty"`
	RelativeHumidity2m  []float64 `json:"relative_humidity_2m,omitempty"`
	DewPoint2m          []float64 `json:"dew_point_2m,omitempty"`
	ApparentTemperature []float64 `json:"apparent_temperature,omitempty"`

	// Precipitation
	Precipitation []float64 `json:"precipitation,omitempty"`
	Rain          []float64 `json:"rain,omitempty"`
	Snowfall      []float64 `json:"snowfall,omitempty"`

	// Weather code
	WeatherCode []WeatherCode `json:"weather_code,omitempty"`

	// Cloud cover
	CloudCover     []float64 `json:"cloud_cover,omitempty"`
	CloudCoverLow  []float64 `json:"cloud_cover_low,omitempty"`
	CloudCoverMid  []float64 `json:"cloud_cover_mid,omitempty"`
	CloudCoverHigh []float64 `json:"cloud_cover_high,omitempty"`

	// Wind
	WindSpeed10m     []float64 `json:"wind_speed_10m,omitempty"`
	WindSpeed80m     []float64 `json:"wind_speed_80m,omitempty"`
	WindDirection10m []float64 `json:"wind_direction_10m,omitempty"`
	WindDirection80m []float64 `json:"wind_direction_80m,omitempty"`
	WindGusts10m     []float64 `json:"wind_gusts_10m,omitempty"`

	// Solar radiation
	ShortwaveRadiation     []float64 `json:"shortwave_radiation,omitempty"`
	DirectRadiation        []float64 `json:"direct_radiation,omitempty"`
	DirectNormalIrradiance []float64 `json:"direct_normal_irradiance,omitempty"`
	DiffuseRadiation       []float64 `json:"diffuse_radiation,omitempty"`
	GlobalTiltedIrradiance []float64 `json:"global_tilted_irradiance,omitempty"`

	// Other
	Visibility               []float64 `json:"visibility,omitempty"`
	Evapotranspiration       []float64 `json:"evapotranspiration,omitempty"`
	ET0FAOEvapotranspiration []float64 `json:"et0_fao_evapotranspiration,omitempty"`
	VapourPressureDeficit    []float64 `json:"vapour_pressure_deficit,omitempty"`
	Cape                     []float64 `json:"cape,omitempty"`
	FreezingLevelHeight      []float64 `json:"freezing_level_height,omitempty"`
	SunshineDuration         []float64 `json:"sunshine_duration,omitempty"`
}

// HourlyData contains hourly weather data.
type HourlyData struct {
	BaseMetrics // embedded shared fields

	// Pressure
	PressureMSL     []float64 `json:"pressure_msl,omitempty"`
	SurfacePressure []float64 `json:"surface_pressure,omitempty"`

	// Additional wind levels
	WindSpeed120m     []float64 `json:"wind_speed_120m,omitempty"`
	WindSpeed180m     []float64 `json:"wind_speed_180m,omitempty"`
	WindDirection120m []float64 `json:"wind_direction_120m,omitempty"`
	WindDirection180m []float64 `json:"wind_direction_180m,omitempty"`

	// Snow and precipitation
	SnowDepth                []float64 `json:"snow_depth,omitempty"`
	PrecipitationProbability []float64 `json:"precipitation_probability,omitempty"`
	Showers                  []float64 `json:"showers,omitempty"`

	// Is day (1 = day, 0 = night)
	IsDay []int `json:"is_day,omitempty"`

	// Soil temperature
	SoilTemperature0cm  []float64 `json:"soil_temperature_0cm,omitempty"`
	SoilTemperature6cm  []float64 `json:"soil_temperature_6cm,omitempty"`
	SoilTemperature18cm []float64 `json:"soil_temperature_18cm,omitempty"`
	SoilTemperature54cm []float64 `json:"soil_temperature_54cm,omitempty"`

	// Soil moisture
	SoilMoisture0to1cm   []float64 `json:"soil_moisture_0_to_1cm,omitempty"`
	SoilMoisture1to3cm   []float64 `json:"soil_moisture_1_to_3cm,omitempty"`
	SoilMoisture3to9cm   []float64 `json:"soil_moisture_3_to_9cm,omitempty"`
	SoilMoisture9to27cm  []float64 `json:"soil_moisture_9_to_27cm,omitempty"`
	SoilMoisture27to81cm []float64 `json:"soil_moisture_27_to_81cm,omitempty"`

	// Pressure level: Temperature
	Temperature1000hPa []float64 `json:"temperature_1000hPa,omitempty"`
	Temperature975hPa  []float64 `json:"temperature_975hPa,omitempty"`
	Temperature950hPa  []float64 `json:"temperature_950hPa,omitempty"`
	Temperature925hPa  []float64 `json:"temperature_925hPa,omitempty"`
	Temperature900hPa  []float64 `json:"temperature_900hPa,omitempty"`
	Temperature850hPa  []float64 `json:"temperature_850hPa,omitempty"`
	Temperature800hPa  []float64 `json:"temperature_800hPa,omitempty"`
	Temperature700hPa  []float64 `json:"temperature_700hPa,omitempty"`
	Temperature600hPa  []float64 `json:"temperature_600hPa,omitempty"`
	Temperature500hPa  []float64 `json:"temperature_500hPa,omitempty"`
	Temperature400hPa  []float64 `json:"temperature_400hPa,omitempty"`
	Temperature300hPa  []float64 `json:"temperature_300hPa,omitempty"`
	Temperature250hPa  []float64 `json:"temperature_250hPa,omitempty"`
	Temperature200hPa  []float64 `json:"temperature_200hPa,omitempty"`
	Temperature150hPa  []float64 `json:"temperature_150hPa,omitempty"`
	Temperature100hPa  []float64 `json:"temperature_100hPa,omitempty"`
	Temperature70hPa   []float64 `json:"temperature_70hPa,omitempty"`
	Temperature50hPa   []float64 `json:"temperature_50hPa,omitempty"`
	Temperature30hPa   []float64 `json:"temperature_30hPa,omitempty"`

	// Pressure level: Relative Humidity
	RelativeHumidity1000hPa []float64 `json:"relative_humidity_1000hPa,omitempty"`
	RelativeHumidity975hPa  []float64 `json:"relative_humidity_975hPa,omitempty"`
	RelativeHumidity950hPa  []float64 `json:"relative_humidity_950hPa,omitempty"`
	RelativeHumidity925hPa  []float64 `json:"relative_humidity_925hPa,omitempty"`
	RelativeHumidity900hPa  []float64 `json:"relative_humidity_900hPa,omitempty"`
	RelativeHumidity850hPa  []float64 `json:"relative_humidity_850hPa,omitempty"`
	RelativeHumidity800hPa  []float64 `json:"relative_humidity_800hPa,omitempty"`
	RelativeHumidity700hPa  []float64 `json:"relative_humidity_700hPa,omitempty"`
	RelativeHumidity600hPa  []float64 `json:"relative_humidity_600hPa,omitempty"`
	RelativeHumidity500hPa  []float64 `json:"relative_humidity_500hPa,omitempty"`
	RelativeHumidity400hPa  []float64 `json:"relative_humidity_400hPa,omitempty"`
	RelativeHumidity300hPa  []float64 `json:"relative_humidity_300hPa,omitempty"`
	RelativeHumidity250hPa  []float64 `json:"relative_humidity_250hPa,omitempty"`
	RelativeHumidity200hPa  []float64 `json:"relative_humidity_200hPa,omitempty"`
	RelativeHumidity150hPa  []float64 `json:"relative_humidity_150hPa,omitempty"`
	RelativeHumidity100hPa  []float64 `json:"relative_humidity_100hPa,omitempty"`
	RelativeHumidity70hPa   []float64 `json:"relative_humidity_70hPa,omitempty"`
	RelativeHumidity50hPa   []float64 `json:"relative_humidity_50hPa,omitempty"`
	RelativeHumidity30hPa   []float64 `json:"relative_humidity_30hPa,omitempty"`

	// Pressure level: Dew Point
	DewPoint1000hPa []float64 `json:"dew_point_1000hPa,omitempty"`
	DewPoint975hPa  []float64 `json:"dew_point_975hPa,omitempty"`
	DewPoint950hPa  []float64 `json:"dew_point_950hPa,omitempty"`
	DewPoint925hPa  []float64 `json:"dew_point_925hPa,omitempty"`
	DewPoint900hPa  []float64 `json:"dew_point_900hPa,omitempty"`
	DewPoint850hPa  []float64 `json:"dew_point_850hPa,omitempty"`
	DewPoint800hPa  []float64 `json:"dew_point_800hPa,omitempty"`
	DewPoint700hPa  []float64 `json:"dew_point_700hPa,omitempty"`
	DewPoint600hPa  []float64 `json:"dew_point_600hPa,omitempty"`
	DewPoint500hPa  []float64 `json:"dew_point_500hPa,omitempty"`
	DewPoint400hPa  []float64 `json:"dew_point_400hPa,omitempty"`
	DewPoint300hPa  []float64 `json:"dew_point_300hPa,omitempty"`
	DewPoint250hPa  []float64 `json:"dew_point_250hPa,omitempty"`
	DewPoint200hPa  []float64 `json:"dew_point_200hPa,omitempty"`
	DewPoint150hPa  []float64 `json:"dew_point_150hPa,omitempty"`
	DewPoint100hPa  []float64 `json:"dew_point_100hPa,omitempty"`
	DewPoint70hPa   []float64 `json:"dew_point_70hPa,omitempty"`
	DewPoint50hPa   []float64 `json:"dew_point_50hPa,omitempty"`
	DewPoint30hPa   []float64 `json:"dew_point_30hPa,omitempty"`

	// Pressure level: Cloud Cover
	CloudCover1000hPa []float64 `json:"cloud_cover_1000hPa,omitempty"`
	CloudCover975hPa  []float64 `json:"cloud_cover_975hPa,omitempty"`
	CloudCover950hPa  []float64 `json:"cloud_cover_950hPa,omitempty"`
	CloudCover925hPa  []float64 `json:"cloud_cover_925hPa,omitempty"`
	CloudCover900hPa  []float64 `json:"cloud_cover_900hPa,omitempty"`
	CloudCover850hPa  []float64 `json:"cloud_cover_850hPa,omitempty"`
	CloudCover800hPa  []float64 `json:"cloud_cover_800hPa,omitempty"`
	CloudCover700hPa  []float64 `json:"cloud_cover_700hPa,omitempty"`
	CloudCover600hPa  []float64 `json:"cloud_cover_600hPa,omitempty"`
	CloudCover500hPa  []float64 `json:"cloud_cover_500hPa,omitempty"`
	CloudCover400hPa  []float64 `json:"cloud_cover_400hPa,omitempty"`
	CloudCover300hPa  []float64 `json:"cloud_cover_300hPa,omitempty"`
	CloudCover250hPa  []float64 `json:"cloud_cover_250hPa,omitempty"`
	CloudCover200hPa  []float64 `json:"cloud_cover_200hPa,omitempty"`
	CloudCover150hPa  []float64 `json:"cloud_cover_150hPa,omitempty"`
	CloudCover100hPa  []float64 `json:"cloud_cover_100hPa,omitempty"`
	CloudCover70hPa   []float64 `json:"cloud_cover_70hPa,omitempty"`
	CloudCover50hPa   []float64 `json:"cloud_cover_50hPa,omitempty"`
	CloudCover30hPa   []float64 `json:"cloud_cover_30hPa,omitempty"`

	// Pressure level: Wind Speed
	WindSpeed1000hPa []float64 `json:"wind_speed_1000hPa,omitempty"`
	WindSpeed975hPa  []float64 `json:"wind_speed_975hPa,omitempty"`
	WindSpeed950hPa  []float64 `json:"wind_speed_950hPa,omitempty"`
	WindSpeed925hPa  []float64 `json:"wind_speed_925hPa,omitempty"`
	WindSpeed900hPa  []float64 `json:"wind_speed_900hPa,omitempty"`
	WindSpeed850hPa  []float64 `json:"wind_speed_850hPa,omitempty"`
	WindSpeed800hPa  []float64 `json:"wind_speed_800hPa,omitempty"`
	WindSpeed700hPa  []float64 `json:"wind_speed_700hPa,omitempty"`
	WindSpeed600hPa  []float64 `json:"wind_speed_600hPa,omitempty"`
	WindSpeed500hPa  []float64 `json:"wind_speed_500hPa,omitempty"`
	WindSpeed400hPa  []float64 `json:"wind_speed_400hPa,omitempty"`
	WindSpeed300hPa  []float64 `json:"wind_speed_300hPa,omitempty"`
	WindSpeed250hPa  []float64 `json:"wind_speed_250hPa,omitempty"`
	WindSpeed200hPa  []float64 `json:"wind_speed_200hPa,omitempty"`
	WindSpeed150hPa  []float64 `json:"wind_speed_150hPa,omitempty"`
	WindSpeed100hPa  []float64 `json:"wind_speed_100hPa,omitempty"`
	WindSpeed70hPa   []float64 `json:"wind_speed_70hPa,omitempty"`
	WindSpeed50hPa   []float64 `json:"wind_speed_50hPa,omitempty"`
	WindSpeed30hPa   []float64 `json:"wind_speed_30hPa,omitempty"`

	// Pressure level: Wind Direction
	WindDirection1000hPa []float64 `json:"wind_direction_1000hPa,omitempty"`
	WindDirection975hPa  []float64 `json:"wind_direction_975hPa,omitempty"`
	WindDirection950hPa  []float64 `json:"wind_direction_950hPa,omitempty"`
	WindDirection925hPa  []float64 `json:"wind_direction_925hPa,omitempty"`
	WindDirection900hPa  []float64 `json:"wind_direction_900hPa,omitempty"`
	WindDirection850hPa  []float64 `json:"wind_direction_850hPa,omitempty"`
	WindDirection800hPa  []float64 `json:"wind_direction_800hPa,omitempty"`
	WindDirection700hPa  []float64 `json:"wind_direction_700hPa,omitempty"`
	WindDirection600hPa  []float64 `json:"wind_direction_600hPa,omitempty"`
	WindDirection500hPa  []float64 `json:"wind_direction_500hPa,omitempty"`
	WindDirection400hPa  []float64 `json:"wind_direction_400hPa,omitempty"`
	WindDirection300hPa  []float64 `json:"wind_direction_300hPa,omitempty"`
	WindDirection250hPa  []float64 `json:"wind_direction_250hPa,omitempty"`
	WindDirection200hPa  []float64 `json:"wind_direction_200hPa,omitempty"`
	WindDirection150hPa  []float64 `json:"wind_direction_150hPa,omitempty"`
	WindDirection100hPa  []float64 `json:"wind_direction_100hPa,omitempty"`
	WindDirection70hPa   []float64 `json:"wind_direction_70hPa,omitempty"`
	WindDirection50hPa   []float64 `json:"wind_direction_50hPa,omitempty"`
	WindDirection30hPa   []float64 `json:"wind_direction_30hPa,omitempty"`

	// Pressure level: Geopotential Height
	GeopotentialHeight1000hPa []float64 `json:"geopotential_height_1000hPa,omitempty"`
	GeopotentialHeight975hPa  []float64 `json:"geopotential_height_975hPa,omitempty"`
	GeopotentialHeight950hPa  []float64 `json:"geopotential_height_950hPa,omitempty"`
	GeopotentialHeight925hPa  []float64 `json:"geopotential_height_925hPa,omitempty"`
	GeopotentialHeight900hPa  []float64 `json:"geopotential_height_900hPa,omitempty"`
	GeopotentialHeight850hPa  []float64 `json:"geopotential_height_850hPa,omitempty"`
	GeopotentialHeight800hPa  []float64 `json:"geopotential_height_800hPa,omitempty"`
	GeopotentialHeight700hPa  []float64 `json:"geopotential_height_700hPa,omitempty"`
	GeopotentialHeight600hPa  []float64 `json:"geopotential_height_600hPa,omitempty"`
	GeopotentialHeight500hPa  []float64 `json:"geopotential_height_500hPa,omitempty"`
	GeopotentialHeight400hPa  []float64 `json:"geopotential_height_400hPa,omitempty"`
	GeopotentialHeight300hPa  []float64 `json:"geopotential_height_300hPa,omitempty"`
	GeopotentialHeight250hPa  []float64 `json:"geopotential_height_250hPa,omitempty"`
	GeopotentialHeight200hPa  []float64 `json:"geopotential_height_200hPa,omitempty"`
	GeopotentialHeight150hPa  []float64 `json:"geopotential_height_150hPa,omitempty"`
	GeopotentialHeight100hPa  []float64 `json:"geopotential_height_100hPa,omitempty"`
	GeopotentialHeight70hPa   []float64 `json:"geopotential_height_70hPa,omitempty"`
	GeopotentialHeight50hPa   []float64 `json:"geopotential_height_50hPa,omitempty"`
	GeopotentialHeight30hPa   []float64 `json:"geopotential_height_30hPa,omitempty"`
}
