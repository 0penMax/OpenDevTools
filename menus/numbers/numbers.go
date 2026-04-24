package numbers

import (
	"fmt"
	"openDevTools/menus/utils"
	"openDevTools/numbers"

	"github.com/pterm/pterm"
)

func Menu() {

	utils.ShowMenu(
		"Numbers converter",
		"Convert your number in binary, octal, hex and decimal presentation, use the examples from the table below",
		nil,
	)

	tableHeader := pterm.TableData{
		{"name", "base", "represent"},
	}

	var data pterm.TableData

	data = append(data, []string{"binary", "2", "0b11111111"})
	data = append(data, []string{"octal", "8", "0o377"})
	data = append(data, []string{"hex", "16", "0xff"})
	data = append(data, []string{"decimal", "10", "255"})

	pterm.Println()

	utils.ShowTable(tableHeader, data)

	isColorCorrect := false
	var r []numbers.Number
	var err error

	for !isColorCorrect {
		textInput := pterm.DefaultInteractiveTextInput

		text, _ := textInput.Show()

		r, err = numbers.Parse(text)
		if err != nil {
			pterm.Error.Println(err)
			continue
		}
		isColorCorrect = true
	}

	pterm.Println()
	pterm.Println("Results:")

	data = pterm.TableData{}

	for _, row := range r {
		data = append(data,
			[]string{row.BaseName, fmt.Sprint(row.Base), row.Repr},
		)
	}

	utils.ShowTable(tableHeader, data)
}
