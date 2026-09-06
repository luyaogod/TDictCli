---
title: "DBMONEY"
source: "fgl-topics/c_fgl_EnvVariables_DBMONEY.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > DBMONEY"
type: "concept"
---

# DBMONEY

> Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined.

The DBMONEY environment variable defines the currency symbol and the decimal separator when
converting character strings to/from numeric values.

- When defined, the [DBFORMAT](0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") environment
  variable takes precedence over DBMONEY.
- The DBMONEY environment variable is also used by the IBM®
  Informix® database client and server to make date to
  string conversions.

The value of a DBMONEY variable must use the following syntax:

```
  {.|,}
| front{.|,}
| {.|,}back
```

1. front is a character string representing a leading currency
   symbol that precedes the value.
2. back is a character string representing a trailing currency
   symbol that follows the value.
3. The decimal separator is mandatory and can only be a dot ( `.` ) or a comma (
   `,` )

If neither DBMONEY, nor DBFORMAT are defined, the default numeric
formatting depends on the type of platform where the runtime system executes:

- On desktop/server platforms, the default numeric format defines the ( `,`
  ) comma as thousands separator, the ( `.` ) dot as decimal separator, and the (
  `$` ) dollar sign as front currency symbol for MONEY values. This corresponds to
  `DBMONEY="$."`, or `DBFORMAT="$:,:.:"`.
- On mobile platforms, the numeric format defaults to the regional settings defined on the
  device. Normally, there is no need to modify these defaults.

DBMONEY can only define the currency symbol and decimal separator characters. To define
the thousands separator, use the DBFORMAT environment variable instead. However, if only DBMONEY is
used, an implicit thousands separator is selected.

The currency symbol in DBMONEY can be up to seven characters long and can contain any
character except a comma or a period. It can be non-ASCII characters if the current
locale supports a code set that defines the non-ASCII characters you use.

DBMONEY is used for implicit data conversion between numeric values and character
strings, for example when using numeric values in form fields and reports. Other areas regarding
string to/from numeric conversion are affected by the DBMONEY. For more details, see [DBFORMAT](0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.").

The position of the currency symbol (relative to the decimal separator) indicates whether
the currency symbol appears before or after the `MONEY` value. When the
currency symbol is positioned in DBMONEY before the decimal separator, it is
displayed before the value ($1234.56). When it is positioned after the decimal
separator, it is displayed after the value (1234.56F).

The runtime system recognizes the period ( . ) and the comma ( , ) as decimal separators.
All other characters are considered to be part of the currency symbol. For example, ",
FR" defines a MONEY format with the comma as decimal separator and the string " FR"
(including the space) as the currency symbol.

Because only its position within a DBMONEY setting indicates whether a symbol is the
front or back currency symbol, the decimal separator is
required. If you use DBMONEY to specify a back symbol, for example, you must
supply a decimal separator (a comma or period). Similarly, if you use DBMONEY to change the decimal
separator from a period to a comma, you must also supply a currency symbol.

To avoid ambiguity in displayed numbers and currency values, do not use the thousands
separator of DBFORMAT as the decimal separator of DBMONEY. For example, specifying comma
as the DBFORMAT thousands separator dictates using the period as the DBMONEY decimal
separator.

When using a graphical front-end, the decimal separator of the numeric keypad will
produce the character defined by this environment variable.

## Related links

**Related concepts**  

[DBFORMAT](0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.")
