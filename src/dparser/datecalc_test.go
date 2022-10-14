package dparser

import (
	"testing"
)

func TestTodayTomorrowYesterday(t *testing.T) {
	assertCalcDiff("today", "tomorrow", 1440, t)
	assertCalcDiff("today", "yesterday", -1440, t)
	assertCalcDiff("yesterday", "tomorrow", 2880, t)
}

func assertCalcDiff(d1, d2 string, expMin float64, t *testing.T) {
	dp := Init(d1, d2, -1, false)
	dp.Parse()
	if dp.Output.Diff.Minutes != expMin {
		t.Errorf(
			"CalcDiff fail: %f != %f,\nd1: %q -> %q,\nd2: %q -> %q",
			dp.Output.Diff.Minutes, expMin,
			d1, dp.Output.Dates[0].Date,
			d2, dp.Output.Dates[1].Date,
		)
	}
}
