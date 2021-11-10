package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func myHyderabad(w fyne.Window) {

	grad := canvas.NewVerticalGradient(color.RGBA{R: 173, G: 216, B: 230, A: 255}, color.White)
	image := canvas.NewImageFromFile("D:\\MyWeatherApp\\background.jpg")
	image.FillMode = canvas.ImageFillContain

	//drop down menu code

	apicity := "https://api.weatherapi.com/v1/current.json?key=0d2ae24d649940b4a9274231212910&q=hyderabad&aqi=yes"

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

	dd := widget.NewSelect([]string{"Mumbai", "Delhi", "Kolkata", "Hyderabad", "Pune", "Bhopal", "Lucknow", "Shimla"}, func(s string) {
		if s == "Mumbai" {
			myMumbai(w)
		} else if s == "Bhopal" {
			myBhopal(w)
		} else if s == "Delhi" {
			myDelhi(w)
		} else if s == "Kolkata" {
			myKolkata(w)
		} else if s == "Hyderabad" {
			myHyderabad(w)
		} else if s == "Pune" {
			myPune(w)
		} else if s == "Lucknow" {
			myLucknow(w)
		} else if s == "Shimla" {
			myShimla(w)
		}
	})

	dd.PlaceHolder = "Hyderabad"

	details := container.NewCenter(container.NewGridWithRows(9, dd, layout.NewSpacer(), label1, layout.NewSpacer(), label2, label3, label4, label5, label6))

	myTheme := container.NewMax(grad, image, details)

	w.SetContent(
		container.NewGridWithColumns(1, myTheme),
	)

	w.Show()

}
