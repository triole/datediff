package dparser

import (
	"log"
	"time"
)

type tDiff struct {
	NanoSeconds int64   `json:"nanoseconds"`
	Seconds     float64 `json:"seconds"`
	Minutes     float64 `json:"minutes"`
	Hours       float64 `json:"hours"`
	Days        float64 `json:"days"`
	Readable    string  `json:"readable"`
}

func (dp *DParser) Parse() {
	now := dp.now(nil)
	for i := 0; i <= len(dp.Output.Dates)-1; i++ {
		if dp.Output.Dates[i].String == "" || dp.Output.Dates[i].String == "now" {
			dp.Output.Dates[i] = now
		} else {
			dp.Output.Dates[i].Layout = dp.detectLayout(dp.Output.Dates[i].String)
			if dp.Output.Dates[i].Layout == "01-02" {
				dp.Output.Dates[i].Layout = "2006-01-02"
				dp.Output.Dates[i].String = dp.now("2006").String + "-" + dp.Output.Dates[i].String
			}
			dp.Output.Dates[i].Date = dp.stringToDate(dp.Output.Dates[i])
		}
	}
	dp.calcDiff(0)
}

func (dp DParser) stringToDate(inp tDate) (tim time.Time) {
	var err error
	tim, err = time.ParseInLocation(
		inp.Layout, inp.String, time.Local,
	)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (dp *DParser) calcDiff(i int) {
	diff := dp.Output.Dates[i+1].Date.Sub(dp.Output.Dates[i].Date)
	dp.Output.Diff.NanoSeconds = diff.Nanoseconds()
	dp.Output.Diff.Seconds = diff.Seconds()
	dp.Output.Diff.Minutes = dp.Output.Diff.Seconds / 60
	dp.Output.Diff.Hours = dp.Output.Diff.Seconds / 3600
	dp.Output.Diff.Days = dp.Output.Diff.Seconds / 86400
	dp.Output.Diff.Readable = diff.String()
}
