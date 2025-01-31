package Unixtime

import (
	"testing"
	"time"
)

func TestGetRelativeTime(t *testing.T) {
	// Current time for reference
	now := time.Now()

	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "Less than a minute ago",
			input:    now.Add(-30 * time.Second),
			expected: "Less than a minute ago",
		},
		{
			name:     "3 minutes ago",
			input:    now.Add(-3 * time.Minute),
			expected: "3 minutes ago",
		},
		{
			name:     "45 minutes ago",
			input:    now.Add(-45 * time.Minute),
			expected: "45 minutes ago",
		},
		{
			name:     "2 hours ago",
			input:    now.Add(-2 * time.Hour),
			expected: "2 hours ago",
		},
		{
			name:     "Yesterday",
			input:    now.Add(-24 * time.Hour),
			expected: "Yesterday",
		},
		{
			name:     "5 days ago",
			input:    now.Add(-5 * 24 * time.Hour),
			expected: "5 days ago",
		},
		{
			name:     "Next month",
			input:    now.Add(25 * 24 * time.Hour), // 25 days from now, still in the next month
			expected: "Next month",
		},
		{
			name:     "Next year",
			input:    now.Add(370 * 24 * time.Hour), // 370 days from now, next year
			expected: "Next year",
		},
		{
			name:     "One year ago",
			input:    now.Add(-365 * 24 * time.Hour), // One year ago
			expected: "One year ago",
		},
		{
			name:     "In 5 hours",
			input:    now.Add(5 * time.Hour),
			expected: "In 5 hours",
		},
		{
			name:     "Tomorrow",
			input:    now.Add(25 * time.Hour),
			expected: "Tomorrow",
		},
		{
			name:     "In less than an hour",
			input:    now.Add(30 * time.Minute),
			expected: "In less than an hour",
		},
		{
			name:     "In about an hour",
			input:    now.Add(90 * time.Minute),
			expected: "In about an hour",
		},
		{
			name:     "In about an hour",
			input:    now.Add(70 * time.Minute),
			expected: "In about an hour",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getRelativeTime(tt.input)
			if actual != tt.expected {
				t.Errorf("For %s, expected %s, but got %s", tt.name, tt.expected, actual)
			}
		})
	}
}
