# omgo - Open-Meteo Go Client

A Go client for the [Open-Meteo](https://open-meteo.com) weather API. Supports weather forecasts and historical data. Outputs are typed and can handle all available metrics as of December 2025.

> [!IMPORTANT]
> v0.2.x is a significant refactor of v0.1.x and not backwards compatible. Once stable, this is likely to be promoted to v1.0.0. Please see the migration guide at the bottom of this readme.

## Installation

```bash
go get github.com/hectormalot/omgo
```

Requires Go 1.21 or later.

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/hectormalot/omgo"
)

func main() {
    client := omgo.NewClient()

    // Create a forecast request for Amsterdam
    req, err := omgo.NewForecastRequest(52.3738, 4.8910)
    if err != nil {
        log.Fatal(err)
    }

    // Request temperature and precipitation
    req.WithHourly(omgo.HourlyTemperature2m, omgo.HourlyPrecipitation).
        WithDaily(omgo.DailyTemperature2mMax, omgo.DailyTemperature2mMin).
        WithTimezone("Europe/Berlin")

    weather, err := client.Forecast(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Current temperature: %.1f%s\n",
        weather.Hourly.Temperature2m[0],
        weather.HourlyUnits.Temperature2m)
}
```

## Features

- **Type-safe metrics**: All weather metrics are typed constants with autocomplete support
- **Builder pattern**: Fluent API for building requests
- **15-minutely data**: High-resolution data for supported regions
- **Historical data**: Access to historical weather archives
- **Units**: Full control over temperature, wind speed, and precipitation units

## Usage Examples

### Current Weather

```go
req, _ := omgo.NewForecastRequest(40.7128, -74.0060) // New York
req.WithCurrent(
    omgo.CurrentTemperature2m,
    omgo.CurrentWeatherCode,
    omgo.CurrentWindSpeed10m,
)

weather, _ := client.Forecast(context.Background(), req)

fmt.Printf("Temperature: %.1f°C\n", *weather.Current.Temperature2m)
fmt.Printf("Conditions: %s\n", weather.Current.WeatherCode.String())
fmt.Printf("Is daytime: %v\n", weather.Current.IsDaytime())
```

### Hourly Forecast

```go
req, _ := omgo.NewForecastRequest(52.52, 13.41)
req.WithHourly(
    omgo.HourlyTemperature2m,
    omgo.HourlyRelativeHumidity2m,
    omgo.HourlyPrecipitation,
    omgo.HourlyWeatherCode,
).WithForecastDays(7).
  WithTimezone("Europe/Berlin")

weather, _ := client.Forecast(context.Background(), req)

for i, t := range weather.Hourly.Times {
    fmt.Printf("%s: %.1f°C, %s\n",
        t.Format("Mon 15:04"),
        weather.Hourly.Temperature2m[i],
        weather.Hourly.WeatherCode[i].String())
}
```

### Daily Forecast with Sunrise/Sunset

```go
req, _ := omgo.NewForecastRequest(52.52, 13.41)
req.WithDaily(
    omgo.DailyTemperature2mMax,
    omgo.DailyTemperature2mMin,
    omgo.DailySunrise,
    omgo.DailySunset,
    omgo.DailyPrecipitationSum,
).WithTimezone("Europe/Berlin")

weather, _ := client.Forecast(context.Background(), req)

for i, t := range weather.Daily.Times {
    fmt.Printf("%s: %.0f°C to %.0f°C, Sunrise: %s\n",
        t.Format("Mon Jan 2"),
        weather.Daily.Temperature2mMin[i],
        weather.Daily.Temperature2mMax[i],
        weather.Daily.Sunrise[i].Format("15:04"))
}
```

### 15-Minutely Data

```go
req, _ := omgo.NewForecastRequest(52.52, 13.41)
req.WithMinutely15(
    omgo.Minutely15Temperature2m,
    omgo.Minutely15Precipitation,
)

weather, _ := client.Forecast(context.Background(), req)

// High-resolution data for the next hours
for i, t := range weather.Minutely15.Times {
    fmt.Printf("%s: %.1f°C\n", t.Format("15:04"), weather.Minutely15.Temperature2m[i])
}
```

### Historical Data

```go
req, _ := omgo.NewHistoricalRequest(52.52, 13.41, "2023-06-01", "2023-06-30")
req.WithHourly(omgo.HourlyTemperature2m, omgo.HourlyPrecipitation).
    WithDaily(omgo.DailyTemperature2mMax).
    WithTimezone("Europe/Berlin")

weather, _ := client.Historical(context.Background(), req)
```

### Custom Units

```go
req, _ := omgo.NewForecastRequest(40.7128, -74.0060) // New York
req.WithHourly(omgo.HourlyTemperature2m, omgo.HourlyWindSpeed10m, omgo.HourlyPrecipitation).
    WithTemperatureUnit(omgo.Fahrenheit).
    WithWindSpeedUnit(omgo.MilesPerHour).
    WithPrecipitationUnit(omgo.Inches)

weather, _ := client.Forecast(context.Background(), req)

fmt.Printf("Temperature: %.1f%s\n",
    weather.Hourly.Temperature2m[0],
    weather.HourlyUnits.Temperature2m) // "°F"
```

### Commercial API Access

```go
client := omgo.NewClient(
    omgo.WithForecastURL("https://api.open-meteo.com/v1/forecast"),
    omgo.WithAPIKey("your-api-key"),
)
```

## Available Metrics

### Current Weather

- Temperature, apparent temperature, humidity
- Precipitation, rain, showers, snowfall
- Weather code, cloud cover
- Pressure (MSL and surface)
- Wind speed, direction, gusts
- Is day/night

### Hourly Metrics

- Basic: temperature, humidity, dew point, apparent temperature
- Precipitation: rain, snow, showers, probability
- Cloud cover: total, low, mid, high
- Wind: speed and direction at multiple heights (10m, 80m, 120m, 180m)
- Solar radiation: shortwave, direct, diffuse, global tilted
- Soil: temperature and moisture at multiple depths
- Pressure levels: temperature, humidity, wind, cloud cover at 19 pressure levels

### Daily Metrics

- Temperature: max, min, mean
- Apparent temperature: max, min, mean
- Precipitation: sum, hours, probability
- Sun: sunrise, sunset, sunshine duration, daylight duration
- Wind: max speed, max gusts, dominant direction
- UV index

### 15-Minutely Metrics

- Temperature, humidity, apparent temperature
- Precipitation, rain, snow
- Solar radiation
- Wind speed and direction
- Lightning potential (regional)

## Weather Codes

The `WeatherCode` type includes all WMO weather interpretation codes:

```go
code := omgo.RainModerate // 63
fmt.Println(code.String())      // "Moderate rain"
fmt.Println(code.Description()) // "Moderate rain"
```

## Error Handling

```go
weather, err := client.Forecast(context.Background(), req)
if err != nil {
    var apiErr *omgo.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API error %d: %s\n", apiErr.StatusCode, apiErr.Reason)
    }
}
```

## Migration from v0.1.x

Version 0.2.0 is a complete rewrite with breaking changes:

### Before (v0.1.x)

```go
c, _ := omgo.NewClient()
loc, _ := omgo.NewLocation(52.52, 13.41)
opts := omgo.Options{
    HourlyMetrics: []string{"temperature_2m", "precipitation"},
}
res, _ := c.Forecast(ctx, loc, &opts)
temps := res.HourlyMetrics["temperature_2m"]
```

### After (v0.2.0)

```go
c := omgo.NewClient()
req, _ := omgo.NewForecastRequest(52.52, 13.41)
req.WithHourly(omgo.HourlyTemperature2m, omgo.HourlyPrecipitation)
weather, _ := c.Forecast(ctx, req)
temps := weather.Hourly.Temperature2m
unit := weather.HourlyUnits.Temperature2m // "°C"
```

### Key Changes

- `NewClient()` no longer returns an error
- `Options` struct replaced by builder pattern (`NewForecastRequest`, `NewHistoricalRequest`)
- `Forecast` response type renamed to `Weather`
- Metrics are now typed constants instead of strings
- Response data accessed via typed struct fields instead of maps
- Units available via parallel `*Units` structs

## License

MIT License.

## Acknowledgments

This library uses the free [Open-Meteo API](https://open-meteo.com). Consider supporting their work if you use it commercially.
