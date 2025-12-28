package omgo

import "fmt"

// Location represents a geographical coordinate for weather queries.
type Location struct {
	Latitude  float64
	Longitude float64
	Elevation *float64 // optional elevation override in meters
}

// NewLocation creates a new Location with the given latitude and longitude.
// Latitude must be between -90 and 90.
// Longitude must be between -180 and 180.
func NewLocation(lat, lon float64) (Location, error) {
	if lat < -90 || lat > 90 {
		return Location{}, fmt.Errorf("latitude must be between -90 and 90, got %f", lat)
	}
	if lon < -180 || lon > 180 {
		return Location{}, fmt.Errorf("longitude must be between -180 and 180, got %f", lon)
	}
	return Location{Latitude: lat, Longitude: lon}, nil
}

// WithElevation returns a copy of the Location with the specified elevation override.
// Elevation is in meters above sea level.
func (l Location) WithElevation(meters float64) Location {
	l.Elevation = &meters
	return l
}
