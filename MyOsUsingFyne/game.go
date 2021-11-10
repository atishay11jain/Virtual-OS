package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2/canvas"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var wGame fyne.Window

func myGame() {

	wGame = myApp.NewWindow("Dice Game")

	wGame.Resize(fyne.NewSize(600, 400))
	myimage := canvas.NewImageFromFile("D:\\DiceGameUsingFyne\\dice1.jpg")
	myimage.FillMode = canvas.ImageFillOriginal

	c1 := canvas.NewCircle(color.RGBA{R: 0, G: 0, B: 0, A: 255})
	c1.Resize(fyne.NewSize(50, 50))
	c2 := canvas.NewCircle(color.RGBA{R: 0, G: 0, B: 0, A: 255})
	c2.Resize(fyne.NewSize(50, 50))
	c3 := canvas.NewCircle(color.RGBA{R: 0, G: 0, B: 0, A: 255})
	c3.Resize(fyne.NewSize(50, 50))

	btn := widget.NewButton("Roll Dice", func() {
		randomnum := rand.Intn(6) + 1

		Reset(c1, c2, c3)

		if randomnum <= 2 {
			c1.FillColor = color.RGBA{R: 0, G: 255, B: 0, A: 255}
			c1.Refresh()
		} else if randomnum > 2 && randomnum <= 5 {
			c2.FillColor = color.RGBA{R: 0, G: 0, B: 255, A: 255}
			c2.Refresh()
		} else {
			c3.FillColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
			c3.Refresh()
		}

		myimage.File = fmt.Sprintf("D:\\DiceGameUsingFyne\\dice%d.jpg", randomnum)
		myimage.Refresh()

	})

	recRed := canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	recRed.Resize(fyne.NewSize(1, 1))
	recGreen := canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 0, A: 255})
	recBlue := canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 255, A: 255})

	lefttemp := container.NewGridWithRows(2,
		container.NewCenter(myimage),
		btn,
	)

	left := container.NewVBox(lefttemp,
		recGreen,
		widget.NewLabel("If number is 1 or 2 then green light will glow"),
		recBlue,
		widget.NewLabel("If number is 3  4 or 5 then blue light will glow"),
		recRed,
		widget.NewLabel("If number is 6 then red light will glow"),
	)
	right := container.NewGridWithRows(7,
		layout.NewSpacer(),
		c1,
		layout.NewSpacer(),
		c2,
		layout.NewSpacer(),
		c3,
		layout.NewSpacer(),
	)
	content := container.NewHSplit(left, right)
	content.Offset = 0.4
	wGame.SetContent(content)

	wGame.Show()
}

func Reset(c1, c2, c3 *canvas.Circle) {

	c1.FillColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}

	c2.FillColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}

	c3.FillColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}

	c1.Refresh()
	c2.Refresh()
	c3.Refresh()
}
