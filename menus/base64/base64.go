package base64

import (
	"openDevTools/base64"
	"openDevTools/menus/io"
	"openDevTools/menus/models"
	"openDevTools/menus/utils"
)

func StringMenu() {
	utils.ShowMenu(
		"Base64 string",
		"Encode/decode base64 string.",
		[]models.NavItem{
			{
				Name: "Encode",
				Do:   showBase64EncodeDialog,
			},
			{
				Name: "Decode",
				Do:   showBase64DecodeDialog,
			},
		},
	)

}

func showBase64EncodeDialog() {
	io.ShowInputMenu(io.NilErrorWrapper(base64.Encode))
}

func showBase64DecodeDialog() {
	io.ShowInputMenu(func(s string) (string, error) {
		b, err := base64.Decode(s)
		return string(b), err
	})
}
