package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	pb "proto/proto" // Adjust this path based on your module name and directory structure

	"google.golang.org/protobuf/proto"
)

func main() {
	resp, err := http.Get("http://localhost:8082/weather/Gwalior")
	if err != nil {
		log.Fatalf("Failed to get weather: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var weather pb.WeatherResponse
	err = proto.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalf("Failed to unmarshal protobuf: %v", err)
	}

	fmt.Printf("Weather in %s, %s, %s:\n", weather.Location.Name, weather.Location.Region, weather.Location.Country)
	fmt.Printf("Latitude: %.2f, Longitude: %.2f\n", weather.Location.Lat, weather.Location.Lon)
	fmt.Printf("Time Zone: %s\n", weather.Location.TzId)
	fmt.Printf("Local Time: %s\n\n", weather.Location.Localtime)

	fmt.Printf("Current Weather:\n")
	fmt.Printf("Last Updated: %s\n", weather.Current.LastUpdated)
	fmt.Printf("Temperature: %.2f°C / %.2f°F\n", weather.Current.TempC, weather.Current.TempF)
	fmt.Printf("Is Day: %d\n", weather.Current.IsDay)
	fmt.Printf("Condition: %s (Icon: %s, Code: %d)\n", weather.Current.Condition.Text, weather.Current.Condition.Icon, weather.Current.Condition.Code)
	fmt.Printf("Wind: %.2f mph / %.2f kph, Degree: %d, Direction: %s\n", weather.Current.WindMph, weather.Current.WindKph, weather.Current.WindDegree, weather.Current.WindDir)
	fmt.Printf("Pressure: %.2f mb / %.2f in\n", weather.Current.PressureMb, weather.Current.PressureIn)
	fmt.Printf("Precipitation: %.2f mm / %.2f in\n", weather.Current.PrecipMm, weather.Current.PrecipIn)
	fmt.Printf("Humidity: %d%%\n", weather.Current.Humidity)
	fmt.Printf("Cloud: %d%%\n", weather.Current.Cloud)
	fmt.Printf("Feels Like: %.2f°C / %.2f°F\n", weather.Current.FeelslikeC, weather.Current.FeelslikeF)
	fmt.Printf("Wind Chill: %.2f°C / %.2f°F\n", weather.Current.WindchillC, weather.Current.WindchillF)
	fmt.Printf("Heat Index: %.2f°C / %.2f°F\n", weather.Current.HeatindexC, weather.Current.HeatindexF)
	fmt.Printf("Dew Point: %.2f°C / %.2f°F\n", weather.Current.DewpointC, weather.Current.DewpointF)
	fmt.Printf("Visibility: %.2f km / %.2f miles\n", weather.Current.VisKm, weather.Current.VisMiles)
	fmt.Printf("UV Index: %.2f\n", weather.Current.Uv)
	fmt.Printf("Gust: %.2f mph / %.2f kph\n", weather.Current.GustMph, weather.Current.GustKph)
}
