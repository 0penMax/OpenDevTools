package numbers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Number struct {
	BaseName string // "binary","octal","hex","decimal"
	Base     int    // numeric base 2,8,16,10
	Repr     string // string representation in that base
}

// Parse parses a single non-negative integer string (no floats). Non-decimal bases must have prefixes:
// 0b... for binary, 0o... for octal, 0x... for hex. Plain digits parse as decimal.
// It returns a slice with four Number entries (binary, octal, hex, decimal) representing the same numeric value.
func Parse(value string) ([]Number, error) {
	if value == "" {
		return nil, errors.New("empty input")
	}
	s := strings.TrimSpace(value)
	if s == "" {
		return nil, errors.New("empty input")
	}

	low := strings.ToLower(s)

	var srcBase int
	var digits string

	switch {
	case strings.HasPrefix(low, "0b"):
		srcBase = 2
		digits = s[2:]
	case strings.HasPrefix(low, "0o"):
		srcBase = 8
		digits = s[2:]
	case strings.HasPrefix(low, "0x"):
		srcBase = 16
		digits = s[2:]
	default:
		srcBase = 10
		digits = s
	}

	if digits == "" {
		return nil, fmt.Errorf("no digits after prefix in %q", s)
	}
	if strings.ContainsAny(digits, ".,") {
		return nil, fmt.Errorf("floats not allowed: %q", s)
	}
	if strings.HasPrefix(digits, "+") || strings.HasPrefix(digits, "-") {
		return nil, fmt.Errorf("signed values not allowed: %q", s)
	}

	// Parse using the detected source base
	val, err := strconv.ParseUint(digits, srcBase, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %q as base %d: %w", s, srcBase, err)
	}

	// Build representations (without prefixes). Use lowercase hex letters.
	binRepr := "0b" + strconv.FormatUint(val, 2)
	octRepr := "0o" + strconv.FormatUint(val, 8)
	hexRepr := "0x" + strings.ToLower(strconv.FormatUint(val, 16))
	decRepr := strconv.FormatUint(val, 10)

	results := []Number{
		{BaseName: "binary", Base: 2, Repr: binRepr},
		{BaseName: "octal", Base: 8, Repr: octRepr},
		{BaseName: "hex", Base: 16, Repr: hexRepr},
		{BaseName: "decimal", Base: 10, Repr: decRepr},
	}

	return results, nil
}
