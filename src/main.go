package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renatorrocha/weather-api-cached/src/config"
	"github.com/renatorrocha/weather-api-cached/src/modules/weather"
)

func main() {
	config.Load()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})

	// inject dependencies
	weatherDeps := weather.WeatherDeps{
		Service: weather.NewWeatherService(config.GetEnv("OPENWEATHER_API_KEY", "")),
	}

	// register routes
	weather.RegisterRoutes(app, weatherDeps)

	// start server
	app.Listen(":3000")
}
