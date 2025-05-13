package util

import (
	"strconv"
	"strings"
	"time"
)

// StartEnd returns the start and end dates based on input values.
// If both start and end are provided, they are returned as is.
// If either is missing, it returns the date range for the last week.
func StartEnd(start, end string) (string, string) {
	if start != "" && end != "" {
		return start, end
	}
	return GetLastWeek()
}

// GetLastWeek returns the start and end dates for the last 7 days.
// The end date is yesterday and the start date is 6 days before that.
// Dates are returned in YYYY-MM-DD format.
func GetLastWeek() (string, string) {
	t := time.Now()
	last := t.AddDate(0, 0, -1)     // Yesterday
	first := last.AddDate(0, 0, -6) // 7 days ago

	return first.Format("2006-01-02"), last.Format("2006-01-02")
}

// ToJst converts an RFC3339 timestamp to Japan Standard Time (JST).
// Returns the time formatted as YYYY-MM-DDThh:mm:ss-0700.
func ToJst(t string) string {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t1, _ := time.Parse(time.RFC3339, t)
	return t1.In(loc).Format("2006-01-02T15:04:05-0700")
}

func ServiceName(service string) string {
	switch service {
	case "Amazon Elastic Compute Cloud - Compute":
		return "EC2"
	case "Amazon Relational Database Service":
		return "RDS"
	case "Amazon ElastiCache":
		return "ElastiCache"
	case "Amazon OpenSearch Service":
		return "OpenSearch"
	default:
		return "Unknown"
	}
}

func ToInt(s string) int {
	f, _ := strconv.ParseFloat(s, 64)
	return int(f)
}
func To2dp(s string) string {
	f, _ := strconv.ParseFloat(s, 64)
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func ToLowers(s []string) []string {

	var lowerStr []string
	for _, v := range s {
		lowerStr = append(lowerStr, strings.ToLower(v))
	}
	return lowerStr
}
