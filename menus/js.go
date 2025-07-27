package menus

import (
	jsb "openDevTools/js/beautifier"
	jsm "openDevTools/js/minify"
)

func showJsMenu() {
	ClearScreen()
	var m Menu

	m.title = "JS"
	m.desc = ""

	m.navItems = append(m.navItems, navItem{
		name: "Beautifier",
		do:   showJsBeautifier,
	})

	m.navItems = append(m.navItems, navItem{
		name: "Minify",
		do:   showJsMinify,
	})

	m.show()
}

func showJsBeautifier() {
	showInputMenu(jsb.Beautify)
}

func showJsMinify() {
	showInputMenu(jsm.Minify)
}
