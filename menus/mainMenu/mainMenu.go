package mainMenu

import (
	"flag"
	"openDevTools/menus/base64"
	"openDevTools/menus/colors"
	"openDevTools/menus/hash"
	"openDevTools/menus/js"
	"openDevTools/menus/jwt"
	"openDevTools/menus/models"
	"openDevTools/menus/pgp"
	"openDevTools/menus/qr"
	stringsMenu "openDevTools/menus/strings"
	"openDevTools/menus/unixTime"
	"openDevTools/menus/utils"
	"strings"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

var mm models.Menu

func ProcessShortcuts() bool {

	// maps for lookup
	nameToHandler := make(map[string]func())
	flags := make(map[string]*bool)

	for _, it := range mm.NavItems {
		n := buildShortName(it.Name)
		nameToHandler[n] = it.Do
		flags[n] = flag.Bool(n, false, "Run "+it.Name+" menu")
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

	mm.NavItems = append(mm.NavItems, models.NavItem{
		Name:  "JS",
		Do:    js.Menu,
		MType: models.NavigationMenu,
	})

	mm.NavItems = append(mm.NavItems, models.NavItem{
		Name:  "Hash Generator",
		Do:    hash.Menu,
		MType: models.NavigationMenu,
	})

	mm.NavItems = append(mm.NavItems, models.NavItem{
		Name:  "Unixtime",
		Do:    unixTime.Menu,
		MType: models.NavigationMenu,
	})

	mm.NavItems = append(mm.NavItems, models.NavItem{
		Name:  "Strings",
		Do:    stringsMenu.Menu,
		MType: models.NavigationMenu,
	})

	mm.NavItems = append(mm.NavItems, models.NavItem{
		Name:  "Base64 image",
		Do:    base64.ImgMenu,
		MType: models.NavigationMenu,
	})

	mm.NavItems = append(mm.NavItems, models.NavItem{
		Name:  "Color convertor",
		Do:    colors.Menu,
		MType: models.NavigationMenu,
	})

	mm.NavItems = append(mm.NavItems, models.NavItem{
		Name:  "JWT",
		Do:    jwt.Menu,
		MType: models.NavigationMenu,
	})

	mm.NavItems = append(mm.NavItems, models.NavItem{
		Name:  "PGP",
		Do:    pgp.Menu,
		MType: models.NavigationMenu,
	})
	mm.NavItems = append(mm.NavItems, models.NavItem{
		Name:  "QR Codes",
		Do:    qr.Menu,
		MType: models.NavigationMenu,
	})

}

func ShowMainMenu() {
	utils.ClearScreen()
	pterm.Println()
	width := pterm.GetTerminalWidth()
	if width >= 115 {
		_ = pterm.DefaultBigText.WithLetters(putils.LettersFromString("OPEN DEV TOOLS")).Render()
	} else {
		_ = pterm.DefaultBigText.WithLetters(putils.LettersFromString("ODT")).Render()
	}

	mm.Show()
	showDoYouWant2Continue()
}

func showDoYouWant2Continue() {

	var m models.Menu

	m.Title = ""
	m.Desc = "Do you want to continue?"

	m.NavItems = append(m.NavItems, models.NavItem{
		Name:  "Last menu",
		Do:    models.LastNavMenu,
		MType: models.NavigationMenu,
	})

	m.NavItems = append(m.NavItems, models.NavItem{
		Name:  "Main menu",
		Do:    ShowMainMenu,
		MType: models.NavigationMenu,
	})

	m.NavItems = append(m.NavItems, models.NavItem{
		Name: "Exit",
		Do: func() {
			return
		},
	})

	m.Show()
}
