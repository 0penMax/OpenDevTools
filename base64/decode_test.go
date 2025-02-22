package base64

import (
	"bytes"
	"os"
	"testing"
)

func TestDecodeImage(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []byte
		wantErr  bool
	}{
		{
			name:     "RawDataURI",
			input:    "data:image/jpeg;base64,SGVsbG8sIHdvcmxkIQ==",
			expected: []byte("Hello, world!"),
			wantErr:  false,
		},
		{
			name:     "HTMLDoubleQuotes",
			input:    `<img src="data:image/jpeg;base64,SGVsbG8sIHdvcmxkIQ==" alt="test">`,
			expected: []byte("Hello, world!"),
			wantErr:  false,
		},
		{
			name:     "HTMLSingleQuotes",
			input:    `<img src='data:image/jpeg;base64,SGVsbG8sIHdvcmxkIQ==' alt="test">`,
			expected: []byte("Hello, world!"),
			wantErr:  false,
		},
		{
			name:    "MissingSrc",
			input:   `<img alt="test">`,
			wantErr: true,
		},
		{
			name:    "MissingComma",
			input:   "data:image/jpeg;base64SGVsbG8sIHdvcmxkIQ==",
			wantErr: true,
		},
		{
			name:    "InvalidBase64",
			input:   "data:image/jpeg;base64,invalid-base64-string",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a temporary directory and file for the output.
			tempDir := t.TempDir()
			outputPath := tempDir + "/output.bin"

			err := DecodeImage(tc.input, outputPath)
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error: %v, got: %v", tc.wantErr, err)
			}

			// If no error was expected, read and compare the output.
			if !tc.wantErr {
				data, err := os.ReadFile(outputPath)
				if err != nil {
					t.Fatalf("failed to read output file: %v", err)
				}
				if !bytes.Equal(data, tc.expected) {
					t.Errorf("output mismatch: expected %q, got %q", tc.expected, data)
				}
			}
		})
	}
}
