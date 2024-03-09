package main

import (
	"fmt"
	"time"
)

func TimeAgo(pastTime time.Time) string {
	duration := time.Since(pastTime)

	if duration < time.Minute {
		return "just now"
	} else if duration < time.Hour {
		minutes := int(duration.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		} else {
			return fmt.Sprintf("%d minutes ago", minutes)
		}
	} else if duration < 24*time.Hour {
		hours := int(duration.Hours())
		if hours == 1 {
			return "1 hour ago"
		} else {
			return fmt.Sprintf("%d hours ago", hours)
		}
	} else if duration < 30*24*time.Hour {
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		} else {
			return fmt.Sprintf("%d days ago", days)
		}
	} else if duration < 365*24*time.Hour {
		months := int(duration.Hours() / 24 / 30)
		if months == 1 {
			return "1 month ago"
		} else {
			return fmt.Sprintf("%d months ago", months)
		}
	} else {
		years := int(duration.Hours() / 24 / 365)
		if years == 1 {
			return "1 year ago"
		} else {
			return fmt.Sprintf("%d years ago", years)
		}
	}
}
