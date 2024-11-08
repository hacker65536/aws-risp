package util

import (
	"strconv"
	"time"
)

func GetCurrentWeek() (string, string) {
	t := time.Now()
	last := t.AddDate(0, 0, -2)
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

func Platform(s string) string {
	switch s {
	case "Aurora PostgreSQL":
		return "Aurora-PostgreSQL"
	default:
		if s == "" {
			return "null"
		}
		return s
	}
}
