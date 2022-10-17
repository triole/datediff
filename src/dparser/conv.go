package dparser

import (
	"log"
	"time"
)

func (dp DParser) parserDateToTime(inp tDate) (tim time.Time) {
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
