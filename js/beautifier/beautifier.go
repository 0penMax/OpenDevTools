// code based on github.com/ditashi/jsbeautifier-go
// Big thanks author Ditashi Sayomi

package beautifier

import (
	"openDevTools/js/beautifier/jsbeautifier"
)

func Beautify(text string) (string, error) {
	return jsbeautifier.Beautify(&text, jsbeautifier.DefaultOptions())
}
