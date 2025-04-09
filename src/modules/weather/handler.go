package weather

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetWeatherHandler(service WeatherService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		city := c.Query("city")

		log.Println(city)

		data, err := service.Get(city)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.SendString(data)
	}
}
