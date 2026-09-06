---
title: "DBFORMAT"
source: "fgl-topics/c_fgl_EnvVariables_DBFORMAT.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > DBFORMAT"
type: "concept"
---

# DBFORMAT

> Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.

The DBFORMAT environment variable defines the front currency symbol, the thousands separator,
the decimal separator and the back currency symbol, when converting character strings to/from
numeric values.

- When defined, the DBFORMAT environment variable takes precedence over [DBMONEY](0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.").
- The DBFORMAT environment variable is also used by the IBM®
  Informix® database client and server to make date to
  string conversions.
- When using a graphical front-end, the decimal separator of the numeric keypad will produce the
  character defined by the DBFORMAT (or DBMONEY) environment variables.

The value of a DBFORMAT variable must use the following syntax:

```
[front]:[thousands]:decimal:[back]
```

1. front is the leading currency symbol, can be an asterisk ( \* ) to avoid the
   front symbol. The front symbol can be omitted when not used.
2. thousands is a character that you specify as a valid thousands separator, can
   be an asterisk ( \* ) to avoid the thousands separator. The thousands symbol can
   be omitted when not used.
3. decimal is a character that you specify as a valid decimal separator.
   Using the asterisk ( `*` ) will have the same result as when using a dot.
4. back is the trailing currency symbol, can be an asterisk ( \* ) to avoid the
   trailing symbol. The back symbol can be omitted when not used.

If neither DBMONEY, nor DBFORMAT are defined, the default numeric
formatting depends on the type of platform where the runtime system executes:

- On desktop/server platforms, the default numeric format defines the ( `,`
  ) comma as thousands separator, the ( `.` ) dot as decimal separator, and the (
  `$` ) dollar sign as front currency symbol for MONEY values. This corresponds to
  `DBMONEY="$."`, or `DBFORMAT="$:,:.:"`.
- On mobile platforms, the numeric format defaults to the regional settings defined on the
  device. Normally, there is no need to modify these defaults.

DBFORMAT affects the string to/from numeric conversions, by defining the thousands separator,
decimal separator and (front or back) currency symbol, depending on the numeric type:

- [`MONEY`](../08_language-basics/0564-money-p-s.md "The MONEY data type is provided to store currency amounts with exact decimal storage."): Thousands separator
  (when using a format string), decimal separator and currency symbol.
- [`DECIMAL`](../08_language-basics/0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage."), [`SMALLFLOAT`](../08_language-basics/0565-smallfloat.md "The SMALLFLOAT data type stores values as single-precision floating-point binary numbers with up to 8 significant digits."), [`FLOAT`](../08_language-basics/0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.") : Thousands separator (when using a
  format string) and decimal separator.
- [`SMALLINT`](../08_language-basics/0566-smallint.md "The SMALLINT data type is used for storing small whole numbers."), [`INTEGER`](../08_language-basics/0562-integer.md "The INTEGER data type is used for storing large whole numbers."), [`BIGINT`](../08_language-basics/0554-bigint.md "The BIGINT data type is used for storing very large whole numbers."): Thousands separator (when using a
  format string)

DBFORMAT can specify the leading and trailing currency symbols (but not their default
positions within a monetary value) and the decimal and thousands separators. The decimal and
thousands separators defined by DBFORMAT apply to both monetary and other numeric data.

DBFORMAT is typically used in conjunction with following language elements:

- The [`FORMAT`](../11_user-interface/1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.") field attribute
  in form files, to specify a formatting string with placeholders for the symbols defined by
  DBFORMAT.
- The [`USING`](../08_language-basics/0634-using.md "The USING operator converts date and numeric values to a string based on a formatting mask.") operator, to
  specify a formatting string with placeholders for the symbols defined by DBFORMAT.
- When using numeric expressions in [`DISPLAY`](../11_user-interface/1878-display-to-stdout.md "The DISPLAY instruction displays text in line mode to the standard output channel.") and [`PRINT`](../12_reports/2491-print.md "Formats and prints a row of data in a report routine.") instructions, when default formatting applies (when no
  `USING` operator is specified)
- Areas where character string to/from numeric conversion is needed, for example in a [`LET`](../08_language-basics/0701-let.md "The LET statement assigns values to variables.") statement, where string expression is
  assigned a monetary or number variable, or in [`LOAD` and `UNLOAD`](../10_sql-support/1175-sql-load-and-unload.md "Describes the instructions to export/import information from/to a database.") instructions, to convert the numeric
  values in the flat files to pass data to or from the database.

With the default formatting, the thousands separator is not added to the result. To get the
thousands separator, specify a format string including comma (the placeholder for the thousands separator), for example
with `USING "---,---,--&.&&"`.

Specify only one character for decimal or thousands
symbols. The front or back currency symbol can be defined with
more than one character, for example: `DBFORMAT="*:,:.:FR"`.

Any printable character that your locale supports is valid for the
thousands separator or for the decimal separator, except
`0-9` digits, `<`, `>,`, `|`,
`?`, `!`, `=`, `[` and
`]`.

The asterisk ( \* ) specifies that a symbol or separator is not applicable; it is the default for
any front, thousands, or back term that you
do not define, for example with `DBFORMAT="$:*:.:*"`.

The same character cannot be both the thousands and
decimal separator. A blank space (ASCII 32) can be the
thousands separator (and is conventionally used for this purpose in some
locales). The asterisk ( \* ) symbol is valid as the decimal separator, but is not
valid as the thousands separator.

Enclosing the DBFORMAT specification in a pair of single quotation marks is recommended to
prevent the shell from attempting to interpret (or execute) any of the DBFORMAT
characters.

The setting in DBFORMAT affects how formatting masks of the [`FORMAT`](../11_user-interface/1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.") attribute and
`USING` operator are interpreted. In formatting masks of `FORMAT` and
`USING`, these symbols are not literal characters but are placeholders for what
DBFORMAT specifies:

- The dollar ($) sign is a placeholder for the front currency symbol.
- The comma (,) is a placeholder for the thousands separator.
- The period (.) is a placeholder for the decimal separator.
- The at (@) sign is a placeholder for the back currency symbol.

This table illustrates the results of different combinations of DBFORMAT setting and format
string on the same value.

| Value | Format String | DBFORMAT | Result |
| --- | --- | --- | --- |
| 1234.56 | $#,###.## | $:,:.: | $1,234.56 |
| 1234.56 | $#,###.## | :.:,:DM | 1.234,56 |
| 1234.56 | #,###.##@ | $:,:.: | 1,234.56 |
| 1234.56 | #,###.##@ | :.:,:DM | 1.234,56DM |
| 1234.56 | $#,###.##@ | *:*:.:* | 1234.56 |

When the user enters numeric or currency values in fields, the runtime system behaves as
follows:

- If a symbol is entered that was defined as a decimal separator in DBFORMAT, it is
  interpreted as the decimal separator.
- For `MONEY` fields, it disregards any *front* (leading) or *back*(trailing) currency symbol and any thousands separators that the user enters.
- For `DECIMAL` fields, the user must enter values without currency
  symbols.

When the runtime system displays or prints values:

- The DBFORMAT-defined leading or trailing currency symbol is displayed for
  `MONEY` values.
- If a leading or trailing currency symbol is specified by the `FORMAT`
  attribute for non-MONEY data types, the symbol is displayed.
- The thousands separator is not displayed unless it is included in a formatting mask of the
  `FORMAT` attribute or of the USING operator.

When `MONEY` values are converted to character strings by the
`LET` statement, both automatic data type conversion and explicit conversion
with a `USING` clause insert the DBFORMAT-defined separators and currency
symbol into the converted strings.

For example, suppose DBFORMAT is set as follows:

```
*:.:,:SFr
```

The value 1234.56 will print or display as follows:

```
1234,56SFr
```

Here `SFr` stands for the Swiss Franc currency symbol. Values input by the
user into a screen form are expected to contain commas, not periods, as their decimal
separator because DBFORMAT has `*:.:,:SFr` as its setting in
this example.

## Related links

**Related concepts**  

[Formatting DATE values](../08_language-basics/0582-formatting-date-values.md "Date values must be formatted when converted to strings.")

[Date, numeric and monetary formats](../09_advanced-features/0890-date-numeric-and-monetary-formats.md "This section describes how Genero BDL handles date, time, numeric and monetary formats.")

[Type conversions](../08_language-basics/0576-type-conversions.md "Explains primitive data type conversion rules of the language.")

[FORMAT attribute](../11_user-interface/1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.")
