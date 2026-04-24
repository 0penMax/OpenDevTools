package utils

import (
	mModels "openDevTools/menus/models"
	"openDevTools/models"
	"os"
	"os/exec"
	"runtime"

	"github.com/pterm/pterm"
)

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

func ShowMenu(title, desc string, items []mModels.NavItem) {
	ClearScreen()
	var m mModels.Menu

	m.Title = title
	m.Desc = desc
	m.NavItems = items

	m.Show()
}

func ShowTable(tableHeader pterm.TableData, tableData pterm.TableData) {
	tableHeader = append(tableHeader, tableData...)
	err := pterm.DefaultTable.WithHasHeader().WithData(tableHeader).Render()
	if err != nil {
		pterm.Error.Println(err)
	}
}

func ParseResultItems4Table(tableData []models.ResultItem) pterm.TableData {
	var table pterm.TableData
	for _, r := range tableData {
		table = append(table, []string{r.Name, r.Value})
	}
	return table
}
