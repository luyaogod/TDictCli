---
title: "util.Datetime.fromSecondsSinceEpoch"
source: "fgl-topics/c_fgl_ext_util_Datetime_fromSecondsSinceEpoch.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Datetime class > util.Datetime methods > util.Datetime.fromSecondsSinceEpoch"
type: "concept"
---

# util.Datetime.fromSecondsSinceEpoch

> Converts a number of seconds since Epoch to a date/time.

## Syntax

```
util.Datetime.fromSecondsSinceEpoch(
  t FLOAT
 )
  RETURNS DATETIME q1 TO q2
```

1. t is the number of seconds since Epoch. This can be a whole integer or a
   decimal, in the target datetime.

## Usage

The `util.Datetime.fromSecondsSinceEpoch()` method converts the number of seconds
since the UNIX® Epoch (1970-01-01 00:00:00 GMT)
passed as parameter, to a `DATETIME` value in local time.

> **Important:**
>
> If the number of seconds passed as parameter is a floating point
> number including a fraction of seconds, the result will be a `DATETIME YEAR
> TO FRACTION(N)`, otherwise it is `DATETIME YEAR TO SECOND`.

## Example

```
IMPORT util
MAIN
    DEFINE dt DATETIME YEAR TO SECOND
    LET dt = util.Datetime.fromSecondsSinceEpoch( 9876234 )
    DISPLAY dt
END MAIN
```

## Related links

**Related concepts**  

[util.Datetime.toSecondsSinceEpoch](3493-util-datetime-tosecondssinceepoch.md "Converts a date/time to a number of seconds since Epoch.")
