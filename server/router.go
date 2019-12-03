package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/viktorstaikov/weather-dashboard-go/controllers"
)

// RootRouter ...
func RootRouter() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	controller := new(controllers.WeatherController)

	api := router.Group("/api")
	{
		weatherRouter := api.Group("weather")
		{
			weatherRouter.GET("temp_series", controller.TempSeries)
			weatherRouter.GET("humidity_series", controller.HumiditySeries)
			weatherRouter.GET("pressure_series", controller.PressureSeries)
			weatherRouter.GET("rain_series", controller.RainSeries)
			weatherRouter.GET("forecast", controller.Forecast)
		}
	}

	return router
}
