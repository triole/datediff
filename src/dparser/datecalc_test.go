package dparser

import (
	"testing"
	"time"
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

func TestNextWeekDay(t *testing.T) {
	assertNextWeekday("2022-10-01", "sunday", "2022-10-02", t)
	assertNextWeekday("2022-12-31", "wednesday", "2023-01-04", t)
}

func assertNextWeekday(dat, wd, exp string, t *testing.T) {
	dp := Init("today", "tomorrow", -1, false)
	inp := tDate{
		String: dat,
		Layout: dp.detectLayout(dat),
	}
	res := dp.nextWeekDay(dp.parserDateToTime(inp), wd)
	expt, _ := time.ParseInLocation("2006-01-02", exp, time.Local)
	if res != expt {
		t.Errorf("NextWeekDay fail: %s != %s", res, expt)
	}
}

func TestNextWeekDayEvenOrOdd(t *testing.T) {
	assertNextWeekdayEvenOrOdd("2022-10-01", "sunday", false, "2022-10-09", t)
	assertNextWeekdayEvenOrOdd("2022-12-31", "wednesday", true, "2023-01-04", t)
}

func assertNextWeekdayEvenOrOdd(dat, wd string, even bool, exp string, t *testing.T) {
	dp := Init("today", "tomorrow", -1, false)
	inp := tDate{
		String: dat,
		Layout: dp.detectLayout(dat),
	}
	res := dp.nextWeekDayEvenOrOdd(dp.parserDateToTime(inp), wd, even)
	expt, _ := time.ParseInLocation("2006-01-02", exp, time.Local)
	if res != expt {
		t.Errorf("NextWeekDay fail: %s != %s", res, expt)
	}
}
