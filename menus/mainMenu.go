package menus

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"os"
	"os/exec"
	"runtime"
)

type Menu struct {
	title    string
	desc     string
	navItems []navItem
}

type navItem struct {
	name string
	do   func()
}

func (m *Menu) show() {
	if m.title != "" {
		pterm.DefaultHeader.WithFullWidth().Println(m.title)
		pterm.Println()
	}
	if m.desc != "" {
		pterm.Info.Println(m.desc)
	}
	if len(m.navItems) > 0 {
		navMap := make(map[string]func())
		var menu []string
		for _, n := range m.navItems {
			navMap[n.name] = n.do
			menu = append(menu, n.name)
		}

		// Use PTerm's interactive select feature to present the options to the user and capture their selection
		// The Show() method displays the options and waits for the user's input
		selectedOption, _ := pterm.DefaultInteractiveSelect.WithMaxHeight(10).WithOptions(menu).Show()

		f := navMap[selectedOption]
		f()
	}

}

func ShowMainMenu() {
	ClearScreen()
	pterm.Println()
	// Generate BigLetters and store in 's'
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("Open Dev Utils")).Srender()
	// Print the BigLetters 's' centered in the terminal
	pterm.DefaultCenter.Println(s)
	// Print a block of text centered in the terminal

	var m Menu

	m.navItems = append(m.navItems, navItem{
		name: "Hash Generator",
		do:   showHashMenu,
	})

	m.navItems = append(m.navItems, navItem{
		name: "Unixtime",
		do:   showUnixTimeMenu,
	})

	m.navItems = append(m.navItems, navItem{
		name: "Base64 string",
		do:   showBase64StringMenu,
	})

	m.navItems = append(m.navItems, navItem{
		name: "Base64 image",
		do:   showBase64ImgMenu,
	})

	m.navItems = append(m.navItems, navItem{
		name: "Color convertor",
		do:   showColorTransformMenu,
	})

	m.navItems = append(m.navItems, navItem{
		name: "lorem generator",
		do:   showLoremMenu,
	})

	m.show()
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
