package menus

import (
	"github.com/nixinwang/dialog"
	"github.com/pterm/pterm"
	"os"
)

func showOutputMenu(value string) {
	ClearScreen()
	pterm.DefaultHeader.WithFullWidth().Println("Select output method")
	pterm.Println()
	navMap := make(map[string]func(value string))
	var menu []string

	os := "On screen"
	navMap[os] = showResult
	menu = append(menu, os)

	clp := "Save 2 clipboard"
	navMap[clp] = save2Clipboard
	menu = append(menu, clp)

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithMaxHeight(10).WithOptions(menu).Show()

	f := navMap[selectedOption]
	ClearScreen()
	f(value)
}

func showResult(value string) {
	pterm.Info.Println("Result:")
	pterm.Println(value)
}

func showImgOutputMenu(data []byte) {
	ClearScreen()
	pterm.DefaultHeader.WithFullWidth().Println("Select output method")
	pterm.Println()
	navMap := make(map[string]func(value []byte))
	var menu []string

	sf := "Save file"
	navMap[sf] = showSaveFile
	menu = append(menu, sf)

	clp := "Save 2 clipboard"
	navMap[clp] = saveImg2Clipboard
	menu = append(menu, clp)

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithMaxHeight(10).WithOptions(menu).Show()

	f := navMap[selectedOption]
	ClearScreen()
	f(data)
}

func showSaveFile(data []byte) {
	ClearScreen()
	filepath, err := dialog.File().Save()

	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	if err = os.WriteFile(filepath, data, 0644); err != nil {
		pterm.Warning.Println(err)
		return
	}

	pterm.Println("saved file:", filepath)
}
