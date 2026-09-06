---
title: "Using the Ming Guo date format"
source: "fgl-topics/c_fgl_localization_028.html"
breadcrumb: "Advanced features > Localization > Application locale > Using the Ming Guo date format"
type: "concept"
---

# Using the Ming Guo date format

> Genero BDL can be configured to use the The Ming Guo calendar.

The Ming Guo (or Minguo) calendar is still used in some Asian regions like Taiwan. This calendar
is equivalent to the Gregorian calendar, except that the years are numbered with a different base.
In the Ming Guo calendar, the first year (1) corresponds to the Gregorian year 1912, the year the
Republic of China was founded.

Digit-based year Ming Guo date format can be enabled by adding the C1 modifier at the end of the
value set for the DBDATE environment variable:

```
$ DBDATE="Y3MD/C1"
$ export DBDATE
```

With this [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") setting, dates will be displayed
with a year following the Ming Guo calendar, and date input will also be interpreted based on that
calendar. For example, if the user enters 90/3/24, it is equivalent to an input of 2002/3/24 when
using the Gregorian calendar. Basically, the runtime system will subtract 1912 or add 1912 respectively
when displaying or reading date values).

When using the C1 modifier, the possible values for the Yn symbol are Y4, Y3, Y2.

The [`MDY()`](../08_language-basics/0668-mdy-function.md "The MDY() operator creates a DATE from month, day and year units.") operator is sensitive to
the C1 modifier usage in DBDATE. For example, if DBDATE=Y3MD/C1, MDY(3,24,1) will build a date the
corresponds in the Gregorian to MDY(3,24,1912).

The [`USING`](../08_language-basics/0634-using.md "The USING operator converts date and numeric values to a string based on a formatting mask.") operator supports the
c1 modifier as well. The c1 modifier must be specified at the end of the format. You can for example
use the following format string: "yyyy-mm-ddc1".

The C2 modifier to use Era names is not supported.

Unlike Informix® 4gl, when using negative years, the minus
sign is placed over the left-most zero of the year, to avoid miss-aligned dates.

For example, if DBDATE=Y3MD/C1:

```
MDY(3,2, 1) USING "yyy/mm/ddc1" 
MDY(3,2,-1) USING "yyy/mm/ddc1"
```

Will align properly as follows:

```
0001/03/02
-001/03/02
```

Front-ends may not support the Ming Guo calendar for widgets like [`DATEEDIT`](../11_user-interface/1688-dateedit-item-type.md "Defines a line-edit with a calendar widget to pick a date.").

## Related links

**Related concepts**  

[Date, numeric and monetary formats](0890-date-numeric-and-monetary-formats.md "This section describes how Genero BDL handles date, time, numeric and monetary formats.")
