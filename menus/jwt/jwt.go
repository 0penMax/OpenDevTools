package jwt

import (
	"openDevTools/jwt"
	"openDevTools/menus/models"
	"openDevTools/menus/utils"

	"github.com/pterm/pterm"
)

func Menu() {
	utils.ShowMenu(
		"JWT",
		"Work with jwt tokens.",
		[]models.NavItem{
			{Name: "Read", Do: readJwtToken},
		})
}

func readJwtToken() {
	utils.ClearScreen()

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
	utils.ShowTable(tableHeader, result.Header)
	pterm.Println("")
	pterm.Bold.Println("Payload:")
	utils.ShowTable(tableHeader, result.Payload)
	pterm.Println("")
	pterm.Bold.Println("Signature:")
	pterm.Println(result.Signature)

}
