---
title: "util.Datetime.toUTC"
source: "fgl-topics/c_fgl_ext_util_Datetime_toUTC.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Datetime class > util.Datetime methods > util.Datetime.toUTC"
type: "concept"
---

# util.Datetime.toUTC

> Converts a date/time value to the UTC date/time.

## Syntax

```
util.Datetime.toUTC(
   t DATETIME q1 TO q2
)
  RETURNS DATETIME q1 TO q2
```

1. t is the local timezone datetime value.

## Usage

The `util.Datetime.toUTC()` method converts the local timezone
`DATETIME` value passed as parameter to the "Coordinated
Universal Time" (UTC), also known as "Greenwich Mean Time" (GMT).

The `toUTC()` method uses local timezone settings of the system.

## Datetime values without full date part

When the source `DATETIME` value to be converted does not contain a complete date
part starting from year, the conversion function uses the current system date to fill the missing
parts. The resulting UTC time value can consequently be different depending on the timezone /
daylight saving time.

In the next example, a `DATETIME HOUR TO MINUTE` is converted to UTC using the
current system date:

```
IMPORT util
MAIN
    DEFINE dt_h2s DATETIME HOUR TO SECOND -- local time
    LET dt_h2s = DATETIME(12:00:00) HOUR TO SECOND
    DISPLAY "Local time : ", dt_h2s
    DISPLAY "Current UTC: ", util.Datetime.getCurrentAsUTC()
    DISPLAY "EXTEND Y2S : ", EXTEND(dt_h2s, YEAR TO SECOND) -- uses current date
    DISPLAY "toUTC      : ", util.Datetime.toUTC(dt_h2s)    -- uses current date
END MAIN
```

Output:

```
Local time : 12:00:00
Current UTC: 2026-05-05 13:55:05.23072
EXTEND Y2S : 2026-05-05 12:00:00
toUTC      : 10:00:00
```

## Fall/Autumn daylight saving time transition period

> **Important:**
>
> The `toUTC()` function cannot determine if the local
> date/time value represents a time before or after the daylight saving time change, when the
> value is in the hour of the daylight saving time transition period in the fall (this is for example,
> the hour 02:00 PM to 03:00 PM on the last Sunday of October in Europe and first Sunday of November
> in the USA). Depending on the operating system, the `toUTC()` method can interpret
> the local time as summer time or as winter time. In order to get the current system time in UTC, use
> the [`util.Datetime.getCurrentAsUTC()`](3490-util-datetime-getcurrentasutc.md "Returns the current date/time in UTC.")
> method.

The `DATETIME` value passed as parameter to the
`toUTC()` method is the date/time in the local timezone. However, this value does not
contain the GMT offset indicator or daylight saving time information.

When passing local date/time values in the hour of the daylight saving time transition period
in the fall (when clocks roll back one hour), the `toUTC()`
function cannot determine if the local date/time value represents a point in time before or after
the daylight saving time transition occurred. Depending on the operating system, the
`toUTC()` method can interpret the local time as summer time or as a winter time. As
a result, the conversion to the UTC time can be misinterpreted.

For example, in Europe, the fall daylight saving time changes on the 25 of October,
at 3:00 PM. The ambiguous period is between 2:00 PM and 3:00 PM (local time). If you pass, for
example, the date/time value 2015-10-25 02:34:11 to the `toUTC()` method, there is no
way for the method to know if this local time is the time before (CEST / UTC+2h) or after (CET /
UTC+1h) the daylight saving time change.

This behavior can be illustrated with the following code example:

```
IMPORT util

MAIN
    DISPLAY "Original UTC         Local time (Paris)   toUTC(local-time)     ( toUTC() - Orig UTC )"
    CALL test( "2015-10-24 23:59:59" )
    CALL test( "2015-10-25 00:59:59" )
    CALL test( "2015-10-25 01:59:59" )
    CALL test( "2015-10-25 02:59:59" )
END MAIN

FUNCTION test(utc)
    DEFINE utc, loc, utc2 DATETIME YEAR TO SECOND
    LET loc = util.Datetime.toLocalTime(utc)
    LET utc2 = util.Datetime.toUTC(loc)
    DISPLAY SFMT("%1  %2  %3  %4", utc,loc,utc2,utc2-utc)
END FUNCTION
```

The above code will produce the following output on Linux®,
with `TZ='Europe/Paris'`:

```
Original UTC         Local time (Paris)   toUTC(local-time)     ( toUTC() - Orig UTC )
2015-10-24 23:59:59  2015-10-25 01:59:59  2015-10-24 23:59:59          0 00:00:00
2015-10-25 00:59:59  2015-10-25 02:59:59  2015-10-25 00:59:59          0 00:00:00
2015-10-25 01:59:59  2015-10-25 02:59:59  2015-10-25 00:59:59         -0 01:00:00
2015-10-25 02:59:59  2015-10-25 03:59:59  2015-10-25 02:59:59          0 00:00:00
```

As you can see, the local time 2015-10-25 02:59:59 is always converted to UTC 2015-10-25 00:59:59.

## Example

```
IMPORT util
MAIN
    DEFINE utc DATETIME YEAR TO SECOND
    LET utc = util.Datetime.toUTC( DATETIME(2015-08-22 15:34:56) YEAR TO SECOND )
    DISPLAY "UTC: ", utc
END MAIN
```

## Related links

**Related concepts**  

[util.Datetime.toLocalTime](3492-util-datetime-tolocaltime.md "Converts a UTC date/time to the local time.")

[TZ](../07_configuration/0504-tz.md "Defines the timezone for date/time values handling.")
