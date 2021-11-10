package main

import (
	"encoding/json"
	"fmt"

	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var wWeather fyne.Window

func myWeather() {

	wWeather = myApp.NewWindow("Weather App")
	wWeather.Resize(fyne.NewSize(250, 400))
	grad := canvas.NewVerticalGradient(color.RGBA{R: 173, G: 216, B: 230, A: 255}, color.White)
	image := canvas.NewImageFromFile("D:\\MyWeatherApp\\background.jpg")
	image.FillMode = canvas.ImageFillContain

	//drop down menu code

	apicity := "https://api.weatherapi.com/v1/current.json?key=0d2ae24d649940b4a9274231212910&q=bhopal&aqi=yes"

	data, err := http.Get(apicity)

	if err != nil {
		fmt.Println(err)
	}

	defer data.Body.Close()

	body, err := ioutil.ReadAll(data.Body)

	if err != nil {
		fmt.Println(err)
	}

	godata, err := UnmarshalWeather(body)

	if err != nil {
		fmt.Println(err)
	}

	label1 := canvas.NewText("Weather Details", color.White)
	label1.TextStyle.Bold = true
	label1.TextSize = 20
	label2 := canvas.NewText(fmt.Sprintf("Country : %s", godata.Location.Country), color.Black)
	label3 := canvas.NewText(fmt.Sprintf("Region : %s", godata.Location.Region), color.Black)

	label4 := canvas.NewText(fmt.Sprintf("City : %s", godata.Location.Name), color.Black)
	label5 := canvas.NewText(fmt.Sprintf("Temperature : %s", godata.Current.TempC), color.Black)
	label6 := canvas.NewText(fmt.Sprintf("Humidity : %s", godata.Current.WindKph), color.Black)

	dd := widget.NewSelect([]string{"Mumbai", "Delhi", "Kolkata", "Hyderabad", "Pune", "Lucknow", "Shimla", "Bhopal"}, func(s string) {
		if s == "Mumbai" {
			myMumbai(wWeather)
		} else if s == "Bhopal" {
			myBhopal(wWeather)
		} else if s == "Delhi" {
			myDelhi(wWeather)
		} else if s == "Kolkata" {
			myKolkata(wWeather)
		} else if s == "Hyderabad" {
			myHyderabad(wWeather)
		} else if s == "Pune" {
			myPune(wWeather)
		} else if s == "Lucknow" {
			myLucknow(wWeather)
		} else if s == "Shimla" {
			myShimla(wWeather)
		}
	})

	dd.PlaceHolder = "Bhopal"

	details := container.NewCenter(container.NewGridWithRows(9, dd, layout.NewSpacer(), label1, layout.NewSpacer(), label2, label3, label4, label5, label6))

	myTheme := container.NewMax(grad, image, details)

	wWeather.SetContent(
		container.NewGridWithColumns(1, myTheme),
	)

	wWeather.Show()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Current struct {
	LastUpdatedEpoch int64              `json:"last_updated_epoch"`
	LastUpdated      string             `json:"last_updated"`
	TempC            json.Number        `json:"temp_c"`
	TempF            float64            `json:"temp_f"`
	IsDay            int64              `json:"is_day"`
	Condition        Condition          `json:"condition"`
	WindMph          float64            `json:"wind_mph"`
	WindKph          json.Number        `json:"wind_kph"`
	WindDegree       int64              `json:"wind_degree"`
	WindDir          string             `json:"wind_dir"`
	PressureMB       int64              `json:"pressure_mb"`
	PressureIn       float64            `json:"pressure_in"`
	PrecipMm         int64              `json:"precip_mm"`
	PrecipIn         int64              `json:"precip_in"`
	Humidity         int64              `json:"humidity"`
	Cloud            int64              `json:"cloud"`
	FeelslikeC       float64            `json:"feelslike_c"`
	FeelslikeF       float64            `json:"feelslike_f"`
	VisKM            int64              `json:"vis_km"`
	VisMiles         int64              `json:"vis_miles"`
	Uv               int64              `json:"uv"`
	GustMph          float64            `json:"gust_mph"`
	GustKph          int64              `json:"gust_kph"`
	AirQuality       map[string]float64 `json:"air_quality"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int64  `json:"code"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}
