package base64

import (
	"openDevTools/base64"
	"openDevTools/menus/io"
	"openDevTools/menus/models"
	"openDevTools/menus/utils"

	"github.com/pterm/pterm"
	"github.com/sqweek/dialog"
	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

func ImgMenu() {
	utils.ShowMenu(
		"Base64 image",
		"Encode/decode base64 image.",
		[]models.NavItem{
			{Name: "Encode to clipboard", Do: showBase64EncodeImgDialog},
			{Name: "Encode to clipboard as html img tag", Do: showBase64EncodeImg2HtmlDialog},
			{Name: "Decode to clipboard", Do: showBase64DecodeImgDialog},
		})
}

func showBase64EncodeImgDialog() {
	utils.ClearScreen()

	filepath, ok := io.OpenFileDialog([]string{"*.bmp", "*.ico", "*.webp", "*.svg", "*.gif", "*.png", "*.png", "*.jpeg", "*.jpg"})

	if !ok {
		pterm.Warning.Println("openFileDialog cancelled")
		return
	}
	pterm.Info.Println("filepath:", filepath)

	result, err := base64.EncodeImage(filepath)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	clipboard.Write(clipboard.FmtText, []byte(result))

	pterm.Println("copied to clipboard")

}

func showBase64EncodeImg2HtmlDialog() {
	utils.ClearScreen()

	filepath, err := dialog.File().Load()

	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	pterm.Info.Println("filepath:", filepath)

	result, err := base64.EncodeImage2HTML(filepath)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	clipboard.Write(clipboard.FmtText, []byte(result))

	pterm.Println("copied to clipboard")

}

func showBase64DecodeImgDialog() {
	utils.ClearScreen()
	filepath, err := dialog.File().Save()

	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	pterm.Info.Println("filepath:", filepath)

	clipboardText := clipboard.Read(clipboard.FmtText)

	err = base64.DecodeImage(string(clipboardText), filepath)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	pterm.Println("decoded")

}
