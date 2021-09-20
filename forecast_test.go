package omgo_test

import (
	"context"
	"testing"

	"github.com/hectormalot/omgo"
	"github.com/stretchr/testify/require"
)

func TestForecast(t *testing.T) {
	c, err := omgo.NewClient()
	require.NoError(t, err)

	loc, err := omgo.NewLocation(52.3738, 4.8910) // Amsterdam
	require.NoError(t, err)

	opts := omgo.Options{
		TemperatureUnit:   "celsius",
		WindspeedUnit:     "kmh",
		PrecipitationUnit: "mm",
		Timezone:          "UTC",
		PastDays:          0,
		Metrics:           []string{"temperature_2m"},
	}
	res, err := c.Forecast(context.Background(), loc, &opts)
	require.NoError(t, err)

	require.False(t, res.CurrentWeather.Time.IsZero())
	require.True(t, len(res.HourlyTimes) > 0)
}
