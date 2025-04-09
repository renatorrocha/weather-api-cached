package weather

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, deps WeatherDeps) {
	weatherGroup := app.Group("/weather")
	weatherGroup.Get("/", GetWeatherHandler(deps.Service))
}
