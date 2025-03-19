package omgo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestForecastUnmarshalWithHourlyValues(t *testing.T) {
	body := []byte(`{"latitude": 52.52,
		"longitude": 13.419,
		"elevation": 44.812,
		"timezone": "UTC",
		"generationtime_ms": 2.2119,
		"hourly": {
			"time": ["2021-08-28T00:00", "2021-08-28T01:00", "2021-08-28T02:00"],
			"temperature_2m": [13, 12.7, 12.7]
		  },
		"hourly_units": {
		  "temperature_2m": "°C"
		},
		"current_weather": {
		  "time": "2021-08-28T09:00",
		  "temperature": 13.3,
		  "is_day": 1,
		  "weathercode": 3,
		  "windspeed": 10.3,
		  "winddirection": 262
		}
	  }`)

	fc, err := ParseBody(body)
	require.NoError(t, err)
	require.Equal(t, []float64{13, 12.7, 12.7}, fc.HourlyMetrics["temperature_2m"])
	require.Equal(t, float64(262), fc.CurrentWeather.WindDirection)
	require.Equal(t,
		[]time.Time{
			time.Date(2021, time.August, 28, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.August, 28, 1, 0, 0, 0, time.UTC),
			time.Date(2021, time.August, 28, 2, 0, 0, 0, time.UTC)},
		fc.HourlyTimes)
	require.Equal(t, float64(1), fc.CurrentWeather.IsDay)
}

func TestForecastUnmarshalWithHourlyValuesInTimezone(t *testing.T) {
	body := []byte(`{"latitude": 52.52,
		"longitude": 13.419,
		"elevation": 44.812,
		"timezone": "US/Pacific",
		"generationtime_ms": 2.2119,
		"hourly": {
			"time": ["2021-08-28T00:00", "2021-08-28T01:00", "2021-08-28T02:00"],
			"temperature_2m": [13, 12.7, 12.7]
		  },
		"hourly_units": {
		  "temperature_2m": "°C"
		},
		"current_weather": {
		  "time": "2021-08-28T09:00",
		  "temperature": 13.3,
		  "is_day": 1,
		  "weathercode": 3,
		  "windspeed": 10.3,
		  "winddirection": 262
		}
	  }`)

	loc, err := time.LoadLocation("US/Pacific")
	require.NoError(t, err)

	fc, err := ParseBody(body)
	require.NoError(t, err)
	require.Equal(t, []float64{13, 12.7, 12.7}, fc.HourlyMetrics["temperature_2m"])
	require.Equal(t, float64(262), fc.CurrentWeather.WindDirection)
	require.Equal(t,
		[]time.Time{
			time.Date(2021, time.August, 28, 0, 0, 0, 0, loc),
			time.Date(2021, time.August, 28, 1, 0, 0, 0, loc),
			time.Date(2021, time.August, 28, 2, 0, 0, 0, loc)},
		fc.HourlyTimes)
	require.Equal(t, float64(1), fc.CurrentWeather.IsDay)
}

func TestForecastUnmarshalWithDailyValues(t *testing.T) {
	body := []byte(`{
		"utc_offset_seconds": 7200,
		"elevation": 44.8125,
		"daily": {
		  "apparent_temperature_max": [14.1, 12.9, 14.8, 15.1, 15, 17.3, 18.5],
		  "time": [
			"2021-09-20",
			"2021-09-21",
			"2021-09-22",
			"2021-09-23",
			"2021-09-24",
			"2021-09-25",
			"2021-09-26"
		  ]
		},
		"daily_units": { "apparent_temperature_max": "°C" },
		"current_weather": {
		  "winddirection": 330.0,
		  "time": "2021-09-20T23:00",
		  "temperature": 12.2,
		  "weathercode": 3,
		  "windspeed": 4
		},
		"longitude": 13.419998,
		"hourly": {
		  "time": [
			"2021-09-20T00:00",
			"2021-09-20T01:00",
			"2021-09-20T02:00",
			"2021-09-20T03:00",
			"2021-09-20T04:00",
			"2021-09-20T05:00",
			"2021-09-20T06:00",
			"2021-09-20T07:00",
			"2021-09-20T08:00",
			"2021-09-20T09:00",
			"2021-09-20T10:00"
		  ],
		  "temperature_2m": [
			11.4, 11.1, 11.6, 11, 10.8, 10.5, 10.5, 10.5, 10.8, 11.3, 12.2
		  ]
		},
		"latitude": 52.52,
		"generationtime_ms": 3.193020820617676,
		"hourly_units": { "temperature_2m": "°C" }
	  }`)

	fc, err := ParseBody(body)
	require.NoError(t, err)
	require.Equal(t, []float64{11.4, 11.1, 11.6, 11, 10.8, 10.5, 10.5, 10.5, 10.8, 11.3, 12.2}, fc.HourlyMetrics["temperature_2m"])
	require.Equal(t, []float64{14.1, 12.9, 14.8, 15.1, 15, 17.3, 18.5}, fc.DailyMetrics["apparent_temperature_max"])
	require.Equal(t,
		[]time.Time{
			time.Date(2021, time.September, 20, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.September, 21, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.September, 22, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.September, 23, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.September, 24, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.September, 25, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.September, 26, 0, 0, 0, 0, time.UTC)},
		fc.DailyTimes)
}
