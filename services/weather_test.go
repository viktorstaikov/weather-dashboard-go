package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/viktorstaikov/weather-dashboard-go/config"
)

func TestGetTempSeries(t *testing.T) {

}

func TestGetForecast(t *testing.T) {
	config.Init("development")
	c := config.GetConfig()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		assert.Equal(t, r.Method, "GET")

		path := r.URL.EscapedPath()
		defaultPath := c.GetString("openWeather.forecastEndpoint")
		assert.Equal(t, defaultPath, path)

		query := r.URL.Query()

		lat := query.Get("lat")
		assert.Equal(t, "42.6979", lat)
		lon := query.Get("lon")
		assert.Equal(t, "23.3222", lon)

		defaultAppID := c.GetString("openWeather.appId")
		appID := query.Get("appid")
		assert.Equal(t, defaultAppID, appID)

		units := query.Get("units")
		assert.Equal(t, "metric", units)

	}))
	defer ts.Close()

	c.Set("openWeather.baseUrl", ts.URL)

	_, err := MakeForecastRequest()
	if err != nil {
		t.Errorf("MakeForecastRequest() returned an error: %s", err)
	}

}
