package main

import (
	"flag"
	"fmt"
	"os"
	"weather-cli/controllers"
)

func main() {
	location := flag.String("location", "london", "Location for the weather forecast")
	language := flag.String("language", "EN", "Language for the weather forecast")
	flag.Parse()

	if err := controllers.GetWeather(*location, *language); err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching weather: %v\n", err)
		os.Exit(1)
	}
}
