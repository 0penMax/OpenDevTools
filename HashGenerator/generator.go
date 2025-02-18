package HashGenerator

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"openDevTools/models"
	"os"
)

func FromString(input string) []models.ResultItem {
	// Generate MD5 hash
	md5Hash := md5.Sum([]byte(input))

	// Generate SHA-1 hash
	sha1Hash := sha1.Sum([]byte(input))

	// Generate SHA-256 hash
	sha256Hash := sha256.Sum256([]byte(input))

	// Generate SHA-384 hash
	sha384Hash := sha512.Sum384([]byte(input))

	// Generate SHA-512 hash
	sha512Hash := sha512.Sum512([]byte(input))

	// Create results as a slice of ResultItem
	results := []models.ResultItem{
		{Name: "MD5", Value: hex.EncodeToString(md5Hash[:])},
		{Name: "SHA-1", Value: hex.EncodeToString(sha1Hash[:])},
		{Name: "SHA-256", Value: hex.EncodeToString(sha256Hash[:])},
		{Name: "SHA-384", Value: hex.EncodeToString(sha384Hash[:])},
		{Name: "SHA-512", Value: hex.EncodeToString(sha512Hash[:])},
	}

	return results
}

func FromFile(filePath string) ([]models.ResultItem, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create hash objects
	md5Hash := md5.New()
	sha1Hash := sha1.New()
	sha256Hash := sha256.New()
	sha384Hash := sha512.New384()
	sha512Hash := sha512.New()

	// Create a multi-writer to write the file content to all hash objects
	writers := io.MultiWriter(md5Hash, sha1Hash, sha256Hash, sha384Hash, sha512Hash)

	// Copy the file content to the writers
	if _, err := io.Copy(writers, file); err != nil {
		return nil, fmt.Errorf("failed to compute hashes: %w", err)
	}

	// Create results as a slice of ResultItem
	results := []models.ResultItem{
		{Name: "MD5", Value: hex.EncodeToString(md5Hash.Sum(nil))},
		{Name: "SHA-1", Value: hex.EncodeToString(sha1Hash.Sum(nil))},
		{Name: "SHA-256", Value: hex.EncodeToString(sha256Hash.Sum(nil))},
		{Name: "SHA-384", Value: hex.EncodeToString(sha384Hash.Sum(nil))},
		{Name: "SHA-512", Value: hex.EncodeToString(sha512Hash.Sum(nil))},
	}

	return results, nil
}
