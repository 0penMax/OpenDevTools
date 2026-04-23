package unicode

import (
	"openDevTools/menus/io"
	"openDevTools/menus/models"
	"openDevTools/menus/utils"
	"openDevTools/unicode"
)

func Menu() {

	utils.ShowMenu(
		"Unicode Decode/encode",
		"Encode/decode unicode string.",
		[]models.NavItem{
			{Name: "Encode", Do: encodeDialog},
			{Name: "Decode", Do: decodeDialog},
		})
}

func decodeDialog() {
	io.ShowInputMenu(unicode.DecodeUnicode)
}

func encodeDialog() {
	io.ShowInputMenu(io.NilErrorWrapper(unicode.EncodeToUnicode))
}
