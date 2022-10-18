package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	// BUILDTAGS are injected ld flags during build
	BUILDTAGS      string
	appName        = "datediff"
	appDescription = "a command line date differ"
	appMainversion = "0.1"
)

var CLI struct {
	Date1       string `help:"date1, default 'now'" arg optional default:now`
	Date2       string `help:"date2, default 'tomorrow'" arg optional default:tomorrow`
	Print       string `help:"print format, display diff in a specific unit, can be: nano,sec,min,hours,days" short:p enum:"readable,nano,seconds,minutes,hours,days" default:readable`
	Round       int    `help:"round output values to precision" short:r default:-1`
	JSON        bool   `help:"print full diff in json format" short:j`
	TOML        bool   `help:"print full diff as toml" short:t`
	Formats     bool   `help:"list supported date formats" short:l`
	Verbose     bool   `help:"verbose mode, default format is json, use -t to switch" short:v`
	VersionFlag bool   `help:"display version" short:V`
}

func parseArgs() {
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:      true,
			Summary:      true,
			NoAppSummary: true,
			FlagsLast:    false,
		}),
		kong.Vars{},
	)
	_ = ctx.Run()

	if CLI.VersionFlag == true {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	// ctx.FatalIfErrorf(err)
}

type tPrinter []tPrinterEl
type tPrinterEl struct {
	Key string
	Val string
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "version: "+appMainversion+".", -1)
	fmt.Printf("\n%s\n%s\n\n", appName, appDescription)
	arr := strings.Split(s, "\n")
	var pr tPrinter
	var maxlen int
	for _, line := range arr {
		if strings.Contains(line, ":") {
			l := strings.Split(line, ":")
			if len(l[0]) > maxlen {
				maxlen = len(l[0])
			}
			pr = append(pr, tPrinterEl{l[0], strings.Join(l[1:], ":")})
		}
	}
	for _, el := range pr {
		fmt.Printf("%"+strconv.Itoa(maxlen)+"s\t%s\n", el.Key, el.Val)
	}
	fmt.Printf("\n")
}

func getBindir() (s string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	s = filepath.Dir(ex)
	return
}
