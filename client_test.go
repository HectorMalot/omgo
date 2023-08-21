package omgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestURLBuilder(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)

	loc, err := NewLocation(52.3738, 4.8910) // Amsterdam
	require.NoError(t, err)

	opts := Options{
		TemperatureUnit:   "celsius",
		WindspeedUnit:     "kmh",
		PrecipitationUnit: "mm",
		Timezone:          "UTC",
		PastDays:          1,
		HourlyMetrics:     []string{"temperature_2m", "cloudcover", "direct_radiation", "diffuse_radiation", "precipitation", "windspeed_10m"},
		DailyMetrics:      []string{"temperature_2m_max"},
	}

	url := urlFromOptions(c.URL, loc, &opts)
	require.Equal(t, "https://api.open-meteo.com/v1/forecast?latitude=52.373800&longitude=4.891000&current_weather=true&temperature_unit=celsius&windspeed_unit=kmh&precipitation_unit=mm&timezone=UTC&past_days=1&hourly=temperature_2m,cloudcover,direct_radiation,diffuse_radiation,precipitation,windspeed_10m&daily=temperature_2m_max", url)
}

func TestUrlFromHistoricalOptions(t *testing.T) {
	hc, err := NewHistoricalClient()
	require.NoError(t, err)

	loc, err := NewLocation(52.5161, 13.4104) // Berlin
	require.NoError(t, err)

	hcOpts := HistoricalOptions{
		TemperatureUnit:   "celsius",
		WindspeedUnit:     "kmh",
		PrecipitationUnit: "mm",
		Timezone:          "UTC",
		StartDate:         "2023-07-25",
		EndDate:           "2023-08-08",
		HourlyMetrics:     []string{"temperature_2m", "rain"},
		DailyMetrics:      []string{"temperature_2m_max", "temperature_2m_mean"},
	}
	url, _ := urlFromHistoricalOptions(hc.URL, loc, &hcOpts)
	require.Equal(t, "https://archive-api.open-meteo.com/v1/archive?latitude=52.516100&longitude=13.410400&start_date=2023-07-25&end_date=2023-08-08&temperature_unit=celsius&windspeed_unit=kmh&precipitation_unit=mm&timezone=UTC&hourly=temperature_2m,rain&daily=temperature_2m_max,temperature_2m_mean", url)
}
