package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()

var myWin fyne.Window = myApp.NewWindow("MY OS")

func main() {

	myWin.Resize(fyne.NewSize(1210, 720))

	pi1, _ := LoadResourceFromPath("D:\\MyOsUsingFyne\\weathericon.jpg")

	ic1 := widget.NewButtonWithIcon("", pi1, func() {
		myWeather()
	})

	ic1.Resize(fyne.NewSize(50, 50))
	ic1.Move(fyne.NewPos(320, 650))

	pi2, _ := LoadResourceFromPath("D:\\AppsUsingFyne\\galleryicon.jpg")

	ic2 := widget.NewButtonWithIcon("", pi2, func() {
		myGallery()
	})

	ic2.Resize(fyne.NewSize(50, 50))
	ic2.Move(fyne.NewPos(400, 650))

	pi3, _ := LoadResourceFromPath("D:\\MyOsUsingFyne\\calculatoricon.jpg")

	ic3 := widget.NewButtonWithIcon("", pi3, func() {
		myCalculator()
	})

	ic3.Resize(fyne.NewSize(50, 50))
	ic3.Move(fyne.NewPos(480, 650))

	pi4, _ := LoadResourceFromPath("D:\\MyOsUsingFyne\\notepadicon.jpg")

	ic4 := widget.NewButtonWithIcon("", pi4, func() {
		myNotepad()
	})

	ic4.Resize(fyne.NewSize(50, 50))
	ic4.Move(fyne.NewPos(560, 650))

	pi5, _ := LoadResourceFromPath("D:\\MyOsUsingFyne\\diceicon.jpg")

	ic5 := widget.NewButtonWithIcon("", pi5, func() {
		myGame()
	})

	ic5.Resize(fyne.NewSize(50, 50))
	ic5.Move(fyne.NewPos(640, 650))

	pi6, _ := LoadResourceFromPath("D:\\MyOsUsingFyne\\musicicon.jpg")

	ic6 := widget.NewButtonWithIcon("", pi6, func() {
		myAudioPlayer()
	})

	ic6.Resize(fyne.NewSize(50, 50))
	ic6.Move(fyne.NewPos(720, 650))

	pi7, _ := LoadResourceFromPath("D:\\MyOsUsingFyne\\info.png")

	ic7 := widget.NewButtonWithIcon("", pi7, func() {
		myValid()
	})

	ic7.Resize(fyne.NewSize(50, 50))
	ic7.Move(fyne.NewPos(800, 650))

	img := canvas.NewImageFromFile("D:\\MyOsUsingFyne\\windows1.jpg")
	img.Resize(fyne.NewSize(1200, 640))
	img.Move(fyne.NewPos(0, 0))

	myWin.SetContent(container.NewWithoutLayout(img, ic1, ic2, ic3, ic4, ic5, ic6, ic7))
	myWin.ShowAndRun()
}

type Resource interface {
	Name() string
	Content() []byte
}

// StaticResource is a bundled resource compiled into the application.
// These resources are normally generated by the fyne_bundle command included in
// the Fyne toolkit.
type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

// Name returns the unique name of this resource, usually matching the file it
// was generated from.
func (r *StaticResource) Name() string {
	return r.StaticName
}

// Content returns the bytes of the bundled resource, no compression is applied
// but any compression on the resource is retained.
func (r *StaticResource) Content() []byte {
	return r.StaticContent
}

// NewStaticResource returns a new static resource object with the specified
// name and content. Creating a new static resource in memory results in
// sharable binary data that may be serialised to the location returned by
// CachePath().
func NewStaticResource(name string, content []byte) *StaticResource {
	return &StaticResource{
		StaticName:    name,
		StaticContent: content,
	}
}

// LoadResourceFromPath creates a new StaticResource in memory using the contents of the specified file.
func LoadResourceFromPath(path string) (Resource, error) {
	bytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	name := filepath.Base(path)
	return NewStaticResource(name, bytes), nil
}

// LoadResourceFromURLString creates a new StaticResource in memory using the body of the specified URL.
func LoadResourceFromURLString(urlStr string) (Resource, error) {
	res, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	name := filepath.Base(urlStr)
	return NewStaticResource(name, bytes), nil
}