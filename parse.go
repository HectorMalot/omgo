package omgo

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"
)

// rawResponse represents the raw JSON response from the API.
// This is used as an intermediate step for parsing.
type rawResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	Elevation            float64 `json:"elevation"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	UTCOffsetSeconds     int     `json:"utc_offset_seconds"`
	GenerationTimeMs     float64 `json:"generationtime_ms"`

	Current      json.RawMessage `json:"current,omitempty"`
	CurrentUnits *CurrentUnits   `json:"current_units,omitempty"`

	Hourly      json.RawMessage `json:"hourly,omitempty"`
	HourlyUnits json.RawMessage `json:"hourly_units,omitempty"`

	Minutely15      json.RawMessage `json:"minutely_15,omitempty"`
	Minutely15Units json.RawMessage `json:"minutely_15_units,omitempty"`

	Daily      json.RawMessage `json:"daily,omitempty"`
	DailyUnits json.RawMessage `json:"daily_units,omitempty"`
}

// rawCurrent represents the raw current weather data.
type rawCurrent struct {
	Time                string       `json:"time"`
	Interval            int          `json:"interval"`
	Temperature2m       *float64     `json:"temperature_2m,omitempty"`
	RelativeHumidity2m  *float64     `json:"relative_humidity_2m,omitempty"`
	ApparentTemperature *float64     `json:"apparent_temperature,omitempty"`
	IsDay               *int         `json:"is_day,omitempty"`
	Precipitation       *float64     `json:"precipitation,omitempty"`
	Rain                *float64     `json:"rain,omitempty"`
	Showers             *float64     `json:"showers,omitempty"`
	Snowfall            *float64     `json:"snowfall,omitempty"`
	WeatherCode         *WeatherCode `json:"weather_code,omitempty"`
	CloudCover          *float64     `json:"cloud_cover,omitempty"`
	PressureMSL         *float64     `json:"pressure_msl,omitempty"`
	SurfacePressure     *float64     `json:"surface_pressure,omitempty"`
	WindSpeed10m        *float64     `json:"wind_speed_10m,omitempty"`
	WindDirection10m    *float64     `json:"wind_direction_10m,omitempty"`
	WindGusts10m        *float64     `json:"wind_gusts_10m,omitempty"`
}

// rawHourly represents the raw hourly data with time as strings.
type rawHourly struct {
	Time []string `json:"time"`
	// All other fields will be unmarshaled directly into HourlyData
}

// rawDaily represents the raw daily data with time and sun times as strings.
type rawDaily struct {
	Time    []string `json:"time"`
	Sunrise []string `json:"sunrise,omitempty"`
	Sunset  []string `json:"sunset,omitempty"`
	// All other fields will be unmarshaled directly into DailyData
}

// rawMinutely15 represents the raw 15-minutely data with time as strings.
type rawMinutely15 struct {
	Time []string `json:"time"`
	// All other fields will be unmarshaled directly into Minutely15Data
}

// parseWeatherResponse parses the API response into a Weather struct.
func parseWeatherResponse(body []byte) (*Weather, error) {
	var raw rawResponse
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	// Load timezone for proper time parsing
	var loc *time.Location
	if raw.Timezone != "" {
		var err error
		loc, err = time.LoadLocation(raw.Timezone)
		if err != nil {
			// Fall back to UTC if timezone is invalid
			loc = time.UTC
		}
	}

	weather := &Weather{
		Latitude:             raw.Latitude,
		Longitude:            raw.Longitude,
		Elevation:            raw.Elevation,
		Timezone:             raw.Timezone,
		TimezoneAbbreviation: raw.TimezoneAbbreviation,
		UTCOffsetSeconds:     raw.UTCOffsetSeconds,
		GenerationTimeMs:     raw.GenerationTimeMs,
		CurrentUnits:         raw.CurrentUnits,
	}

	// Parse current weather
	if len(raw.Current) > 0 {
		current, err := parseCurrent(raw.Current, loc)
		if err != nil {
			return nil, err
		}
		weather.Current = current
	}

	if err := assignUnits(raw.HourlyUnits, parseHourlyUnits, func(units *HourlyUnits, byModel map[string]*HourlyUnits) {
		weather.HourlyUnits = units
		weather.HourlyUnitsByModel = byModel
	}); err != nil {
		return nil, err
	}

	if err := assignData(raw.Hourly, loc, parseHourly, func(data *HourlyData, byModel map[string]*HourlyData) {
		weather.Hourly = data
		weather.HourlyByModel = byModel
	}); err != nil {
		return nil, err
	}

	if err := assignUnits(raw.Minutely15Units, parseMinutely15Units, func(units *Minutely15Units, byModel map[string]*Minutely15Units) {
		weather.Minutely15Units = units
		weather.Minutely15UnitsByModel = byModel
	}); err != nil {
		return nil, err
	}

	if err := assignData(raw.Minutely15, loc, parseMinutely15, func(data *Minutely15Data, byModel map[string]*Minutely15Data) {
		weather.Minutely15 = data
		weather.Minutely15ByModel = byModel
	}); err != nil {
		return nil, err
	}

	if err := assignUnits(raw.DailyUnits, parseDailyUnits, func(units *DailyUnits, byModel map[string]*DailyUnits) {
		weather.DailyUnits = units
		weather.DailyUnitsByModel = byModel
	}); err != nil {
		return nil, err
	}

	if err := assignData(raw.Daily, loc, parseDaily, func(data *DailyData, byModel map[string]*DailyData) {
		weather.Daily = data
		weather.DailyByModel = byModel
	}); err != nil {
		return nil, err
	}

	return weather, nil
}

func assignUnits[U any](
	raw json.RawMessage,
	parse func(json.RawMessage) (*U, map[string]*U, error),
	assign func(*U, map[string]*U),
) error {
	if len(raw) == 0 {
		return nil
	}
	units, byModel, err := parse(raw)
	if err != nil {
		return err
	}
	assign(units, byModel)
	return nil
}

func assignData[T any](
	raw json.RawMessage,
	loc *time.Location,
	parse func(json.RawMessage, *time.Location) (*T, map[string]*T, error),
	assign func(*T, map[string]*T),
) error {
	if len(raw) == 0 {
		return nil
	}
	data, byModel, err := parse(raw, loc)
	if err != nil {
		return err
	}
	assign(data, byModel)
	return nil
}

// parseCurrent parses current weather data.
func parseCurrent(data json.RawMessage, loc *time.Location) (*CurrentData, error) {
	var raw rawCurrent
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	t, err := parseDateTime(raw.Time, loc)
	if err != nil {
		return nil, err
	}

	return &CurrentData{
		Time:                t,
		Interval:            raw.Interval,
		Temperature2m:       raw.Temperature2m,
		RelativeHumidity2m:  raw.RelativeHumidity2m,
		ApparentTemperature: raw.ApparentTemperature,
		IsDay:               raw.IsDay,
		Precipitation:       raw.Precipitation,
		Rain:                raw.Rain,
		Showers:             raw.Showers,
		Snowfall:            raw.Snowfall,
		WeatherCode:         raw.WeatherCode,
		CloudCover:          raw.CloudCover,
		PressureMSL:         raw.PressureMSL,
		SurfacePressure:     raw.SurfacePressure,
		WindSpeed10m:        raw.WindSpeed10m,
		WindDirection10m:    raw.WindDirection10m,
		WindGusts10m:        raw.WindGusts10m,
	}, nil
}

// parseHourly parses hourly weather data.
func parseHourly(data json.RawMessage, loc *time.Location) (*HourlyData, map[string]*HourlyData, error) {
	obj, err := unmarshalJSONObject(data)
	if err != nil {
		return nil, nil, err
	}

	times, err := parseDateTimeArrayFromRaw(obj["time"], loc)
	if err != nil {
		return nil, nil, err
	}

	info := getJSONTagInfo(reflect.TypeOf(HourlyData{}))
	baseObj, modelObjs := splitJSONObjectByModel(obj, info)

	hourly := &HourlyData{}
	if len(baseObj) > 0 {
		if err := json.Unmarshal(mustMarshalObject(baseObj), hourly); err != nil {
			return nil, nil, err
		}
	}
	hourly.Times = times

	var byModel map[string]*HourlyData
	if len(modelObjs) > 0 {
		byModel = make(map[string]*HourlyData, len(modelObjs))
		for model, modelObj := range modelObjs {
			modelObj["time"] = obj["time"]
			h := &HourlyData{}
			if err := json.Unmarshal(mustMarshalObject(modelObj), h); err != nil {
				return nil, nil, err
			}
			h.Times = times
			byModel[model] = h
		}
		if !hasMetricData(hourly) {
			first := firstSortedKey(byModel)
			if first != "" {
				hourly = byModel[first]
			}
		}
	}

	return hourly, byModel, nil
}

// parseMinutely15 parses 15-minutely weather data.
func parseMinutely15(data json.RawMessage, loc *time.Location) (*Minutely15Data, map[string]*Minutely15Data, error) {
	obj, err := unmarshalJSONObject(data)
	if err != nil {
		return nil, nil, err
	}

	times, err := parseDateTimeArrayFromRaw(obj["time"], loc)
	if err != nil {
		return nil, nil, err
	}

	info := getJSONTagInfo(reflect.TypeOf(Minutely15Data{}))
	baseObj, modelObjs := splitJSONObjectByModel(obj, info)

	minutely15 := &Minutely15Data{}
	if len(baseObj) > 0 {
		if err := json.Unmarshal(mustMarshalObject(baseObj), minutely15); err != nil {
			return nil, nil, err
		}
	}
	minutely15.Times = times

	var byModel map[string]*Minutely15Data
	if len(modelObjs) > 0 {
		byModel = make(map[string]*Minutely15Data, len(modelObjs))
		for model, modelObj := range modelObjs {
			modelObj["time"] = obj["time"]
			m := &Minutely15Data{}
			if err := json.Unmarshal(mustMarshalObject(modelObj), m); err != nil {
				return nil, nil, err
			}
			m.Times = times
			byModel[model] = m
		}
		if !hasMetricData(minutely15) {
			first := firstSortedKey(byModel)
			if first != "" {
				minutely15 = byModel[first]
			}
		}
	}

	return minutely15, byModel, nil
}

// parseDaily parses daily weather data.
func parseDaily(data json.RawMessage, loc *time.Location) (*DailyData, map[string]*DailyData, error) {
	obj, err := unmarshalJSONObject(data)
	if err != nil {
		return nil, nil, err
	}

	times, err := parseDateArrayFromRaw(obj["time"], loc)
	if err != nil {
		return nil, nil, err
	}

	var sunrise []time.Time
	if len(obj["sunrise"]) > 0 {
		sunrise, err = parseDateTimeArrayFromRaw(obj["sunrise"], loc)
		if err != nil {
			return nil, nil, err
		}
	}
	var sunset []time.Time
	if len(obj["sunset"]) > 0 {
		sunset, err = parseDateTimeArrayFromRaw(obj["sunset"], loc)
		if err != nil {
			return nil, nil, err
		}
	}

	info := getJSONTagInfo(reflect.TypeOf(DailyData{}))
	baseObj, modelObjs := splitJSONObjectByModel(obj, info)

	daily := &DailyData{}
	if len(baseObj) > 0 {
		if err := json.Unmarshal(mustMarshalObject(baseObj), daily); err != nil {
			return nil, nil, err
		}
	}
	daily.Times = times
	daily.Sunrise = sunrise
	daily.Sunset = sunset

	var byModel map[string]*DailyData
	if len(modelObjs) > 0 {
		byModel = make(map[string]*DailyData, len(modelObjs))
		for model, modelObj := range modelObjs {
			modelObj["time"] = obj["time"]
			if len(obj["sunrise"]) > 0 {
				modelObj["sunrise"] = obj["sunrise"]
			}
			if len(obj["sunset"]) > 0 {
				modelObj["sunset"] = obj["sunset"]
			}
			d := &DailyData{}
			if err := json.Unmarshal(mustMarshalObject(modelObj), d); err != nil {
				return nil, nil, err
			}
			d.Times = times
			d.Sunrise = sunrise
			d.Sunset = sunset
			byModel[model] = d
		}
		if !hasMetricData(daily) {
			first := firstSortedKey(byModel)
			if first != "" {
				daily = byModel[first]
			}
		}
	}

	return daily, byModel, nil
}

type jsonTagInfo struct {
	set       map[string]struct{}
	byLenDesc []string
}

var jsonTagInfoCache sync.Map

func getJSONTagInfo(t reflect.Type) *jsonTagInfo {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if v, ok := jsonTagInfoCache.Load(t); ok {
		return v.(*jsonTagInfo)
	}

	set := make(map[string]struct{})
	var names []string
	collectJSONTags(t, set, &names)

	byLenDesc := make([]string, 0, len(names))
	byLenDesc = append(byLenDesc, names...)
	sort.Slice(byLenDesc, func(i, j int) bool {
		if len(byLenDesc[i]) != len(byLenDesc[j]) {
			return len(byLenDesc[i]) > len(byLenDesc[j])
		}
		return byLenDesc[i] < byLenDesc[j]
	})

	info := &jsonTagInfo{
		set:       set,
		byLenDesc: byLenDesc,
	}
	jsonTagInfoCache.Store(t, info)
	return info
}

func collectJSONTags(t reflect.Type, set map[string]struct{}, names *[]string) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			collectJSONTags(f.Type, set, names)
			continue
		}
		tag := f.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		name := tag
		if idx := strings.IndexByte(tag, ','); idx >= 0 {
			name = tag[:idx]
		}
		if name == "" {
			continue
		}
		if _, ok := set[name]; ok {
			continue
		}
		set[name] = struct{}{}
		*names = append(*names, name)
	}
}

func unmarshalJSONObject(data json.RawMessage) (map[string]json.RawMessage, error) {
	if len(data) == 0 {
		return map[string]json.RawMessage{}, nil
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}
	if obj == nil {
		obj = map[string]json.RawMessage{}
	}
	return obj, nil
}

func mustMarshalObject(obj map[string]json.RawMessage) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return b
}

func splitJSONObjectByModel(
	obj map[string]json.RawMessage,
	info *jsonTagInfo,
) (map[string]json.RawMessage, map[string]map[string]json.RawMessage) {
	baseObj := make(map[string]json.RawMessage)
	modelObjs := make(map[string]map[string]json.RawMessage)

	for k, v := range obj {
		if k == "time" || k == "sunrise" || k == "sunset" {
			baseObj[k] = v
			continue
		}
		if _, ok := info.set[k]; ok {
			baseObj[k] = v
			continue
		}
		base, model, ok := matchModelSuffix(k, info.byLenDesc)
		if !ok || model == "" {
			continue
		}
		m := modelObjs[model]
		if m == nil {
			m = make(map[string]json.RawMessage)
			modelObjs[model] = m
		}
		m[base] = v
	}

	if len(modelObjs) == 0 {
		return baseObj, nil
	}
	return baseObj, modelObjs
}

func matchModelSuffix(key string, basesByLenDesc []string) (base string, model string, ok bool) {
	for _, b := range basesByLenDesc {
		prefix := b + "_"
		if strings.HasPrefix(key, prefix) {
			return b, strings.TrimPrefix(key, prefix), true
		}
	}
	return "", "", false
}

func parseDateTimeArrayFromRaw(data json.RawMessage, loc *time.Location) ([]time.Time, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var s []string
	if err := json.Unmarshal(data, &s); err == nil {
		return parseDateTimeArray(s, loc)
	}
	var i []int64
	if err := json.Unmarshal(data, &i); err == nil {
		return unixSecondsToTimes(i, loc), nil
	}
	return nil, fmt.Errorf("unsupported time format: %s", string(data))
}

func parseDateArrayFromRaw(data json.RawMessage, loc *time.Location) ([]time.Time, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var s []string
	if err := json.Unmarshal(data, &s); err == nil {
		return parseDateArray(s, loc)
	}
	var i []int64
	if err := json.Unmarshal(data, &i); err == nil {
		return unixSecondsToTimes(i, loc), nil
	}
	return nil, fmt.Errorf("unsupported date format: %s", string(data))
}

func unixSecondsToTimes(arr []int64, loc *time.Location) []time.Time {
	if loc == nil {
		loc = time.UTC
	}
	result := make([]time.Time, len(arr))
	for idx, sec := range arr {
		result[idx] = time.Unix(sec, 0).In(loc)
	}
	return result
}

func hasMetricData(v any) bool {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	return hasMetricDataValue(rv)
}

func hasMetricDataValue(rv reflect.Value) bool {
	if rv.Kind() != reflect.Struct {
		return false
	}
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)

		if ft.Anonymous && fv.Kind() == reflect.Struct {
			if hasMetricDataValue(fv) {
				return true
			}
			continue
		}

		if ft.Name == "Times" || ft.Name == "Sunrise" || ft.Name == "Sunset" {
			continue
		}

		if fv.Kind() == reflect.Slice && fv.Len() > 0 {
			return true
		}
	}
	return false
}

func hasUnitData(v any) bool {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	return hasUnitDataValue(rv)
}

func hasUnitDataValue(rv reflect.Value) bool {
	if rv.Kind() != reflect.Struct {
		return false
	}
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)

		if ft.Anonymous && fv.Kind() == reflect.Struct {
			if hasUnitDataValue(fv) {
				return true
			}
			continue
		}

		if fv.Kind() == reflect.String && fv.String() != "" {
			return true
		}
	}
	return false
}

func firstSortedKey[V any](m map[string]V) string {
	if len(m) == 0 {
		return ""
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys[0]
}

func parseHourlyUnits(data json.RawMessage) (*HourlyUnits, map[string]*HourlyUnits, error) {
	obj, err := unmarshalJSONObject(data)
	if err != nil {
		return nil, nil, err
	}

	info := getJSONTagInfo(reflect.TypeOf(HourlyUnits{}))
	baseObj, modelObjs := splitJSONObjectByModel(obj, info)

	units := &HourlyUnits{}
	if len(baseObj) > 0 {
		if err := json.Unmarshal(mustMarshalObject(baseObj), units); err != nil {
			return nil, nil, err
		}
	}

	var byModel map[string]*HourlyUnits
	if len(modelObjs) > 0 {
		byModel = make(map[string]*HourlyUnits, len(modelObjs))
		for model, modelObj := range modelObjs {
			u := &HourlyUnits{}
			if err := json.Unmarshal(mustMarshalObject(modelObj), u); err != nil {
				return nil, nil, err
			}
			byModel[model] = u
		}
		if !hasUnitData(units) {
			first := firstSortedKey(byModel)
			if first != "" {
				units = byModel[first]
			}
		}
	}

	return units, byModel, nil
}

func parseMinutely15Units(data json.RawMessage) (*Minutely15Units, map[string]*Minutely15Units, error) {
	obj, err := unmarshalJSONObject(data)
	if err != nil {
		return nil, nil, err
	}

	info := getJSONTagInfo(reflect.TypeOf(Minutely15Units{}))
	baseObj, modelObjs := splitJSONObjectByModel(obj, info)

	units := &Minutely15Units{}
	if len(baseObj) > 0 {
		if err := json.Unmarshal(mustMarshalObject(baseObj), units); err != nil {
			return nil, nil, err
		}
	}

	var byModel map[string]*Minutely15Units
	if len(modelObjs) > 0 {
		byModel = make(map[string]*Minutely15Units, len(modelObjs))
		for model, modelObj := range modelObjs {
			u := &Minutely15Units{}
			if err := json.Unmarshal(mustMarshalObject(modelObj), u); err != nil {
				return nil, nil, err
			}
			byModel[model] = u
		}
		if !hasUnitData(units) {
			first := firstSortedKey(byModel)
			if first != "" {
				units = byModel[first]
			}
		}
	}

	return units, byModel, nil
}

func parseDailyUnits(data json.RawMessage) (*DailyUnits, map[string]*DailyUnits, error) {
	obj, err := unmarshalJSONObject(data)
	if err != nil {
		return nil, nil, err
	}

	info := getJSONTagInfo(reflect.TypeOf(DailyUnits{}))
	baseObj, modelObjs := splitJSONObjectByModel(obj, info)

	units := &DailyUnits{}
	if len(baseObj) > 0 {
		if err := json.Unmarshal(mustMarshalObject(baseObj), units); err != nil {
			return nil, nil, err
		}
	}

	var byModel map[string]*DailyUnits
	if len(modelObjs) > 0 {
		byModel = make(map[string]*DailyUnits, len(modelObjs))
		for model, modelObj := range modelObjs {
			u := &DailyUnits{}
			if err := json.Unmarshal(mustMarshalObject(modelObj), u); err != nil {
				return nil, nil, err
			}
			byModel[model] = u
		}
		if !hasUnitData(units) {
			first := firstSortedKey(byModel)
			if first != "" {
				units = byModel[first]
			}
		}
	}

	return units, byModel, nil
}
