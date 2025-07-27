package menus

import (
	"github.com/pterm/pterm"
	"golang.design/x/clipboard"
)

func processFromClipboard(processF func(string) (string, error)) {
	ClearScreen()
	pterm.Info.Println("Read from clipboard")
	clipboardText := clipboard.Read(clipboard.FmtText)

	r, err := processF(string(clipboardText))
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	clipboard.Write(clipboard.FmtText, []byte(r))

	pterm.Info.Println("Result copied to clipboard")
}
