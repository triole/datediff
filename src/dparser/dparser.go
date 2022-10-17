package dparser

import (
	"strconv"
	"strings"
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
		// TODO: improve calling logic
		dp.Output.Dates[i].String = strings.Replace(
			dp.Output.Dates[i].String, "next_year",
			strconv.Itoa(
				dp.nextYear(dp.Output.Dates[i].Date).Year(),
			), -1,
		)
		dp.Output.Dates[i].String = strings.Replace(
			dp.Output.Dates[i].String, "last_year",
			strconv.Itoa(
				dp.lastYear(dp.Output.Dates[i].Date).Year(),
			), -1,
		)

		switch dp.Output.Dates[i].String {
		case "now":
			dp.Output.Dates[i] = now
		case "yesterday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.yesterday(), nil)
		case "today":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.today(), nil)
		case "tomorrow":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.tomorrow(), nil)
		case "next_monday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDay(dp.today(), "monday"), nil)
		case "next_tuesday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDay(dp.today(), "tuesday"), nil)
		case "next_wednesday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDay(dp.today(), "wednesday"), nil)
		case "next_thursday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDay(dp.today(), "thursday"), nil)
		case "next_friday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDay(dp.today(), "friday"), nil)
		case "next_saturday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDay(dp.today(), "saturday"), nil)
		case "next_sunday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDay(dp.today(), "sunday"), nil)
		case "next_even_monday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "monday", false), nil)
		case "next_even_tuesday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "tuesday", false), nil)
		case "next_even_wednesday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "wednesday", false), nil)
		case "next_even_thursday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "thursday", false), nil)
		case "next_even_friday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "friday", false), nil)
		case "next_even_saturday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "saturday", false), nil)
		case "next_even_sunday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "sunday", false), nil)
		case "next_odd_monday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "monday", true), nil)
		case "next_odd_tuesday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "tuesday", true), nil)
		case "next_odd_wednesday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "wednesday", true), nil)
		case "next_odd_thursday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "thursday", true), nil)
		case "next_odd_friday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "friday", true), nil)
		case "next_odd_saturday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "saturday", true), nil)
		case "next_odd_sunday":
			dp.Output.Dates[i] = dp.timeToParserDate(dp.nextWeekDayEvenOrOdd(dp.today(), "sunday", true), nil)
		default:
			dp.Output.Dates[i].Layout = dp.detectLayout(dp.Output.Dates[i].String)
			if dp.Output.Dates[i].Layout == "01-02" {
				dp.Output.Dates[i].Layout = "2006-01-02"
				dp.Output.Dates[i].String = dp.timeToParserDate(
					dp.now(), "2006",
				).String + "-" + dp.Output.Dates[i].String
			}
			dp.Output.Dates[i].Date = dp.parserDateToTime(dp.Output.Dates[i])
		}
		dp.Output.Dates[i].Unix = dp.Output.Dates[i].Date.Unix()
	}
	dp.Output.Diff = dp.calcDiff(
		dp.Output.Dates[0].Date, dp.Output.Dates[1].Date,
	)
}
