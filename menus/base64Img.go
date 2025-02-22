package menus

import (
	"github.com/pterm/pterm"
	"github.com/sqweek/dialog"
	"golang.design/x/clipboard"
	"openDevTools/base64"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

func showBase64ImgMenu() {
	ClearScreen()
	var m Menu

	m.title = "Base64 image"
	m.desc = "Encode/decode base64 image."

	m.navItems = append(m.navItems, navItem{
		name: "Encode to clipboard",
		do:   showBase64EncodeImgDialog,
	})

	m.navItems = append(m.navItems, navItem{
		name: "Encode to clipboard as html img tag",
		do:   showBase64EncodeImg2HtmlDialog,
	})

	m.navItems = append(m.navItems, navItem{
		name: "Decode from clipboard",
		do:   showBase64DecodeImgDialog,
	})

	m.show()
}

func showBase64EncodeImgDialog() {
	ClearScreen()

	filepath, err := dialog.File().Filter("Image", "bmp", "ico", "webp", "svg", "gif", "png", "jpeg", "jpg").Load()

	if err != nil {
		pterm.Warning.Println(err)
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
	ClearScreen()

	filepath, err := dialog.File().Filter("Image", "bmp", "ico", "webp", "svg", "gif", "png", "jpeg", "jpg").Load()

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
	ClearScreen()
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
