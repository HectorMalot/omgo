package omgo

// Weather contains the response from the Open-Meteo API.
// This is the main response type for both Forecast and Historical requests.
type Weather struct {
	// Location information
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Elevation float64 `json:"elevation"`

	// Timezone information
	Timezone             string `json:"timezone"`
	TimezoneAbbreviation string `json:"timezone_abbreviation"`
	UTCOffsetSeconds     int    `json:"utc_offset_seconds"`

	// Generation time for performance monitoring
	GenerationTimeMs float64 `json:"generationtime_ms"`

	// Current weather conditions
	Current      *CurrentData  `json:"-"` // parsed separately
	CurrentUnits *CurrentUnits `json:"current_units,omitempty"`

	// Hourly data
	Hourly      *HourlyData  `json:"-"` // parsed separately
	HourlyUnits *HourlyUnits `json:"hourly_units,omitempty"`

	// 15-minutely data
	Minutely15      *Minutely15Data  `json:"-"` // parsed separately
	Minutely15Units *Minutely15Units `json:"minutely_15_units,omitempty"`

	// Daily data
	Daily      *DailyData  `json:"-"` // parsed separately
	DailyUnits *DailyUnits `json:"daily_units,omitempty"`
}

