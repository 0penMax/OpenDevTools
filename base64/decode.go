package base64

import (
	"encoding/base64"
	"errors"
	"os"
	"strings"
)

func Decode(input string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

// DecodeImage decodes the Base64-encoded image from either a raw Data URI
// (e.g. "data:image/jpeg;base64,...") or an HTML <img> tag containing that Data URI,
// and writes the output to the given file path.
func DecodeImage(input string, outputFilePath string) error {
	// If input is an HTML tag, extract the src attribute value.
	if strings.Contains(input, "<img") {
		var src string
		if idx := strings.Index(input, `src="`); idx != -1 {
			start := idx + len(`src="`)
			end := strings.Index(input[start:], `"`)
			if end == -1 {
				return errors.New("invalid HTML tag: missing closing quote in src attribute")
			}
			src = input[start : start+end]
		} else if idx := strings.Index(input, `src='`); idx != -1 {
			start := idx + len(`src='`)
			end := strings.Index(input[start:], `'`)
			if end == -1 {
				return errors.New("invalid HTML tag: missing closing quote in src attribute")
			}
			src = input[start : start+end]
		} else {
			return errors.New("invalid HTML tag: missing src attribute")
		}
		input = src
	}

	// If input starts with "data:", then assume it is a Data URI.
	if strings.HasPrefix(input, "data:") {
		// A Data URI looks like: data:[<MIME-type>][;charset=<encoding>][;base64],<encoded-data>
		// We need to strip everything up to and including the comma.
		commaIndex := strings.Index(input, ",")
		if commaIndex == -1 {
			return errors.New("invalid data URI: missing comma separator")
		}
		input = input[commaIndex+1:]
	}

	// Decode the Base64 string.
	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return err
	}

	// Write the decoded data to the output file.
	if err = os.WriteFile(outputFilePath, data, 0644); err != nil {
		return err
	}

	return nil
}
