package main

import (
	"github.com/renatorrocha/weather-api-cached/src/config"
	"github.com/renatorrocha/weather-api-cached/src/services"
)

func main() {
	config.Load()

	services.GetWeather("Sao Paulo")
}
