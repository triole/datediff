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
	Output           tOutput
	TimeZoneLocation *time.Location
	DiffObj          time.Duration
	Layouts          tLayouts
}

type tOutput struct {
	LocalTimeZone string `json:"local_timezone"`
	Dates         tDates `json:"dates"`
	Diff          tDiff  `json:"diff"`
}

type tDates []tDate

type tDate struct {
	DateString string    `json:"string"`
	Layout     string    `json:"layout"`
	Date       time.Time `json:"date"`
}

func Init(date1, date2 string) (dp DParser) {
	dp.Output.Dates = addDateString(date1, dp.Output.Dates)
	dp.Output.Dates = addDateString(date2, dp.Output.Dates)
	zone, err := timezone.Name()
	if err != nil {
		log.Fatal(err)
	}
	dp.Output.LocalTimeZone = zone
	var loc *time.Location
	loc, err = time.LoadLocation(dp.Output.LocalTimeZone)
	if err != nil {
		log.Fatal(err)
	}
	dp.TimeZoneLocation = loc
	err = json.Unmarshal(layouts, &dp.Layouts)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func addDateString(str string, tdates tDates) tDates {
	td := tDate{
		DateString: str,
	}
	tdates = append(tdates, td)
	return tdates
}
