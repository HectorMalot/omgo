# Open-Meteo-Go

A simple go client for the [open meteo](https://open-meteo.com) API. It supports all options of the API as of Sept 20 2021.

## Usage

Simple example:

```go
package main

import (
	"context"
	"fmt"

	"github.com/hectormalot/omgo"
)

func main() {
	c, _ := omgo.NewClient()

	// Get the current weather for amsterdam
	loc, _ := omgo.NewLocation(52.3738, 4.8910)
	res, _ := c.CurrentWeather(context.Background(), loc, nil)
	fmt.Println("The temperature in Amsterdam is: ", res.Temperature)

	// Get the humidity and cloud cover forecast for berlin, 
	// including the last 2 days and non-metric units
	loc, _ := omgo.NewLocation(52.5235, 13.4115)
	opts := omgo.Options{
		TemperatureUnit:   "fahrenheit",
		WindspeedUnit:     "mph",
		PrecipitationUnit: "inch",
		Timezone:          "US/Eastern",
		PastDays:          2,
		HourlyMetrics:     []string{"cloudcover", "relativehumidity_2m"},
		DailyMetrics:      []string{"temperature_2m_max"},
	}
	
	res, _ := c.Forecast(context.Background(), loc, &opts)
	fmt.Println(res)
	// res.HourlyMetrics["cloudcover"] contains an array of cloud coverage predictions
	// res.HourlyMetrics["relativehumidity_2m"] contains an array of relative humidity predictions
	// res.HourlyTimes contains the timestamps for each prediction
	// res.DailyMetrics["temperature_2m_max"] contains daily maximum values for the temperature_2m metric
	// res.DailyTimes contains the timestamps for all daily predictions
}


```
