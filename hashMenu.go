package main

import (
	"fmt"
	"github.com/pterm/pterm"
	"openDevTools/HashGenerator"
)

func showHashMenu() {
	ClearScreen()
	var m Menu

	m.title = "Hash Generator"
	m.desc = "Create MD5, SHA-1, SHA-256, SHA-384, SHA-512 from data."

	m.navItems = append(m.navItems, navItem{
		name: "From string",
		do:   showHashInputAndWResult,
	})

	m.navItems = append(m.navItems, navItem{
		name: "From file",
		do:   func() { fmt.Println("make this part") },
	})

	m.show()
}

func showHashInputAndWResult() {
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

	tableData := pterm.TableData{
		{"hash name", "value"},
	}

	for _, r := range result {
		tableData = append(tableData, []string{r.Name, r.Value})
	}

	showTable(tableData)

}

func showTable(tableData pterm.TableData) {
	// Create a table with the defined data.
	// The table has a header and is boxed.
	pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
}
