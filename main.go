package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")

	content := widget.NewButton("click me", func() {
		myWindow.SetContent(widget.NewLabel("Changed"))
	})
	content2 := widget.NewButton("click me", func() {
		log.Println("tapped")
	})

	hbox := container.New(layout.NewHBoxLayout(), content, content2)

	//content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	//})

	myWindow.SetContent(hbox)
	myWindow.ShowAndRun()
}
