---
title: "base.Channel.setDelimiter"
source: "fgl-topics/c_fgl_ClassChannel_setDelimiter.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.setDelimiter"
type: "concept"
---

# base.Channel.setDelimiter

> Define the value delimiter for a channel.

## Syntax

```
setDelimiter(
   delimiter STRING )
```

1. delimiter defines the delimiter to be used to separate values. Possible
   options are:
   - "sep": Use single-char delimiter formatting rules, with sep as separator.
   - `"CSV"` : Use DSV
     delimiter formatting rules, with comma as separator.
   - `"TSV"` : Use DSV delimiter formatting rules, with TAB as separator.
   - `"DSV=sep"` : Use DSV delimiter formatting rules, with
     sep as separator.

## Usage

After creating the channel object, define the field value delimiter with the `setDelimiter()`
method.

```
CALL ch.setDelimiter("^")
```

The default delimiter is defined by the [DBDELIMITER](../07_configuration/0509-dbdelimiter.md "Defines the value separator for unload data files.")
environment variable, or a pipe (`|`) if DBDELIMITER is not defined.

Specify `CSV` to read/write in Comma Separated Value
format:

```
CALL ch.setDelimiter("CSV")
```

Specify `TSV` to read/write in TAB Separated Value
format:

```
CALL ch.setDelimiter("TSV")
```

Specify `DSV=sep` to read/write in Delimiter Separated Value
format, using the specified character as separator. For example, to define a semicolon as
separator:

```
CALL ch.setDelimiter("DSV=;")
```

Setting a `NULL` delimiter is allowed for backward compatibility, but must be
avoided. This was a workaround to read/write complete lines. If the delimiter is set to
`NULL`, the `read()` and `write()` methods do not use
the backslash (`\`) escape character. As a result, data with special characters like
backslash, delimiter or line-feed will be written as is, and reading data will ignore escaped
characters in the source stream. If you need to read or write non-formatted data, it is recommended
that you use the [`readLine()`](2994-base-channel-readline.md "Read a complete line from the channel.")/[`writeLine()`](2998-base-channel-writeline.md "Write a complete line to the channel.") methods instead. These methods do not use a delimiter, nor do they use
the backslash escape character.

## Rules for single-char delimiter format

Single-char
delimiter formatting rules apply when a unique character is specified as delimiter. These rules are
slightly different as when using a DSV delimiter.

Character strings must be encoded in the [current application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."): Not charset conversion is
done.

Trailing blanks are dropped from `CHAR` and `TEXT`, but not
from `VARCHAR` values.

Character-type data need a backslash ( \ ) before any
literal backslash or delimiter character and before a NEWLINE character in a character value. When
reading character-type data, values can have more characters than the declared maximum length of the
column, but any extra characters are ignored. Blank values can be represented as one or more blank
characters between delimiters, but leading blanks must not precede other `CHAR`,
`VARCHAR`, or `TEXT` values.

Numeric-typed data representation
depends on `DBFORMAT/DBMONEY` environment variables.

`DATE`
value formatting is based on the DBDATE environment variable. The day and month must be a 2-digit
number, and the year must be a 4-digit number.

When reading for `MONEY` types, values can include currency symbols, but these
are not required.

`DATETIME` values are represented in the format
`year-month-day hour:minute:second.fraction` or a contiguous subset. Time units
outside the reference type precision are omitted. The `year` must be a four-digit
number; all other time units (except `fraction`) require two
digits.

`INTERVAL` values are formatted `year-month` or
`day hour:minute:second.fraction` or a contiguous subset. Time units outside the
reference type precision are omitted.

`BYTE` values must be ASCII-hexadecimals;
without leading or trailing blanks.

`NULL` values of any data type are
represented by consecutive delimiters, without any characters between the delimiter
symbols.

The backslash symbol (`\`) serves as an escape character to indicate
that the next character in a data value is a literal that needs to be escaped, such as a backslash,
NEWLINE, or the delimiter character.

## Rules for Delimiter Separated Values format

The Delimiter Separated Values (DSV) formatting rules apply when using `"CSV"`
(Comma Separated Values), `"TSV"` (TAB Separated Values), or
`"DSV=sep"` as delimiter.

DSV serialization and de-serialization rules are similar to the single-char delimiter formatting
rules when using a regular single-character delimiter, with the following differences:

- Leading and trailing blanks are kept (no truncation).
- No ending delimiter is expected at the end of the input record.
- Values might be surrounded with `"` double quotes, if the data contains
  characters such as the `"` double quote, `\` backslash, new-line, or
  the delimiter character, and that character is not escaped.

## Related links

**Related concepts**  

[Read and write record data](3001-read-and-write-record-data.md "Read and write record data")
