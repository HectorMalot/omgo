package omgo

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	URL       string
	UserAgent string
	Client    *http.Client
}

type HistClient struct {
	URL       string
	UserAgent string
	Client    *http.Client
}

const DefaultUserAgent = "Open-Meteo_Go_Client"

func NewClient() (Client, error) {
	return Client{
		URL:       "https://api.open-meteo.com/v1/forecast",
		UserAgent: DefaultUserAgent,
		Client:    http.DefaultClient,
	}, nil
}

func NewHistoricalClient() (HistClient, error) {
	return HistClient{
		URL:       "https://api.open-meteo.com/v1/archive",
		UserAgent: DefaultUserAgent,
		Client:    http.DefaultClient,
	}, nil
}

type Location struct {
	lat, lon float64
}

func NewLocation(lat, lon float64) (Location, error) {
	return Location{lat: lat, lon: lon}, nil
}

type Options struct {
	TemperatureUnit   string   // Default "celsius"
	WindspeedUnit     string   // Default "kmh",
	PrecipitationUnit string   // Default "mm"
	Timezone          string   // Default "UTC"
	PastDays          int      // Default 0
	HourlyMetrics     []string // Lists required hourly metrics, see https://open-meteo.com/en/docs for valid metrics
	DailyMetrics      []string // Lists required daily metrics, see https://open-meteo.com/en/docs for valid metrics
}

type HistoricalOptions struct {
	TemperatureUnit   string   // Default "celsius"
	WindspeedUnit     string   // Default "kmh",
	PrecipitationUnit string   // Default "mm"
	Timezone          string   // Default "UTC"
	PastDays          int      // Default 0
	StartDate         string   // Format yyyy-mm-dd ISO8601 date
	EndDate           string   // Format yyyy-mm-dd ISO8601 date
	HourlyMetrics     []string // Lists required hourly metrics, see https://open-meteo.com/en/docs for valid metrics
	DailyMetrics      []string // Lists required daily metrics, see https://open-meteo.com/en/docs for valid metrics
}

func urlFromOptions(baseURL string, loc Location, opts *Options) string {
	// TODO: Validate the Options
	url := fmt.Sprintf(`%s?latitude=%f&longitude=%f&current_weather=true`, baseURL, loc.lat, loc.lon)
	if opts == nil {
		return url
	}

	if opts.TemperatureUnit != "" {
		url = fmt.Sprintf(`%s&temperature_unit=%s`, url, opts.TemperatureUnit)
	}
	if opts.WindspeedUnit != "" {
		url = fmt.Sprintf(`%s&windspeed_unit=%s`, url, opts.WindspeedUnit)
	}
	if opts.PrecipitationUnit != "" {
		url = fmt.Sprintf(`%s&precipitation_unit=%s`, url, opts.PrecipitationUnit)
	}
	if opts.Timezone != "" {
		url = fmt.Sprintf(`%s&timezone=%s`, url, opts.Timezone)
	}

	if opts.HourlyMetrics != nil && len(opts.HourlyMetrics) > 0 {
		metrics := strings.Join(opts.HourlyMetrics, ",")
		url = fmt.Sprintf(`%s&hourly=%s`, url, metrics)
	}

	if opts.DailyMetrics != nil && len(opts.DailyMetrics) > 0 {
		metrics := strings.Join(opts.DailyMetrics, ",")
		url = fmt.Sprintf(`%s&daily=%s`, url, metrics)
	}

	return url
}

func urlFromHistoricalOptions(baseURL string, loc Location, opts *HistoricalOptions) (string, error) {
	url := fmt.Sprintf(`%s?latitude=%f&longitude=%f`, baseURL, loc.lat, loc.lon)

	if opts.StartDate == "" || opts.EndDate == "" {
		return "", fmt.Errorf("please, provide a start date and end date")
	}

	if opts == nil {
		return url, nil
	}

	// TODO: if end_date is equal or less than end_date throw an error
	url = fmt.Sprintf(`%s&start_date=%s&end_date=%s`, url, opts.StartDate, opts.EndDate)

	if opts.TemperatureUnit != "" {
		url = fmt.Sprintf(`%s&temperature_unit=%s`, url, opts.TemperatureUnit)
	}
	if opts.WindspeedUnit != "" {
		url = fmt.Sprintf(`%s&windspeed_unit=%s`, url, opts.WindspeedUnit)
	}
	if opts.PrecipitationUnit != "" {
		url = fmt.Sprintf(`%s&precipitation_unit=%s`, url, opts.PrecipitationUnit)
	}
	if opts.Timezone != "" {
		url = fmt.Sprintf(`%s&timezone=%s`, url, opts.Timezone)
	}
	if opts.PastDays != 0 {
		url = fmt.Sprintf(`%s&past_days=%d`, url, opts.PastDays)
	}

	if opts.HourlyMetrics != nil && len(opts.HourlyMetrics) > 0 {
		metrics := strings.Join(opts.HourlyMetrics, ",")
		url = fmt.Sprintf(`%s&hourly=%s`, url, metrics)
	}

	if opts.DailyMetrics != nil && len(opts.DailyMetrics) > 0 {
		metrics := strings.Join(opts.DailyMetrics, ",")
		url = fmt.Sprintf(`%s&daily=%s`, url, metrics)
	}

	return url, nil
}

func (c Client) Get(ctx context.Context, loc Location, opts *Options) ([]byte, error) {
	url := urlFromOptions(c.URL, loc, opts)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("%s - %s", res.Status, body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (hc HistClient) Get(ctx context.Context, loc Location, opts *HistoricalOptions) ([]byte, error) {
	url, err := urlFromHistoricalOptions(hc.URL, loc, opts)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to form url from historical options")
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", hc.UserAgent)

	res, err := hc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("%s - %s", res.Status, body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
