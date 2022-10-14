package dparser

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func (dp DParser) rxMatch(rx string, str string) (b bool) {
	re, _ := regexp.Compile(rx)
	b = re.MatchString(str)
	return
}

func (dp DParser) pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "    ")
	fmt.Println(string(s))
}
