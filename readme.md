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
   "NanoSeconds": 1209600000000000,
   "Seconds": 1209600,
   "Minutes": 20160,
   "Hours": 336,
   "Days": 14,
   "Readable": "336h0m0s"
}
```

## Help

```go mdox-exec="r -h"

a command line date differ

Arguments:
  [<date-1>]    date1
  [<date-2>]    date2

Flags:
  -h, --help            Show context-sensitive help.
  -j, --json            print diff output in json format
  -l, --formats         list supported date formats
  -V, --version-flag    display version
```
