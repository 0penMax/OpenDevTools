package strings

import (
	"fmt"
	"openDevTools/lorem"
	"openDevTools/menus/base64"
	"openDevTools/menus/io"
	"openDevTools/menus/models"
	"openDevTools/menus/unicode"
	"openDevTools/menus/utils"
	"strconv"

	"github.com/pterm/pterm"
)

func Menu() {

	utils.ShowMenu(
		"Strings",
		"Working with strings.",
		[]models.NavItem{
			{
				Name: "Length",
				Do: func() {
					io.ShowInputMenu(io.NilErrorWrapper(
						func(value string) string {
							return fmt.Sprint(len(value)) //calculate string length
						},
					))
				},
			},
			{Name: "Lorem generator", Do: showLorem},
			{Name: "Unicode string", Do: unicode.Menu},
			{Name: "Base64 string", Do: base64.StringMenu},
		})
}

func showLorem() {
	utils.ClearScreen()

	textInput := pterm.DefaultInteractiveTextInput.WithDefaultText("enter word count")
	text, _ := textInput.Show()

	count, err := strconv.Atoi(text)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	io.ShowOutputMenu([]byte(lorem.Generate(count)))
}
