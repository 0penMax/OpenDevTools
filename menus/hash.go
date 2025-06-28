package menus

import (
	"github.com/nixinwang/dialog"
	"github.com/pterm/pterm"
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
	tableHeader := pterm.TableData{
		{"hash name", "value"},
	}

	showTable(tableHeader, hashResult)
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

func showTable(tableHeader pterm.TableData, tableData []models.ResultItem) {
	for _, r := range tableData {
		tableHeader = append(tableHeader, []string{r.Name, r.Value})
	}
	pterm.DefaultTable.WithHasHeader().WithData(tableHeader).Render()
}
