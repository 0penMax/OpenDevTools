package menus

import (
	"fmt"
	"openDevTools/lorem"
	"strconv"

	"github.com/pterm/pterm"
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
		mType: navigationMenu,
	})

	m.navItems = append(m.navItems, navItem{
		name:  "Lorem generator",
		do:    showLorem,
		mType: navigationMenu,
	})

	m.navItems = append(m.navItems, navItem{
		name:  "Unicode string",
		do:    showUnicodeDEMenu,
		mType: navigationMenu,
	})

	m.navItems = append(m.navItems, navItem{
		name:  "Base64 string",
		do:    showBase64StringMenu,
		mType: navigationMenu,
	})

	m.show()
}

func showLorem() {
	ClearScreen()

	textInput := pterm.DefaultInteractiveTextInput.WithDefaultText("enter word count")
	text, _ := textInput.Show()

	count, err := strconv.Atoi(text)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	showOutputMenu(lorem.Generate(count))
}
