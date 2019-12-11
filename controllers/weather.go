package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viktorstaikov/weather-dashboard-go/config"
	"github.com/viktorstaikov/weather-dashboard-go/services"
	"github.com/viktorstaikov/weather-dashboard-go/services/openweatherapi"
)

// WeatherController ...
type WeatherController struct {
	weatherService *services.WeatherService
}

var ws *services.WeatherService

// MakeWeatherController ...
func MakeWeatherController() *WeatherController {
	c := new(WeatherController)

	config.Init("development")
	cnfg := config.GetConfig()
	api := openweatherapi.MakeOpenWeatherAPI(cnfg)
	c.weatherService = services.MakeWeatherService(api)

	return c
}

// TempSeries ...
func (h WeatherController) TempSeries(c *gin.Context) {
	resp, err := h.weatherService.GetTempSeries()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not get Temprature series", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, resp)
}

// RainSeries ...
func (h WeatherController) RainSeries(c *gin.Context) {
	resp, err := h.weatherService.GetRainSeries()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not get Rain series", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, resp)
}

// PressureSeries ...
func (h WeatherController) PressureSeries(c *gin.Context) {
	resp, err := h.weatherService.GetPressureSeries()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not get Pressure series", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, resp)
}

// HumiditySeries ...
func (h WeatherController) HumiditySeries(c *gin.Context) {
	resp, err := h.weatherService.GetHumiditySeries()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not get Pressure series", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Forecast ...
func (h WeatherController) Forecast(c *gin.Context) {
	query := c.Request.URL.Query()
	dateStr := query.Get("date")

	date, dateErr := time.Parse(time.RFC3339, dateStr)
	if dateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid date", "error": dateErr.Error()})
		c.Abort()
		return
	}
	resp, err := h.weatherService.GetForecast(&date)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not get Forecast for given date", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resp)
}
