package Unixtime

import (
	"fmt"
	"openDevTools/models"
	"time"
)

// ParseStr parse datetime string in format "dd/mm/yyyy hh:mm:ss" to unixtime
func ParseStr(date string) (int64, error) {
	layout := "02/01/2006 15:04:05" // This is the format for "dd/mm/yyyy hh:mm:ss"
	t, err := time.Parse(layout, date)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// ParseUnixTime parse unixtime in string in "dd/mm/yyyy hh:mm:ss" format
func ParseUnixTime(utime int64) ([]models.ResultItem, error) {
	var result []models.ResultItem
	// Convert Unix time to UTC time
	utcTime := time.Unix(utime, 0).UTC()
	result = append(result, models.ResultItem{
		Name:  "UTC",
		Value: utcTime.Format("2006-01-02 15:04:05"),
	})

	// Convert Unix time to local time
	localTime := time.Unix(utime, 0).Local()
	result = append(result, models.ResultItem{
		Name:  "Local Time",
		Value: localTime.Format("2006-01-02 15:04:05"),
	})

	// Convert Unix time to relative time
	t := time.Unix(utime, 0)
	relativeTimeString := getRelativeTime(t)
	result = append(result, models.ResultItem{
		Name:  "Relative Time",
		Value: relativeTimeString,
	})

	return result, nil
}
func getRelativeTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	// Handle future dates first
	if t.After(now) {
		futureDiff := t.Sub(now)
		switch {
		case futureDiff < time.Hour:
			return "In less than an hour"
		case futureDiff < 2*time.Hour:
			return "In about an hour"
		case futureDiff < 24*time.Hour:
			hours := int((futureDiff + time.Hour/2) / time.Hour)
			return fmt.Sprintf("In %d hours", hours)
		case futureDiff < 7*24*time.Hour:
			days := int(futureDiff / (24 * time.Hour))
			if days == 1 {
				return "Tomorrow"
			}
			return fmt.Sprintf("In %d days", days)
		case futureDiff < 30*24*time.Hour: // within a month
			return "Next month"
		default:
			// For dates farther than a month
			if t.Year() == now.Year()+1 {
				return "Next year"
			}
			return t.Format("2006-01-02") // format date for future years
		}
	}

	// Handle past dates
	switch {
	case diff < time.Minute:
		return "Less than a minute ago"
	case diff < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(diff/time.Minute))
	case diff < 2*time.Hour:
		return "About an hour ago"
	case diff < 24*time.Hour:
		hours := int(diff / time.Hour)
		return fmt.Sprintf("%d hours ago", hours)
	case diff < 7*24*time.Hour: // One week
		days := int(diff / (24 * time.Hour))
		if days == 1 {
			return "Yesterday"
		}
		return fmt.Sprintf("%d days ago", days)
	default:
		// For longer periods in the past
		if now.Year() == t.Year() && int(diff/365) == 1 {
			return "One year ago"
		} else if t.Year()-now.Year() == -1 {
			return "One year ago"
		} else if int(diff/30) < 2 {
			return "Last month"
		} else {
			return t.Format("02/01/2006")
		}
	}
}
