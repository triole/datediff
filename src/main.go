package main

import (
	"datediff/dparser"
	"fmt"
)

func main() {
	parseArgs()

	dp := dparser.Init(CLI.Date1, CLI.Date2, CLI.TOML)

	if CLI.Formats == true {
		dp.ListSupportedFormats()
	} else {

		dp.Parse()

		if CLI.Verbose == true {
			dp.Print(dp.Output)
		} else {
			if CLI.JSON == true || CLI.TOML == true {
				dp.Print(dp.Output.Diff)
			}
		}

		if CLI.Verbose == false && CLI.JSON == false && CLI.TOML == false {
			fmt.Printf("%s\n", dp.Output.Diff.Readable)
		}
	}
}
