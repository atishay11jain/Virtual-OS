package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/widget"
)

var cnt int = 1
var wNote fyne.Window

func myNotepad() {

	wNote = myApp.NewWindow("Notepad")

	wNote.Resize(fyne.NewSize(720, 400))

	data := widget.NewMultiLineEntry()

	fileList := container.NewVBox(
		widget.NewLabel(fmt.Sprintf("File %d", cnt)),
	)
	cnt++
	btn := widget.NewButton("New File", func() {
		data.SetText("")
		fileList.Add(widget.NewLabel(fmt.Sprintf("File %d", cnt)))
		cnt++
	})

	saveBtn := widget.NewButton("Save", func() {
		dbox := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				txtdata := []byte(data.Text)
				uc.Write(txtdata)
			}, wNote)

		dbox.SetFileName("File " + strconv.Itoa(cnt-1) + ".txt")
		dbox.Show()
	})

	openBtn := widget.NewButton("Open", func() {
		openDialog := dialog.NewFileOpen(func(ur fyne.URIReadCloser, _ error) {
			readData, _ := ioutil.ReadAll(ur)

			fixedData := fyne.NewStaticResource("New File", readData)

			place := widget.NewMultiLineEntry()

			place.SetText(string(fixedData.StaticContent))

			w := fyne.CurrentApp().NewWindow(
				string(fixedData.StaticName))
			con := container.NewVSplit(
				container.NewScroll(place),
				widget.NewButton("Save", func() {
					dBox := dialog.NewFileSave(
						func(uc fyne.URIWriteCloser, _ error) {
							fileToSave := []byte(place.Text)
							uc.Write(fileToSave)
						}, w)

					dBox.SetFileName("File " + strconv.Itoa(cnt) + ".txt")
					dBox.Show()
				}),
			)
			con.Offset = 1.25
			w.SetContent(con)

			w.Resize(fyne.NewSize(400, 400))

			w.Show()

		}, wNote)

		openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))

		openDialog.Show()
	})

	right := container.NewVSplit(data, saveBtn)
	right.Offset = 5.25

	leftwithoutbtn := container.NewBorder(btn, nil, nil, nil, fileList)
	left := container.NewVSplit(leftwithoutbtn, openBtn)
	left.Offset = 5.5
	content := container.NewHSplit(left, right)
	content.Offset = 0.25
	wNote.SetContent(content)
	wNote.Show()
}
