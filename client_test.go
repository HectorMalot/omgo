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
		Metrics:           []string{"temperature_2m", "cloudcover", "direct_radiation", "diffuse_radiation", "precipitation", "windspeed_10m"},
	}

	url := urlFromOptions(c.URL, loc, &opts)
	require.Equal(t, "https://api.open-meteo.com/v1/forecast?latitude=52.373800&longitude=4.891000&current_weather=true&temperature_unit=celsius&windspeed_unit=kmh&precipitation_unit=mm&timezone=UTC&past_days=1&hourly=temperature_2m,cloudcover,direct_radiation,diffuse_radiation,precipitation,windspeed_10m", url)
}
