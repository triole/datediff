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
	now := dp.timeToParserDate(dp.now(), nil)
	for i := 0; i <= len(dp.Output.Dates)-1; i++ {
		switch dp.Output.Dates[i].String {
		case "now":
			dp.Output.Dates[i] = now
		case "yesterday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.yesterday(), nil)
		case "today":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.today(), nil)
		case "tomorrow":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.tomorrow(), nil)
		default:
			dp.Output.Dates[i].Layout = dp.detectLayout(dp.Output.Dates[i].String)
			if dp.Output.Dates[i].Layout == "01-02" {
				dp.Output.Dates[i].Layout = "2006-01-02"
				dp.Output.Dates[i].String = dp.timeToParserDate(
					dp.now(), "2006",
				).String + "-" + dp.Output.Dates[i].String
			}
			dp.Output.Dates[i].Date = dp.stringToDate(dp.Output.Dates[i])
		}
	}
	dp.Output.Diff = dp.calcDiff(
		dp.Output.Dates[0].Date, dp.Output.Dates[1].Date,
	)
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

func (dp DParser) timeToParserDate(tim time.Time, lt interface{}) (d tDate) {
	layout := "2006-01-02T15:04:05"
	switch lt.(type) {
	case string:
		layout = lt.(string)
	}
	d.Date = tim
	d.Layout = layout
	d.String = tim.Format(layout)
	return
}
