---
title: "Data type conversion reference"
source: "fgl-topics/c_fgl_DataConversions_003.html"
breadcrumb: "Language basics > Type conversions > Data type conversion reference"
type: "concept"
---

# Data type conversion reference

> This topic lists type conversion rules for all data types.

## Controlling type conversion errors

Type conversion errors can be controlled with the `WHENEVER ANY ERROR` clause, or
with `TRY/CATCH` blocks. For more details, see [Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

## Boolean type conversions

A `BOOLEAN` value is an integer value 1 or 0 and thus can be converted to/from
any other numeric type of the language.

When converting a numeric value to a `BOOLEAN`, any value different from 0
becomes `TRUE`, otherwise (zero) is `FALSE`.

```
DEFINE hasContent BOOLEAN, s STRING
LET s = "abc"
LET hasContent = s.getLength()
```

When converting a string (`CHAR`, `VARCHAR` or `STRING`)
to `BOOLEAN`, the string will be converted to a number first, then the number-to-boolean
conversion applies. If the string value cannot convert to a numeric value (for example, "abc"), the
boolean value becomes `NULL`.

When converting a `BOOLEAN` to a string, the result will be `"1"` or
`"0"` string values, depending on the boolean value.

## Large object (TEXT/BYTE) type conversions

A `BYTE` value cannot be converted to/from any other data type. However, you can
use APIs such as `util.JSON.stringify()`/`util.JSON.parse()` to
serialize/deserizalize a `BYTE` value:

```
IMPORT util
MAIN
    DEFINE by BYTE
    LOCATE by IN MEMORY
    CALL util.JSON.parse('"bGlnaHQgdw=="',by)
    DISPLAY util.JSON.stringify(by)
END MAIN
```

Output:

```
"bGlnaHQgdw=="
```

A `TEXT` value can be converted to/from `CHAR`,
`VARCHAR` or `STRING`:

```
MAIN
    DEFINE s STRING
    DEFINE tx TEXT
    LOCATE tx IN MEMORY
    LET tx = "abc\ndef"
    DISPLAY "tx = ", tx
    LET s = tx
    DISPLAY "s  = ", s
    LET tx = s
    DISPLAY "tx = ", tx
END MAIN
```

Output:

```
tx = abc
def
s  = abc
def
tx = abc
def
```

## Character string to character string types

When assigning a `CHAR/VARCHAR` value to a `CHAR/VARCHAR` variable,
the trailing characters are silently truncated, if the receiving type is not large enouth.

Not runtime error can be detected, even with `WHENEVER ANY ERROR` or
`TRY/CATCH`.

```
MAIN
  DEFINE c6 CHAR(6)
  DEFINE c3 CHAR(3)
  WHENEVER ANY ERROR CONTINUE
  LET c6 = "abcdef"
  LET c3 = c6
  DISPLAY c3, " : ", status
END MAIN
```

Output:

```
abc :           0
```

## Integers to integers types

Converting between `TINYINT`, `SMALLINT`, `INTEGER`
and `BIGINT` is possible, as long as the receiving type can hold the whole number to
be assigned.

If the integer value exceeds the range of the receiving data type, the conversion fails.

```
MAIN
  DEFINE int4 INTEGER
  DEFINE int2 SMALLINT
  WHENEVER ANY ERROR CONTINUE
  LET int4 = 32767
  LET int2 = int4
  DISPLAY int2, " : ", status
  LET int4 = 32768
  LET int2 = int4
  DISPLAY int2, " : ", status
END MAIN
```

Output:

```
 32767 :           0
       :       -1214
```

## Integers to decimal types

`TINYINT`, `SMALLINT`, `INTEGER` and `BIGINT`
values can be converted to `SMALLFLOAT`, `FLOAT`, `DECIMAL` or
`MONEY` as long as the decimal type is defined with sufficient digits to hold the whole
number.

If the integer value exceeds the range of the receiving data type, the conversion fails.

```
MAIN
  DEFINE int4 INTEGER
  DEFINE d051 DECIMAL(5,1)
  WHENEVER ANY ERROR CONTINUE
  LET int4 = 999
  LET d051 = int4
  DISPLAY d051, " : ", status
  LET int4 = 999999
  LET d051 = int4
  DISPLAY d051, " : ", status
END MAIN
```

Output:

```
  999.0 :           0
        :       -1226
```

## Decimal to integer types

When converting a `SMALLFLOAT`, `FLOAT` `DECIMAL` or
`MONEY` to a `TINYINT`, `SMALLINT`, `INTEGER`
or `BIGINT`, the fractional part of the decimal value is truncated.

If the decimal value exceeds the range of the receiving integer data type, the conversion
fails.

```
MAIN
  DEFINE d102 DECIMAL(10,2)
  DEFINE int2 SMALLINT
  WHENEVER ANY ERROR CONTINUE
  LET d102 = 32767.99
  LET int2 = d102
  DISPLAY int2, " : ", status
  LET d102 = 32768.0
  LET int2 = d102
  DISPLAY int2, " : ", status
END MAIN
```

Output:

```
 32767 :           0
       :       -1214
```

## Decimal to decimal types

Converting between `SMALLFLOAT`, `FLOAT`, `DECIMAL`
or `MONEY` types is allowed as long as the receiving type is defined with sufficient
digits to hold the whole part of the original value.

If the original value contains more fractional digits than the receiving data type supports,
low-order digits are discarded.

If the decimal value exceeds the range of the receiving decimal data type, the conversion
fails.

```
MAIN
  DEFINE d102 DECIMAL(10,2)
  DEFINE d051 DECIMAL(5,1)
  WHENEVER ANY ERROR CONTINUE
  LET d102 = 123.45
  LET d051 = d102
  DISPLAY d051, " : ", status
  LET d102 = 99999.99
  LET d051 = d102
  DISPLAY d051, " : ", status
END MAIN
```

Output:

```
  123.5 :           0
        :       -1226
```

## Integers to character string types

Converting `TINYINT`, `SMALLINT`, `INTEGER` or
`BIGINT` values to `CHAR`, `VARCHAR` and
`STRING` is allowed as long as the receiving character variable is large enough.

If the receiving character type is to small to hold the integer value, it is filled with
`*` overflow stars, and the conversion fails.

```
MAIN
  DEFINE int4 INTEGER
  DEFINE vc05 VARCHAR(5)
  WHENEVER ANY ERROR CONTINUE
  LET int4 = -123
  LET vc05 = int4
  DISPLAY vc05, " : ", status
  LET int4 = 999999
  LET vc05 = int4
  DISPLAY vc05, " : ", status
END MAIN
```

Output:

```
-123 :           0
***** :       -1207
```

## Decimal to character string types

Converting `SMALLFLOAT`, `FLOAT`, `DECIMAL` or
`MONEY` values to `CHAR`, `VARCHAR` and `STRING`
implies numeric formatting.

Numeric formatting is controlled by the [DBMONEY](../07_configuration/0512-dbmoney.md "Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined.") and [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") environment variables. Especially, the
decimal separator is defined by these environment variables.

> **Tip:**
>
> To convert numeric data to character strings in ISO format independently from
> DBMONEY/DBFORMAT, use the [`util.JSON.stringify()`](../15_library-reference/3586-util-json-stringify.md "Produces a JSON formatted string from the input, by including empty records and empty arrays.") method:
>
> ```
> IMPORT util
> FUNCTION dec2iso(d DECIMAL)
>     RETURN util.JSON.stringify(d)
> END FUNCTION
> ```

The resulting string is left-aligned (for lossless conversions) or right-aligned (for visual conversions),
depending on the conversion context; and the decimal part is kept depending on the numeric type.

```
MAIN
  DEFINE m MONEY(8,2),
         s VARCHAR(10)
  LET m = 123.45
  LET s = m   -- Lossless conversion "$123.45"
  DISPLAY m   -- Visual conversion   "    $123.45"
END MAIN
```

Fixed point decimals (`DECIMAL(p,s)`) are converted to strings that can fit in a
`CHAR(p+2)`: The string is built with up to p significant digits + 1 character for the
sign + 1 character for the decimal point. The result of a `DECIMAL(p,s)` to string
conversion is never longer than p + 2 characters. For example, a `DECIMAL(5,2)` can
produce "-999.99" (5 + 2 = 7c).

Floating point decimals (`DECIMAL(p)`) are converted to strings that can fit
in a `CHAR(p+7)`: The string is built with up to p significant digits + 1 character
for the sign + 1 character for the decimal point + the length of the exponent of needed ("e+123").
The result of a `DECIMAL(p)` to string conversion is never longer than p + 7. For
example, a `DECIMAL(5)` can produce "-1.2345e-123" (5 + 7 = 12c).

`DECIMAL` to string conversion depends on the context in which the conversion
occurs:

1. Visual conversion: The result of this conversion will typically be presented to
   the end user. This conversion happens in `DISPLAY`, `MESSAGE`,
   `ERROR`, `PRINT`. The result of a visual conversion is right aligned
   (padded with leading blanks). This padding results in the same length for any value for a given
   decimal precision. The length of the result is the maximum possible length as described previously
   (p+2 for `DECIMAL(p,s)`, p+7 for `DECIMAL(p)`).

   Visual conversion
   examples for `DECIMAL(5,2)`:

   ```
      Values     1234567
   ----------------------
      0       | "   0.00"
      -999.99 | "-999.99"
      12.3    | "  12.30"
      12.34   | "  12.34"
   ```

   Visual conversion examples for `DECIMAL(5)`:

   ```
      Values     123456789012
   ---------------------------
      0       | "         0.0"
      -99999  | "    -99999.0"
      12.3    | "        12.3"
      12.34   | "       12.34"
      12.345  | "      12.345"
      1.23e7  | "  12300000.0"
      1e100   | "       1e100"
   ```
2. Form field conversion: This conversion concerns decimal numbers presented in
   form-fields. The result of this conversion is in best case the same as (1). The result of the
   conversion depends on the width of the form-field. If the width of the form-field is smaller than
   the perfect length, automatic rounding and exponential notation might be used.
3. Lossless conversion: Such conversion happens when assigning numbers to string
   variables (`LET`), passing numbers as parameters to functions expecting strings,
   returning numbers from functions to strings, serializing numbers (`UNLOAD`, XML or
   JSON APIs). These conversions must avoid the loss of significant digits. When using floating point
   decimals, this leads to a variable length of the resulting string. A conversion must be reversible:
   decimal to string to decimal must give the original value. If the target variable is shorter then
   the maximum possible length, then automatic rounding will occur.

   Lossless conversion examples of
   `DECIMAL(5,2)`:

   ```
      Values     1234567
   ----------------------
      0       | "0.00"
      -999.99 | "-999.99"
      12.3    | "12.30"
      12.34   | "12.34"
   ```

   Lossless conversion examples of `DECIMAL(5)`:

   ```
      Values     123456789012
   ---------------------------
      0       | "0.0"
      -99999  | "-99999.0"
      12.3    | "12.3"
      12.34   | "12.34"
      12.345  | "12.345"
      1.23e7  | "12300000.0"
      1e100   | "1e100"
   ```

Automatic rounding occurs if the target string variable is shorter than the maximum possible
length of the `DECIMAL` type. Such conversion might loose significant digits:
The runtime system tries to round the value, to fit into the target variable.

```
  Values     Different target sizes
              123456   12345   1234
------------------------------------
  0.98765  | "0.9877" "0.988" "0.99"
  123.45   | "123.45" "123.5" "123"
```

Automatic switch to the exponential notation will occur if the integer part of the decimal
value does not fit into the target string variable. For example, if the source variable is a
`DECIMAL(12)` and the target variable is a `CHAR(9)`:

```
  Values         123456789
----------------------------
  1234567     | "1234567.0"
  12345678    | "12345678"
  123456789   | "123456789"
  1234567890  | "1.2346e10"
  12345678901 | "1.2346e11"
```

The exponential notation will also be used if the absolute value of a floating point decimal
is less than 1e-8 (0.00000001).

Default formatting of floating point `DECIMAL(P)` has been [revised with Genero 2.50](../05_upgrading/0238-floating-point-to-string-conversion.md "The default formatting of a DECIMAL(P), SMALLFLOAT and FLOAT adapts to the significant digits of the value."). If
`DECIMAL(P)`-to-string conversion must round to 2 digits, use the
`fglrun.decToCharScale2` FGLPROFILE entry (this applies to all contexts):

```
fglrun.decToCharScale2 = true
```

Another FGLPROFILE entry can be used, to get the 2-digit rounding of `DECIMAL(P)`
only in the context of the `PRINT` statement in reports. (the
`fglrun.decToCharScale2` and `fglrun.decToCharScale2.print` parameters
are mutually exclusive):

```
fglrun.decToCharScale2.print = true
```

> **Note:**
>
> Do not use the `fglrun.decToCharScale2*` configuration parameters, unless
> you have migration issues.

Formatting a `FLOAT` is the same as `DECIMAL(16)`. Any
`FLOAT` value with up to 15 digits is exact. There is no precision loss
when converting an exact `FLOAT` back and forth to/form a string. Some
`FLOAT` values require 16, in some rare cases 17 digits for an exact string
representation. 16 and 17 digits are not always exact: `"8.000000000000001"`
and `"8.000000000000002"` represent the same float value.

Formatting a `SMALLFLOAT` is the same as `DECIMAL(7)`.
Any `SMALLFLOAT` value with up to 6 digits is exact. There is no precision
loss when converting an exact `SMALLFLOAT` back and forth to/form a string.
Some `SMALLFLOAT` values require 7, in some rare cases 8 digits for an exact
string representation. 7 and 8 digits `SMALLFLOAT` are not always exact:
`"0.0009999901"` and `"0.0009999902"` represent the same
`SMALLFLOAT` value.

If `FLOAT/SMALLFLOAT`-to-string conversion must round to 2 digits, use the
`fglrun.floatToCharScale2` FGLPROFILE entry (this applies to all
contexts):

```
fglrun.floatToCharScale2 = true
```

Another FGLPROFILE entry can be used, to get the 2-digit rounding of
`FLOAT/SMALLFLOAT` only in the context of the `PRINT` statement
in reports. (the `fglrun.floatToCharScale2` and
`fglrun.floatToCharScale2.print` parameters are mutually exclusive):

```
fglrun.floatToCharScale2.print = true
```

> **Note:**
>
> Do not use the `fglrun.floatToCharScale2*` configuration parameters,
> unless you have migration issues.

## Character string to numeric types

A `CHAR`, `VARCHAR` and `STRING` value can be
converted to a `TINYINT`, `SMALLINT`, `INTEGER`,
`BIGINT`, `SMALLFLOAT`, `FLOAT`,
`DECIMAL` or `MONEY` value as long as the character string value
represents a valid number.

Numeric formatting is controlled by the [DBMONEY](../07_configuration/0512-dbmoney.md "Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined.") and [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") environment variables. Especially, the
decimal separator is defined by these environment variables.

> **Tip:**
>
> To convert character strings representing numbers in ISO format to numeric data
> independently from DBMONEY/DBFORMAT, use the [`util.JSON.parse()`](../15_library-reference/3582-util-json-parse.md "Parses a JSON string and fills program variables with the values.") method:
>
> ```
> IMPORT util
> FUNCTION iso2dec(s STRING)
>     DEFINE d DECIMAL
>     CALL util.JSON.parse(s,d)
>     RETURN d
> END FUNCTION
> ```

If the original value contains more significant digits or more fractional digits than the
receiving data type supports, low-order digits are discarded.

Example using a `DECIMAL`:

```
MAIN
  DEFINE d102 DECIMAL(10,2)
  WHENEVER ANY ERROR CONTINUE
  LET d102 = "-123.45"
  DISPLAY d102, " : ", status
  LET d102 = "1234567890123.45"
  DISPLAY d102, " : ", status
  LET d102 = "12345678.999"
  DISPLAY d102, " : ", status
END MAIN
```

Output:

```
     -123.45 :           0
             :       -1226
 12345679.00 :           0
```

Example using an `INTEGER`:

```
MAIN
  DEFINE int4 INTEGER
  WHENEVER ANY ERROR CONTINUE
  LET int4 = "-12345"
  DISPLAY int4, " : ", status
  LET int4 = "999999999999"
  DISPLAY int4, " : ", status
END MAIN
```

Output:

```
     -12345 :           0
            :       -1215
```

## Date time to character string types

Converting `DATE`, `DATETIME` and `INTERVAL` values
to `CHAR`, `VARCHAR` and `STRING` implies date time formatting.

- `DATE` formatting is controlled by the [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") environment variable.
- When converting an `INTERVAL` to a string, either `YYYY-MM` or
  `DD hh:mm:ss.fffff` standard formats are used, depending on the interval class.
- When converting a `DATETIME` to a string, the `YYYY-MM-DD
  hh:mm:ss.fffff` standard format is used.

If the conversion result is longer than the receiving variable, the resulting character string is
filled with `*` overflow stars.

```
MAIN
  DEFINE da DATE
  DEFINE vc20 VARCHAR(20)
  DEFINE vc05 VARCHAR(5)
  WHENEVER ANY ERROR CONTINUE
  LET da = MDY(12,24,2012)
  LET vc20 = da
  DISPLAY vc20, " : ", status
  LET vc05 = da
  DISPLAY vc05, " : ", status
END MAIN
```

Output:

```
12/24/2012 :           0
***** :           0
```

See also [Formatting DATE values](0582-formatting-date-values.md "Date values must be formatted when converted to strings."), [Formatting DATETIME values](0583-formatting-datetime-values.md "Date-time values must be formatted when converted to strings.") and [Formatting INTERVAL values](0584-formatting-interval-values.md "Interval values must be formatted when converted to strings.").

## Character string to date time types

Converting a `CHAR`, `VARCHAR` or `STRING` value to
a `DATE`, `DATETIME` or `INTERVAL` is possible as long
as the character string defines a well formatted date time or interval value.

- When converting a character string to a `DATE`, the string must follow the date
  format defined by the [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") environment
  variable.
- When converting a string to an `INTERVAL`, depending on the interval class of the
  target variable, the source string must be formatted as `[-]YYYY-MM` or `[-]DD
  hh:mm:ss.fffff`. In addition to the legacy `INTERVAL` format, the ISO 8601
  duration format is also supported:
  - For `INTERVAL` of class year-month, the ISO 8601 string must be formatted as
    `[-]nnnnnnnnn[-nn]` such as `-P9834Y10M`.
  - For `INTERVAL` of class day-second, the ISO 8601 string must be formatted as
    "`[-]PnnnnnnnnnDTnnHnnMnn[.nnnnn]S`", such as
    "`-P555DT14H55M59S`".
  - For negative durations, in both standard FGL format and ISO 8601 duration format, the minus sign
    must appear as first character.
- When converting a string to a `DATETIME`, the format must be in the form
  `YYYY-MM-DD hh:mm:ss.fffff` (1), or must follow the RFC 3339 format (2) with the
  `T` separator between the date and time part, with optional timezone indicator
  (`Z` or `+/-hh:mm`):

  - With form (1), the parts in the date/time string must match the target `DATETIME`
    type. For example, if the target type is a `DATETIME YEAR TO MINUTE`, the string must
    have the form `YYYY-MM-DD hh:mm`. For a `DATETIME HOUR TO SECOND`, the
    string must be in the form `hh:mm:ss[.fffff]`.
  - With form (2), the RFC 3339 specification requires full time part including the seconds (with
    optional fractions): `YYYY-MM-DDThh:mm:ss[.fffff]Z` for UTC or
    `YYYY-MM-DDThh:mm:ss[.fffff]±hh:mm` for a local time. When the target variable is a
    `DATETIME YEAR TO HOUR` or `DATETIME YEAR TO MINUTE`, the last units
    of the source string will be ignored.
  - When using seconds, the fraction part (`.fffff`) will be silently lost without
    conversion error, if the `DATETIME` type does not use the `FRACTION`
    precision.
  - The RFC 3339 / ISO 8601 standard requires a `T` separator between the date and
    time part. However, it is possible that date/time values with timezone offset might deviate from ISO
    8601, by using a space instead of the `T` separator. If a timezone offset is
    specified, the timezone conversion rules will apply, even if no `T` separator is
    used.
  - When the source string does not specify a timezone offset, the source date/time is considered to
    be in the same timezone than the current system timezone, and the value is directly assigned to the
    `DATETIME` target variable.
  - When the source string ends with a `"Z"` (UTC) or "`+/-hh:mm`"
    timezone offset, the value is adjusted to the local timezone, by converting the source date/time
    with timezone offset to a local time. For example, `"2012-07-24 11:33+01:00"` is
    converted to UTC `2012-07-24 10:33`, which is local time `2012-07-24
    12:33` in CEST, or local time `2012-07-24 03:33` in the PDT timezone.
  - When a timezone offset is specified after a time-only value (`12:00+01:00`), the
    runtime system uses the current local date for the conversion to a local time. This convention is
    needed, because it's not possible to determine the original date. For example, to convert
    `"12:00+01:00"` to a `DATETIME HOUR TO MINUTE`: The 1st January 2023
    in timezone CET, we get `"2023-01-01 12:00+01:00"`, that will convert to the local
    time `12:00`. On 1st July in timezone CEST, we get `"2023-07-01
    12:00+01:00"`, that will convert to the local time `13:00`.

Example converting character strings to `DATETIME`, on a Linux OS:

```
IMPORT util

MAIN
  DISPLAY "Current Time Zone (TZ):"
  RUN 'date +"%Z %:z"'
  DISPLAY "Current date/time: ", CURRENT, " / UTC:", util.Datetime.getCurrentAsUTC()
  DISPLAY "Source", COLUMN 30, "Local Time", COLUMN 60, "UTC"
  DISPLAY "With December date (winter):"
  CALL test("2012-12-24 11:33")
  CALL test("2012-12-24T11:33+01:00")
  CALL test("2012-12-24 11:33+01:00")
  CALL test("2012-12-24T10:33Z")
  CALL test("2012-12-24 10:33Z")
  DISPLAY "With July date (DST applies):"
  CALL test("2012-07-24 11:33")
  CALL test("2012-07-24T11:33+01:00")
  CALL test("2012-07-24 11:33+01:00")
  CALL test("2012-07-24T10:33Z")
  CALL test("2012-07-24 10:33Z")
END MAIN

FUNCTION test(s STRING)
  DEFINE dt_y2m DATETIME YEAR TO MINUTE
  DEFINE dt_h2m DATETIME HOUR TO MINUTE
  DEFINE h STRING
  LET dt_y2m = s
  DISPLAY s, COLUMN 30, dt_y2m, COLUMN 60, util.Datetime.toUTC(dt_y2m)
  LET h = s.subString(12,s.getLength()) -- keep only time part
  LET dt_h2m = h
  DISPLAY COLUMN 12, h, COLUMN 41, dt_h2m, COLUMN 71, util.Datetime.toUTC(dt_h2m)
END FUNCTION
```

Output example on Linux (note that the current date is
2023-09-07):

```
$ TZ="Europe/Paris" fglrun dtconv.42m
Current Time Zone (TZ):
CEST +02:00
Current date/time: 2023-09-07 15:11:59.338 / UTC:2023-09-07 13:11:59.33864
Source                       Local Time                    UTC
With December date (winter):
2012-12-24 11:33             2012-12-24 11:33              2012-12-24 10:33
           11:33                        11:33                         09:33
2012-12-24T11:33+01:00       2012-12-24 11:33              2012-12-24 10:33
           11:33+01:00                  12:33                         10:33
2012-12-24 11:33+01:00       2012-12-24 11:33              2012-12-24 10:33
           11:33+01:00                  12:33                         10:33
2012-12-24T10:33Z            2012-12-24 11:33              2012-12-24 10:33
           10:33Z                       12:33                         10:33
2012-12-24 10:33Z            2012-12-24 11:33              2012-12-24 10:33
           10:33Z                       12:33                         10:33
With July date (DST applies):
2012-07-24 11:33             2012-07-24 11:33              2012-07-24 09:33
           11:33                        11:33                         09:33
2012-07-24T11:33+01:00       2012-07-24 12:33              2012-07-24 10:33
           11:33+01:00                  12:33                         10:33
2012-07-24 11:33+01:00       2012-07-24 12:33              2012-07-24 10:33
           11:33+01:00                  12:33                         10:33
2012-07-24T10:33Z            2012-07-24 12:33              2012-07-24 10:33
           10:33Z                       12:33                         10:33
2012-07-24 10:33Z            2012-07-24 12:33              2012-07-24 10:33
           10:33Z                       12:33                         10:33

$ TZ="America/Los_Angeles" fglrun dtconv.42m
Current Time Zone (TZ):
PDT -07:00
Current date/time: 2023-09-07 06:12:37.157 / UTC:2023-09-07 13:12:37.15716
Source                       Local Time                    UTC
With December date (winter):
2012-12-24 11:33             2012-12-24 11:33              2012-12-24 19:33
           11:33                        11:33                         18:33
2012-12-24T11:33+01:00       2012-12-24 02:33              2012-12-24 10:33
           11:33+01:00                  03:33                         10:33
2012-12-24 11:33+01:00       2012-12-24 02:33              2012-12-24 10:33
           11:33+01:00                  03:33                         10:33
2012-12-24T10:33Z            2012-12-24 02:33              2012-12-24 10:33
           10:33Z                       03:33                         10:33
2012-12-24 10:33Z            2012-12-24 02:33              2012-12-24 10:33
           10:33Z                       03:33                         10:33
With July date (DST applies):
2012-07-24 11:33             2012-07-24 11:33              2012-07-24 18:33
           11:33                        11:33                         18:33
2012-07-24T11:33+01:00       2012-07-24 03:33              2012-07-24 10:33
           11:33+01:00                  03:33                         10:33
2012-07-24 11:33+01:00       2012-07-24 03:33              2012-07-24 10:33
           11:33+01:00                  03:33                         10:33
2012-07-24T10:33Z            2012-07-24 03:33              2012-07-24 10:33
           10:33Z                       03:33                         10:33
2012-07-24 10:33Z            2012-07-24 03:33              2012-07-24 10:33
           10:33Z                       03:33                         10:33
```

See also [util.Date.parse](../15_library-reference/3485-util-date-parse.md "Converts a string to a DATE value based on a format specification."), [util.Datetime.parse](../15_library-reference/3491-util-datetime-parse.md "Converts a string to a DATETIME value based on a specified format.") and [util.Interval.parse](../15_library-reference/3529-util-interval-parse.md "Converts a string to an INTERVAL value based on a specified format.").

## Converting `DATE` to/from `DATETIME` types

When converting a `DATETIME` to another `DATETIME` with a different
precision, truncation from the left or right can occur. If the target type has more fields than the
source type, the year, month and day fields are filled with the current date.

- When converting a `DATE` to a `DATETIME`, the datetime fields are
  filled with year, month and day from the date value and time fields are set to zero.
- When converting a `DATETIME` to a `DATE`, an implicit
  `EXTEND( datetime-value, YEAR TO DAY )` is performed.

```
MAIN
  DEFINE da DATE
  DEFINE dt1 DATETIME YEAR TO SECOND
  DEFINE dt2 DATETIME HOUR TO MINUTE
  LET da = MDY(12,24,2012)
  LET dt1 = da
  DISPLAY dt1
  LET dt2 = "23:45"
  LET dt1 = dt2
  DISPLAY dt1
END MAIN
```

Output:

```
2012-12-24 00:00:00
2023-12-21 23:45:00
```

## Unsupported type conversions

Other data type conversions not mentioned in this topic are not allowed and will result
in a runtime error.

## Related links

**Related concepts**  

[Primitive Data types](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")
