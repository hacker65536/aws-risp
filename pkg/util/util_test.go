package util

import (
	"testing"
)

func TestGetLasttWeek(t *testing.T) {
	f, l := GetLastWeek()
	t.Logf("fisrt %s last %s", f, l)
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
