package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/renatorrocha/weather-api-cached/src/config"
)

func GetWeather(city string) (string, error) {
	baseUrl := config.GetEnv("OPENWEATHER_URL", "")
	apiKey := config.GetEnv("OPENWEATHER_API_KEY", "")

	if baseUrl == "" || apiKey == "" {
		return "", errors.New("OPENWEATHER_URL or OPENWEATHER_API_KEY is not set")
	}

	cityParam := url.QueryEscape(city)

	fullUrl := fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", baseUrl, cityParam, apiKey)

	response, err := http.Get(fullUrl)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		return "", errors.New("error fetching weather data: " + string(body))

	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
