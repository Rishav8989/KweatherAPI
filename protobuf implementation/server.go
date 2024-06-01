package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"

	pb "proto/proto"
)

const weatherAPIURL = "http://api.weatherapi.com/v1/current.json"
const apiKey = "hidden key" // Replace with your actual API key

type WeatherAPIResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		WindchillC float64 `json:"windchill_c"`
		WindchillF float64 `json:"windchill_f"`
		HeatindexC float64 `json:"heatindex_c"`
		HeatindexF float64 `json:"heatindex_f"`
		DewpointC  float64 `json:"dewpoint_c"`
		DewpointF  float64 `json:"dewpoint_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

func fetchWeather(city string) (*pb.WeatherResponse, error) {
	url := weatherAPIURL + "?key=" + apiKey + "&q=" + city + "&aqi=no"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse WeatherAPIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	weatherResponse := &pb.WeatherResponse{
		Location: &pb.Location{
			Name:           apiResponse.Location.Name,
			Region:         apiResponse.Location.Region,
			Country:        apiResponse.Location.Country,
			Lat:            float32(apiResponse.Location.Lat),
			Lon:            float32(apiResponse.Location.Lon),
			TzId:           apiResponse.Location.TzID,
			LocaltimeEpoch: int32(apiResponse.Location.LocaltimeEpoch),
			Localtime:      apiResponse.Location.Localtime,
		},
		Current: &pb.Current{
			LastUpdatedEpoch: int32(apiResponse.Current.LastUpdatedEpoch),
			LastUpdated:      apiResponse.Current.LastUpdated,
			TempC:            apiResponse.Current.TempC,
			TempF:            apiResponse.Current.TempF,
			IsDay:            int32(apiResponse.Current.IsDay),
			Condition: &pb.Condition{
				Text: apiResponse.Current.Condition.Text,
				Icon: apiResponse.Current.Condition.Icon,
				Code: int32(apiResponse.Current.Condition.Code),
			},
			WindMph:    apiResponse.Current.WindMph,
			WindKph:    apiResponse.Current.WindKph,
			WindDegree: int32(apiResponse.Current.WindDegree),
			WindDir:    apiResponse.Current.WindDir,
			PressureMb: apiResponse.Current.PressureMb,
			PressureIn: apiResponse.Current.PressureIn,
			PrecipMm:   apiResponse.Current.PrecipMm,
			PrecipIn:   apiResponse.Current.PrecipIn,
			Humidity:   int32(apiResponse.Current.Humidity),
			Cloud:      int32(apiResponse.Current.Cloud),
			FeelslikeC: apiResponse.Current.FeelslikeC,
			FeelslikeF: apiResponse.Current.FeelslikeF,
			WindchillC: apiResponse.Current.WindchillC,
			WindchillF: apiResponse.Current.WindchillF,
			HeatindexC: apiResponse.Current.HeatindexC,
			HeatindexF: apiResponse.Current.HeatindexF,
			DewpointC:  apiResponse.Current.DewpointC,
			DewpointF:  apiResponse.Current.DewpointF,
			VisKm:      apiResponse.Current.VisKm,
			VisMiles:   apiResponse.Current.VisMiles,
			Uv:         apiResponse.Current.Uv,
			GustMph:    apiResponse.Current.GustMph,
			GustKph:    apiResponse.Current.GustKph,
		},
	}

	fmt.Printf("WeatherResponse: %+v\n", weatherResponse)

	return weatherResponse, nil
}

func main() {
	port := flag.String("port", "8082", "port number to run the server on")
	flag.Parse()

	router := gin.Default()

	router.GET("/weather/:city", func(c *gin.Context) {
		city := c.Param("city")

		weather, err := fetchWeather(city)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		data, err := proto.Marshal(weather)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Data(http.StatusOK, "application/x-protobuf", data)
	})

	log.Fatal(router.Run(fmt.Sprintf(":%s", *port)))
}
