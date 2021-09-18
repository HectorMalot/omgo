package omgo

import (
	"context"
)

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
