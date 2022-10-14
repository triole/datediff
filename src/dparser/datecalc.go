package dparser

import (
	"time"
)

func (dp DParser) now() (now time.Time) {
	now = time.Now().UTC().In(dp.TimeZoneLocation)
	return
}

func (dp DParser) yesterday() time.Time {
	return dp.addDays(-1, dp.today())
}

func (dp DParser) today() (today time.Time) {
	n := dp.now()
	today = time.Date(
		n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, dp.TimeZoneLocation,
	)
	return
}

func (dp DParser) tomorrow() time.Time {
	return dp.addDays(1, dp.today())
}

func (dp DParser) addDays(days int, tim time.Time) time.Time {
	return tim.AddDate(0, 0, days)
}

func (dp *DParser) calcDiff(d1, d2 time.Time) (diff tDiff) {
	dur := d2.Sub(d1)
	diff.NanoSeconds = dur.Nanoseconds()
	diff.Seconds = dur.Seconds()
	diff.Minutes = dur.Seconds() / 60
	diff.Hours = dur.Seconds() / 3600
	diff.Days = dur.Seconds() / 86400
	diff.Readable = dur.String()
	return
}
