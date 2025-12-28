package omgo

import (
	"fmt"
	"strings"
	"time"
)

// Time format layouts used by the Open-Meteo API
const (
	timeLayoutDateTime = "2006-01-02T15:04"
	timeLayoutDate     = "2006-01-02"
)

// parseDateTime parses a datetime string in ISO8601 format.
func parseDateTime(s string, loc *time.Location) (time.Time, error) {
	s = strings.Trim(s, "\"")
	if s == "" || s == "null" {
		return time.Time{}, nil
	}
	t, err := time.ParseInLocation(timeLayoutDateTime, s, time.UTC)
	if err != nil {
		return time.Time{}, err
	}
	// Apply the location if specified
	if loc != nil && loc != time.UTC {
		y, m, d := t.Date()
		return time.Date(y, m, d, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc), nil
	}
	return t, nil
}

// parseDate parses a date string in ISO8601 format.
func parseDate(s string, loc *time.Location) (time.Time, error) {
	s = strings.Trim(s, "\"")
	if s == "" || s == "null" {
		return time.Time{}, nil
	}
	t, err := time.ParseInLocation(timeLayoutDate, s, time.UTC)
	if err != nil {
		return time.Time{}, err
	}
	// Apply the location if specified
	if loc != nil && loc != time.UTC {
		y, m, d := t.Date()
		return time.Date(y, m, d, 0, 0, 0, 0, loc), nil
	}
	return t, nil
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
