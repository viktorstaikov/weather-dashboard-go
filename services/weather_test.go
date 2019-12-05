package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockAPI struct{}

func (f *MockAPI) MakeForecastRequest() ([]MetaForecast, error) {
	mf := []MetaForecast{
		MetaForecast{
			Timestamp: 1575590400000,
			Temp:      -1,
			TempMin:   -2,
			TempMax:   -1,
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
		},
		MetaForecast{
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
		},
	}
	return mf, nil
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

func TestGetTempSeries(t *testing.T) {
	api := new(MockAPI)
	s := MakeWeatherService(api)

	series, err := s.GetTempSeries()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(series))
}
