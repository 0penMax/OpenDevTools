package menus

import (
	"openDevTools/unicode"
)

func showUnicodeDEMenu() {
	ClearScreen()
	var m Menu

	m.title = "Unicode Decode/encode"
	m.desc = "Encode/decode unicode string."

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
	showInputMenu(unicode.DecodeUnicode)
}

func showUnicodeEncodeDialog() {
	showInputMenu(nilErrorWrapper(unicode.EncodeToUnicode))
}
