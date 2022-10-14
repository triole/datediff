package dparser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func (dp DParser) PrintJSON(obj interface{}) {
	json, err := json.MarshalIndent(obj, "", "   ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(json))
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
