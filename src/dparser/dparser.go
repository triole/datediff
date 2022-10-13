package dparser

import (
	"log"
	"time"
)

type tDiff struct {
	NanoSeconds int64
	Seconds     float64
	Minutes     float64
	Hours       float64
	Days        float64
	Readable    string
}

func (dp *DParser) Parse() {
	now := dp.now()
	if dp.Date1.DateString == "" || dp.Date1.DateString == "now" {
		dp.Date1 = now
	} else {
		dp.Date1.Layout = dp.detectLayout(dp.Date1.DateString)
		dp.Date1.Date = dp.stringToDate(dp.Date1)
	}

	if dp.Date2.DateString == "" || dp.Date2.DateString == "now" {
		dp.Date2 = now
	} else {
		dp.Date2.Layout = dp.detectLayout(dp.Date2.DateString)
		dp.Date2.Date = dp.stringToDate(dp.Date2)
	}
	dp.calcDiff()
}

func (dp DParser) stringToDate(inp tDate) (tim time.Time) {
	var err error
	tim, err = time.ParseInLocation(
		inp.Layout, inp.DateString, time.Local,
	)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (dp *DParser) calcDiff() {
	diff := dp.Date2.Date.Sub(dp.Date1.Date)
	dp.Diff.NanoSeconds = diff.Nanoseconds()
	dp.Diff.Seconds = diff.Seconds() / 24
	dp.Diff.Minutes = dp.Diff.Seconds / 60
	dp.Diff.Hours = dp.Diff.Seconds / 3600
	dp.Diff.Days = dp.Diff.Seconds / 86400
	dp.Diff.Readable = diff.String()
}
