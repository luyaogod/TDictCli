---
title: "Formatting numeric values"
source: "fgl-topics/c_fgl_DataConversions_format_numbers.html"
breadcrumb: "Language basics > Formatting data > Formatting numeric values"
type: "concept"
---

# Formatting numeric values

> Numeric values must be formatted when converted to strings.

## When does numeric formatting take place?

Numeric formatting occurs when converting a number to a string with the [`USING`](0634-using.md "The USING operator converts date and numeric values to a string based on a formatting mask.") operator, for example in a
`LET`, `DISPLAY` or `PRINT` instruction, and when
displaying numeric values in form fields defined with the [`FORMAT`](../11_user-interface/1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.") attribute.

Numeric values can be of type such as [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers."), [`FLOAT`](0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits."), [`DECIMAL`](0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage."), [`MONEY`](0564-money-p-s.md "The MONEY data type is provided to store currency amounts with exact decimal storage."), etc.

This example formats a `DECIMAL(10,2)` value with the `USING`
operator:

```
MAIN
  DEFINE d DECIMAL(10,2)
  LET d = -123456.78
  DISPLAY d USING "-,---,--&.&& @"
END MAIN
```

Front currency symbol, thousands separator, decimal separator and back currency symbol are
defined with the [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") (or [DBMONEY](../07_configuration/0512-dbmoney.md "Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined.")) environment variable. For example, with
DBFORMAT defined as `":.:,:E"`, there is no front currency symbol defined, the
thousands separator is a dot, the decimal separator is a comma and the back currency symbol is
`E`. With this setting, the previous code example will produce the following
output:

```
 -123,456.78 E
```

Default formatting occurs when `USING` or `FORMAT` are not used,
and a numeric value has to be converted to a character string, for example when passing a
`DECIMAL(p,s)` to a function expecting a `VARCHAR(n)`. For more
details about default formatting, see [Data type conversion reference](0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.").

This topic describes the syntax of the format-string in the `USING
"format-string"` operator and `FORMAT =
"format-string"` form field attribute.

## Formatting symbols for numbers

When formatting numeric values, the format-string of the
`USING` operator or `FORMAT` attribute consists of a set of place
holders that represent digits, currency symbols, thousands and decimal separators. For example,
`"###.##@"` defines three places to the left of the decimal point and exactly two to
the right, plus a "back" currency symbol at the end of the string.

The `USING` operator or `FORMAT` attribute are required to display
the thousands separator defined in DBFORMAT.

The format-string must use normalized placeholders described in Table 1. The placeholders will be
replaced by digits, blanks or by the elements defined in the DBFORMAT (or DBMONEY) environment
variables. Any other character will be interpreted as a literal, and can be used at any place in the
format string.

If the numeric value is too large to fit in the number of characters defined by the format, the
result string is filled with a set of star characters (`********`), indicating an
overflow. Note that this includes the `+/-` sign.

If the numeric values can be negative numbers, specify one additional `-` (minus)
placeholder for the sign. Otherwise, if the value is negative and not enough placeholders are
defined, overflow star characters will be displayed.

The minus sign ( `-` ), plus sign ( `+` ), parentheses ( `(
)` ), and dollar sign ( `$` ) float. This means that when you specify
multiple leading occurrences of one of these characters, the result string gets only a single
character immediately to the left of the first digit.

| Placeholder | Description |
| --- | --- |
| `*` | On the left of the decimal separator, the star `*` placeholder fills with asterisks any position that would otherwise be blank. When `*` is used on the right of the decimal separator, a zero digit is displayed instead of blank. |
| `&` | The ampersand `&` placeholder is used to define the position of a digit, and is replaced by a zero if that position would otherwise be blank. |
| `#` | On the left of the decimal separator, the hash `#` placeholder is replaced by a blank, if no digit is to be displayed at that position. On the right of the decimal separator, a zero digit is displayed instead of blank. |
| `<` | The "less than" `<` placeholder aligns the numeric value to the left: Consecutive "less than" characters cause left alignment and define digit positions. |
| `-` | The minus `-` placeholder displays a minus sign if the value is negative, or a blank if the value is positive. When grouping several minus signs in the format string, a single minus sign floats immediately to the left of the first digit. When `-` is used on the right of the decimal separator, a zero digit is displayed instead of blank or minus sign. |
| `+` | The plus `+` placeholder displays a minus sign if the value is negative, or a plus sign if the value is positive. When grouping several plus signs in the format string, a single plus sign floats immediately to the left of the first digit. When `+` is used on the right of the decimal separator, a zero digit is displayed instead of blank or plus/minus sign. |
| `(` | The opening parenthesis `(` is displayed as left parenthesis for negative numbers. It is used to display "accounting parentheses" instead of a minus sign for negative numbers. Consecutive left parentheses display a single left parenthesis to the left of the first digit. |
| `)` | The closing parenthesis `)` is displayed as right parenthesis for negative numbers. This placeholder is used in conjunction with a opening parenthesis to display "accounting parentheses" for negative numbers. |
| `,` (comma) | The comma placeholder is used to define the position for the thousand separator defined in DBFORMAT. The thousand separator will only be displayed if there is a number on the left of it. |
| `.` (period) | The period placeholder is used to define the position for the decimal separator defined in DBFORMAT. You can only have one decimal separator in a number format string. |
| `$` | The dollar sign is the placeholder for the front currency symbol defined in DBFORMAT. When you group several consecutive dollar signs, a single front currency symbol floats immediately to the left of the first digit. The front currency symbol can be defined in DBFORMAT with more than one character (EUR, USD). |
| `@` | The "at" sign is the placeholder for the back currency symbol defined in DBFORMAT. Put several consecutive @ signs at the end of the format string to display a currency symbol defined in DBFORMAT with more than one character. |

| Format String | Value | DBFORMAT | Result string |
| --- | --- | --- | --- |
| `[######.##]` | `0` | `:.:,:` | `[______,00]` |
| `[######.##]` | `-1234.56` | `:.:,:` | `[__1234,56]` (no sign!) |
| `[######.##]` | `-1234567.89` | `:.:,:` | `[*********]` (overflow) |
| `[######.##]` | `+1234.56` | `:.:,:` | `[__1234,56]` |
| `[#####&.&&]` | `0` | `:.:,:` | `[_____0,00]` |
| `[******.**]` | `0` | `:.:,:` | `[******,00]` |
| `[******.**]` | `-12.34` | `:.:,:` | `[****12,34]` (no sign!) |
| `[******.**]` | `+12.34` | `:.:,:` | `[****12,34]` |
| `[<<<<<<.<<]` | `-12.34` | `:.:,:` | `[12,34]` (no sign!) |
| `[<<<<<<.<<]` | `+12.34` | `:.:,:` | `[12,34]` |
| `[---,--&.&&]` | `-1234.56` | `:.:,:` | `[_-1.234,56]` |
| `[+++,++&.&&]` | `-1234.56` | `:.:,:` | `[_-1.234,56]` |
| `[+++,++&.&&]` | `+1234.56` | `:.:,:` | `[_+1.234,56]` |
| `[$---,--&.&&]` | `-1234.56` | `E:.:,:` | `[E_-1.234,56]` |
| `[$---,--&.&&]` | `+1234.56` | `E:.:,:` | `[E__1.234,56]` |
| `[$$$---,--&.&&]` | `+1234.56` | `E:.:,:` | `[_E___1.234,56]` |
| `[$$$---,--&.&&]` | `+1234.56` | `EUR:.:,:` | `[EUR__1.234,56]` |
| `[-,---,-$&.&&]` | `-12.34` | `E:.:,:` | `[_____-E12,34]` |
| `[-,---,-$&.&&]` | `-1234.56` | `E:.:,:` | `[__-E1.234,56]` |
| `[-,-$$,$$&.&&]` | `-12.34` | `E:.:,:` | `[__-___E12,34]` |
| `[-,-$$,$$&.&&]` | `-1234.56` | `E:.:,:` | `[__-E1.234,56]` |
| `[---,--&.&&@]` | `-1234.56` | `:.:,:E` | `[_-1.234,56E]` |
| `[---,--&.&&@]` | `+1234.56` | `:.:,:E` | `[__1.234,56E]` |
| `[---,--&.&&@@@]` | `+1234.56` | `:.:,:EUR` | `[__1.234,56EUR]` |
| `[($---,--&.&&)]` | `-1234.56` | `E:.:,:` | `[(E_-1.234,56)]` |
| `[($###,##&.&&)]` | `-1234.56` | `E:.:,:` | `[(E__1.234,56)]` (no sign!) |
| `[((((,(($.&&)]` | `0` | `E:.:,:` | `[_______E,00_]` |
| `[((((,(($.&&)]` | `-12.34` | `E:.:,:` | `[____(E12,34)]` (no sign!) |
| `[((((,(($.&&)]` | `+12.34` | `E:.:,:` | `[_____E12,34_]` |
| `[((((,(($.&&)]` | `-1234.56` | `E:.:,:` | `[_(E1.234,56)]` (no sign!) |
| `[((((,(($.&&)]` | `+1234.56` | `E:.:,:` | `[__E1.234,56_]` |

> **Note:**
>
> In the first and last columns, the strings are shown with
> opening and closing square brackets. These brackets mark the start and the end of the string. In the
> result column, underscore characters are used to represent spaces. This is important when examining
> the result string, as it shows where spaces exist.

## Related links

**Related concepts**  

[Formatting DATE values](0582-formatting-date-values.md "Date values must be formatted when converted to strings.")

[Formatting DATETIME values](0583-formatting-datetime-values.md "Date-time values must be formatted when converted to strings.")

[Primitive Data types](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")
