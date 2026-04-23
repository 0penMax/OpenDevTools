package hash

import (
	"openDevTools/HashGenerator"
	"openDevTools/menus/io"
	mModels "openDevTools/menus/models"
	"openDevTools/menus/utils"
	"openDevTools/models"

	"github.com/pterm/pterm"
)

func Menu() {

	utils.ShowMenu(
		"Hash generator",
		"Create MD5, SHA-1, SHA-256, SHA-384, SHA-512 from data.",
		[]mModels.NavItem{
			{Name: "From string", Do: showHashInputAndResult},
			{Name: "From file", Do: showSelectFile4HashAndResult},
		},
	)
}

func showSelectFile4HashAndResult() {
	utils.ClearScreen()
	// Create a default interactive text input with multi-line enabled.
	// This allows the user to input multiple lines of text.
	filepath, ok := io.OpenFileDialog(nil)
	if !ok {
		pterm.Warning.Println("openFileDialog cancelled")
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

	utils.ShowTable(tableHeader, hashResult)
}

func showHashInputAndResult() {
	utils.ClearScreen()
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
