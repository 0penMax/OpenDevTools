package menus

import (
	"github.com/pterm/pterm"
	"golang.design/x/clipboard"
	"openDevTools/unicode"
)

func showUnicodeDEMenu() {
	ClearScreen()
	var m Menu

	m.title = "Unicode Decode/encode"
	m.desc = "Encode/decode unicode string from clipboard."

	m.navItems = append(m.navItems, navItem{
		name: "Encode",
		do:   showUnicodeEncodeDialog,
	})

	m.navItems = append(m.navItems, navItem{
		name: "Decode",
		do:   showUnicodeDecodeDialog,
	})

	m.show()
}

func showUnicodeDecodeDialog() {
	ClearScreen()
	pterm.Info.Println("Read from clipboard")
	clipboardText := clipboard.Read(clipboard.FmtText)

	r, err := unicode.DecodeUnicode(string(clipboardText))
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	clipboard.Write(clipboard.FmtText, []byte(r))

	pterm.Info.Println("Result copied to clipboard")
}

func showUnicodeEncodeDialog() {
	ClearScreen()

	pterm.Info.Println("Read from clipboard")
	clipboardText := clipboard.Read(clipboard.FmtText)

	r := unicode.EncodeToUnicode(string(clipboardText))

	clipboard.Write(clipboard.FmtText, []byte(r))

	pterm.Info.Println("Result copied to clipboard")

}
