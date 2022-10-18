# Datediff ![example workflow](https://github.com/triole/datediff/actions/workflows/build.yaml/badge.svg)

<!--- mdtoc: toc begin -->

1. [Synopsis](#synopsis)
2. [Examples](#examples)
3. [Help](#help)<!--- mdtoc: toc end -->

## Synopsis

Running date diff calculations in shell can be tedious. This is where this tiny tool comes in.

## Examples

```go mdox-exec="sh/dd.sh 2022-10-13T12:00:00 2022-10-15T19:00:00"
$ datediff 2022-10-13T12:00:00 2022-10-15T19:00:00

55h0m0s
```

```go mdox-exec="sh/dd.sh -j 2022-10-01 2022-10-15"
$ datediff -j 2022-10-01 2022-10-15

{
   "nanoseconds": 1209600000000000,
   "seconds": 1209600,
   "minutes": 20160,
   "hours": 336,
   "days": 14,
   "readable": "336h0m0s"
}
```

```go mdox-exec="sh/dd.sh now 2022-12-31"
$ datediff now 2022-12-31

1769h1m14.37872469s
```

```go mdox-exec="sh/dd.sh today tomorrow -p seconds -r 0"
$ datediff today tomorrow -p seconds -r 0

86400
```

```go mdox-exec="sh/dd.sh now next_friday -p hours -r 2"
$ datediff now next_friday -p hours -r 2

64.02
```

```go mdox-exec="sh/dd.sh today next_odd_monday -p hours"
$ datediff today next_odd_monday -p hours

144
```

```go mdox-exec="sh/dd.sh today next_year-11-11 -p days"
$ datediff today next_year-11-11 -p days

389.0416666666667
```

## Help

```go mdox-exec="r -h"

a command line date differ

Arguments:
  [<date-1>]    date1, default 'now'
  [<date-2>]    date2, default 'tomorrow'

Flags:
  -h, --help                Show context-sensitive help.
  -p, --print="readable"    print format, display diff in a specific unit,
                            can be: nano,sec,min,hours,days
  -r, --round=-1            round output values to precision
  -j, --json                print full diff in json format
  -t, --toml                print full diff as toml
  -l, --formats             list supported date formats
  -v, --verbose             verbose mode, default format is json, use -t to
                            switch
  -V, --version-flag        display version
```
