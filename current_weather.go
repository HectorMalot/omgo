package omgo

import (
	"context"
)

// CurrentWeather returns the current weather for the provided location
//
// Units and timezones can be provided using an optional `Options` parameter.
// Any requested hourly or daily metrics as part of the options are discarded
func (c Client) CurrentWeather(ctx context.Context, loc Location, opts *Options) (CurrentWeather, error) {

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
