package omgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForecastUnMarshal(t *testing.T) {
	body := []byte(`{"latitude": 52.52,
		"longitude": 13.419,
		"elevation": 44.812,
		"generationtime_ms": 2.2119,
		"hourly": {
			"time": ["2021-08-28T00:00", "2021-08-28T01:00", "2021-08-28T02:00"],
			"temperature_2m": [13, 12.7, 12.7, 12.5, 12.5, 12.8, 13, 12.9, 13.3]
		  },
		"hourly_units": {
		  "temperature_2m": "°C"
		},
		"current_weather": {
		  "time": "2021-08-28T09:00",
		  "temperature": 13.3,
		  "weathercode": 3,
		  "windspeed": 10.3,
		  "winddirection": 262
		}
	  }`)

	fc, err := ParseBody(body)
	require.NoError(t, err)
	require.Equal(t, []float64{13, 12.7, 12.7, 12.5, 12.5, 12.8, 13, 12.9, 13.3}, fc.Metrics["temperature_2m"])

}
