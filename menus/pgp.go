package menus

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/sqweek/dialog"
	"openDevTools/pgp"
	"os"
)

func showPgpMenu() {
	ClearScreen()
	var m Menu

	m.title = "PGP"
	m.desc = ""

	m.navItems = append(m.navItems, navItem{
		name: "Read data from key(public/private)",
		do:   showPgpPublicKeyInputAndResult,
	})

	m.show()
}

func showPgpPublicKeyInputAndResult() {
	ClearScreen()

	filepath, err := dialog.File().Filter("any file", "*").Load()

	if err != nil {
		pterm.Warning.Println(err)
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

	showTable(tableHeader, result)

}
