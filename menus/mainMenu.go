package menus

import (
	"flag"
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
		selectedOption, _ := pterm.DefaultInteractiveSelect.WithMaxHeight(15).WithOptions(menu).Show()

		f := navMap[selectedOption]
		f()
	}

}

var mm Menu

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
		name: "JS",
		do:   showJsMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name: "Hash Generator",
		do:   showHashMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name: "Unixtime",
		do:   showUnixTimeMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name: "Strings",
		do:   showStringsMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name: "Base64 string",
		do:   showBase64StringMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name: "Base64 image",
		do:   showBase64ImgMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name: "Color convertor",
		do:   showColorTransformMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name: "JWT",
		do:   showJwtMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name: "PGP",
		do:   showPgpMenu,
	})

	mm.navItems = append(mm.navItems, navItem{
		name: "Unicode string",
		do:   showUnicodeDEMenu,
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
	// Create an interactive continue prompt with default settings
	// This will pause the program execution until the user presses enter
	// The message displayed is "Press 'Enter' to continue..."
	prompt := pterm.DefaultInteractiveContinue
	prompt.Options = []string{"yes", "no"}

	// Show the prompt and wait for user input
	// The returned result is the user's input (should be empty as it's a continue prompt)
	// The second return value is an error which is ignored here
	result, _ := prompt.Show()

	if result == "yes" {
		ShowMainMenu()
	} else {
		return
	}
}
