package menus

import (
	"github.com/pterm/pterm"
)

func showInputMenu(processF func(string) (string, error)) {
	ClearScreen()

	pterm.DefaultHeader.WithFullWidth().Println("Select input method")
	pterm.Println()

	navMap := make(map[string]func(processF func(string) (string, error)) (string, error))
	var menu []string

	mi := "Multiline input"
	navMap[mi] = showMultilineInput
	menu = append(menu, mi)

	clp := "Clipboard"
	navMap[clp] = processFromClipboard
	menu = append(menu, clp)

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithMaxHeight(10).WithOptions(menu).Show()

	f := navMap[selectedOption]
	r, err := f(processF)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	showOutputMenu(r)

}

func nilErrorWrapper(processF func(string) string) func(string) (string, error) {
	return func(v string) (string, error) { return processF(v), nil }
}

func showMultilineInput(processF func(string) (string, error)) (string, error) {
	ClearScreen()
	// Create a default interactive text input with multi-line enabled.
	// This allows the user to input multiple lines of text.
	textInput := pterm.DefaultInteractiveTextInput.WithMultiLine()

	// Show the text input to the user and store the result.
	// The second return value (an error) is ignored with '_'.
	text, _ := textInput.Show()

	// Print a blank line for better readability in the output.
	pterm.Println()

	return processF(text)

}

func showResult(value string) {
	pterm.Info.Println("Result:")
	pterm.Println(value)
}

func showOutputMenu(value string) {
	ClearScreen()
	pterm.DefaultHeader.WithFullWidth().Println("Select output method")
	pterm.Println()
	navMap := make(map[string]func(value string))
	var menu []string

	os := "On screen"
	navMap[os] = showResult
	menu = append(menu, os)

	clp := "Save 2 clipboard"
	navMap[clp] = save2Clipboard
	menu = append(menu, clp)

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithMaxHeight(10).WithOptions(menu).Show()

	f := navMap[selectedOption]
	ClearScreen()
	f(value)
}
