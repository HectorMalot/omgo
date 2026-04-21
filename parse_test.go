package omgo

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseHourlyForecast(t *testing.T) {
	data, err := os.ReadFile("testdata/forecast_hourly.json")
	require.NoError(t, err)

	weather, err := parseWeatherResponse(data)
	require.NoError(t, err)

	// Check metadata
	assert.InDelta(t, 52.52, weather.Latitude, 0.01)
	assert.InDelta(t, 13.42, weather.Longitude, 0.01)
	assert.Equal(t, "Europe/Berlin", weather.Timezone)
	assert.Equal(t, "CET", weather.TimezoneAbbreviation)
	assert.Equal(t, 3600, weather.UTCOffsetSeconds)

	// Check hourly data
	require.NotNil(t, weather.Hourly)
	assert.Len(t, weather.Hourly.Times, 3)
	assert.Len(t, weather.Hourly.Temperature2m, 3)
	assert.Len(t, weather.Hourly.RelativeHumidity2m, 3)
	assert.Len(t, weather.Hourly.Precipitation, 3)
	assert.Len(t, weather.Hourly.WeatherCode, 3)
	assert.Len(t, weather.Hourly.WindSpeed10m, 3)

	// Check values
	assert.Equal(t, 2.5, weather.Hourly.Temperature2m[0])
	assert.Equal(t, WeatherCode(3), weather.Hourly.WeatherCode[0])
	assert.Equal(t, WeatherCode(61), weather.Hourly.WeatherCode[1])

	// Check time parsing with timezone
	loc, _ := time.LoadLocation("Europe/Berlin")
	expectedTime := time.Date(2024, 1, 15, 0, 0, 0, 0, loc)
	assert.Equal(t, expectedTime, weather.Hourly.Times[0])

	// Check units
	require.NotNil(t, weather.HourlyUnits)
	assert.Equal(t, "°C", weather.HourlyUnits.Temperature2m)
	assert.Equal(t, "mm", weather.HourlyUnits.Precipitation)
	assert.Equal(t, "km/h", weather.HourlyUnits.WindSpeed10m)
}

func TestParseDailyForecast(t *testing.T) {
	data, err := os.ReadFile("testdata/forecast_daily.json")
	require.NoError(t, err)

	weather, err := parseWeatherResponse(data)
	require.NoError(t, err)

	// Check daily data
	require.NotNil(t, weather.Daily)
	assert.Len(t, weather.Daily.Times, 3)
	assert.Len(t, weather.Daily.Temperature2mMax, 3)
	assert.Len(t, weather.Daily.Temperature2mMin, 3)
	assert.Len(t, weather.Daily.Sunrise, 3)
	assert.Len(t, weather.Daily.Sunset, 3)

	// Check values
	assert.Equal(t, 5.2, weather.Daily.Temperature2mMax[0])
	assert.Equal(t, -1.2, weather.Daily.Temperature2mMin[0])

	// Check time parsing
	loc, _ := time.LoadLocation("Europe/Berlin")
	expectedDate := time.Date(2024, 1, 15, 0, 0, 0, 0, loc)
	assert.Equal(t, expectedDate, weather.Daily.Times[0])

	// Check sunrise/sunset parsing
	expectedSunrise := time.Date(2024, 1, 15, 8, 15, 0, 0, loc)
	assert.Equal(t, expectedSunrise, weather.Daily.Sunrise[0])

	expectedSunset := time.Date(2024, 1, 15, 16, 30, 0, 0, loc)
	assert.Equal(t, expectedSunset, weather.Daily.Sunset[0])

	// Check units
	require.NotNil(t, weather.DailyUnits)
	assert.Equal(t, "°C", weather.DailyUnits.Temperature2mMax)
}

func TestParseForecastWithMultipleModels(t *testing.T) {
	data := []byte(`{
  "latitude": 52.52,
  "longitude": 13.419998,
  "elevation": 38.0,
  "generationtime_ms": 0.4,
  "utc_offset_seconds": 3600,
  "timezone": "Europe/Berlin",
  "timezone_abbreviation": "CET",
  "hourly": {
    "time": ["2024-01-15T00:00", "2024-01-15T01:00", "2024-01-15T02:00"],
    "temperature_2m_ecmwf_ifs": [2.5, 2.3, 2.1],
    "temperature_2m_gfs_global": [2.2, 2.0, 1.8],
    "wind_speed_10m_ecmwf_ifs": [12.5, 13.2, 11.8],
    "wind_speed_10m_gfs_global": [10.5, 11.1, 9.8]
  },
  "hourly_units": {
    "time": "iso8601",
    "temperature_2m_ecmwf_ifs": "°C",
    "temperature_2m_gfs_global": "°C",
    "wind_speed_10m_ecmwf_ifs": "km/h",
    "wind_speed_10m_gfs_global": "km/h"
  },
  "daily": {
    "time": ["2024-01-15", "2024-01-16", "2024-01-17"],
    "temperature_2m_max_ecmwf_ifs": [5.2, 6.1, 4.8],
    "temperature_2m_max_gfs_global": [4.9, 5.8, 4.4],
    "temperature_2m_min_ecmwf_ifs": [-1.2, 0.5, -0.3],
    "temperature_2m_min_gfs_global": [-1.6, 0.2, -0.7],
    "sunrise": ["2024-01-15T08:15", "2024-01-16T08:14", "2024-01-17T08:13"],
    "sunset": ["2024-01-15T16:30", "2024-01-16T16:32", "2024-01-17T16:34"]
  },
  "daily_units": {
    "time": "iso8601",
    "temperature_2m_max_ecmwf_ifs": "°C",
    "temperature_2m_max_gfs_global": "°C",
    "temperature_2m_min_ecmwf_ifs": "°C",
    "temperature_2m_min_gfs_global": "°C",
    "sunrise": "iso8601",
    "sunset": "iso8601"
  }
}`)

	weather, err := parseWeatherResponse(data)
	require.NoError(t, err)

	require.NotNil(t, weather.Hourly)
	assert.Len(t, weather.Hourly.Times, 3)
	assert.Len(t, weather.Hourly.Temperature2m, 3)
	assert.Len(t, weather.Hourly.WindSpeed10m, 3)

	require.NotNil(t, weather.HourlyByModel)
	assert.Len(t, weather.HourlyByModel, 2)
	require.NotNil(t, weather.HourlyByModel["ecmwf_ifs"])
	require.NotNil(t, weather.HourlyByModel["gfs_global"])
	assert.Len(t, weather.HourlyByModel["ecmwf_ifs"].Temperature2m, 3)
	assert.Len(t, weather.HourlyByModel["gfs_global"].Temperature2m, 3)

	require.NotNil(t, weather.HourlyUnits)
	require.NotNil(t, weather.HourlyUnitsByModel)
	assert.Equal(t, "°C", weather.HourlyUnitsByModel["ecmwf_ifs"].Temperature2m)
	assert.Equal(t, "°C", weather.HourlyUnitsByModel["gfs_global"].Temperature2m)

	require.NotNil(t, weather.Daily)
	assert.Len(t, weather.Daily.Times, 3)
	assert.Len(t, weather.Daily.Temperature2mMax, 3)
	assert.Len(t, weather.Daily.Temperature2mMin, 3)
	assert.Len(t, weather.Daily.Sunrise, 3)
	assert.Len(t, weather.Daily.Sunset, 3)

	require.NotNil(t, weather.DailyByModel)
	assert.Len(t, weather.DailyByModel, 2)
	assert.Len(t, weather.DailyByModel["ecmwf_ifs"].Temperature2mMax, 3)
	assert.Len(t, weather.DailyByModel["gfs_global"].Temperature2mMax, 3)

	require.NotNil(t, weather.DailyUnits)
	require.NotNil(t, weather.DailyUnitsByModel)
	assert.Equal(t, "°C", weather.DailyUnitsByModel["ecmwf_ifs"].Temperature2mMax)
	assert.Equal(t, "°C", weather.DailyUnitsByModel["gfs_global"].Temperature2mMax)
}

func TestParseCurrentWeather(t *testing.T) {
	data, err := os.ReadFile("testdata/forecast_current.json")
	require.NoError(t, err)

	weather, err := parseWeatherResponse(data)
	require.NoError(t, err)

	// Check current data
	require.NotNil(t, weather.Current)
	assert.Equal(t, 900, weather.Current.Interval)

	require.NotNil(t, weather.Current.Temperature2m)
	assert.Equal(t, 3.5, *weather.Current.Temperature2m)

	require.NotNil(t, weather.Current.IsDay)
	assert.Equal(t, 1, *weather.Current.IsDay)
	assert.True(t, weather.Current.IsDaytime())

	require.NotNil(t, weather.Current.WeatherCode)
	assert.Equal(t, PartlyCloudy, *weather.Current.WeatherCode)

	// Check time
	loc, _ := time.LoadLocation("Europe/Berlin")
	expectedTime := time.Date(2024, 1, 15, 14, 0, 0, 0, loc)
	assert.Equal(t, expectedTime, weather.Current.Time)

	// Check units
	require.NotNil(t, weather.CurrentUnits)
	assert.Equal(t, "°C", weather.CurrentUnits.Temperature2m)
}

func TestParseHistorical(t *testing.T) {
	data, err := os.ReadFile("testdata/historical.json")
	require.NoError(t, err)

	weather, err := parseWeatherResponse(data)
	require.NoError(t, err)

	require.NotNil(t, weather.Hourly)
	assert.Len(t, weather.Hourly.Times, 3)
	assert.Equal(t, 18.5, weather.Hourly.Temperature2m[0])
}

func TestFindOptionsChunkURLFromDocsHTML(t *testing.T) {
	html := `<link href="/_app/immutable/chunks/options.TESTHASH.chunk.js" rel="modulepreload">`
	u, ok := findOptionsChunkURLFromHTML(html)
	require.True(t, ok)
	assert.Equal(t, "https://open-meteo.com/_app/immutable/chunks/options.TESTHASH.chunk.js", u)
}

func TestParseModelsFromOptionsChunkJS(t *testing.T) {
	js := `var x=1,u=[[{value:` + "`best_match`" + `,label:` + "`Best match`" + `},{value:` + "`ecmwf_ifs`" + `,label:` + "`ECMWF IFS`" + `}],[{value:` + "`metno_nordic`" + `,label:` + "`MET Norway Nordic`" + `}]],d=[{value:` + "`7`" + `,label:` + "`7 days`" + `}];`
	models, ok := parseModelsFromOptionsChunkJS(js)
	require.True(t, ok)
	assert.Equal(t, []Model{"best_match", "ecmwf_ifs", "metno_nordic"}, models)
}

func TestParseModelsFromS3ListXML(t *testing.T) {
	doc := `<?xml version="1.0" encoding="UTF-8"?>
<ListBucketResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  <CommonPrefixes><Prefix>data/bom_access_global/</Prefix></CommonPrefixes>
  <CommonPrefixes><Prefix>data/ecmwf_ifs/</Prefix></CommonPrefixes>
  <CommonPrefixes><Prefix>data/metno_nordic_pp/</Prefix></CommonPrefixes>
</ListBucketResult>`

	models, ok := parseModelsFromS3ListXML(doc)
	require.True(t, ok)
	assert.Equal(t, []Model{"bom_access_global", "ecmwf_ifs", "metno_nordic_pp"}, models)
}

type stubHTTPClient struct {
	do func(*http.Request) (*http.Response, error)
}

func (s *stubHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return s.do(req)
}

func TestClientFetchForecastModels(t *testing.T) {
	htmlNoModels := `<html></html>`
	htmlWithModels := `<link href="/_app/immutable/chunks/options.TESTHASH.chunk.js" rel="modulepreload">`
	js := `var x=1,u=[[{value:` + "`best_match`" + `,label:` + "`Best match`" + `},{value:` + "`ecmwf_ifs`" + `,label:` + "`ECMWF IFS`" + `}],[{value:` + "`metno_nordic`" + `,label:` + "`MET Norway Nordic`" + `}]],d=[{value:` + "`7`" + `,label:` + "`7 days`" + `}];`

	hc := &stubHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			switch req.URL.String() {
			case modelUpdatesURL:
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(htmlNoModels)),
					Header:     make(http.Header),
				}, nil
			case docsURL:
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(htmlWithModels)),
					Header:     make(http.Header),
				}, nil
			case "https://open-meteo.com/_app/immutable/chunks/options.TESTHASH.chunk.js":
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(js)),
					Header:     make(http.Header),
				}, nil
			default:
				return &http.Response{
					StatusCode: http.StatusNotFound,
					Body:       io.NopCloser(strings.NewReader("not found")),
					Header:     make(http.Header),
				}, nil
			}
		},
	}

	client := NewClient(WithHTTPClient(hc))
	models, err := client.FetchForecastModels(context.Background())
	require.NoError(t, err)
	assert.Equal(t, []Model{"best_match", "ecmwf_ifs", "metno_nordic"}, models)
}

func TestClientFetchAllModels(t *testing.T) {
	xmlDoc := `<?xml version="1.0" encoding="UTF-8"?>
<ListBucketResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  <CommonPrefixes><Prefix>data/bom_access_global/</Prefix></CommonPrefixes>
  <CommonPrefixes><Prefix>data/ecmwf_ifs/</Prefix></CommonPrefixes>
</ListBucketResult>`

	hc := &stubHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			require.Equal(t, openDataModelsListURL, req.URL.String())
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(xmlDoc)),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := NewClient(WithHTTPClient(hc))
	models, err := client.FetchAllModels(context.Background())
	require.NoError(t, err)
	assert.Equal(t, []Model{"bom_access_global", "ecmwf_ifs"}, models)
}

func TestClientFetchForecastModelsNotFound(t *testing.T) {
	hc := &stubHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("<html/>")),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := NewClient(WithHTTPClient(hc))
	_, err := client.FetchForecastModels(context.Background())
	require.Error(t, err)
}

func TestClientFetchModelMeta(t *testing.T) {
	hc := &stubHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			require.Equal(t, "https://marine-api.open-meteo.com/data/ecmwf_wam/static/meta.json", req.URL.String())
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"data_end_time":1751266800,"last_run_availability_time":1774299092,"last_run_initialisation_time":1750960800,"last_run_modification_time":1750988228,"temporal_resolution_seconds":3600,"update_interval_seconds":21600}`)),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := NewClient(WithHTTPClient(hc))
	meta, err := client.FetchModelMeta(context.Background(), "https://marine-api.open-meteo.com/data/ecmwf_wam/static/meta.json")
	require.NoError(t, err)
	require.NotNil(t, meta.DataEndTime)
	assert.Equal(t, int64(1751266800), *meta.DataEndTime)
	require.NotNil(t, meta.TemporalResolutionSeconds)
	assert.Equal(t, int64(3600), *meta.TemporalResolutionSeconds)
}

func TestForecastRequestModelsAreRepeatedQueryParams(t *testing.T) {
	req, err := NewForecastRequest(51.505, 0.053)
	require.NoError(t, err)

	req.WithModels("b", "a", "a")
	u := req.buildURL("https://api.open-meteo.com/v1/forecast", "")

	parsed, err := url.Parse(u)
	require.NoError(t, err)

	q := parsed.Query()
	assert.Equal(t, "a,b", q.Get("models"))
	assert.Contains(t, parsed.RawQuery, "models=a,b")
}

func TestWeatherCodeString(t *testing.T) {
	tests := []struct {
		code     WeatherCode
		expected string
	}{
		{ClearSky, "Clear sky"},
		{PartlyCloudy, "Partly cloudy"},
		{Fog, "Fog"},
		{RainSlight, "Slight rain"},
		{ThunderstormSlight, "Thunderstorm"},
		{WeatherCode(999), "Unknown (999)"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.code.String())
		})
	}
}
