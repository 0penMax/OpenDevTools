package main

import "openDevTools/menus"

func main() {
	menus.BuildMenu()
	isFlagRun := menus.ProcessShortcuts()
	if !isFlagRun {
		menus.ShowMainMenu()
	}
}
