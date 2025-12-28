package omgo

// Minutely15Data contains 15-minutely weather data.
// This data is based on NOAA HRRR for North America and
// DWD ICON-D2 / Météo-France AROME for Central Europe.
type Minutely15Data struct {
	BaseMetrics // embedded shared fields

	// Lightning potential (HRRR only)
	LightningPotential []float64 `json:"lightning_potential,omitempty"`

	// Snowfall height
	SnowfallHeight []float64 `json:"snowfall_height,omitempty"`

	// Showers (convective precipitation)
	Showers []float64 `json:"showers,omitempty"`

	// Global tilted irradiance instant
	GlobalTiltedIrradianceInstant []float64 `json:"global_tilted_irradiance_instant,omitempty"`
}
