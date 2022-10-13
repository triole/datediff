package dparser

type tLayouts []tLayout

type tLayout struct {
	Layout     string
	Detector   string
	Descriptor string
}

func (dp DParser) layoutDefinitions() (ll tLayouts) {
	ll = append(ll, dp.newLayout(
		"2006-01-02T15:04:05",
		"^[0-9]{4}-[0-9]{2}-[0-9]{2}(T|t)[0-9]{2}:[0-9]{2}:[0-9]{2}$",
		"YYYY-MM-DDTHH:MM:SS [24h]",
	),
	)
	ll = append(ll, dp.newLayout(
		"2006-01-02T15:04",
		"^[0-9]{4}-[0-9]{2}-[0-9]{2}(T|t)[0-9]{2}:[0-9]{2}$",
		"YYYY-MM-DDTHH:MM [24h]",
	),
	)
	ll = append(ll, dp.newLayout(
		"2006-01-02T15",
		"^[0-9]{4}-[0-9]{2}-[0-9]{2}(T|t)[0-9]{2}$",
		"YYYY-MM-DDTHH [24h]",
	),
	)
	ll = append(ll, dp.newLayout(
		"2006-01-02",
		"^[0-9]{4}-[0-9]{2}-[0-9]{2}$",
		"YYYY-MM-DD",
	),
	)
	return
}

func (dp DParser) newLayout(str, det, des string) (l tLayout) {
	l.Layout = str
	l.Detector = det
	l.Descriptor = des
	return
}

func (dp DParser) detectLayout(str string) (r string) {
	for _, el := range dp.layoutDefinitions() {
		if dp.rxMatch(el.Detector, str) == true {
			r = el.Layout
			break
		}
	}
	return
}
