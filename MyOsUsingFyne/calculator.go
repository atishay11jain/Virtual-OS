package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/knetic/govaluate"
)

var str = ""
var content = widget.NewLabel(str)
var wCalc fyne.Window

func cont(txt string) {
	content.SetText(txt)
}

func myCalculator() {

	wCalc = myApp.NewWindow("Calculator")

	wCalc.Resize(fyne.NewSize(400, 300))

	var slice = make([]string, 0)

	var hisstr string = ""

	var history = widget.NewLabel(hisstr)

	var showhis = true

	hisBtn := widget.NewButton("history", func() {
		if showhis {
			showhis = false
			for i := len(slice) - 1; i >= 0; i-- {
				hisstr = hisstr + slice[i]
				hisstr = hisstr + "\n"
			}
			history.SetText(hisstr)
		}
	})

	backBtn := widget.NewButton("back", func() {
		if len(str) > 0 {
			str = str[:len(str)-1]
			cont(str)
		}
	})
	clearBtn := widget.NewButton("clear", func() {
		str = ""
		cont(str)
		hisstr = ""
		history.SetText(hisstr)
	})
	openBtn := widget.NewButton("(", func() {
		str = str + "("
		cont(str)
	})
	closeBtn := widget.NewButton(")", func() {
		str = str + ")"
		cont(str)
	})
	divideBtn := widget.NewButton("/", func() {
		str = str + "/"
		cont(str)
	})
	sevenBtn := widget.NewButton("7", func() {
		str = str + "7"
		cont(str)
	})
	eightBtn := widget.NewButton("8", func() {
		str = str + "8"
		cont(str)
	})
	nineBtn := widget.NewButton("9", func() {
		str = str + "9"
		cont(str)
	})
	mulyiplyBtn := widget.NewButton("*", func() {
		str = str + "*"
		cont(str)
	})
	fourBtn := widget.NewButton("4", func() {
		str = str + "4"
		cont(str)
	})
	fiveBtn := widget.NewButton("5", func() {
		str = str + "5"
		cont(str)
	})

	sixBtn := widget.NewButton("6", func() {
		str = str + "6"
		cont(str)
	})
	minusBtn := widget.NewButton("-", func() {
		str = str + "-"
		cont(str)
	})
	oneBtn := widget.NewButton("1", func() {
		str = str + "1"
		cont(str)
	})
	twoBtn := widget.NewButton("2", func() {
		str = str + "2"
		cont(str)
	})

	threeBtn := widget.NewButton("3", func() {
		str = str + "3"
		cont(str)
	})

	plusBtn := widget.NewButton("+", func() {
		str = str + "+"
		cont(str)
	})

	zeroBtn := widget.NewButton("0", func() {
		str = str + "0"
		cont(str)
	})

	decBtn := widget.NewButton(".", func() {
		str = str + "."
		cont(str)
	})

	equalBtn := widget.NewButton("=", func() {
		temp := str
		exp, err := govaluate.NewEvaluableExpression(str)
		if err == nil {
			result, err := exp.Evaluate(nil)
			if err == nil {
				str = strconv.FormatFloat(result.(float64), 'f', -1, 64)
			} else {
				str = "error"
			}
		} else {
			str = "error"
		}
		slice = append(slice, temp+" = "+str)
		cont(str)
	})

	wCalc.SetContent(container.NewVBox(
		content,
		history,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				hisBtn,
				backBtn,
			),
			container.NewGridWithColumns(4,
				clearBtn,
				openBtn,
				closeBtn,
				divideBtn,
			),
			container.NewGridWithColumns(4,
				sevenBtn,
				eightBtn,
				nineBtn,
				mulyiplyBtn),
			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				minusBtn),
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				plusBtn),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					decBtn,
				),

				equalBtn),
		),
	))

	wCalc.Show()
}
