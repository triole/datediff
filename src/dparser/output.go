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
		fmt.Printf("%s\n", dp.Output.Diff.Readable)
	case "nano":
		fmt.Printf("%d\n", dp.Output.Diff.NanoSeconds)
	case "sec":
		fmt.Printf("%f\n", dp.Output.Diff.Seconds)
	case "min":
		fmt.Printf("%f\n", dp.Output.Diff.Minutes)
	case "hours":
		fmt.Printf("%f\n", dp.Output.Diff.Hours)
	case "days":
		fmt.Printf("%f\n", dp.Output.Diff.Days)
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
