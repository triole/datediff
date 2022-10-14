package dparser

import (
	"encoding/json"
	"fmt"
	"math"
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

func (dp DParser) roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
