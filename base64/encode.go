package base64

import (
	"encoding/base64"
	"fmt"
	"mime"
	"os"
	"path/filepath"
)

func Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// EncodeImage reads the image from the given file path and returns its Base64 encoded string.
func EncodeImage(filePath string) (string, error) {
	// Read the image file into a byte slice.
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Encode the byte slice to a Base64 string.
	base64String := base64.StdEncoding.EncodeToString(data)
	return base64String, nil
}

// getMimeType returns the MIME type based on the file extension in the given path.
// If the MIME type cannot be determined, it defaults to "application/octet-stream".
func getMimeType(filePath string) string {
	ext := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	return mimeType
}

// EncodeImage2HTML generates an HTML <img> tag string with the Base64-encoded image.
// It automatically detects the MIME type based on the file extension.
func EncodeImage2HTML(filePath string) (string, error) {
	// Encode the image to a Base64 string using the provided function.
	encoded, err := EncodeImage(filePath)
	if err != nil {
		return "", err
	}

	// Get the MIME type for the file.
	mimeType := getMimeType(filePath)

	// Format and return the HTML <img> tag with a Data URI.
	htmlTag := fmt.Sprintf(`<img src="data:%s;base64,%s" alt="Embedded Image">`, mimeType, encoded)
	return htmlTag, nil
}
