package main

import (
	"datediff/dparser"
	"fmt"
)

func main() {
	parseArgs()

	dp := dparser.Init(CLI.Date1, CLI.Date2)

	if CLI.Formats == true {
		dp.ListSupportedFormats()
	} else {

		dp.Parse()

		if CLI.JSON == true {
			dp.PrintDiffJSON()
		} else {
			fmt.Printf("%s\n", dp.Diff.Readable)
		}

	}
}
