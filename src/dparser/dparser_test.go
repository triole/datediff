package dparser

import (
	"fmt"
	"testing"
)

func TestDetectLayout(t *testing.T) {
	assertDetectLayout("2022-12-31T12:11:10", "2006-01-02T15:04:05", t)
	assertDetectLayout("2022-12-31T12:11", "2006-01-02T15:04", t)
	assertDetectLayout("2022-12-31T12", "2006-01-02T15", t)
	assertDetectLayout("2022-12-31", "2006-01-02", t)
}

func assertDetectLayout(input, exp string, t *testing.T) {
	dp := Init(input, "")
	dp.Parse()
	fmt.Printf("%+v\n", dp)
	if dp.detectLayout(input) != exp {
		t.Errorf("Fail DetectLayout: %s != %s", input, exp)
	}
}