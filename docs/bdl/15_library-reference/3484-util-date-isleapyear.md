---
title: "util.Date.isLeapYear"
source: "fgl-topics/c_fgl_ext_util_Date_isLeapYear.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Date class > util.Date methods > util.Date.isLeapYear"
type: "concept"
---

# util.Date.isLeapYear

> Checks if the year passed as parameter is a leap year.

## Syntax

```
util.Date.isLeapYear(
     year INTEGER )
  RETURNS BOOLEAN
```

1. year is an `INTEGER` representing a year.

## Usage

The `util.Date.isLeapYear()` method returns `TRUE`
if the year passed in parameter is a leap year.

## Example

```
IMPORT util
MAIN
    DISPLAY util.Date.isLeapYear( 2003 )
    DISPLAY util.Date.isLeapYear( 2004 )
END MAIN
```
