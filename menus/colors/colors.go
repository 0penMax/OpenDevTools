package colors

import (
	"openDevTools/colors"
	"openDevTools/menus/utils"
	"openDevTools/models"

	"github.com/pterm/pterm"
)

func Menu() {
	utils.ShowMenu(
		"Color transformation",
		"Transform one of support colors formats to all others. Support most popular colors format: HEX, RGB, RGBA, HSL, HSLA and colors name. ",
		nil,
	)

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

	utils.ShowTable(tableHeader, r)
}
