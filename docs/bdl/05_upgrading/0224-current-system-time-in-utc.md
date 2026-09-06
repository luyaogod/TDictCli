---
title: "Current system time in UTC"
source: "fgl-topics/c_fgl_Migrate_to_300_current_utc.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > Current system time in UTC"
type: "concept"
---

# Current system time in UTC

> Use the util.Datetime.getCurrentAsUTC() method to get the current system date/time in UTC.

Starting with Genero 3.00, the [util.Datetime.getCurrentAsUTC()](../15_library-reference/3490-util-datetime-getcurrentasutc.md "Returns the current date/time in UTC.")
method is provided to get the current system time in UTC (Coordinated Universal Time).

This method has been added to solve the issue when using `util.Datetime.toUTC(CURRENT)`
during the daylight saving time transition period in the fall, as described in [util.Datetime.toUTC](../15_library-reference/3494-util-datetime-toutc.md "Converts a date/time value to the UTC date/time.").
