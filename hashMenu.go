package main

import (
	"github.com/pterm/pterm"
	"github.com/sqweek/dialog"
	"openDevTools/HashGenerator"
	"openDevTools/models"
)

func showHashMenu() {
	ClearScreen()
	var m Menu

	m.title = "Hash Generator"
	m.desc = "Create MD5, SHA-1, SHA-256, SHA-384, SHA-512 from data."

	m.navItems = append(m.navItems, navItem{
		name: "From string",
		do:   showHashInputAndResult,
	})

	m.navItems = append(m.navItems, navItem{
		name: "From file",
		do:   showSelectFile4HashAndResult,
	})

	m.show()
}

func showSelectFile4HashAndResult() {
	ClearScreen()
	// Create a default interactive text input with multi-line enabled.
	// This allows the user to input multiple lines of text.
	filepath, err := dialog.File().Filter("any file", "*").Load()

	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	pterm.Info.Println("filepath:", filepath)

	result, err := HashGenerator.FromFile(filepath)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	showHashTable(result)

}

func showHashTable(hashResult []models.ResultItem) {
	tableData := pterm.TableData{
		{"hash name", "value"},
	}

	for _, r := range hashResult {
		tableData = append(tableData, []string{r.Name, r.Value})
	}

	showTable(tableData)
}

func showHashInputAndResult() {
	ClearScreen()
	// Create a default interactive text input with multi-line enabled.
	// This allows the user to input multiple lines of text.
	textInput := pterm.DefaultInteractiveTextInput.WithMultiLine()

	// Show the text input to the user and store the result.
	// The second return value (an error) is ignored with '_'.
	text, _ := textInput.Show()

	// Print a blank line for better readability in the output.
	pterm.Println()

	result := HashGenerator.FromString(text)

	showHashTable(result)

}

func showTable(tableData pterm.TableData) {
	// Create a table with the defined data.
	// The table has a header and is boxed.
	pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
}
