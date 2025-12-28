package omgo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockHTTPClient is a mock HTTP client for testing.
type mockHTTPClient struct {
	response *http.Response
	err      error
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.response, m.err
}

func newMockResponse(statusCode int, body []byte) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(bytes.NewReader(body)),
	}
}

func TestClientForecast(t *testing.T) {
	data, err := os.ReadFile("testdata/forecast_hourly.json")
	require.NoError(t, err)

	mock := &mockHTTPClient{
		response: newMockResponse(http.StatusOK, data),
	}

	client := NewClient(WithHTTPClient(mock))

	req, err := NewForecastRequest(52.52, 13.41)
	require.NoError(t, err)
	req.WithHourly(HourlyTemperature2m)

	weather, err := client.Forecast(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, weather)

	assert.InDelta(t, 52.52, weather.Latitude, 0.01)
	require.NotNil(t, weather.Hourly)
	assert.Len(t, weather.Hourly.Temperature2m, 3)
}

func TestClientHistorical(t *testing.T) {
	data, err := os.ReadFile("testdata/historical.json")
	require.NoError(t, err)

	mock := &mockHTTPClient{
		response: newMockResponse(http.StatusOK, data),
	}

	client := NewClient(WithHTTPClient(mock))

	req, err := NewHistoricalRequest(52.52, 13.41, "2023-06-01", "2023-06-01")
	require.NoError(t, err)
	req.WithHourly(HourlyTemperature2m)

	weather, err := client.Historical(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, weather)

	require.NotNil(t, weather.Hourly)
	assert.Equal(t, 18.5, weather.Hourly.Temperature2m[0])
}

func TestClientAPIError(t *testing.T) {
	data, err := os.ReadFile("testdata/error.json")
	require.NoError(t, err)

	mock := &mockHTTPClient{
		response: newMockResponse(http.StatusBadRequest, data),
	}

	client := NewClient(WithHTTPClient(mock))

	req, err := NewForecastRequest(52.52, 13.41)
	require.NoError(t, err)

	_, err = client.Forecast(context.Background(), req)
	require.Error(t, err)

	var apiErr *APIError
	require.ErrorAs(t, err, &apiErr)
	assert.Equal(t, http.StatusBadRequest, apiErr.StatusCode)
	assert.Contains(t, apiErr.Reason, "tempeture_2m")
}

func TestClientOptions(t *testing.T) {
	client := NewClient(
		WithForecastURL("https://custom-api.example.com/forecast"),
		WithHistoricalURL("https://custom-archive.example.com/archive"),
		WithUserAgent("CustomAgent/1.0"),
		WithAPIKey("test-api-key"),
	)

	assert.Equal(t, "https://custom-api.example.com/forecast", client.forecastURL)
	assert.Equal(t, "https://custom-archive.example.com/archive", client.historicalURL)
	assert.Equal(t, "CustomAgent/1.0", client.userAgent)
	assert.Equal(t, "test-api-key", client.apiKey)
}

func TestNewClientDefaults(t *testing.T) {
	client := NewClient()

	assert.Equal(t, forecastBaseURL, client.forecastURL)
	assert.Equal(t, historicalBaseURL, client.historicalURL)
	assert.Equal(t, DefaultUserAgent, client.userAgent)
	assert.Empty(t, client.apiKey)
	assert.NotNil(t, client.httpClient)
}
