package colors

import (
	"strings"
	"testing"
)

func TestConvertColor(t *testing.T) {
	// Each test case specifies an input, whether an error is expected,
	// and the expected CSS colors name (if any).
	testCases := []struct {
		input        string
		expectedErr  bool
		expectedName string
	}{
		// Colors that do not have a corresponding CSS name in our mapping.
		{"#FF5733", false, ""},
		{"rgb(255,87,51)", false, ""},
		{"rgba(255,87,51,0.5)", false, ""},
		{"hsl(9,100%,60%)", false, ""},
		{"hsla(9,100%,60%,0.5)", false, ""},
		// A known CSS named colors (will return "Blue" as title-cased name).
		{"blue", false, "Blue"},
		// Test with a known colors using HEX notation that is mapped.
		{"#0000FF", false, "Blue"},
		// An invalid input should produce an error.
		{"invalid", true, ""},
		{"F53", false, ""},
		// Hex with alpha (semi-transparent) should not return a CSS name.
		{"#FF573380", false, ""},
		// Extra whitespace and mixed case for a named colors.
		{" Blue  ", false, "Blue"},
		// HSL that converts to blue.
		{"hsl(240,100%,50%)", false, "Blue"},
		// HSLA that converts to blue.
		{"hsla(240,100%,50%,1)", false, "Blue"},
		// RGB version of blue.
		{"rgb(0, 0, 255)", false, "Blue"},
		// RGBA with full opacity should return a name.
		{"rgba(0,0,255,1)", false, "Blue"},
		// RGBA with non-full opacity should not return a name.
		{"rgba(0,0,255,0.9)", false, ""},
		// Mixed-case named colors.
		{"ReD", false, "Red"},
		// A named colors with lower-case input.
		{"lightgreen", false, "Lightgreen"},
		// RGB with extra inner whitespace.
		{" rgb( 255 , 87 , 51 ) ", false, ""},
		// 6-digit hex without "#" should be recognized.
		{"123456", false, ""},
		// Uppercase function name for RGB.
		{"RGB(255,87,51)", false, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			results, err := ConvertColor(tc.input)
			if tc.expectedErr {
				if err == nil {
					t.Errorf("expected error for input %q but got none", tc.input)
				}
				// No further checking if error is expected.
				return
			}
			if err != nil {
				t.Errorf("unexpected error for input %q: %v", tc.input, err)
				return
			}

			// Verify that all required formats are present.
			requiredFormats := []string{"HEX", "RGB", "RGBA", "HSL", "HSLA"}
			for _, format := range requiredFormats {
				found := false
				for _, item := range results {
					if item.Name == format {
						found = true
						// Also check that the output seems valid (e.g. HEX starts with "#")
						switch format {
						case "HEX":
							if !strings.HasPrefix(item.Value, "#") {
								t.Errorf("expected HEX value to start with '#', got %q", item.Value)
							}
						case "RGB":
							if !strings.HasPrefix(item.Value, "rgb(") {
								t.Errorf("expected RGB value to start with 'rgb(', got %q", item.Value)
							}
						case "RGBA":
							if !strings.HasPrefix(item.Value, "rgba(") {
								t.Errorf("expected RGBA value to start with 'rgba(', got %q", item.Value)
							}
						case "HSL":
							if !strings.HasPrefix(item.Value, "hsl(") {
								t.Errorf("expected HSL value to start with 'hsl(', got %q", item.Value)
							}
						case "HSLA":
							if !strings.HasPrefix(item.Value, "hsla(") {
								t.Errorf("expected HSLA value to start with 'hsla(', got %q", item.Value)
							}
						}
					}
				}
				if !found {
					t.Errorf("expected format %q in results for input %q", format, tc.input)
				}
			}

			// Verify the "Name" result (if any).
			var nameFound string
			for _, item := range results {
				if item.Name == "Name" {
					nameFound = item.Value
				}
			}
			if nameFound != tc.expectedName {
				t.Errorf("expected Name %q, got %q for input %q", tc.expectedName, nameFound, tc.input)
			}
		})
	}
}
