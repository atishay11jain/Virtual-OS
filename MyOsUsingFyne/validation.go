package main

import (
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

var wValid fyne.Window

func myValid() {

	wValid = myApp.NewWindow("Settings")

	wValid.Resize(fyne.NewSize(400, 400))

	label := widget.NewLabel("")

	form := widget.NewForm(
		widget.NewFormItem("Username", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)

	form.OnCancel = func() {

	}

	form.OnSubmit = func() {
		myIP()
	}

	wValid.SetContent(container.NewVBox(form, label))
	wValid.Show()
}
