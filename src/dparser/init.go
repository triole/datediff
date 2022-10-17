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
	Round            int
	TOML             bool
}

type tOutput struct {
	LocalTimeZone string `json:"local_timezone"`
	Dates         tDates `json:"dates"`
	Diff          tDiff  `json:"diff"`
}

type tDates []tDate

type tDate struct {
	Input  string    `json:"input"`
	String string    `json:"string"`
	Layout string    `json:"layout"`
	Date   time.Time `json:"date"`
	Unix   int64     `json:"unix"`
}

func Init(date1, date2 string, round int, toml bool) (dp DParser) {
	dp.Output.Dates = addToOutputDates(date1, dp.Output.Dates)
	dp.Output.Dates = addToOutputDates(date2, dp.Output.Dates)
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
	dp.Round = round
	dp.TOML = toml
	return
}

func addToOutputDates(str string, tdates tDates) tDates {
	td := tDate{Input: str, String: str}
	tdates = append(tdates, td)
	return tdates
}
