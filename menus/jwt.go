package menus

import (
	"github.com/pterm/pterm"
	"openDevTools/jwt"
)

func showJwtMenu() {
	ClearScreen()
	var m Menu

	m.title = "JWT"
	m.desc = "Work with jwt tokens."

	m.navItems = append(m.navItems, navItem{
		name: "read",
		do:   readJwtToken,
	})

	m.show()
}

func readJwtToken() {
	ClearScreen()

	textInput, err := pterm.DefaultInteractiveTextInput.WithMultiLine().WithDefaultText("jwt token").Show()
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	result, err := jwt.Read(textInput)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	tableHeader := pterm.TableData{
		{"Name", "Value"},
	}

	pterm.Bold.Println("Header:")
	showTable(tableHeader, result.Header)
	pterm.Println("")
	pterm.Bold.Println("Payload:")
	showTable(tableHeader, result.Payload)
	pterm.Println("")
	pterm.Bold.Println("Signature:")
	pterm.Println(result.Signature)

}
