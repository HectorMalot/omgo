package omgo

import "strconv"

// WeatherCode represents WMO weather interpretation codes.
// See: https://open-meteo.com/en/docs for full documentation.
type WeatherCode int

// WMO Weather interpretation codes
const (
	ClearSky              WeatherCode = 0
	MainlyClear           WeatherCode = 1
	PartlyCloudy          WeatherCode = 2
	Overcast              WeatherCode = 3
	Fog                   WeatherCode = 45
	DepositingRimeFog     WeatherCode = 48
	DrizzleLight          WeatherCode = 51
	DrizzleModerate       WeatherCode = 53
	DrizzleDense          WeatherCode = 55
	FreezingDrizzleLight  WeatherCode = 56
	FreezingDrizzleDense  WeatherCode = 57
	RainSlight            WeatherCode = 61
	RainModerate          WeatherCode = 63
	RainHeavy             WeatherCode = 65
	FreezingRainLight     WeatherCode = 66
	FreezingRainHeavy     WeatherCode = 67
	SnowFallSlight        WeatherCode = 71
	SnowFallModerate      WeatherCode = 73
	SnowFallHeavy         WeatherCode = 75
	SnowGrains            WeatherCode = 77
	RainShowersSlight     WeatherCode = 80
	RainShowersModerate   WeatherCode = 81
	RainShowersViolent    WeatherCode = 82
	SnowShowersSlight     WeatherCode = 85
	SnowShowersHeavy      WeatherCode = 86
	ThunderstormSlight    WeatherCode = 95
	ThunderstormWithHailSlight WeatherCode = 96
	ThunderstormWithHailHeavy  WeatherCode = 99
)

// String returns a human-readable description of the weather code.
func (w WeatherCode) String() string {
	switch w {
	case ClearSky:
		return "Clear sky"
	case MainlyClear:
		return "Mainly clear"
	case PartlyCloudy:
		return "Partly cloudy"
	case Overcast:
		return "Overcast"
	case Fog:
		return "Fog"
	case DepositingRimeFog:
		return "Depositing rime fog"
	case DrizzleLight:
		return "Light drizzle"
	case DrizzleModerate:
		return "Moderate drizzle"
	case DrizzleDense:
		return "Dense drizzle"
	case FreezingDrizzleLight:
		return "Light freezing drizzle"
	case FreezingDrizzleDense:
		return "Dense freezing drizzle"
	case RainSlight:
		return "Slight rain"
	case RainModerate:
		return "Moderate rain"
	case RainHeavy:
		return "Heavy rain"
	case FreezingRainLight:
		return "Light freezing rain"
	case FreezingRainHeavy:
		return "Heavy freezing rain"
	case SnowFallSlight:
		return "Slight snow fall"
	case SnowFallModerate:
		return "Moderate snow fall"
	case SnowFallHeavy:
		return "Heavy snow fall"
	case SnowGrains:
		return "Snow grains"
	case RainShowersSlight:
		return "Slight rain showers"
	case RainShowersModerate:
		return "Moderate rain showers"
	case RainShowersViolent:
		return "Violent rain showers"
	case SnowShowersSlight:
		return "Slight snow showers"
	case SnowShowersHeavy:
		return "Heavy snow showers"
	case ThunderstormSlight:
		return "Thunderstorm"
	case ThunderstormWithHailSlight:
		return "Thunderstorm with slight hail"
	case ThunderstormWithHailHeavy:
		return "Thunderstorm with heavy hail"
	default:
		return "Unknown (" + strconv.Itoa(int(w)) + ")"
	}
}

// Description returns a longer description of the weather code.
func (w WeatherCode) Description() string {
	switch w {
	case ClearSky:
		return "Clear sky with no clouds"
	case MainlyClear:
		return "Mainly clear skies with minimal cloud cover"
	case PartlyCloudy:
		return "Partly cloudy skies"
	case Overcast:
		return "Overcast with full cloud cover"
	case Fog:
		return "Foggy conditions with reduced visibility"
	case DepositingRimeFog:
		return "Fog depositing rime ice on surfaces"
	case DrizzleLight:
		return "Light drizzle with fine water droplets"
	case DrizzleModerate:
		return "Moderate drizzle"
	case DrizzleDense:
		return "Dense drizzle with heavier water droplets"
	case FreezingDrizzleLight:
		return "Light freezing drizzle that may cause ice"
	case FreezingDrizzleDense:
		return "Dense freezing drizzle with significant icing risk"
	case RainSlight:
		return "Slight rain"
	case RainModerate:
		return "Moderate rain"
	case RainHeavy:
		return "Heavy rain with high precipitation"
	case FreezingRainLight:
		return "Light freezing rain that may cause ice accumulation"
	case FreezingRainHeavy:
		return "Heavy freezing rain with significant ice accumulation"
	case SnowFallSlight:
		return "Slight snow fall"
	case SnowFallModerate:
		return "Moderate snow fall"
	case SnowFallHeavy:
		return "Heavy snow fall with significant accumulation"
	case SnowGrains:
		return "Snow grains - small, white, opaque ice particles"
	case RainShowersSlight:
		return "Slight rain showers"
	case RainShowersModerate:
		return "Moderate rain showers"
	case RainShowersViolent:
		return "Violent rain showers with intense precipitation"
	case SnowShowersSlight:
		return "Slight snow showers"
	case SnowShowersHeavy:
		return "Heavy snow showers"
	case ThunderstormSlight:
		return "Thunderstorm with lightning"
	case ThunderstormWithHailSlight:
		return "Thunderstorm with slight hail"
	case ThunderstormWithHailHeavy:
		return "Thunderstorm with heavy hail - take shelter"
	default:
		return "Unknown weather condition (code " + strconv.Itoa(int(w)) + ")"
	}
}

// IsDay returns true if the value represents daytime.
// This is a convenience for the is_day field which is 0 or 1.
func IsDay(value int) bool {
	return value == 1
}
