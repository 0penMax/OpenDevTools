package main

import (
	"fmt"
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
	}
	if m.desc != "" {
		pterm.Info.Println(m.desc)
	}
	navMap := make(map[string]func())
	menu := make([]string, len(m.navItems))
	for _, n := range m.navItems {
		navMap[n.name] = n.do
		menu = append(menu, n.name)
	}

	// Use PTerm's interactive select feature to present the options to the user and capture their selection
	// The Show() method displays the options and waits for the user's input
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(menu).Show()

	f := navMap[selectedOption]
	f()
}

func showMainMenu() {
	// Generate BigLetters and store in 's'
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("Open Dev Utils")).Srender()
	// Print the BigLetters 's' centered in the terminal
	pterm.DefaultCenter.Println(s)
	// Print a block of text centered in the terminal

	var m Menu

	m.navItems = append(m.navItems, navItem{
		name: "Hash Generator",
		do:   func() { fmt.Println("hello Hash Generator") },
	})

	m.navItems = append(m.navItems, navItem{
		name: "Unixtime",
		do:   func() { fmt.Println("hello Unixtime") },
	})

	m.show()
}
