package io

import (
	"openDevTools/menus/utils"

	"github.com/pterm/pterm"
	"golang.design/x/clipboard"
)

func processFromClipboard(processF func(string) (string, error)) (string, error) {
	utils.ClearScreen()
	pterm.Info.Println("Read from clipboard")
	clipboardText := clipboard.Read(clipboard.FmtText)

	return processF(string(clipboardText))
}
func ReadImgFromClipboard() ([]byte, error) {
	pterm.Info.Println("Read from clipboard")
	data := clipboard.Read(clipboard.FmtImage)

	return data, nil
}

func save2Clipboard(value []byte) {
	clipboard.Write(clipboard.FmtText, []byte(value))

	pterm.Info.Println("Result copied to clipboard")
}

func saveImg2Clipboard(value []byte) {
	clipboard.Write(clipboard.FmtImage, value)

	pterm.Info.Println("Img copied to clipboard")
}
