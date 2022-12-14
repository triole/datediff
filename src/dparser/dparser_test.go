package dparser

import (
	"testing"
)

func TestDetectLayout(t *testing.T) {
	assertDetectLayout("2022-12-31T12:11:10", "2006-01-02T15:04:05", t)
	assertDetectLayout("2022-12-31T12:11", "2006-01-02T15:04", t)
	assertDetectLayout("2022-12-31T12", "2006-01-02T15", t)
	assertDetectLayout("2022-12-31", "2006-01-02", t)
	assertDetectLayout("12-31", "01-02", t)
}

func assertDetectLayout(d1, exp string, t *testing.T) {
	dp := Init(d1, "", -1, false)
	dp.Parse()
	if dp.detectLayout(d1) != exp {
		t.Errorf("Fail DetectLayout: %s != %s", d1, exp)
	}
}
