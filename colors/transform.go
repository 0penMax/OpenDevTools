package colors

import (
	"fmt"
	"math"
	"openDevTools/models"
	"regexp"
	"strconv"
	"strings"
)

// Our internal color type stores R, G, B values (0–255) and an Alpha channel.
type color struct {
	R, G, B int
	A       float64
}

// --- Parsing Functions ---

// parseHexColor parses hex strings such as "#FF5733", "#F53" or "#FF573380" (with alpha).
func parseHexColor(s string) (*color, error) {
	s = strings.TrimPrefix(s, "#")
	var r, g, b, a int64
	a = 255 // default alpha (opaque)

	switch len(s) {
	case 3: // shorthand format: "F53" → "FF", "55", "33"
		r, _ = strconv.ParseInt(strings.Repeat(string(s[0]), 2), 16, 0)
		g, _ = strconv.ParseInt(strings.Repeat(string(s[1]), 2), 16, 0)
		b, _ = strconv.ParseInt(strings.Repeat(string(s[2]), 2), 16, 0)
	case 6:
		var err error
		r, err = strconv.ParseInt(s[0:2], 16, 0)
		if err != nil {
			return nil, err
		}
		g, err = strconv.ParseInt(s[2:4], 16, 0)
		if err != nil {
			return nil, err
		}
		b, err = strconv.ParseInt(s[4:6], 16, 0)
		if err != nil {
			return nil, err
		}
	case 8: // with alpha: "RRGGBBAA"
		var err error
		r, err = strconv.ParseInt(s[0:2], 16, 0)
		if err != nil {
			return nil, err
		}
		g, err = strconv.ParseInt(s[2:4], 16, 0)
		if err != nil {
			return nil, err
		}
		b, err = strconv.ParseInt(s[4:6], 16, 0)
		if err != nil {
			return nil, err
		}
		a, err = strconv.ParseInt(s[6:8], 16, 0)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid hex format")
	}

	return &color{R: int(r), G: int(g), B: int(b), A: float64(a) / 255.0}, nil
}

// parseRGB parses strings like "rgb(255, 87, 51)".
func parseRGB(s string) (*color, error) {
	re := regexp.MustCompile(`(?i)^rgb\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*\)$`)
	matches := re.FindStringSubmatch(s)
	if len(matches) != 4 {
		return nil, fmt.Errorf("invalid rgb format")
	}
	r, _ := strconv.Atoi(matches[1])
	g, _ := strconv.Atoi(matches[2])
	b, _ := strconv.Atoi(matches[3])
	return &color{R: r, G: g, B: b, A: 1.0}, nil
}

// parseRGBA parses strings like "rgba(255, 87, 51, 0.5)".
func parseRGBA(s string) (*color, error) {
	re := regexp.MustCompile(`(?i)^rgba\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*,\s*([01]?\.?\d+)\s*\)$`)
	matches := re.FindStringSubmatch(s)
	if len(matches) != 5 {
		return nil, fmt.Errorf("invalid rgba format")
	}
	r, _ := strconv.Atoi(matches[1])
	g, _ := strconv.Atoi(matches[2])
	b, _ := strconv.Atoi(matches[3])
	a, err := strconv.ParseFloat(matches[4], 64)
	if err != nil {
		return nil, err
	}
	return &color{R: r, G: g, B: b, A: a}, nil
}

// parseHSL parses strings like "hsl(9, 100%, 60%)".
// It converts the HSL values into an RGB representation.
func parseHSL(s string) (*color, error) {
	re := regexp.MustCompile(`(?i)^hsl\(\s*([\d.]+)\s*,\s*(\d+)%\s*,\s*(\d+)%\s*\)$`)
	matches := re.FindStringSubmatch(s)
	if len(matches) != 4 {
		return nil, fmt.Errorf("invalid hsl format")
	}
	hue, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return nil, err
	}
	sat, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}
	light, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, err
	}
	sVal := float64(sat) / 100.0
	lVal := float64(light) / 100.0
	r, g, b := hslToRGB(hue, sVal, lVal)
	return &color{R: r, G: g, B: b, A: 1.0}, nil
}

// parseHSLA parses strings like "hsla(9, 100%, 60%, 0.5)".
func parseHSLA(s string) (*color, error) {
	re := regexp.MustCompile(`(?i)^hsla\(\s*([\d.]+)\s*,\s*(\d+)%\s*,\s*(\d+)%\s*,\s*([01]?\.?\d+)\s*\)$`)
	matches := re.FindStringSubmatch(s)
	if len(matches) != 5 {
		return nil, fmt.Errorf("invalid hsla format")
	}
	hue, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return nil, err
	}
	sat, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}
	light, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, err
	}
	a, err := strconv.ParseFloat(matches[4], 64)
	if err != nil {
		return nil, err
	}
	sVal := float64(sat) / 100.0
	lVal := float64(light) / 100.0
	r, g, b := hslToRGB(hue, sVal, lVal)
	return &color{R: r, G: g, B: b, A: a}, nil
}

// --- HSL to RGB Conversion ---

// hue2rgb is a helper for converting HSL to RGB.
func hue2rgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

// hslToRGB converts HSL values (hue in degrees, saturation and lightness as fractions)
// into RGB values (each 0–255).
func hslToRGB(h, s, l float64) (int, int, int) {
	// Convert hue from degrees to a fraction
	h = h / 360.0
	var r, g, b float64
	if s == 0 {
		r, g, b = l, l, l // achromatic (gray)
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q
		r = hue2rgb(p, q, h+1.0/3.0)
		g = hue2rgb(p, q, h)
		b = hue2rgb(p, q, h-1.0/3.0)
	}
	return int(math.Round(r * 255)), int(math.Round(g * 255)), int(math.Round(b * 255))
}

// --- Conversion Methods ---

// ToHex returns the hex string. If fully opaque, returns "#RRGGBB"; otherwise, "#RRGGBBAA".
func (c color) ToHex() string {
	if c.A >= 1.0 {
		return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
	}
	return fmt.Sprintf("#%02X%02X%02X%02X", c.R, c.G, c.B, int(c.A*255))
}

// ToRGB returns the RGB string.
func (c color) ToRGB() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
}

// ToRGBA returns the RGBA string.
func (c color) ToRGBA() string {
	return fmt.Sprintf("rgba(%d, %d, %d, %.2f)", c.R, c.G, c.B, c.A)
}

// ToHSL returns the HSL string.
func (c color) ToHSL() string {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0
	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	l := (max + min) / 2
	var h, s float64
	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}
		switch max {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h /= 6
	}
	return fmt.Sprintf("hsl(%d, %d%%, %d%%)", int(h*360), int(s*100), int(l*100))
}

// ToHSLA returns the HSLA string.
func (c color) ToHSLA() string {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0
	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	l := (max + min) / 2
	var h, s float64
	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}
		switch max {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h /= 6
	}
	return fmt.Sprintf("hsla(%d, %d%%, %d%%, %.2f)", int(h*360), int(s*100), int(l*100), c.A)
}

// --- Named Colors ---

// A simple mapping for some CSS named colors.
// Expanded CSS named colors mapping.
var cssNamedColors = map[string]string{
	"aliceblue":            "#F0F8FF",
	"antiquewhite":         "#FAEBD7",
	"aqua":                 "#00FFFF",
	"aquamarine":           "#7FFFD4",
	"azure":                "#F0FFFF",
	"beige":                "#F5F5DC",
	"bisque":               "#FFE4C4",
	"black":                "#000000",
	"blanchedalmond":       "#FFEBCD",
	"blue":                 "#0000FF",
	"blueviolet":           "#8A2BE2",
	"brown":                "#A52A2A",
	"burlywood":            "#DEB887",
	"cadetblue":            "#5F9EA0",
	"chartreuse":           "#7FFF00",
	"chocolate":            "#D2691E",
	"coral":                "#FF7F50",
	"cornflowerblue":       "#6495ED",
	"cornsilk":             "#FFF8DC",
	"crimson":              "#DC143C",
	"cyan":                 "#00FFFF",
	"darkblue":             "#00008B",
	"darkcyan":             "#008B8B",
	"darkgoldenrod":        "#B8860B",
	"darkgray":             "#A9A9A9",
	"darkgreen":            "#006400",
	"darkkhaki":            "#BDB76B",
	"darkmagenta":          "#8B008B",
	"darkolivegreen":       "#556B2F",
	"darkorange":           "#FF8C00",
	"darkorchid":           "#9932CC",
	"darkred":              "#8B0000",
	"darksalmon":           "#E9967A",
	"darkseagreen":         "#8FBC8F",
	"darkslateblue":        "#483D8B",
	"darkslategray":        "#2F4F4F",
	"darkturquoise":        "#00CED1",
	"darkviolet":           "#9400D3",
	"deeppink":             "#FF1493",
	"deepskyblue":          "#00BFFF",
	"dimgray":              "#696969",
	"dodgerblue":           "#1E90FF",
	"firebrick":            "#B22222",
	"floralwhite":          "#FFFAF0",
	"forestgreen":          "#228B22",
	"fuchsia":              "#FF00FF",
	"gainsboro":            "#DCDCDC",
	"ghostwhite":           "#F8F8FF",
	"gold":                 "#FFD700",
	"goldenrod":            "#DAA520",
	"gray":                 "#808080",
	"green":                "#008000",
	"greenyellow":          "#ADFF2F",
	"honeydew":             "#F0FFF0",
	"hotpink":              "#FF69B4",
	"indianred":            "#CD5C5C",
	"indigo":               "#4B0082",
	"ivory":                "#FFFFF0",
	"khaki":                "#F0E68C",
	"lavender":             "#E6E6FA",
	"lavenderblush":        "#FFF0F5",
	"lawngreen":            "#7CFC00",
	"lemonchiffon":         "#FFFACD",
	"lightblue":            "#ADD8E6",
	"lightcoral":           "#F08080",
	"lightcyan":            "#E0FFFF",
	"lightgoldenrodyellow": "#FAFAD2",
	"lightgrey":            "#D3D3D3",
	"lightgreen":           "#90EE90",
	"lightpink":            "#FFB6C1",
	"lightsalmon":          "#FFA07A",
	"lightseagreen":        "#20B2AA",
	"lightskyblue":         "#87CEFA",
	"lightslategray":       "#778899",
	"lightsteelblue":       "#B0C4DE",
	"lightyellow":          "#FFFFE0",
	"lime":                 "#00FF00",
	"limegreen":            "#32CD32",
	"linen":                "#FAF0E6",
	"magenta":              "#FF00FF",
	"maroon":               "#800000",
	"mediumaquamarine":     "#66CDAA",
	"mediumblue":           "#0000CD",
	"mediumorchid":         "#BA55D3",
	"mediumpurple":         "#9370D8",
	"mediumseagreen":       "#3CB371",
	"mediumslateblue":      "#7B68EE",
	"mediumspringgreen":    "#00FA9A",
	"mediumturquoise":      "#48D1CC",
	"mediumvioletred":      "#C71585",
	"midnightblue":         "#191970",
	"mintcream":            "#F5FFFA",
	"mistyrose":            "#FFE4E1",
	"moccasin":             "#FFE4B5",
	"navajowhite":          "#FFDEAD",
	"navy":                 "#000080",
	"oldlace":              "#FDF5E6",
	"olive":                "#808000",
	"olivedrab":            "#6B8E23",
	"orange":               "#FFA500",
	"orangered":            "#FF4500",
	"orchid":               "#DA70D6",
	"palegoldenrod":        "#EEE8AA",
	"palegreen":            "#98FB98",
	"paleturquoise":        "#AFEEEE",
	"palevioletred":        "#D87093",
	"papayawhip":           "#FFEFD5",
	"peachpuff":            "#FFDAB9",
	"peru":                 "#CD853F",
	"pink":                 "#FFC0CB",
	"plum":                 "#DDA0DD",
	"powderblue":           "#B0E0E6",
	"purple":               "#800080",
	"red":                  "#FF0000",
	"rosybrown":            "#BC8F8F",
	"royalblue":            "#4169E1",
	"saddlebrown":          "#8B4513",
	"salmon":               "#FA8072",
	"sandybrown":           "#F4A460",
	"seagreen":             "#2E8B57",
	"seashell":             "#FFF5EE",
	"sienna":               "#A0522D",
	"silver":               "#C0C0C0",
	"skyblue":              "#87CEEB",
	"slateblue":            "#6A5ACD",
	"slategray":            "#708090",
	"snow":                 "#FFFAFA",
	"springgreen":          "#00FF7F",
	"steelblue":            "#4682B4",
	"tan":                  "#D2B48C",
	"teal":                 "#008080",
	"thistle":              "#D8BFD8",
	"tomato":               "#FF6347",
	"turquoise":            "#40E0D0",
	"violet":               "#EE82EE",
	"wheat":                "#F5DEB3",
	"white":                "#FFFFFF",
	"whitesmoke":           "#F5F5F5",
	"yellow":               "#FFFF00",
	"yellowgreen":          "#9ACD32",
}

// getColorName returns the proper CSS color name if the color's hex value (fully opaque)
// matches one in the cssNamedColors mapping. Otherwise, it returns an empty string.
func getColorName(c *color) string {
	if c.A < 1.0 {
		return ""
	}
	// Normalize the hex code to lower-case.
	hex := strings.ToLower(c.ToHex())
	for name, hexVal := range cssNamedColors {
		if strings.ToLower(hexVal) == hex {
			// Optionally title-case the name before returning.
			return strings.Title(name)
		}
	}
	return ""
}

// --- Updated ConvertColor Function ---

// ConvertColor takes an input colors string (HEX, RGB, RGBA, HSL, HSLA, or a CSS named colors),
// detects its format, and returns the colors converted to all supported formats.
func ConvertColor(input string) ([]models.ResultItem, error) {
	trimmed := strings.TrimSpace(input)
	lower := strings.ToLower(trimmed)
	var col *color
	var err error

	// Check for HEX (with or without "#")
	hexRegex := regexp.MustCompile(`^(#)?([0-9a-fA-F]{3}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$`)
	if strings.HasPrefix(lower, "#") || hexRegex.MatchString(trimmed) {
		col, err = parseHexColor(trimmed)
		if err != nil {
			return nil, err
		}
	} else if strings.HasPrefix(lower, "rgb(") {
		col, err = parseRGB(trimmed)
		if err != nil {
			return nil, err
		}
	} else if strings.HasPrefix(lower, "rgba(") {
		col, err = parseRGBA(trimmed)
		if err != nil {
			return nil, err
		}
	} else if strings.HasPrefix(lower, "hsl(") {
		col, err = parseHSL(trimmed)
		if err != nil {
			return nil, err
		}
	} else if strings.HasPrefix(lower, "hsla(") {
		col, err = parseHSLA(trimmed)
		if err != nil {
			return nil, err
		}
	} else if val, ok := cssNamedColors[lower]; ok {
		// If the input is a CSS named colors, look it up and parse its hex value.
		col, err = parseHexColor(val)
		if err != nil {
			return nil, err
		}
	}

	if col != nil {
		results := []models.ResultItem{
			{"HEX", col.ToHex()},
			{"RGB", col.ToRGB()},
			{"RGBA", col.ToRGBA()},
			{"HSL", col.ToHSL()},
			{"HSLA", col.ToHSLA()},
		}
		// Append the colors name if one exists.
		if name := getColorName(col); name != "" {
			results = append(results, models.ResultItem{"Name", name})
		}
		return results, nil
	}

	return nil, fmt.Errorf("unrecognized colors format")
}
