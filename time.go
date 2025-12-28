package omgo

import (
	"fmt"
	"time"
)

// Time format layouts used by the Open-Meteo API
const (
	timeLayoutDateTime = "2006-01-02T15:04"
	timeLayoutDate     = "2006-01-02"
)

// parseDateTime parses a datetime string in ISO8601 format.
// If loc is nil, defaults to UTC.
func parseDateTime(s string, loc *time.Location) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}
	if loc == nil {
		loc = time.UTC
	}
	return time.ParseInLocation(timeLayoutDateTime, s, loc)
}

// parseDate parses a date string in ISO8601 format.
// If loc is nil, defaults to UTC.
func parseDate(s string, loc *time.Location) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}
	if loc == nil {
		loc = time.UTC
	}
	return time.ParseInLocation(timeLayoutDate, s, loc)
}

// parseDateTimeArray parses an array of datetime strings.
func parseDateTimeArray(arr []string, loc *time.Location) ([]time.Time, error) {
	result := make([]time.Time, len(arr))
	for i, s := range arr {
		t, err := parseDateTime(s, loc)
		if err != nil {
			return nil, fmt.Errorf("parsing time at index %d: %w", i, err)
		}
		result[i] = t
	}
	return result, nil
}

// parseDateArray parses an array of date strings.
func parseDateArray(arr []string, loc *time.Location) ([]time.Time, error) {
	result := make([]time.Time, len(arr))
	for i, s := range arr {
		t, err := parseDate(s, loc)
		if err != nil {
			return nil, fmt.Errorf("parsing date at index %d: %w", i, err)
		}
		result[i] = t
	}
	return result, nil
}
