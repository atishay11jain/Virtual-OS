package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

var win fyne.Window

func myIP() {

	win = myApp.NewWindow("Information")
	win.Resize(fyne.NewSize(400, 400))
	label1 := widget.NewLabel("FInd IP Address")
	label2 := widget.NewLabel("My Ip Address ...")
	label_IP := widget.NewLabel("...")
	label_Country := widget.NewLabel("...")
	label_Region := widget.NewLabel("...")
	label_city := widget.NewLabel("...")

	btn := widget.NewButton("Find Details", func() {
		fmt.Println(myIp())
		label_IP.Text = myIp()[0]
		label_IP.Refresh()
		label_Country.Text = myIp()[1]
		label_Country.Refresh()
		label_Region.Text = myIp()[2]
		label_Region.Refresh()
		label_city.Text = myIp()[3]
		label_city.Refresh()
	})

	win.SetContent(container.NewVBox(
		label1,
		label2,
		label_IP,
		label_Country,
		label_Region,
		label_city,
		btn))

	win.Show()

}

func myIp() []string {

	res, err := http.Get("http://ip-api.com/json/?fields=61439")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	godata, err := UnmarshalIP(body)

	slice := make([]string, 0)
	slice = append(slice, godata.Query, godata.Country, godata.RegionName, godata.City)

	return slice
}

func UnmarshalIP(data []byte) (IP, error) {
	var r IP
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *IP) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type IP struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}
