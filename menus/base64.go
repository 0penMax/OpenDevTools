package menus

import (
	"github.com/pterm/pterm"
	"openDevTools/base64"
)

func showBase64StringMenu() {
	ClearScreen()
	var m Menu

	m.title = "Base64 string"
	m.desc = "Encode/decode base64  string."

	m.navItems = append(m.navItems, navItem{
		name: "Encode",
		do:   showBase64EncodeDialog,
	})

	m.navItems = append(m.navItems, navItem{
		name: "Decode",
		do:   showBase64DecodeDialog,
	})

	m.show()
}

func showBase64EncodeDialog() {
	ClearScreen()

	//pterm.Println("Write your date in format - dd/mm/yyyy hh:mm:ss")
	textInput := pterm.DefaultInteractiveTextInput

	// Show the text input to the user and store the result.
	// The second return value (an error) is ignored with '_'.
	text, _ := textInput.Show()

	// Print a blank line for better readability in the output.
	pterm.Println()

	r := base64.Encode(text)

	pterm.Println(r)

}

func showBase64DecodeDialog() {
	ClearScreen()

	pterm.Println("Write your unixtime:")
	textInput := pterm.DefaultInteractiveTextInput

	// Show the text input to the user and store the result.
	// The second return value (an error) is ignored with '_'.
	text, _ := textInput.Show()

	// Print a blank line for better readability in the output.
	pterm.Println()

	r, err := base64.Decode(text)
	if err != nil {
		pterm.Error.Println(err)
		return
	}
	pterm.Println(r)

}
