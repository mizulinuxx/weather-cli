package controllers

import (
	"weather-cli/models"
	"weather-cli/views"
)

func GetWeather(location, language string) error {
	weatherData, err := models.FetchWeather(location, language)
	
	if err != nil {
		return err
	}

	views.DisplayWeather(weatherData)
	return nil
}
