---
title: "FORMAT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_FORMAT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > FORMAT attribute"
type: "concept"
---

# FORMAT attribute

> The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.

## Syntax

```
FORMAT = "format"
```

1. format is a string of characters that specifies a data format.
2. For details about possible format strings, see [Formatting data](../08_language-basics/0580-formatting-data.md "Explains data to string conversion options of the language.").

## Usage

The `FORMAT` attribute can be specified to define input and display format for
fields bound to numeric and date variables.

- With fields bound to numeric variables such as `MONEY`, `DECIMAL`,
  `FLOAT` and `SMALLFLOAT`, the `FORMAT` attribute
  defines the formatting string with placeholders for the thousands, the decimal separator and the
  front or back currency symbols. The symbols for these elements can then be specified (localized)
  with the [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") (or [DBMONEY](../07_configuration/0512-dbmoney.md "Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined.")) environment variable. When the
  `FORMAT` attribute is not used, the numeric format is defined globally by the
  DBFORMAT (or DBMONEY) environment variable.
- With fields bound to `DATE` variables, the `FORMAT` attribute can
  be used to define the position of the year, month and day, with or without separators. When the
  `FORMAT` attribute is not used, the date format is defined globally by the [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") environment variable.
- The [`FORMAT`](1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.") attribute does not apply to form fields that are bound to a
  `DATETIME` or `INTERVAL` variable.

> **Tip:**
>
> Consider using the `FORMAT` attribute only in specific cases, and use the DBDATE
> and DBFORMAT format settings for most of your date and numeric form fields.

The format specification is used when converting the input buffer to the program variable, and
when displaying program variable data to form fields. For example, when defining a
`FORMAT="yyyy-mm-dd"` for a form field bound to a program variable defined as a
`DATE`, the user can input a date as `2013-12-24`, and the date value
will be displayed in the same manner.

Do not mix `PICTURE` and `FORMAT` attributes in field definitions:
The `PICTURE` attribute is used to define an input mask for character string fields,
such as vehicle registration numbers, phone numbers. Use either `PICTURE`, or use
`FORMAT`.

If the format string is smaller than the field width, you get a compile-time warning, but the
form is usable.

If necessary to satisfy the format specification, number values are rounded before they are
displayed. Consider using the format specification that matches the precision and scale of the
`DECIMAL` or `MONEY` variable bound to the field.

The format string can be any valid string expression using formatting characters as described
in [Formatting numeric values](../08_language-basics/0581-formatting-numeric-values.md "Numeric values must be formatted when converted to strings.") and [Formatting DATE values](../08_language-basics/0582-formatting-date-values.md "Date values must be formatted when converted to strings.").

## Example

```
EDIT f003 = order.totamount, FORMAT = "---,--&.&&@";
EDIT f005 = order.thedate, FORMAT = "mm/dd/yyyy";
```

## Related links

**Related concepts**  

[DATE](../08_language-basics/0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.")

[DECIMAL(p,s)](../08_language-basics/0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.")

[MONEY(p,s)](../08_language-basics/0564-money-p-s.md "The MONEY data type is provided to store currency amounts with exact decimal storage.")
