---
title: "DBDATE"
source: "fgl-topics/c_fgl_EnvVariables_DBDATE.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > DBDATE"
type: "concept"
---

# DBDATE

> Defines the default display and input format for DATE values.

The DBDATE environment variable defines the default display and input format for
[`DATE`](../08_language-basics/0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") values.

> **Important:**
>
> The DBDATE environment variable is also used by the IBM®
> Informix® database client and server to
> make date to string conversions.

DBDATE defines the order of the month, day, and year time units within a string
representing a date with numeric month and day such as `"24/04/2014"`.

Values of DBDATE must be a restricted combination of symbols representing the
position of the year (Yn), month (M) and day (D), the separator and some optional
configuration options. For example, `DMY4/` defines a date format with the day
unit at the first position, followed by the month and the year (on 4 digits):
"`dd/mm/yyyy`".

The separator always goes at the end of the format string (for example, DMY2/). If no
separator or an invalid character is specified, the slash ( / ) character is the default.
Specifying a `0` (zero) as separator indicates that no separator is used.

The default value of DBDATE depends on the type of platform: On desktop/server platforms, the
default setting for DBDATE is: `MDY4/`. On mobile platforms, DBDATE defaults to
the regional settings defined on the device.

| Symbol | Meaning in DBDATE format string |
| --- | --- |
| `D` | Day of month (can be one or two digits on input) |
| `M` | Month (can be one or two digits on input) |
| `Y2` | Year as two digits (can be one or two digits on input) |
| `Y3` | Year as three digits (Ming Guo format only) |
| `Y4` | Year as four digits (can be 1, 2, 3 or 4 digits on input) |
| `/` | Default time-unit separator for the default locale |
| `C1` | Ming Guo format modifier (years as digits) |
| `-` | Hyphen time-unit separator |
| `.` | Period time-unit separator |
| `0` | Indicates no time-unit separator |

The combinations must follow a specific order:

```
  { DM | MD } { Y2 | Y3 | Y4 } { / | - | . | 0 } [C1]
  { Y2 | Y3 | Y4 } { DM | MD } { / | - | . | 0 } [C1]
```

When a form field and its corresponding variable are defined with the
`DATE` type, values will be displayed depending on the DBDATE format, except if a
[`FORMAT`](../11_user-interface/1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.") attribute is defined.

The DBDATE format is also used to automatically convert a character string
to/from a `DATE` value in programs.

Note that DBDATE takes also effect when fetching `DATE` values
from the database into `CHAR`/`VARCHAR` program variables. However, it
is not recommended to fetch date information into string variables, it is recommended that you use
`DATE` or [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.")
variables instead.

The C1 modifier can be used at the end of the DBDATE value in order to use Ming
Guo date format with digit-based years. When using C1, you can use one of
the Y4, Y3 or Y2 symbols for the year.

A Gregorian date format can look like `"DMY4/"`, while a Ming Guo
date format would look like `"Y3MD/C1"`.

Date formatting specified in a [`USING`](../08_language-basics/0634-using.md "The USING operator converts date and numeric values to a string based on a formatting mask.") clause or `FORMAT` attribute overrides the formatting
specified in DBDATE.

## Related links

**Related concepts**  

[Formatting DATE values](../08_language-basics/0582-formatting-date-values.md "Date values must be formatted when converted to strings.")

[Date, numeric and monetary formats](../09_advanced-features/0890-date-numeric-and-monetary-formats.md "This section describes how Genero BDL handles date, time, numeric and monetary formats.")

[Using the Ming Guo date format](../09_advanced-features/0891-using-the-ming-guo-date-format.md "Genero BDL can be configured to use the The Ming Guo calendar.")
