package main

import (
	"image/color"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var w fyne.Window

func myGallery() {

	path := "C:\\Users\\atish\\OneDrive\\Pictures\\temp"

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}
	tabs := container.NewAppTabs()

	tabs.SetTabLocation(container.TabLocationLeading)
	for _, file := range files {
		if !file.IsDir() {
			strArr := strings.Split(file.Name(), ".")

			if strArr[1] == "png" || strArr[1] == "jpeg" || strArr[1] == "jpg" {
				image := canvas.NewImageFromFile(path + "\\" + file.Name())
				image.FillMode = canvas.ImageFillContain
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}

	}
	back := canvas.NewLinearGradient(color.White, color.RGBA{R: 202, G: 236, B: 241, A: 255}, 45)
	w = myApp.NewWindow("Gallery")

	w.Resize(fyne.NewSize(1080, 720))

	w.SetContent(container.NewMax(back, tabs))

	w.Show()
}
