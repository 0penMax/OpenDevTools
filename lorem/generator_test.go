package lorem

import (
	"strconv"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		wordCount     int
		expectedCount int
	}{
		{wordCount: 0, expectedCount: 0},
		{wordCount: 1, expectedCount: 1},
		{wordCount: 10, expectedCount: 10},
		{wordCount: 20, expectedCount: 20},
		{wordCount: 50, expectedCount: 50},
	}

	for _, tc := range testCases {
		t.Run("wordCount_"+strconv.Itoa(tc.wordCount), func(t *testing.T) {
			result := Generate(tc.wordCount)
			// Split the generated string into words.
			words := strings.Fields(result)
			if len(words) != tc.expectedCount {
				t.Errorf("Generate(%d) returned %d words; expected %d", tc.wordCount, len(words), tc.expectedCount)
			}
		})
	}
}
