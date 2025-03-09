package menus

import (
	"github.com/pterm/pterm"
	"golang.design/x/clipboard"
	"openDevTools/lorem"
	"strconv"
)

func showLoremMenu() {
	ClearScreen()

	textInput := pterm.DefaultInteractiveTextInput.WithDefaultText("enter word count")
	text, _ := textInput.Show()

	count, err := strconv.Atoi(text)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	clipboard.Write(clipboard.FmtText, []byte(lorem.Generate(count)))

	pterm.Println("copied to clipboard")

}
