package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"openDevTools/base64"
	"openDevTools/models"
	"strings"
)

type Token struct {
	Header    []models.ResultItem
	Payload   []models.ResultItem
	Signature string
}

func Read(jwt string) (Token, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return Token{}, errors.New("jwt format error")
	}
	header64, payload64, signature := parts[0], parts[1], parts[2]

	header, err := base64.Decode(header64)
	if err != nil {
		return Token{}, err
	}
	headerItems, err := readJson(header)
	if err != nil {
		return Token{}, err
	}

	payload, err := base64.Decode(payload64)
	if err != nil {
		return Token{}, err
	}
	payloadItems, err := readJson(payload)
	if err != nil {
		return Token{}, err
	}

	return Token{
		Header:    headerItems,
		Payload:   payloadItems,
		Signature: signature,
	}, nil

}

func readJson(jsonData []byte) ([]models.ResultItem, error) {
	// Unmarshal into a temporary map
	var temp map[string]interface{}
	if err := json.Unmarshal(jsonData, &temp); err != nil {
		return nil, err
	}

	// Convert every value to a string using fmt.Sprintf
	var result []models.ResultItem
	for key, value := range temp {
		result = append(result, models.ResultItem{
			Name:  key,
			Value: fmt.Sprintf("%v", value),
		})
	}
	return result, nil
}
