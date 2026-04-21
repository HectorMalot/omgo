package omgo

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
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
const DefaultUserAgent = "omgo/0.2.0"

// DefaultTimeout is the default timeout for HTTP requests.
// Users can override this by providing a custom HTTP client via WithHTTPClient.
const DefaultTimeout = 30 * time.Second

// defaultHTTPClient is the default HTTP client with a sensible timeout.
var defaultHTTPClient = &http.Client{
	Timeout: DefaultTimeout,
}

// NewClient creates a new Open-Meteo API client.
func NewClient(opts ...Option) *Client {
	c := &Client{
		forecastURL:   forecastBaseURL,
		historicalURL: historicalBaseURL,
		httpClient:    defaultHTTPClient,
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
	url := req.buildURL(c.forecastURL, c.apiKey)

	body, err := c.doRequest(ctx, url)
	if err != nil {
		return nil, err
	}

	return parseWeatherResponse(body)
}

// Historical retrieves historical weather data for the given request.
func (c *Client) Historical(ctx context.Context, req *HistoricalRequest) (*Weather, error) {
	url := req.buildURL(c.historicalURL, c.apiKey)

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

func (c *Client) FetchForecastModels(ctx context.Context) ([]Model, error) {
	var lastErr error
	for _, pageURL := range []string{modelUpdatesURL, docsURL} {
		html, err := c.doRequest(ctx, pageURL)
		if err != nil {
			lastErr = err
			continue
		}
		optionsChunkURL, ok := findOptionsChunkURLFromHTML(string(html))
		if !ok {
			lastErr = fmt.Errorf("could not locate models list in %s", pageURL)
			continue
		}

		js, err := c.doRequest(ctx, optionsChunkURL)
		if err != nil {
			lastErr = err
			continue
		}

		models, ok := parseModelsFromOptionsChunkJS(string(js))
		if !ok {
			lastErr = fmt.Errorf("could not parse models list from %s", optionsChunkURL)
			continue
		}

		return models, nil
	}

	if lastErr == nil {
		lastErr = fmt.Errorf("could not fetch models list")
	}
	return nil, lastErr
}

func (c *Client) FetchAllModels(ctx context.Context) ([]Model, error) {
	return c.FetchOpenDataModels(ctx)
}

func findOptionsChunkURLFromHTML(html string) (string, bool) {
	re := regexp.MustCompile(`href=['"](/_app/immutable/chunks/options\.[^'"]+\.js)['"]`)
	m := re.FindStringSubmatch(html)
	if len(m) < 2 {
		return "", false
	}
	return "https://open-meteo.com" + m[1], true
}

func parseModelsFromOptionsChunkJS(js string) ([]Model, bool) {
	start := strings.Index(js, "u=[[")
	if start == -1 {
		return nil, false
	}
	start += len("u=[[")
	end := strings.Index(js[start:], "]],d=")
	if end == -1 {
		return nil, false
	}
	section := js[start : start+end]

	reVal := regexp.MustCompile("value:`([a-z0-9_]+)`")
	matches := reVal.FindAllStringSubmatch(section, -1)
	if len(matches) == 0 {
		return nil, false
	}

	set := make(map[string]struct{}, len(matches))
	for _, m := range matches {
		if len(m) < 2 {
			continue
		}
		set[m[1]] = struct{}{}
	}

	out := make([]Model, 0, len(set))
	for m := range set {
		out = append(out, Model(m))
	}
	sort.Slice(out, func(i, j int) bool { return string(out[i]) < string(out[j]) })
	return out, true
}

func (c *Client) FetchOpenDataModels(ctx context.Context) ([]Model, error) {
	body, err := c.doRequest(ctx, openDataModelsListURL)
	if err != nil {
		return nil, err
	}
	models, ok := parseModelsFromS3ListXML(string(body))
	if !ok {
		return nil, fmt.Errorf("could not parse models list from %s", openDataModelsListURL)
	}
	return models, nil
}

type s3ListBucketResult struct {
	CommonPrefixes []struct {
		Prefix string `xml:"Prefix"`
	} `xml:"CommonPrefixes"`
}

func parseModelsFromS3ListXML(doc string) ([]Model, bool) {
	var res s3ListBucketResult
	if err := xml.Unmarshal([]byte(doc), &res); err != nil {
		return nil, false
	}

	if len(res.CommonPrefixes) == 0 {
		return nil, false
	}

	set := make(map[string]struct{}, len(res.CommonPrefixes))
	for _, p := range res.CommonPrefixes {
		s := strings.TrimSpace(p.Prefix)
		s = strings.TrimPrefix(s, "data/")
		s = strings.TrimSuffix(s, "/")
		if s == "" {
			continue
		}
		set[s] = struct{}{}
	}

	if len(set) == 0 {
		return nil, false
	}

	out := make([]Model, 0, len(set))
	for m := range set {
		out = append(out, Model(m))
	}
	sort.Slice(out, func(i, j int) bool { return string(out[i]) < string(out[j]) })
	return out, true
}

type ModelMeta struct {
	ChunkTimeLength           *int64  `json:"chunk_time_length,omitempty"`
	CRSWKT                    *string `json:"crs_wkt,omitempty"`
	DataEndTime               *int64  `json:"data_end_time,omitempty"`
	LastRunAvailabilityTime   *int64  `json:"last_run_availability_time,omitempty"`
	LastRunInitialisationTime *int64  `json:"last_run_initialisation_time,omitempty"`
	LastRunModificationTime   *int64  `json:"last_run_modification_time,omitempty"`
	TemporalResolutionSeconds *int64  `json:"temporal_resolution_seconds,omitempty"`
	UpdateIntervalSeconds     *int64  `json:"update_interval_seconds,omitempty"`
}

func (c *Client) FetchModelMeta(ctx context.Context, metaURL string) (*ModelMeta, error) {
	body, err := c.doRequest(ctx, metaURL)
	if err != nil {
		return nil, err
	}
	var v ModelMeta
	if err := json.Unmarshal(body, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
