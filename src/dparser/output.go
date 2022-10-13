package dparser

import (
	"encoding/json"
	"fmt"
	"log"
)

func (dp DParser) PrintDiffJSON() {
	json, err := json.MarshalIndent(dp.Diff, "", "   ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(json))
}

func (dp DParser) ListSupportedFormats() {
	fmt.Printf("\nThe following date layouts are supported\n\n")
	for _, el := range dp.layoutDefinitions() {
		fmt.Printf("%22s |%s\n", el.Layout, el.Descriptor)
	}
	fmt.Printf("\n")
}
