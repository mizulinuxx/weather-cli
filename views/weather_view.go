package views

import (
	"fmt"
	"weather-cli/models"
)

func DisplayWeather(weather models.WeatherResponse) {
	fmt.Printf("Current weather for %s:\n", weather.Location)
	fmt.Printf("Temperature: %.2f\n", weather.Current.Temperature)
	fmt.Printf("Condition: %s\n", weather.Current.Condition)
	
	fmt.Println("Forecast for the next few days:")
	
	for _, forecast := range weather.Forecast {
		fmt.Printf("%s: %.2f %s, %s\n", forecast.Date, forecast.Temperature, forecast.Condition)
	}
}
