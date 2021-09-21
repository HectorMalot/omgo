package omgo

import (
	"context"
)

// CurrentWeather returns the current weather for the provided location
//
// Units and timezones can be provided using an optional `Options` parameter.
// Any requested hourly or daily metrics as part of the options are discarded
func (c Client) CurrentWeather(ctx context.Context, loc Location, opts *Options) (CurrentWeather, error) {
	// Discard requested daily/hourly metrics, as they are not returned as part of the current weather
	if opts != nil {
		opts.DailyMetrics = nil
		opts.HourlyMetrics = nil
	}

	body, err := c.Get(ctx, loc, opts)
	if err != nil {
		return CurrentWeather{}, err
	}

	fc, err := ParseBody(body)
	if err != nil {
		return CurrentWeather{}, err
	}

	return fc.CurrentWeather, err
}
