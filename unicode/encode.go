package unicode

import (
	"fmt"
	"strings"
)

// EncodeToUnicode takes a string and returns its Unicode escape sequence representation.
func EncodeToUnicode(input string) string {
	var result strings.Builder

	for _, r := range input {
		if r > 127 { // Only encode characters outside the ASCII range
			result.WriteString(fmt.Sprintf("\\u%04x", r))
		} else {
			result.WriteRune(r) // Write ASCII characters as they are
		}
	}

	return result.String()
}
