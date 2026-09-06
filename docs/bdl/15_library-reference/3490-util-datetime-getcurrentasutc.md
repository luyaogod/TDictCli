---
title: "util.Datetime.getCurrentAsUTC"
source: "fgl-topics/c_fgl_ext_util_Datetime_getCurrentAsUTC.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Datetime class > util.Datetime methods > util.Datetime.getCurrentAsUTC"
type: "concept"
---

# util.Datetime.getCurrentAsUTC

> Returns the current date/time in UTC.

## Syntax

```
util.Datetime.getCurrentAsUTC( )
  RETURNS DATETIME YEAR TO FRACTION(5)
```

## Usage

The `util.Datetime.getCurrentAsUTC()` method returns the current
system date/time in UTC (Coordinated Universal Time).

This method is provided to solve the daylight saving time transition issue of the
[`util.Datetime.toUTC()`](3494-util-datetime-toutc.md "Converts a date/time value to the UTC date/time.") method.

> **Note:**
>
> The precision of the value returned by this method is a `DATETIME YEAR TO
> FRACTION(5)`. Note that this precision is different from the default
> `CURRENT` precision when no qualifiers are specified.

## Example

```
IMPORT util
MAIN
    DEFINE utc DATETIME YEAR TO FRACTION(5)
    LET utc = util.Datetime.getCurrentAsUTC( )
    DISPLAY "Current UTC: ", utc
END MAIN
```

## Related links

**Related concepts**  

[util.Datetime.toLocalTime](3492-util-datetime-tolocaltime.md "Converts a UTC date/time to the local time.")
