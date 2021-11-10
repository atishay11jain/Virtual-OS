package main

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2/canvas"

	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var f *os.File
var format beep.Format
var streamer beep.StreamSeekCloser
var pause bool = false
var wAudio fyne.Window

func myAudioPlayer() {

	go func(msg string) {
		fmt.Println(msg)

		if streamer == nil {

		} else {
			fmt.Println(fmt.Sprint(streamer.Len()))
		}
	}("going")

	time.Sleep(time.Second)

	wAudio = myApp.NewWindow("Audio Player")

	wAudio.Resize(fyne.NewSize(430, 400))

	img := canvas.NewImageFromFile("D:\\MyOsUsingFyne\\music.jpg")
	img.Resize(fyne.NewSize(400, 200))
	img.Move(fyne.NewPos(10, 0))

	tools := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
		}),

		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if !pause {
				pause = true
				speaker.Lock()
			} else if pause {
				pause = false
				speaker.Unlock()
			}
		}),

		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()
		}),
		widget.NewToolbarSpacer(),
	)

	txt1 := widget.NewLabel("Audio Player")
	txt1.Alignment = fyne.TextAlignCenter

	txt2 := widget.NewLabel("Play MP3....")
	txt2.Alignment = fyne.TextAlignCenter

	browse := widget.NewButton("Browse", func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {
			streamer, format, _ = mp3.Decode(uc)
			txt2.Text = uc.URI().Name()
			txt2.Refresh()
		}, wAudio)

		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
	})

	con := container.NewWithoutLayout(container.NewVBox(txt1, browse, txt2, tools))
	con.Move(fyne.NewPos(150, 200))

	wAudio.SetContent(container.NewWithoutLayout(con, img))

	wAudio.Show()

}
