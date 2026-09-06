---
title: "util.Datetime methods"
source: "fgl-topics/c_fgl_ext_util_Datetime_methods.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Datetime class > util.Datetime methods"
type: "concept"
---

# util.Datetime methods

> Methods for the util.Datetime class.

| Name | Description |
| --- | --- |
| util.Datetime.format( t DATETIME q1 TO q2, format STRING ) RETURNS STRING | Formats a date/time value based on a specified format. |
| util.Datetime.fromSecondsSinceEpoch( t FLOAT ) RETURNS DATETIME q1 TO q2 | Converts a number of seconds since Epoch to a date/time. |
| util.Datetime.getCurrentAsUTC( ) RETURNS DATETIME YEAR TO FRACTION(5) | Returns the current date/time in UTC. |
| util.Datetime.parse( s STRING, format STRING ) RETURNS DATETIME q1 TO q2 | Converts a string to a `DATETIME` value based on a specified format. |
| util.Datetime.toLocalTime( t DATETIME q1 TO q2 ) RETURNS DATETIME q1 TO q2 | Converts a UTC date/time to the local time. |
| util.Datetime.toSecondsSinceEpoch( t DATETIME q1 TO q2 ) RETURNS FLOAT | Converts a date/time to a number of seconds since Epoch. |
| util.Datetime.toUTC( t DATETIME q1 TO q2 ) RETURNS DATETIME q1 TO q2 | Converts a date/time value to the UTC date/time. |
