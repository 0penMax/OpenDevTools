package main

import (
	"github.com/pterm/pterm"
	"openDevTools/Unixtime"
	"openDevTools/models"
	"strconv"
)

func showUnixTimeMenu() {
	ClearScreen()
	var m Menu

	m.title = "Unixtime"
	m.desc = "Parse from/to unixtime."

	m.navItems = append(m.navItems, navItem{
		name: "To unixtime",
		do:   showToUnixtimeDialog,
	})

	m.navItems = append(m.navItems, navItem{
		name: "From unixtime",
		do:   showFromUnixtimeDialog,
	})

	m.show()
}

func showToUnixtimeDialog() {
	ClearScreen()

	pterm.Println("Parser use UTC timezone.")
	pterm.Println("Write your date in format - dd/mm/yyyy hh:mm:ss")
	textInput := pterm.DefaultInteractiveTextInput

	// Show the text input to the user and store the result.
	// The second return value (an error) is ignored with '_'.
	text, _ := textInput.Show()

	// Print a blank line for better readability in the output.
	pterm.Println()

	r, err := Unixtime.ParseStr(text)
	if err != nil {
		pterm.Error.Println(err)
		return
	}
	pterm.Println("result:")

	pterm.Println(r)

}

func showFromUnixtimeDialog() {
	ClearScreen()

	pterm.Println("Write your unixtime:")
	textInput := pterm.DefaultInteractiveTextInput

	// Show the text input to the user and store the result.
	// The second return value (an error) is ignored with '_'.
	text, _ := textInput.Show()

	// Print a blank line for better readability in the output.
	pterm.Println()

	n, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		pterm.Error.Println(err)
		return
	}

	r, err := Unixtime.ParseUnixTime(n)
	if err != nil {
		pterm.Error.Println(err)
		return
	}
	pterm.Println("result:")

	showParseUnixtimeTable(r)

}

func showParseUnixtimeTable(hashResult []models.ResultItem) {
	tableData := pterm.TableData{
		{"name", "value"},
	}

	for _, r := range hashResult {
		tableData = append(tableData, []string{r.Name, r.Value})
	}

	showTable(tableData)
}
