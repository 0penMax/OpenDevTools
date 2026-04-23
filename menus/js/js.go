package js

import (
	jsb "openDevTools/js/beautifier"
	jsm "openDevTools/js/minify"
	"openDevTools/menus/io"
	"openDevTools/menus/models"
	"openDevTools/menus/utils"
)

func Menu() {
	utils.ShowMenu(
		"JS", "",
		[]models.NavItem{
			{Name: "Beautifier", Do: showJsBeautifier},
			{Name: "Minify", Do: showJsMinify},
		},
	)

}

func showJsBeautifier() {
	io.ShowInputMenu(jsb.Beautify)
}

func showJsMinify() {
	io.ShowInputMenu(jsm.Minify)
}
