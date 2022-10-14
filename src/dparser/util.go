package dparser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

func (dp DParser) rxMatch(rx string, str string) (b bool) {
	re, _ := regexp.Compile(rx)
	b = re.MatchString(str)
	return
}

func (dp DParser) now(lt interface{}) (d tDate) {
	layout := "2006-01-02T15:04:05"
	switch lt.(type) {
	case string:
		layout = lt.(string)
	}
	now := time.Now().UTC().In(dp.TimeZoneLocation)
	d.Date = now
	d.Layout = layout
	d.String = now.Format(layout)
	return
}

func (dp DParser) pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "    ")
	fmt.Println(string(s))
}
