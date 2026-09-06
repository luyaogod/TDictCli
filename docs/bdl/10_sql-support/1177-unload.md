---
title: "UNLOAD"
source: "fgl-topics/c_fgl_InOutSql_UNLOAD.html"
breadcrumb: "SQL support > SQL LOAD and UNLOAD > UNLOAD"
type: "concept"
---

# UNLOAD

> Copies data from the database tables into a file.

## Syntax

```
UNLOAD TO filename [ DELIMITER delimiter]
{
  select-statement
|
  select-string
}
```

1. filename is a string expression containing the name of the file
   the data is written to.
2. delimiter is the character used as the value delimiter. When not specified,
   default is `|` pipe, or [DBDELIMITER](../07_configuration/0509-dbdelimiter.md "Defines the value separator for unload data files.") environment variable when set. The `DELIMITER` clause can also
   specify "CSV" or "TSV", to get respectively a Comma Separated Values format, or a TAB Separated
   Values format.
3. select-statement is static SELECT statement.
4. select-string is string expression containing
   the `SELECT` statement.

## Usage

The `UNLOAD` instruction serializes into a file the SQL data produced
by a `SELECT` statement.

A data file used by `LOAD` or `UNLOAD`
instructions looks like this (when the delimiter is the pipe), and the record has an
`INTEGER`, `VARCHAR` and `DATE`
fields:

```
102|Mike|12/24/2020|
192|Tom|04/29/2019|
```

The `UNLOAD` command cannot be used in a [`PREPARE`](1142-prepare-sql-statement.md "Prepares an SQL statement for execution.") statement. However,
the `UNLOAD` command accepts a string literal in place of a static
`SELECT` statement:

```
UNLOAD TO file-name
    select-string
```

The filename after the `TO` keyword identifies an
output file in which to store the rows retrieved from the database by the
`SELECT` statement. In the default (U.S. English) locale, this file
contains only ASCII characters. (In other locales, output from `UNLOAD`
can contain characters from the codeset of the locale.)

The `UNLOAD` statement must include a `SELECT`
statement (directly, or in a variable) to specify what rows to copy into
filename. `UNLOAD` does not delete the copied
data.

A single character delimiter instruct `UNLOAD` to write data in the default
format. When using "CSV" or "TSV" as delimiter specification, the `UNLOAD`
instruction will write the data in CSV or TSV format.

If the `DELIMITER` clause is not specified, the delimiter
is defined by the [DBDELIMITER](../07_configuration/0509-dbdelimiter.md "Defines the value separator for unload data files.") environment variable. If the DBDELIMITER environment variable is not set, the
default is a `|` pipe. The field delimiter can be a blank character.
It cannot be backslash or any hexadecimal digit (`0-9`, `A-F`,
`a-f`). If the delimiter specified in the `DELIMITER` clause is
`NULL`, the runtime system will use the default delimiter or DBDELIMITER if the
variable is defined.

When using a select-string, do not attempt to substitute question
marks (`?`) in place of host variables to make the `SELECT`
statement dynamic, because this usage has binding problems.

At this time, data type description of the output file fields is implicit; in order
to create the fetch buffers to hold the column values, the `UNLOAD`
instruction uses the current database connection to get the column data types of the
generated result set. Those data types depend on the type of database server. For example,
IBM® Informix®
`INTEGER` columns are integers of 4 bytes, while the Oracle `INTEGER`
data type is actually a `NUMBER(10,0)` type. Therefore, be aware when using this
instruction that if your application connects to different kinds of database servers, you may
get data conversion errors.

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

## Example

```
MAIN
  DEFINE x INTEGER
  DATABASE stores
  LET x = 123
  UNLOAD TO "items.unl"
    SELECT * FROM items WHERE item_num > x 
END MAIN
```

## Related links

**Related concepts**  

[LOAD](1176-load.md "Inserts data from a file into an existing database table.")
