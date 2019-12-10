package services

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockAPI struct{}

func (f *MockAPI) GetForecast(d *time.Time) (*MetaForecast, error) {
	forecast := &MetaForecast{
		Timestamp: 1575590400000,
		Temp:      2,
		TempMin:   1,
		TempMax:   2,
		Pressure:  1030,
		Humidity:  89,
		Weather: WeatherCondition{
			ID:          800,
			Main:        "Clear",
			Description: "clear sky",
			Icon:        "01n",
		},
		Clouds: 0,
		Wind: WindCondition{
			Speed: 1.07,
			Deg:   234,
		},
		Rain:    0,
		Snow:    0,
		UVIndex: 0,
	}
	return forecast, nil
}
func (f *MockAPI) GetHumiditySeries() ([]StatsData, error) {
	data := []StatsData{
		StatsData{
			Timestamp: 1575590400000,
			Value:     89,
		},
		StatsData{
			Timestamp: 1575590420000,
			Value:     91,
		},
		StatsData{
			Timestamp: 1575590440000,
			Value:     50,
		},
	}
	return data, nil
}
func (f *MockAPI) GetPressureSeries() ([]StatsData, error) {
	data := []StatsData{
		StatsData{
			Timestamp: 1575590400000,
			Value:     1030,
		},
		StatsData{
			Timestamp: 1575590420000,
			Value:     1060,
		},
		StatsData{
			Timestamp: 1575590440000,
			Value:     1056,
		},
	}
	return data, nil
}
func (f *MockAPI) GetRainSeries() ([]StatsData, error) {
	data := []StatsData{
		StatsData{
			Timestamp: 1575590400000,
			Value:     0,
		},
		StatsData{
			Timestamp: 1575590420000,
			Value:     3,
		},
		StatsData{
			Timestamp: 1575590440000,
			Value:     10,
		},
	}
	return data, nil
}
func (f *MockAPI) GetTempSeries() ([]TempData, error) {
	td := []TempData{
		TempData{
			Timestamp: 1575590400000,
			TempMin:   -2,
			TempMax:   -1,
		},
		TempData{
			Timestamp: 1575590420000,
			TempMin:   1,
			TempMax:   2,
		},
	}
	return td, nil
}

type MockFailingAPI struct{}

func (f *MockFailingAPI) GetForecast(d *time.Time) (*MetaForecast, error) {
	return nil, errors.New("some error")
}
func (f *MockFailingAPI) GetHumiditySeries() ([]StatsData, error) {
	return nil, errors.New("some error")
}
func (f *MockFailingAPI) GetPressureSeries() ([]StatsData, error) {
	return nil, errors.New("some error")
}
func (f *MockFailingAPI) GetRainSeries() ([]StatsData, error) {
	return nil, errors.New("some error")
}
func (f *MockFailingAPI) GetTempSeries() ([]TempData, error) {
	return nil, errors.New("some error")
}

func TestMakeWeatherService(t *testing.T) {
	api := new(MockAPI)
	s := MakeWeatherService(api)

	assert.NotNil(t, s)
	assert.NotNil(t, s.api)
}

func TestToTempData(t *testing.T) {
	m := &MetaForecast{
		Timestamp: 1575590400000,
		Temp:      -1.29,
		TempMin:   -2.84,
		TempMax:   -1.29,
		Pressure:  1030,
		Humidity:  89,
		Weather: WeatherCondition{
			ID:          800,
			Main:        "Clear",
			Description: "clear sky",
			Icon:        "01n",
		},
		Clouds: 0,
		Wind: WindCondition{
			Speed: 1.07,
			Deg:   234,
		},
		Rain:    0,
		Snow:    0,
		UVIndex: 0,
	}
	d := m.ToTempData()

	assert.Equal(t, m.Timestamp, d.Timestamp)
	assert.Equal(t, m.TempMin, d.TempMin)
	assert.Equal(t, m.TempMax, d.TempMax)
}

func TestGetForecast(t *testing.T) {
	api := new(MockAPI)
	s := MakeWeatherService(api)

	date, _ := time.Parse(time.RFC3339, "2019-12-06T00:00:00.000Z")
	forecast, err := s.GetForecast(&date)

	assert.Nil(t, err)
	assert.NotNil(t, forecast)
}

func TestGetForecastFail(t *testing.T) {
	api := new(MockFailingAPI)
	s := MakeWeatherService(api)

	date, _ := time.Parse(time.RFC3339, "2019-12-06T00:00:00.000Z")
	forecast, err := s.GetForecast(&date)

	assert.NotNil(t, err)
	assert.Nil(t, forecast)
}

func TestGetHumidity(t *testing.T) {
	api := new(MockAPI)
	s := MakeWeatherService(api)

	series, err := s.GetHumiditySeries()

	assert.Nil(t, err)
	assert.Equal(t, 3, len(series))
}

func TestGetHumidityFail(t *testing.T) {
	api := new(MockFailingAPI)
	s := MakeWeatherService(api)

	data, err := s.GetHumiditySeries()

	assert.NotNil(t, err)
	assert.Nil(t, data)
}

func TestGetPressure(t *testing.T) {
	api := new(MockAPI)
	s := MakeWeatherService(api)

	series, err := s.GetPressureSeries()

	assert.Nil(t, err)
	assert.Equal(t, 3, len(series))
}

func TestGetPressureFail(t *testing.T) {
	api := new(MockFailingAPI)
	s := MakeWeatherService(api)

	data, err := s.GetPressureSeries()

	assert.NotNil(t, err)
	assert.Nil(t, data)
}

func TestGetRain(t *testing.T) {
	api := new(MockAPI)
	s := MakeWeatherService(api)

	series, err := s.GetRainSeries()

	assert.Nil(t, err)
	assert.Equal(t, 3, len(series))
}

func TestGetRainFail(t *testing.T) {
	api := new(MockFailingAPI)
	s := MakeWeatherService(api)

	data, err := s.GetRainSeries()

	assert.NotNil(t, err)
	assert.Nil(t, data)
}

func TestGetTempSeries(t *testing.T) {
	api := new(MockAPI)
	s := MakeWeatherService(api)

	series, err := s.GetTempSeries()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(series))
}

func TestGetTempSeriesFail(t *testing.T) {
	api := new(MockFailingAPI)
	s := MakeWeatherService(api)

	data, err := s.GetTempSeries()

	assert.NotNil(t, err)
	assert.Nil(t, data)
}
