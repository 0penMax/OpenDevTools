package unicode

import "fmt"

// DecodeUnicode takes a string with Unicode escape sequences and returns the decoded string.
func DecodeUnicode(input string) (string, error) {
	var result []rune
	for i := 0; i < len(input); i++ {
		if input[i] == '\\' && i+1 < len(input) && input[i+1] == 'u' {
			// Extract the Unicode code point
			if i+6 <= len(input) {
				var codePoint int
				_, err := fmt.Sscanf(input[i+2:i+6], "%x", &codePoint)
				if err != nil {
					return "", err
				}
				result = append(result, rune(codePoint))
				i += 5 // Move past the Unicode escape sequence
			} else {
				return "", fmt.Errorf("incomplete Unicode escape sequence")
			}
		} else {
			result = append(result, rune(input[i]))
		}
	}
	return string(result), nil
}
