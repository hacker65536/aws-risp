package util

import (
	"strconv"
	"strings"
	"time"
)

func StartEnd(start, end string) (string, string) {
	// 2022-09-25T16:08:04+0900 > 2025-09-25T16:08:04+0900
	if start != "" && end != "" {
		return start, end
	}
	return GetLastWeek()
}

func GetLastWeek() (string, string) {
	t := time.Now()
	last := t.AddDate(0, 0, -1)
	first := last.AddDate(0, 0, -6)

	return first.Format("2006-01-02"), last.Format("2006-01-02")
}

func ToJst(t string) string {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t1, _ := time.Parse(time.RFC3339, t)
	//  2006-01-02T15:04:05-0700 > 2025-09-25T16:08:04+0900
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
