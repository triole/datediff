package dparser

import (
	"log"
	"time"

	timezone "github.com/gandarez/go-olson-timezone"
)

type DParser struct {
	Date1         tDate
	Date2         tDate
	Diff          tDiff
	DiffObj       time.Duration
	LocalTimeZone string
}

type tDate struct {
	DateString string
	Layout     string
	Date       time.Time
}

func Init(date1, date2 string) (dp DParser) {
	dp.Date1.DateString = date1
	dp.Date2.DateString = date2
	zone, err := timezone.Name()
	if err != nil {
		log.Fatal(err)
	}
	dp.LocalTimeZone = zone
	return
}
