package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"weatherapp/config"
	"weatherapp/pkg/logging"
)

type OpenWeatherMapClient interface {
	GetWeatherData(cityName string) (string, error)
}

//go:generate mockgen -package=external -source=open_weather_map_client.go -destination=open_weather_map_client_mock.go

type openWeatherMapClient struct {
	apiKey   string
	baseURL  string
	logger   logging.Logger
	client   *http.Client //TODO: Can have a wrapper around HTTP client and expose it as interface
	hostname string
}

func NewOpenWeatherMapClient(config config.OpenWeatherMapClient, logger logging.Logger) OpenWeatherMapClient {
	client := &http.Client{
		Timeout: config.Timeout * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:    config.MaxIdleConns, // conn pooling
			IdleConnTimeout: config.IdleConnTimeout * time.Second,
		},
	}

	return &openWeatherMapClient{
		apiKey:   config.APIKey,
		baseURL:  config.BaseURL,
		logger:   logger,
		client:   client,
		hostname: config.Hostname,
	}
}

func (c *openWeatherMapClient) GetWeatherData(cityName string) (string, error) {

	requestURL := fmt.Sprintf("%s%s?q=%s&appid=%s", c.hostname, c.baseURL, url.QueryEscape(cityName), c.apiKey)

	resp, err := c.client.Get(requestURL)
	if err != nil {
		c.logger.LogError(err.Error(), map[string]interface{}{"client": "openWeatherMapClient", "method": "GetWeatherData"})
		return "", err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var weatherData map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&weatherData)
		if err != nil {
			c.logger.LogError(err.Error(), map[string]interface{}{"client": "openWeatherMapClient", "method": "GetWeatherData"})
			return "", err
		}
		rawData, _ := json.Marshal(weatherData)
		return string(rawData), nil

	case http.StatusUnauthorized:
		c.logger.LogError("API key is invalid or missing", map[string]interface{}{"client": "openWeatherMapClient", "method": "GetWeatherData"})
		return "", errors.New("API key is invalid or missing")

	case http.StatusNotFound:
		c.logger.LogError("city not found", map[string]interface{}{"client": "openWeatherMapClient", "method": "GetWeatherData"})
		return "", errors.New("city not found")

	case http.StatusTooManyRequests:
		c.logger.LogError("rate limit exceeded", map[string]interface{}{"client": "openWeatherMapClient", "method": "GetWeatherData"})
		return "", errors.New("rate limit exceeded")

	case http.StatusInternalServerError:
		c.logger.LogError("server encountered an internal error", map[string]interface{}{"client": "openWeatherMapClient", "method": "GetWeatherData"})
		return "", errors.New("server encountered an internal error")

	default:
		return "", fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}
}
