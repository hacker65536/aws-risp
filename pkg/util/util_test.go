package util

import (
	"testing"
	"time"
)

func TestGetLastWeek(t *testing.T) {
	first, last := GetLastWeek()

	// Parse the returned dates
	firstDate, err1 := time.Parse("2006-01-02", first)
	lastDate, err2 := time.Parse("2006-01-02", last)

	if err1 != nil || err2 != nil {
		t.Fatalf("GetLastWeek() returned invalid date format: (%s, %s)", first, last)
	}

	// Check that the dates are 6 days apart
	diff := lastDate.Sub(firstDate)
	if diff.Hours() != 144 { // 6 days = 144 hours
		t.Errorf("GetLastWeek() returned dates %s and %s, which are %.0f hours apart; want 144 hours",
			first, last, diff.Hours())
	}

	// Check that the end date is yesterday
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1).Format("2006-01-02")

	if last != yesterday {
		t.Errorf("GetLastWeek() end date = %s; want %s (yesterday)", last, yesterday)
	}

	t.Logf("GetLastWeek() = %s to %s", first, last)
}

func TestStartEnd(t *testing.T) {
	// Test with provided start and end dates
	start, end := "2023-01-01", "2023-01-31"
	resultStart, resultEnd := StartEnd(start, end)

	if resultStart != start || resultEnd != end {
		t.Errorf("StartEnd(%s, %s) = (%s, %s); want (%s, %s)",
			start, end, resultStart, resultEnd, start, end)
	}

	// Test with empty values (should return last week)
	resultStart, resultEnd = StartEnd("", "")

	// Verify the dates are in the correct format
	_, err1 := time.Parse("2006-01-02", resultStart)
	_, err2 := time.Parse("2006-01-02", resultEnd)

	if err1 != nil || err2 != nil {
		t.Errorf("StartEnd(\"\", \"\") returned invalid date format: (%s, %s)",
			resultStart, resultEnd)
	}
}

func TestToJst(t *testing.T) {
	t1 := ToJst("2025-09-25T07:08:04.000Z")
	t.Logf("t1 %s", t1)
}

func TestServiceName(t *testing.T) {
	s := ServiceName("Amazon Elastic Compute Cloud - Compute")
	t.Logf("s %s", s)
}

func TestToInt(t *testing.T) {
	s := "87.1339690972"
	i := ToInt(s)
	t.Logf("i %d", i)
}

func TestTo2dp(t *testing.T) {
	s := "87.1339690972"
	i := To2dp(s)
	t.Logf("i %s", i)
}

func TestToLowers(t *testing.T) {
	s := []string{"A", "B", "C"}
	i := ToLowers(s)
	t.Logf("i %v", i)
}

func BenchmarkToLowers(b *testing.B) {
	s := []string{"A", "B", "C"}
	for i := 0; i < b.N; i++ {
		ToLowers(s)
	}
}
