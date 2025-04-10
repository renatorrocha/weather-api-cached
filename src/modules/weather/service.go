package weather

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type WeatherService interface {
	Get(city string) (string, error)
}

type weatherServiceImpl struct {
	apiKey string
	cache  map[string]cachedData
}

type cachedData struct {
	Data      string
	ExpiresAt time.Time
}

func NewWeatherService(apiKey string) WeatherService {
	return &weatherServiceImpl{
		apiKey: apiKey,
		cache:  make(map[string]cachedData),
	}
}

func (s *weatherServiceImpl) Get(city string) (string, error) {

	cityParam := url.QueryEscape(city)
	if val, ok := s.cache[cityParam]; ok && time.Now().Before(val.ExpiresAt) {
		return val.Data, nil
	}

	fullUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", cityParam, s.apiKey)

	log.Println("Fetching weather data for", city)

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

	s.cache[cityParam] = cachedData{
		Data:      string(body),
		ExpiresAt: time.Now().Add(time.Minute * 10),
	}

	return string(body), nil
}
