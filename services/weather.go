package services

import (
	"time"
)

// WeatherAPI ...
type WeatherAPI interface {
	MakeForecastRequest() ([]MetaForecast, error)
	GetForecast(*time.Time) (*MetaForecast, error)
}

// WeatherService ...
type WeatherService struct {
	api WeatherAPI
}

// WeatherCondition ...
type WeatherCondition struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// WindCondition ...
type WindCondition struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

// MetaForecast ...
type MetaForecast struct {
	Timestamp uint             `json:"timestamp"`
	Temp      float64          `json:"temp"`
	TempMin   float64          `json:"temp_min"`
	TempMax   float64          `json:"temp_max"`
	Pressure  int              `json:"pressure"`
	Humidity  int              `json:"humidity"`
	Weather   WeatherCondition `json:"weather"`
	Clouds    int              `json:"clouds"`
	Wind      WindCondition    `json:"wind"`
	Rain      float64          `json:"rain"`
	Snow      float64          `json:"snow"`
	UVIndex   float64          `json:"uvindex"`
}

// TempData represents temperature forecast at given time
type TempData struct {
	Timestamp uint    `json:"timestamp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
}

// StatsData ...
type StatsData struct {
	Timestamp uint    `json:"timestamp"`
	Value     float64 `json:"value"`
}

// MakeWeatherService ...
func MakeWeatherService(weatherAPI WeatherAPI) *WeatherService {
	s := new(WeatherService)
	s.api = weatherAPI

	return s
}

// ToTempData ...
func (m MetaForecast) ToTempData() TempData {
	var item TempData
	item.TempMax = m.TempMax
	item.TempMin = m.TempMin
	item.Timestamp = m.Timestamp
	return item
}

// GetTempSeries returns data series for min and max temperature
func (w *WeatherService) GetTempSeries() ([]TempData, error) {

	meta, err := w.api.MakeForecastRequest()

	if err != nil {
		return nil, err
	}

	var mapped []TempData
	for _, item := range meta {
		entry := item.ToTempData()
		mapped = append(mapped, entry)
	}
	return mapped, nil
}

// GetRainSeries ...
func (w *WeatherService) GetRainSeries() ([]StatsData, error) {
	meta, err := w.api.MakeForecastRequest()

	if err != nil {
		return nil, err
	}

	var mapped []StatsData
	for _, item := range meta {
		var entry StatsData
		entry.Timestamp = item.Timestamp
		entry.Value = item.Rain
		mapped = append(mapped, entry)
	}
	return mapped, nil
}

// GetPressureSeries ...
func (w *WeatherService) GetPressureSeries() ([]StatsData, error) {
	meta, err := w.api.MakeForecastRequest()

	if err != nil {
		return nil, err
	}

	var mapped []StatsData
	for _, item := range meta {
		var entry StatsData
		entry.Timestamp = item.Timestamp
		entry.Value = float64(item.Pressure)
		mapped = append(mapped, entry)
	}
	return mapped, nil
}

// GetHumiditySeries ...
func (w *WeatherService) GetHumiditySeries() ([]StatsData, error) {
	meta, err := w.api.MakeForecastRequest()

	if err != nil {
		return nil, err
	}

	var mapped []StatsData
	for _, item := range meta {
		var entry StatsData
		entry.Timestamp = item.Timestamp
		entry.Value = float64(item.Humidity)
		mapped = append(mapped, entry)
	}
	return mapped, nil
}

// GetForecast returns full forecast for given day
func (w *WeatherService) GetForecast(date *time.Time) (*MetaForecast, error) {
	// return nil, errors.New("not implemented")
	forecast, err := w.api.GetForecast(date)
	if err != nil {
		return nil, err
	}
	return forecast, nil
}
