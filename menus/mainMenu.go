package menus

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type Menu struct {
	title    string
	desc     string
	navItems []navItem
}

type navItem struct {
	name  string
	do    func()
	mType menuType
}

type menuType string

const (
	navigationMenu menuType = "navigationMenu"
	functionMenu   menuType = "functionMenu"
)

func (m *Menu) show() {
	if m.title != "" {
		pterm.DefaultHeader.WithFullWidth().Println(m.title)
		pterm.Println()
	}
	if m.desc != "" {
		pterm.Info.Println(m.desc)
	}
	if len(m.navItems) > 0 {
		navMap := make(map[string]navItem)
		var menu []string
		for _, n := range m.navItems {
			navMap[n.name] = n
			menu = append(menu, n.name)
		}

		// Use PTerm's interactive select feature to present the options to the user and capture their selection
		// The Show() method displays the options and waits for the user's input
		selectedOption, _ := pterm.DefaultInteractiveSelect.WithMaxHeight(15).WithOptions(menu).Show()
		item := navMap[selectedOption]
		if item.mType == navigationMenu {
			lastNavMenu = item.do
		}

		item.do()
	}

}

var mm Menu

var lastNavMenu func() = func() { fmt.Println("last function not found") }

func ProcessShortcuts() bool {

	// maps for lookup
	nameToHandler := make(map[string]func())
	flags := make(map[string]*bool)

	for _, it := range mm.navItems {
		n := buildShortName(it.name)
		nameToHandler[n] = it.do
		flags[n] = flag.Bool(n, false, "Run "+it.name+" menu")
	}

	flag.Parse()

	// search for fist lags
	for n, ptr := range flags {
		if ptr != nil && *ptr {
			nameToHandler[n]()
			return true
		}
	}
	return false
}

func buildShortName(fullName string) string {
	fullName = strings.ToLower(fullName)
	if len(fullName) < 4 {
		return fullName
	}
	words := strings.Split(fullName, " ")
	name := ""
	for _, word := range words {
		name += string(word[0])
	}
	return name
}

func BuildMenu() {

	mm.navItems = append(mm.navItems, navItem{
		name:  "JS",
		do:    showJsMenu,
		mType: navigationMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name:  "Hash Generator",
		do:    showHashMenu,
		mType: navigationMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name:  "Unixtime",
		do:    showUnixTimeMenu,
		mType: navigationMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name:  "Strings",
		do:    showStringsMenu,
		mType: navigationMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name:  "Base64 image",
		do:    showBase64ImgMenu,
		mType: navigationMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name:  "Color convertor",
		do:    showColorTransformMenu,
		mType: navigationMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name:  "JWT",
		do:    showJwtMenu,
		mType: navigationMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name:  "PGP",
		do:    showPgpMenu,
		mType: navigationMenu,
	})
	mm.navItems = append(mm.navItems, navItem{
		name:  "QR Codes",
		do:    showQRMenu,
		mType: navigationMenu,
	})

}

func ShowMainMenu() {
	ClearScreen()
	pterm.Println()
	width := pterm.GetTerminalWidth()
	if width >= 115 {
		_ = pterm.DefaultBigText.WithLetters(putils.LettersFromString("OPEN DEV TOOLS")).Render()
	} else {
		_ = pterm.DefaultBigText.WithLetters(putils.LettersFromString("ODT")).Render()
	}

	mm.show()
	showDoYouWant2Continue()
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

func showDoYouWant2Continue() {

	var m Menu

	m.title = ""
	m.desc = "Do you want to continue?"

	m.navItems = append(m.navItems, navItem{
		name:  "Last menu",
		do:    lastNavMenu,
		mType: navigationMenu,
	})

	m.navItems = append(m.navItems, navItem{
		name:  "Main menu",
		do:    ShowMainMenu,
		mType: navigationMenu,
	})

	m.navItems = append(m.navItems, navItem{
		name: "Exit",
		do: func() {
			return
		},
	})

	m.show()
}
