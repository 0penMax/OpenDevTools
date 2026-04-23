package models

import (
	"fmt"

	"github.com/pterm/pterm"
)

type Menu struct {
	Title    string
	Desc     string
	NavItems []NavItem
}

type NavItem struct {
	Name  string
	Do    func()
	MType MenuType
}

type MenuType string

const (
	NavigationMenu MenuType = "NavigationMenu"
	FunctionMenu   MenuType = "FunctionMenu"
)

var LastNavMenu func() = func() { fmt.Println("last function not found") }

func (m *Menu) Show() {
	if m.Title != "" {
		pterm.DefaultHeader.WithFullWidth().Println(m.Title)
		pterm.Println()
	}
	if m.Desc != "" {
		pterm.Info.Println(m.Desc)
	}
	if len(m.NavItems) > 0 {
		navMap := make(map[string]NavItem)
		var menu []string
		for _, n := range m.NavItems {
			navMap[n.Name] = n
			menu = append(menu, n.Name)
		}

		// Use PTerm's interactive select feature to present the options to the user and capture their selection
		// The Show() method displays the options and waits for the user's input
		selectedOption, _ := pterm.DefaultInteractiveSelect.WithMaxHeight(15).WithOptions(menu).Show()
		item := navMap[selectedOption]
		if item.MType == NavigationMenu {
			LastNavMenu = item.Do
		}

		item.Do()
	}

}
