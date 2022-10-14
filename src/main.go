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

		if CLI.Verbose == true {
			dp.PrintJSON(dp.Output)
		}

		if CLI.JSON == true && CLI.Verbose == false {
			dp.PrintJSON(dp.Output.Diff)
		}

		if CLI.JSON == false && CLI.Verbose == false {
			fmt.Printf("%s\n", dp.Output.Diff.Readable)
		}
	}
}
