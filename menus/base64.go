package menus

import (
	"openDevTools/base64"
)

func showBase64StringMenu() {
	ClearScreen()
	var m Menu

	m.title = "Base64 string"
	m.desc = "Encode/decode base64 string."

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
	showInputMenu(nilErrorWrapper(base64.Encode))
}

func showBase64DecodeDialog() {
	showInputMenu(func(s string) (string, error) {
		b, err := base64.Decode(s)
		return string(b), err
	})
}
