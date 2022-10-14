package dparser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/table"
	toml "github.com/pelletier/go-toml/v2"
)

func (dp DParser) Print(obj interface{}) {
	if dp.TOML == true {
		dp.printTOML(obj)
	} else {
		dp.printJSON(obj)
	}
}

func (dp DParser) printJSON(obj interface{}) {
	json, err := json.MarshalIndent(obj, "", "   ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(json))
}

func (dp DParser) printTOML(obj interface{}) {
	toml, err := toml.Marshal(obj)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(toml))
}

func (dp DParser) PrintShort(format string) {
	switch format {
	case "readable":
		dp.printVal(dp.Output.Diff.Readable)
	case "nano":
		dp.printVal(dp.Output.Diff.NanoSeconds)
	case "sec":
		dp.printVal(dp.Output.Diff.Seconds)
	case "min":
		dp.printVal(dp.Output.Diff.Minutes)
	case "hours":
		dp.printVal(dp.Output.Diff.Hours)
	case "days":
		dp.printVal(dp.Output.Diff.Days)
	}
}

func (dp DParser) printVal(itf interface{}) {
	switch val := itf.(type) {
	case int:
		fmt.Printf("%d\n", val)
	case float64:
		if dp.Round == 0 {
			fmt.Printf("%d\n", int(val))
		} else {
			fmt.Printf("%f\n", val)
		}
	default:
		fmt.Printf("%s\n", val)
	}
}

func (dp DParser) ListSupportedFormats() {
	fmt.Printf("\nThe following date layouts are supported\n\n")
	t := initTable()
	t.AppendHeader(table.Row{"Format", "Layout", "Comment"})
	for _, el := range dp.Layouts {
		t.AppendRow(
			[]interface{}{
				el.Desc, el.Layout, el.Comment,
			},
		)
	}
	fmt.Printf("\n")
	t.Render()
}

func initTable() (t table.Writer) {
	t = table.NewWriter()
	t.SetStyle(table.Style{
		Name: "myNewStyle",
		Box: table.BoxStyle{
			BottomLeft:       "\\",
			BottomRight:      "/",
			BottomSeparator:  "v",
			Left:             "[",
			LeftSeparator:    "{",
			MiddleHorizontal: "-",
			MiddleSeparator:  "+",
			MiddleVertical:   "|",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			Right:            " ]",
			RightSeparator:   "}",
			TopLeft:          "(",
			TopRight:         ")",
			TopSeparator:     "^",
			UnfinishedRow:    " ~~~",
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: true,
			SeparateFooter:  true,
			SeparateHeader:  true,
			SeparateRows:    false,
		},
	})
	t.SetOutputMirror(os.Stdout)
	return
}
