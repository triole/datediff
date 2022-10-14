package dparser

import (
	_ "embed"
)

type tLayouts []tLayout

type tLayout struct {
	Layout  string
	Matcher string
	Desc    string
	Comment string
}

func (dp DParser) detectLayout(str string) (r string) {
	for _, el := range dp.Layouts {
		if dp.rxMatch(el.Matcher, str) == true {
			r = el.Layout
			break
		}
	}
	return
}
