package unixTime

import (
	"openDevTools/Unixtime"
	"openDevTools/menus/io"
	"openDevTools/menus/models"
	"openDevTools/menus/utils"
	"strconv"

	"github.com/pterm/pterm"
)

func Menu() {
	utils.ShowMenu(
		"Unixtime",
		"Parse from/to unixtime.",
		[]models.NavItem{
			{Name: "Now", Do: showNowDialog},
			{Name: "To unixtime", Do: showToDialog},
			{Name: "From unixtime", Do: showFromDialog},
		},
	)
}

func showToDialog() {
	utils.ClearScreen()

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

	pterm.Println(r)
}

func showNowDialog() {
	io.ShowOutputMenu([]byte(Unixtime.Now()))
}

func showFromDialog() {
	utils.ClearScreen()

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

	tableHeader := pterm.TableData{
		{"name", "value"},
	}

	utils.ShowTable(tableHeader, utils.ParseResultItems4Table(r))

}
