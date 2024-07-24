package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherData struct {
	Temperature float64 `json:"temp"`
	Condition   string  `json:"description"`
}

type Forecast struct {
	Date        string  `json:"date"`
	Temperature float64 `json:"temp"`
	Condition   string  `json:"description"`
}

type WeatherResponse struct {
	Location string     `json:"name"`
	Current  WeatherData `json:"main"`
	Forecast []Forecast `json:"list"`
}

const (
	apiKey  = "dfc55dfdc5mshf9befcbab108ad6p11b1b2jsn8796c9c7c1bb"
	apiHost = "open-weather13.p.rapidapi.com"
)

func FetchWeather(location, language string) (WeatherResponse, error) {
	var weatherResponse WeatherResponse
	url := fmt.Sprintf("https://%s/city/%s/%s", apiHost, location, language)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", apiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return weatherResponse, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return weatherResponse, err
	}

	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return weatherResponse, err
	}

	return weatherResponse, nil
}

