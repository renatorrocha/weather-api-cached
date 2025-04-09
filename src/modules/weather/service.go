package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WeatherService interface {
	Get(city string) (string, error)
}

type weatherServiceImpl struct {
	apiKey string
}

func NewWeatherService(apiKey string) WeatherService {
	return &weatherServiceImpl{
		apiKey: apiKey,
	}
}

func (s *weatherServiceImpl) Get(city string) (string, error) {
	cityParam := url.QueryEscape(city)

	fullUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", cityParam, s.apiKey)

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
