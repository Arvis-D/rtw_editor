package main

type MainFragment struct {
	app *App
}

func CreateMainFragment(app *App) *MainFragment {
	mainFragment := &MainFragment{}
	mainFragment.app = app

	return mainFragment
}
