---
title: "Formatting DATE values"
source: "fgl-topics/c_fgl_DataConversions_format_dates.html"
breadcrumb: "Language basics > Formatting data > Formatting DATE values"
type: "concept"
---

# Formatting DATE values

> Date values must be formatted when converted to strings.

## When does DATE formatting take place?

Date formatting occurs when converting a [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") to a string with
the [`USING`](0634-using.md "The USING operator converts date and numeric values to a string based on a formatting mask.") operator, for example in a
`LET`, `DISPLAY` or `PRINT` instruction, and when
displaying date values in form fields defined with the [`FORMAT`](../11_user-interface/1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.") attribute.

This example formats a `DATE` value with the `USING`
operator:

```
MAIN
  DEFINE d DATE
  LET d = MDY(12,24,2014)
  DISPLAY d USING "mmm ddd yyyy"
END MAIN
```

This
code example produces the following output:

```
Dec Wed 2014
```

Default formatting occurs when `USING` or `FORMAT` are not used,
and a date value has to be converted to a character string, for example when passing a
`DATE` to a function expecting a `VARCHAR(n)`.

Default date formatting is based on the date format defined with the [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") environment variable. For more
details about default formatting, see [Data type conversion reference](0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.").

This topic describes the syntax of the format-string in the
`USING "format-string"` operator and `FORMAT =
"format-string"` form field attribute.

## Formatting symbols for DATE values

When formatting `DATE` values, the format-string of
the `USING` operator or `FORMAT` attribute consists of a set of place
holders that represent date parts as described in the following table:

| Placeholder | Description |
| --- | --- |
| `dd` | Day of the month as a 2-digit integer. |
| `ddd` | Three-letter English-language abbreviation of the day of the week. For example: Mon, Tue. |
| `mm` | Month as a 2-digit integer. |
| `mmm` | Three-letter English-language abbreviation of the month. For example: Jan, Feb. |
| `yy` | Year, as a 2-digits integer representing the 2 trailing digits. |
| `yyy` | Year as a 3-digit number (Ming Guo format only) |
| `yyyy` | Year as a 4-digit number. |
| `c1` | Ming Guo format modifier, see [Using the Ming Guo date format](../09_advanced-features/0891-using-the-ming-guo-date-format.md "Genero BDL can be configured to use the The Ming Guo calendar."). |

Any character different from the date-formatting placeholders is interpreted as a
literal and will appear as-is in the resulting string.

The calendar used for date formatting is the Gregorian calendar. The
`c1` placeholder is a formatting symbol used to adapt the date to the
Ming Guo calendar.

| Format String | Date value | Result string |
| --- | --- | --- |
| `dd/mm/yyyy` | `2011-10-24` | `24/10/2011` |
| `[dd/mm/yy]` | `2011-10-24` | `[24/10/11]` |
| `[ddd-mmm-yyyy]` | `0141-10-24` | `[Tue-Oct-0141]` |
| `(ddd.) mmm. dd, yyyy` | `1999-09-23` | `(Thu.) Sep. 23, 1999` |
| `mmm dd'yy` | `2020-12-05` | `Dec 05'20` |

## Related links

**Related concepts**  

[The util.Date class](../15_library-reference/3482-the-util-date-class.md "The util.Date class provides DATE data-type related utility methods.")

[Using the Ming Guo date format](../09_advanced-features/0891-using-the-ming-guo-date-format.md "Genero BDL can be configured to use the The Ming Guo calendar.")
