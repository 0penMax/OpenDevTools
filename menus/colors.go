package menus

import (
	"github.com/pterm/pterm"
	"openDevTools/colors"
	"openDevTools/models"
)

func showColorTransformMenu() {
	ClearScreen()
	var m Menu

	m.title = "Color transformation"
	m.desc = "Transform one of support colors formats to all others. Support most popular colors format: HEX, RGB, RGBA, HSL, HSLA and colors name. "

	m.show()

	isColorCorrect := false
	var r []models.ResultItem
	var err error

	for !isColorCorrect {
		textInput := pterm.DefaultInteractiveTextInput

		text, _ := textInput.Show()

		r, err = colors.ConvertColor(text)
		if err != nil {
			pterm.Error.Println(err)
			continue
		}
		isColorCorrect = true
	}

	pterm.Println()

	tableHeader := pterm.TableData{
		{"name", "value"},
	}

	showTable(tableHeader, r)
}
