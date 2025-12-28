package omgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HTTPClient is an interface for making HTTP requests.
// This allows for easy mocking in tests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the Open-Meteo API client.
type Client struct {
	forecastURL   string
	historicalURL string
	httpClient    HTTPClient
	userAgent     string
	apiKey        string
}

// Option is a functional option for configuring the Client.
type Option func(*Client)

// DefaultUserAgent is the default User-Agent header sent with requests.
const DefaultUserAgent = "omgo/0.1.0"

// NewClient creates a new Open-Meteo API client.
func NewClient(opts ...Option) *Client {
	c := &Client{
		forecastURL:   forecastBaseURL,
		historicalURL: historicalBaseURL,
		httpClient:    http.DefaultClient,
		userAgent:     DefaultUserAgent,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// WithForecastURL sets a custom base URL for the Forecast API.
func WithForecastURL(url string) Option {
	return func(c *Client) {
		c.forecastURL = url
	}
}

// WithHistoricalURL sets a custom base URL for the Historical API.
func WithHistoricalURL(url string) Option {
	return func(c *Client) {
		c.historicalURL = url
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(hc HTTPClient) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// WithUserAgent sets a custom User-Agent header.
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = ua
	}
}

// WithAPIKey sets the API key for commercial access.
// Note: When using an API key, you should also set custom URLs
// with the "customer-" prefix using WithForecastURL/WithHistoricalURL.
func WithAPIKey(key string) Option {
	return func(c *Client) {
		c.apiKey = key
	}
}

// Forecast retrieves weather forecast data for the given request.
func (c *Client) Forecast(ctx context.Context, req *ForecastRequest) (*Weather, error) {
	url := req.buildURL(c.forecastURL)
	if c.apiKey != "" {
		url += "&apikey=" + c.apiKey
	}

	body, err := c.doRequest(ctx, url)
	if err != nil {
		return nil, err
	}

	return parseWeatherResponse(body)
}

// Historical retrieves historical weather data for the given request.
func (c *Client) Historical(ctx context.Context, req *HistoricalRequest) (*Weather, error) {
	url := req.buildURL(c.historicalURL)
	if c.apiKey != "" {
		url += "&apikey=" + c.apiKey
	}

	body, err := c.doRequest(ctx, url)
	if err != nil {
		return nil, err
	}

	return parseWeatherResponse(body)
}

// doRequest performs an HTTP GET request and returns the response body.
func (c *Client) doRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		// Try to parse the error response
		var errResp apiErrorResponse
		if json.Unmarshal(body, &errResp) == nil && errResp.Error {
			return nil, &APIError{
				StatusCode: resp.StatusCode,
				Reason:     errResp.Reason,
			}
		}
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Reason:     string(body),
		}
	}

	return body, nil
}
