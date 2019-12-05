package server

import (
	"github.com/viktorstaikov/weather-dashboard-go/config"
)

// Init ...
func Init() {
	config := config.GetConfig()
	port := config.GetString("port")

	r := RootRouter()
	r.Run(port)
}
