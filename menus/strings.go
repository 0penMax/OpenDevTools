package menus

import (
	"fmt"
)

func showStringsMenu() {
	ClearScreen()
	var m Menu

	m.title = "Strings"
	m.desc = "Working with string."

	m.navItems = append(m.navItems, navItem{
		name: "Length",
		do: func() {
			showInputMenu(nilErrorWrapper(
				func(value string) string {
					return fmt.Sprint(len(value)) //calculate string length
				},
			))
		},
	})

	m.show()
}
