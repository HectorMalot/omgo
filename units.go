package omgo

// TemperatureUnit specifies the unit for temperature values.
type TemperatureUnit string

const (
	Celsius    TemperatureUnit = "celsius"
	Fahrenheit TemperatureUnit = "fahrenheit"
)

// WindSpeedUnit specifies the unit for wind speed values.
type WindSpeedUnit string

const (
	KilometersPerHour WindSpeedUnit = "kmh"
	MetersPerSecond   WindSpeedUnit = "ms"
	MilesPerHour      WindSpeedUnit = "mph"
	Knots             WindSpeedUnit = "kn"
)

// PrecipitationUnit specifies the unit for precipitation values.
type PrecipitationUnit string

const (
	Millimeters PrecipitationUnit = "mm"
	Inches      PrecipitationUnit = "inch"
)

// CellSelection specifies how grid-cells are selected.
type CellSelection string

const (
	CellSelectionLand    CellSelection = "land"
	CellSelectionSea     CellSelection = "sea"
	CellSelectionNearest CellSelection = "nearest"
)

// BaseUnits contains unit strings for metrics shared between hourly and minutely15 data.
type BaseUnits struct {
	Temperature2m            string `json:"temperature_2m,omitempty"`
	RelativeHumidity2m       string `json:"relative_humidity_2m,omitempty"`
	DewPoint2m               string `json:"dew_point_2m,omitempty"`
	ApparentTemperature      string `json:"apparent_temperature,omitempty"`
	Precipitation            string `json:"precipitation,omitempty"`
	Rain                     string `json:"rain,omitempty"`
	Snowfall                 string `json:"snowfall,omitempty"`
	WeatherCode              string `json:"weather_code,omitempty"`
	CloudCover               string `json:"cloud_cover,omitempty"`
	CloudCoverLow            string `json:"cloud_cover_low,omitempty"`
	CloudCoverMid            string `json:"cloud_cover_mid,omitempty"`
	CloudCoverHigh           string `json:"cloud_cover_high,omitempty"`
	WindSpeed10m             string `json:"wind_speed_10m,omitempty"`
	WindSpeed80m             string `json:"wind_speed_80m,omitempty"`
	WindDirection10m         string `json:"wind_direction_10m,omitempty"`
	WindDirection80m         string `json:"wind_direction_80m,omitempty"`
	WindGusts10m             string `json:"wind_gusts_10m,omitempty"`
	ShortwaveRadiation       string `json:"shortwave_radiation,omitempty"`
	DirectRadiation          string `json:"direct_radiation,omitempty"`
	DirectNormalIrradiance   string `json:"direct_normal_irradiance,omitempty"`
	DiffuseRadiation         string `json:"diffuse_radiation,omitempty"`
	GlobalTiltedIrradiance   string `json:"global_tilted_irradiance,omitempty"`
	Visibility               string `json:"visibility,omitempty"`
	Evapotranspiration       string `json:"evapotranspiration,omitempty"`
	ET0FAOEvapotranspiration string `json:"et0_fao_evapotranspiration,omitempty"`
	VapourPressureDeficit    string `json:"vapour_pressure_deficit,omitempty"`
	Cape                     string `json:"cape,omitempty"`
	FreezingLevelHeight      string `json:"freezing_level_height,omitempty"`
	SunshineDuration         string `json:"sunshine_duration,omitempty"`
}

// HourlyUnits contains unit strings for all hourly metrics.
type HourlyUnits struct {
	BaseUnits

	// Pressure
	PressureMSL     string `json:"pressure_msl,omitempty"`
	SurfacePressure string `json:"surface_pressure,omitempty"`

	// Additional wind levels
	WindSpeed120m     string `json:"wind_speed_120m,omitempty"`
	WindSpeed180m     string `json:"wind_speed_180m,omitempty"`
	WindDirection120m string `json:"wind_direction_120m,omitempty"`
	WindDirection180m string `json:"wind_direction_180m,omitempty"`

	// Snow and precipitation
	SnowDepth                  string `json:"snow_depth,omitempty"`
	PrecipitationProbability   string `json:"precipitation_probability,omitempty"`
	Showers                    string `json:"showers,omitempty"`

	// Soil
	SoilTemperature0cm   string `json:"soil_temperature_0cm,omitempty"`
	SoilTemperature6cm   string `json:"soil_temperature_6cm,omitempty"`
	SoilTemperature18cm  string `json:"soil_temperature_18cm,omitempty"`
	SoilTemperature54cm  string `json:"soil_temperature_54cm,omitempty"`
	SoilMoisture0to1cm   string `json:"soil_moisture_0_to_1cm,omitempty"`
	SoilMoisture1to3cm   string `json:"soil_moisture_1_to_3cm,omitempty"`
	SoilMoisture3to9cm   string `json:"soil_moisture_3_to_9cm,omitempty"`
	SoilMoisture9to27cm  string `json:"soil_moisture_9_to_27cm,omitempty"`
	SoilMoisture27to81cm string `json:"soil_moisture_27_to_81cm,omitempty"`

	// Is day
	IsDay string `json:"is_day,omitempty"`

	// Pressure level variables (selected levels)
	Temperature1000hPa       string `json:"temperature_1000hPa,omitempty"`
	Temperature975hPa        string `json:"temperature_975hPa,omitempty"`
	Temperature950hPa        string `json:"temperature_950hPa,omitempty"`
	Temperature925hPa        string `json:"temperature_925hPa,omitempty"`
	Temperature900hPa        string `json:"temperature_900hPa,omitempty"`
	Temperature850hPa        string `json:"temperature_850hPa,omitempty"`
	Temperature800hPa        string `json:"temperature_800hPa,omitempty"`
	Temperature700hPa        string `json:"temperature_700hPa,omitempty"`
	Temperature600hPa        string `json:"temperature_600hPa,omitempty"`
	Temperature500hPa        string `json:"temperature_500hPa,omitempty"`
	Temperature400hPa        string `json:"temperature_400hPa,omitempty"`
	Temperature300hPa        string `json:"temperature_300hPa,omitempty"`
	Temperature250hPa        string `json:"temperature_250hPa,omitempty"`
	Temperature200hPa        string `json:"temperature_200hPa,omitempty"`
	Temperature150hPa        string `json:"temperature_150hPa,omitempty"`
	Temperature100hPa        string `json:"temperature_100hPa,omitempty"`
	Temperature70hPa         string `json:"temperature_70hPa,omitempty"`
	Temperature50hPa         string `json:"temperature_50hPa,omitempty"`
	Temperature30hPa         string `json:"temperature_30hPa,omitempty"`

	RelativeHumidity1000hPa  string `json:"relative_humidity_1000hPa,omitempty"`
	RelativeHumidity975hPa   string `json:"relative_humidity_975hPa,omitempty"`
	RelativeHumidity950hPa   string `json:"relative_humidity_950hPa,omitempty"`
	RelativeHumidity925hPa   string `json:"relative_humidity_925hPa,omitempty"`
	RelativeHumidity900hPa   string `json:"relative_humidity_900hPa,omitempty"`
	RelativeHumidity850hPa   string `json:"relative_humidity_850hPa,omitempty"`
	RelativeHumidity800hPa   string `json:"relative_humidity_800hPa,omitempty"`
	RelativeHumidity700hPa   string `json:"relative_humidity_700hPa,omitempty"`
	RelativeHumidity600hPa   string `json:"relative_humidity_600hPa,omitempty"`
	RelativeHumidity500hPa   string `json:"relative_humidity_500hPa,omitempty"`
	RelativeHumidity400hPa   string `json:"relative_humidity_400hPa,omitempty"`
	RelativeHumidity300hPa   string `json:"relative_humidity_300hPa,omitempty"`
	RelativeHumidity250hPa   string `json:"relative_humidity_250hPa,omitempty"`
	RelativeHumidity200hPa   string `json:"relative_humidity_200hPa,omitempty"`
	RelativeHumidity150hPa   string `json:"relative_humidity_150hPa,omitempty"`
	RelativeHumidity100hPa   string `json:"relative_humidity_100hPa,omitempty"`
	RelativeHumidity70hPa    string `json:"relative_humidity_70hPa,omitempty"`
	RelativeHumidity50hPa    string `json:"relative_humidity_50hPa,omitempty"`
	RelativeHumidity30hPa    string `json:"relative_humidity_30hPa,omitempty"`

	CloudCover1000hPa        string `json:"cloud_cover_1000hPa,omitempty"`
	CloudCover975hPa         string `json:"cloud_cover_975hPa,omitempty"`
	CloudCover950hPa         string `json:"cloud_cover_950hPa,omitempty"`
	CloudCover925hPa         string `json:"cloud_cover_925hPa,omitempty"`
	CloudCover900hPa         string `json:"cloud_cover_900hPa,omitempty"`
	CloudCover850hPa         string `json:"cloud_cover_850hPa,omitempty"`
	CloudCover800hPa         string `json:"cloud_cover_800hPa,omitempty"`
	CloudCover700hPa         string `json:"cloud_cover_700hPa,omitempty"`
	CloudCover600hPa         string `json:"cloud_cover_600hPa,omitempty"`
	CloudCover500hPa         string `json:"cloud_cover_500hPa,omitempty"`
	CloudCover400hPa         string `json:"cloud_cover_400hPa,omitempty"`
	CloudCover300hPa         string `json:"cloud_cover_300hPa,omitempty"`
	CloudCover250hPa         string `json:"cloud_cover_250hPa,omitempty"`
	CloudCover200hPa         string `json:"cloud_cover_200hPa,omitempty"`
	CloudCover150hPa         string `json:"cloud_cover_150hPa,omitempty"`
	CloudCover100hPa         string `json:"cloud_cover_100hPa,omitempty"`
	CloudCover70hPa          string `json:"cloud_cover_70hPa,omitempty"`
	CloudCover50hPa          string `json:"cloud_cover_50hPa,omitempty"`
	CloudCover30hPa          string `json:"cloud_cover_30hPa,omitempty"`

	WindSpeed1000hPa         string `json:"wind_speed_1000hPa,omitempty"`
	WindSpeed975hPa          string `json:"wind_speed_975hPa,omitempty"`
	WindSpeed950hPa          string `json:"wind_speed_950hPa,omitempty"`
	WindSpeed925hPa          string `json:"wind_speed_925hPa,omitempty"`
	WindSpeed900hPa          string `json:"wind_speed_900hPa,omitempty"`
	WindSpeed850hPa          string `json:"wind_speed_850hPa,omitempty"`
	WindSpeed800hPa          string `json:"wind_speed_800hPa,omitempty"`
	WindSpeed700hPa          string `json:"wind_speed_700hPa,omitempty"`
	WindSpeed600hPa          string `json:"wind_speed_600hPa,omitempty"`
	WindSpeed500hPa          string `json:"wind_speed_500hPa,omitempty"`
	WindSpeed400hPa          string `json:"wind_speed_400hPa,omitempty"`
	WindSpeed300hPa          string `json:"wind_speed_300hPa,omitempty"`
	WindSpeed250hPa          string `json:"wind_speed_250hPa,omitempty"`
	WindSpeed200hPa          string `json:"wind_speed_200hPa,omitempty"`
	WindSpeed150hPa          string `json:"wind_speed_150hPa,omitempty"`
	WindSpeed100hPa          string `json:"wind_speed_100hPa,omitempty"`
	WindSpeed70hPa           string `json:"wind_speed_70hPa,omitempty"`
	WindSpeed50hPa           string `json:"wind_speed_50hPa,omitempty"`
	WindSpeed30hPa           string `json:"wind_speed_30hPa,omitempty"`

	WindDirection1000hPa     string `json:"wind_direction_1000hPa,omitempty"`
	WindDirection975hPa      string `json:"wind_direction_975hPa,omitempty"`
	WindDirection950hPa      string `json:"wind_direction_950hPa,omitempty"`
	WindDirection925hPa      string `json:"wind_direction_925hPa,omitempty"`
	WindDirection900hPa      string `json:"wind_direction_900hPa,omitempty"`
	WindDirection850hPa      string `json:"wind_direction_850hPa,omitempty"`
	WindDirection800hPa      string `json:"wind_direction_800hPa,omitempty"`
	WindDirection700hPa      string `json:"wind_direction_700hPa,omitempty"`
	WindDirection600hPa      string `json:"wind_direction_600hPa,omitempty"`
	WindDirection500hPa      string `json:"wind_direction_500hPa,omitempty"`
	WindDirection400hPa      string `json:"wind_direction_400hPa,omitempty"`
	WindDirection300hPa      string `json:"wind_direction_300hPa,omitempty"`
	WindDirection250hPa      string `json:"wind_direction_250hPa,omitempty"`
	WindDirection200hPa      string `json:"wind_direction_200hPa,omitempty"`
	WindDirection150hPa      string `json:"wind_direction_150hPa,omitempty"`
	WindDirection100hPa      string `json:"wind_direction_100hPa,omitempty"`
	WindDirection70hPa       string `json:"wind_direction_70hPa,omitempty"`
	WindDirection50hPa       string `json:"wind_direction_50hPa,omitempty"`
	WindDirection30hPa       string `json:"wind_direction_30hPa,omitempty"`

	GeopotentialHeight1000hPa string `json:"geopotential_height_1000hPa,omitempty"`
	GeopotentialHeight975hPa  string `json:"geopotential_height_975hPa,omitempty"`
	GeopotentialHeight950hPa  string `json:"geopotential_height_950hPa,omitempty"`
	GeopotentialHeight925hPa  string `json:"geopotential_height_925hPa,omitempty"`
	GeopotentialHeight900hPa  string `json:"geopotential_height_900hPa,omitempty"`
	GeopotentialHeight850hPa  string `json:"geopotential_height_850hPa,omitempty"`
	GeopotentialHeight800hPa  string `json:"geopotential_height_800hPa,omitempty"`
	GeopotentialHeight700hPa  string `json:"geopotential_height_700hPa,omitempty"`
	GeopotentialHeight600hPa  string `json:"geopotential_height_600hPa,omitempty"`
	GeopotentialHeight500hPa  string `json:"geopotential_height_500hPa,omitempty"`
	GeopotentialHeight400hPa  string `json:"geopotential_height_400hPa,omitempty"`
	GeopotentialHeight300hPa  string `json:"geopotential_height_300hPa,omitempty"`
	GeopotentialHeight250hPa  string `json:"geopotential_height_250hPa,omitempty"`
	GeopotentialHeight200hPa  string `json:"geopotential_height_200hPa,omitempty"`
	GeopotentialHeight150hPa  string `json:"geopotential_height_150hPa,omitempty"`
	GeopotentialHeight100hPa  string `json:"geopotential_height_100hPa,omitempty"`
	GeopotentialHeight70hPa   string `json:"geopotential_height_70hPa,omitempty"`
	GeopotentialHeight50hPa   string `json:"geopotential_height_50hPa,omitempty"`
	GeopotentialHeight30hPa   string `json:"geopotential_height_30hPa,omitempty"`

	DewPoint1000hPa          string `json:"dew_point_1000hPa,omitempty"`
	DewPoint975hPa           string `json:"dew_point_975hPa,omitempty"`
	DewPoint950hPa           string `json:"dew_point_950hPa,omitempty"`
	DewPoint925hPa           string `json:"dew_point_925hPa,omitempty"`
	DewPoint900hPa           string `json:"dew_point_900hPa,omitempty"`
	DewPoint850hPa           string `json:"dew_point_850hPa,omitempty"`
	DewPoint800hPa           string `json:"dew_point_800hPa,omitempty"`
	DewPoint700hPa           string `json:"dew_point_700hPa,omitempty"`
	DewPoint600hPa           string `json:"dew_point_600hPa,omitempty"`
	DewPoint500hPa           string `json:"dew_point_500hPa,omitempty"`
	DewPoint400hPa           string `json:"dew_point_400hPa,omitempty"`
	DewPoint300hPa           string `json:"dew_point_300hPa,omitempty"`
	DewPoint250hPa           string `json:"dew_point_250hPa,omitempty"`
	DewPoint200hPa           string `json:"dew_point_200hPa,omitempty"`
	DewPoint150hPa           string `json:"dew_point_150hPa,omitempty"`
	DewPoint100hPa           string `json:"dew_point_100hPa,omitempty"`
	DewPoint70hPa            string `json:"dew_point_70hPa,omitempty"`
	DewPoint50hPa            string `json:"dew_point_50hPa,omitempty"`
	DewPoint30hPa            string `json:"dew_point_30hPa,omitempty"`
}

// Minutely15Units contains unit strings for 15-minutely metrics.
type Minutely15Units struct {
	BaseUnits

	LightningPotential string `json:"lightning_potential,omitempty"`
	SnowfallHeight     string `json:"snowfall_height,omitempty"`
}

// DailyUnits contains unit strings for daily metrics.
type DailyUnits struct {
	WeatherCode                string `json:"weather_code,omitempty"`
	Temperature2mMax           string `json:"temperature_2m_max,omitempty"`
	Temperature2mMin           string `json:"temperature_2m_min,omitempty"`
	Temperature2mMean          string `json:"temperature_2m_mean,omitempty"`
	ApparentTemperatureMax     string `json:"apparent_temperature_max,omitempty"`
	ApparentTemperatureMin     string `json:"apparent_temperature_min,omitempty"`
	ApparentTemperatureMean    string `json:"apparent_temperature_mean,omitempty"`
	Sunrise                    string `json:"sunrise,omitempty"`
	Sunset                     string `json:"sunset,omitempty"`
	SunshineDuration           string `json:"sunshine_duration,omitempty"`
	DaylightDuration           string `json:"daylight_duration,omitempty"`
	PrecipitationSum           string `json:"precipitation_sum,omitempty"`
	RainSum                    string `json:"rain_sum,omitempty"`
	ShowersSum                 string `json:"showers_sum,omitempty"`
	SnowfallSum                string `json:"snowfall_sum,omitempty"`
	PrecipitationHours         string `json:"precipitation_hours,omitempty"`
	PrecipitationProbabilityMax  string `json:"precipitation_probability_max,omitempty"`
	PrecipitationProbabilityMin  string `json:"precipitation_probability_min,omitempty"`
	PrecipitationProbabilityMean string `json:"precipitation_probability_mean,omitempty"`
	WindSpeed10mMax            string `json:"wind_speed_10m_max,omitempty"`
	WindGusts10mMax            string `json:"wind_gusts_10m_max,omitempty"`
	WindDirection10mDominant   string `json:"wind_direction_10m_dominant,omitempty"`
	ShortwaveRadiationSum      string `json:"shortwave_radiation_sum,omitempty"`
	ET0FAOEvapotranspiration   string `json:"et0_fao_evapotranspiration,omitempty"`
	UVIndexMax                 string `json:"uv_index_max,omitempty"`
	UVIndexClearSkyMax         string `json:"uv_index_clear_sky_max,omitempty"`
}

// CurrentUnits contains unit strings for current weather metrics.
type CurrentUnits struct {
	Time                     string `json:"time,omitempty"`
	Interval                 string `json:"interval,omitempty"`
	Temperature2m            string `json:"temperature_2m,omitempty"`
	RelativeHumidity2m       string `json:"relative_humidity_2m,omitempty"`
	ApparentTemperature      string `json:"apparent_temperature,omitempty"`
	IsDay                    string `json:"is_day,omitempty"`
	Precipitation            string `json:"precipitation,omitempty"`
	Rain                     string `json:"rain,omitempty"`
	Showers                  string `json:"showers,omitempty"`
	Snowfall                 string `json:"snowfall,omitempty"`
	WeatherCode              string `json:"weather_code,omitempty"`
	CloudCover               string `json:"cloud_cover,omitempty"`
	PressureMSL              string `json:"pressure_msl,omitempty"`
	SurfacePressure          string `json:"surface_pressure,omitempty"`
	WindSpeed10m             string `json:"wind_speed_10m,omitempty"`
	WindDirection10m         string `json:"wind_direction_10m,omitempty"`
	WindGusts10m             string `json:"wind_gusts_10m,omitempty"`
}
