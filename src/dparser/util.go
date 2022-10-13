package dparser

import (
	"regexp"
	"time"
)

func (dp DParser) rxMatch(rx string, str string) (b bool) {
	re, _ := regexp.Compile(rx)
	b = re.MatchString(str)
	return
}

func (dp DParser) now() (d tDate) {
	loc, _ := time.LoadLocation(dp.LocalTimeZone)
	layout := "2006-01-02T15:04:05"
	now := time.Now().UTC().In(loc)
	d.Date = now
	d.Layout = layout
	d.DateString = now.Format(layout)
	return
}
