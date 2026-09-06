---
title: "util.Datetime.toSecondsSinceEpoch"
source: "fgl-topics/c_fgl_ext_util_Datetime_toSecondsSinceEpoch.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Datetime class > util.Datetime methods > util.Datetime.toSecondsSinceEpoch"
type: "concept"
---

# util.Datetime.toSecondsSinceEpoch

> Converts a date/time to a number of seconds since Epoch.

## Syntax

```
util.Datetime.toSecondsSinceEpoch(
  t DATETIME q1 TO q2
 )
  RETURNS FLOAT
```

1. t is the local datetime value.

## Usage

The `util.Datetime.toSecondsSinceEpoch()` method converts the [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") value passed as parameter to a
number of seconds since the UNIX® Epoch
(1970-01-01 00:00:00 GMT)

> **Important:**
>
> The result is a whole number when the source is a `DATETIME
> YEAR TO SECOND`, but will be a floating point number when the source is a `DATETIME
> YEAR TO FRACTION(N)`, to include the fractional part.

## Example

```
IMPORT util
MAIN
    DEFINE sec INTEGER, loc DATETIME YEAR TO SECOND
    LET loc = CURRENT YEAR TO SECOND
    LET sec = util.Datetime.toSecondsSinceEpoch( loc )
    DISPLAY sec
END MAIN
```

## Related links

**Related concepts**  

[util.Datetime.fromSecondsSinceEpoch](3489-util-datetime-fromsecondssinceepoch.md "Converts a number of seconds since Epoch to a date/time.")
