---
title: "DBCENTURY"
source: "fgl-topics/c_fgl_EnvVariables_DBCENTURY.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > DBCENTURY"
type: "concept"
---

# DBCENTURY

> Specifies the expansion for the century in DATE and DATETIME values.

The DBCENTURY environment variable specifies how to expand abbreviated one- and two-digit
*year* specifications within [`DATE`](../08_language-basics/0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") and [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") values, especially during field input.

> **Important:**
>
> The DBCENTURY environment variable is also used by the IBM® Informix® database client and
> server to make date to string conversions.

Default value is "R" (prefix the entered value with the first two digits of the current year).

Values are case sensitive; only the four uppercase letters are valid.

| Symbol | Algorithm for Expanding Abbreviated Years |
| --- | --- |
| `C` | Use the past, future, or current year closest to the current date. |
| `F` | Use the nearest year in the future to expand the entered value. |
| `P` | Use the nearest year in the past to expand the entered value. |
| `R` | Prefix the entered value with the first two digits of the current year. |

If a year is entered as a single digit, it is first expanded to two digits by prefixing it with
a zero; DBCENTURY then expands this value to four digits.

Three-digit years are not expanded.

Years before 99 AD (or CE) require leading zeros (to avoid expansion).

If the database server and the client system have different settings for DBCENTURY, the client
system setting takes precedence for abbreviations of years in dates entered through the application.
Expansion is sensitive to the time of execution and to the accuracy of the system clock-calendar.
You can avoid the need to rely on DBCENTURY by requiring the user to enter four-digit years or by
setting the [CENTURY](../11_user-interface/1765-century-attribute.md "The CENTURY attribute defines expansion of the year in a DATE or DATETIME field.") attribute in the form
specification of DATE and DATETIME fields.

## Related links

**Related concepts**  

[Formatting DATE values](../08_language-basics/0582-formatting-date-values.md "Date values must be formatted when converted to strings.")

[Date, numeric and monetary formats](../09_advanced-features/0890-date-numeric-and-monetary-formats.md "This section describes how Genero BDL handles date, time, numeric and monetary formats.")
