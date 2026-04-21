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
	Hourly             *HourlyData             `json:"-"` // parsed separately
	HourlyUnits        *HourlyUnits            `json:"hourly_units,omitempty"`
	HourlyByModel      map[string]*HourlyData  `json:"-"`
	HourlyUnitsByModel map[string]*HourlyUnits `json:"-"`

	// 15-minutely data
	Minutely15             *Minutely15Data             `json:"-"` // parsed separately
	Minutely15Units        *Minutely15Units            `json:"minutely_15_units,omitempty"`
	Minutely15ByModel      map[string]*Minutely15Data  `json:"-"`
	Minutely15UnitsByModel map[string]*Minutely15Units `json:"-"`

	// Daily data
	Daily             *DailyData             `json:"-"` // parsed separately
	DailyUnits        *DailyUnits            `json:"daily_units,omitempty"`
	DailyByModel      map[string]*DailyData  `json:"-"`
	DailyUnitsByModel map[string]*DailyUnits `json:"-"`
}

func (w *Weather) HourlyForModel(model string) (*HourlyData, bool) {
	if w == nil {
		return nil, false
	}
	if model != "" && w.HourlyByModel != nil {
		if data, ok := w.HourlyByModel[model]; ok && data != nil {
			return data, true
		}
	}
	return w.Hourly, model == ""
}

func (w *Weather) HourlyUnitsForModel(model string) (*HourlyUnits, bool) {
	if w == nil {
		return nil, false
	}
	if model != "" && w.HourlyUnitsByModel != nil {
		if units, ok := w.HourlyUnitsByModel[model]; ok && units != nil {
			return units, true
		}
	}
	return w.HourlyUnits, model == ""
}

func (w *Weather) Minutely15ForModel(model string) (*Minutely15Data, bool) {
	if w == nil {
		return nil, false
	}
	if model != "" && w.Minutely15ByModel != nil {
		if data, ok := w.Minutely15ByModel[model]; ok && data != nil {
			return data, true
		}
	}
	return w.Minutely15, model == ""
}

func (w *Weather) Minutely15UnitsForModel(model string) (*Minutely15Units, bool) {
	if w == nil {
		return nil, false
	}
	if model != "" && w.Minutely15UnitsByModel != nil {
		if units, ok := w.Minutely15UnitsByModel[model]; ok && units != nil {
			return units, true
		}
	}
	return w.Minutely15Units, model == ""
}

func (w *Weather) DailyForModel(model string) (*DailyData, bool) {
	if w == nil {
		return nil, false
	}
	if model != "" && w.DailyByModel != nil {
		if data, ok := w.DailyByModel[model]; ok && data != nil {
			return data, true
		}
	}
	return w.Daily, model == ""
}

func (w *Weather) DailyUnitsForModel(model string) (*DailyUnits, bool) {
	if w == nil {
		return nil, false
	}
	if model != "" && w.DailyUnitsByModel != nil {
		if units, ok := w.DailyUnitsByModel[model]; ok && units != nil {
			return units, true
		}
	}
	return w.DailyUnits, model == ""
}
