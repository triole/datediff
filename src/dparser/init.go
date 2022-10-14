package dparser

import (
	_ "embed"
	"encoding/json"
	"log"
	"time"

	timezone "github.com/gandarez/go-olson-timezone"
)

var (
	//go:embed embed/layouts.json
	layouts []byte
)

type DParser struct {
	Date1         tDate
	Date2         tDate
	Diff          tDiff
	DiffObj       time.Duration
	LocalTimeZone string
	Layouts       tLayouts
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
	err = json.Unmarshal(layouts, &dp.Layouts)
	if err != nil {
		log.Fatal(err)
	}
	return
}
