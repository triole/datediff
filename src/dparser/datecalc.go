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
