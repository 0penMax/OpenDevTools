package main

import (
	menu "openDevTools/menus/mainMenu"
)

func main() {
	menu.BuildMenu()
	isFlagRun := menu.ProcessShortcuts()
	if !isFlagRun {
		menu.ShowMainMenu()
	}
}
