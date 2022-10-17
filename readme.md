# Datediff ![example workflow](https://github.com/triole/datediff/actions/workflows/build.yaml/badge.svg)

<!--- mdtoc: toc begin -->

1. [Synopsis](#synopsis)
2. [Examples](#examples)
3. [Help](#help)<!--- mdtoc: toc end -->

## Synopsis

Running date diff calculations in shell can be tedious. This is where this tiny tool comes in.

## Examples

$ datediff 2022-10-13T12:00:00 2022-10-15T19:00:00

```go mdox-exec="r 2022-10-13T12:00:00 2022-10-15T19:00:00"
55h0m0s
```

$ datediff -j 2022-10-01 2022-10-15

```go mdox-exec="r -j 2022-10-01 2022-10-15"
{
   "nanoseconds": 1209600000000000,
   "seconds": 1209600,
   "minutes": 20160,
   "hours": 336,
   "days": 14,
   "readable": "336h0m0s"
}
```

$ datediff now 2022-12-31

```go mdox-exec="r now 2022-12-31"
1787h54m16.614944778s
```

```go mdox-exec="r today tomorrow"
24h0m0s
```

```go mdox-exec="r today next_friday"
109h5m43.961155549s
```

```go mdox-exec="r today next_odd_monday"
181h5m44.22773963s
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
