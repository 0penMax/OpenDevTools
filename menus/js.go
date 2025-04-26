package menus

import (
	"github.com/pterm/pterm"
	"golang.design/x/clipboard"
	jsb "openDevTools/js/beautifier"
)

func showJsMenu() {
	ClearScreen()
	var m Menu

	m.title = "JS"
	m.desc = "All option use your clipboard for get value and for set result"

	m.navItems = append(m.navItems, navItem{
		name: "Beautifier",
		do:   showJsBeautifier,
	})

	m.show()
}

func showJsBeautifier() {
	ClearScreen()

	pterm.Info.Println("Read from clipboard")
	clipboardText := clipboard.Read(clipboard.FmtText)

	r, err := jsb.Beautify(string(clipboardText))
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	clipboard.Write(clipboard.FmtText, []byte(r))

	pterm.Info.Println("Beautified. Result copied to clipboard")

}
