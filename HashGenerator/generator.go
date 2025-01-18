package HashGenerator

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"openDevTools/models"
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
