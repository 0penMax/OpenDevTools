package minify

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

func Minify(data string) (string, error) {
	m := minify.New()
	m.AddFunc("application/javascript", js.Minify)
	return m.String("application/javascript", data)
}
