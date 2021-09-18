package omgo

import (
	"context"
)

func (c Client) Forecast(ctx context.Context, loc Location, opts *Options) (*Forecast, error) {
	body, err := c.Get(ctx, loc, opts)
	if err != nil {
		return nil, err
	}

	return ParseBody(body)
}
