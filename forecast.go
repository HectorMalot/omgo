package omgo

import (
	"context"
)

// Forecast retreives the 7 day (by default) weather forecast for the provided location.
//
// Use `Options` to specify which metrics to retrieve. The response is a Forecast
// struct that will contains the current weather, all requested hourly predictions and
// all requested daily predictions
func (c Client) Forecast(ctx context.Context, loc Location, opts *Options) (*Forecast, error) {
	body, err := c.Get(ctx, loc, opts)
	if err != nil {
		return nil, err
	}

	return ParseBody(body)
}
