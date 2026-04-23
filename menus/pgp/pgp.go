package pgp

import (
	"fmt"
	"openDevTools/menus/io"
	"openDevTools/menus/models"
	"openDevTools/menus/utils"
	"openDevTools/pgp"
	"os"

	"github.com/pterm/pterm"
)

func Menu() {
	utils.ShowMenu(
		"PGP", "",
		[]models.NavItem{
			{Name: "Read data from key(public/private)", Do: showPgpPublicKeyInputAndResult},
		},
	)

}

func showPgpPublicKeyInputAndResult() {
	utils.ClearScreen()

	filepath, ok := io.OpenFileDialog(nil)
	if !ok {
		pterm.Warning.Println("openFileDialog cancelled")
		return
	}

	pterm.Info.Println("filepath:", filepath)

	file, err := os.Open(filepath)
	if err != nil {
		pterm.Error.Println(fmt.Errorf("failed to open file: %w", err))
		return
	}
	defer file.Close()

	// Print a blank line for better readability in the output.
	pterm.Println()

	result, err := pgp.ReadPublicKeyData(file)
	if err != nil {
		pterm.Error.Println(err)
		return
	}

	tableHeader := pterm.TableData{
		{"name", "value"},
	}

	utils.ShowTable(tableHeader, result)

}
