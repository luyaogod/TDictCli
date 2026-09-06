---
title: "util.Datetime.toLocalTime"
source: "fgl-topics/c_fgl_ext_util_Datetime_toLocalTime.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Datetime class > util.Datetime methods > util.Datetime.toLocalTime"
type: "concept"
---

# util.Datetime.toLocalTime

> Converts a UTC date/time to the local time.

## Syntax

```
util.Datetime.toLocalTime(
  t DATETIME q1 TO q2
 )
  RETURNS DATETIME q1 TO q2
```

1. t is the datetime value in UTC.

## Usage

The `util.Datetime.toLocalTime()` method converts a `DATETIME`
value from UTC (Coordinated Universal Time), to the local timezone date/time.

## Example

```
IMPORT util
MAIN
    DEFINE loc DATETIME YEAR TO SECOND
    LET loc = util.Datetime.toLocalTime( DATETIME(2015-08-22 15:34:56) YEAR TO SECOND )
    DISPLAY "LOC: ", loc
END MAIN
```

## Related links

**Related concepts**  

[util.Datetime.toUTC](3494-util-datetime-toutc.md "Converts a date/time value to the UTC date/time.")

[util.Datetime.getCurrentAsUTC](3490-util-datetime-getcurrentasutc.md "Returns the current date/time in UTC.")

[TZ](../07_configuration/0504-tz.md "Defines the timezone for date/time values handling.")
