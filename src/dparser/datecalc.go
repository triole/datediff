package dparser

import (
	"strings"
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

func (dp DParser) nextWeekDay(dat time.Time, weekday string) (nwd time.Time) {
	for i := 1; i <= 20; i++ {
		nextDate := dp.addDays(i, dat)
		if strings.ToLower(nextDate.Weekday().String()) == strings.ToLower(weekday) {
			nwd = nextDate
			break
		}
	}
	return
}

func (dp DParser) nextWeekDayEvenOrOdd(dat time.Time, weekday string, odd bool) (nwd time.Time) {
	nwd = dp.nextWeekDay(dat, weekday)
	_, weekno := nwd.ISOWeek()
	if odd == (weekno%2 == 0) {
		nwd = dp.nextWeekDay(nwd, weekday)
		_, weekno = nwd.ISOWeek()
	}
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
	diff.Readable = dur.String()

	arr := []float64{
		dur.Seconds(),
		dur.Seconds() / 60,
		dur.Seconds() / 3600,
		dur.Seconds() / 86400,
	}

	if dp.Round > -1 {
		for i := 0; i <= len(arr)-1; i++ {
			arr[i] = dp.roundFloat(arr[i], uint(dp.Round))
		}
	}

	diff.Seconds = arr[0]
	diff.Minutes = arr[1]
	diff.Hours = arr[2]
	diff.Days = arr[3]

	return
}
