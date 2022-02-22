package main

//"fyne.io/fyne/v2/theme"

func main() {
	app := CreateApp()
	app.Run(CreateMainFragment(app.Context()))
}
