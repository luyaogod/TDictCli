---
title: "PRINT"
source: "fgl-topics/c_fgl_reports_PRINT.html"
breadcrumb: "Reports > Report instructions > PRINT"
type: "concept"
---

# PRINT

> Formats and prints a row of data in a report routine.

## Syntax

```
PRINT
 { 
    expression 
  | COLUMN left-offset
  | PAGENO
  | LINENO
  | num-spaces SPACES
  | [GROUP] COUNT(*) [ WHERE condition ]
  | [GROUP] PERCENT(*) [ WHERE condition]
  | [GROUP] AVG( variable ) [ WHERE condition]
  | [GROUP] SUM( variable ) [ WHERE condition]
  | [GROUP] MIN( variable ) [ WHERE condition]
  | [GROUP] MAX( variable ) [ WHERE condition]
  | char-expression WORDWRAP [ RIGHT MARGIN rm ]
  | FILE file-name
 } [,...]
 [ ; ]
```

1. expression is any legal language expression.
2. left-offset is described in `COLUMN`.
3. num-spaces is described in `SPACES`.
4. char-expression is a string expression or a `TEXT` variable.
5. file-name is a string expression , or a quoted string, that specifies the
   name of a text file to include in the output from the report.

## Usage

The `PRINT` instruction is used in a report routine to output a line of data.

The `PRINT` statement can include character data in the form of an ASCII file, a
`TEXT` variable, or a comma-separated expression list of character expressions in
the output of the report. (For `TEXT` variable or file name, you cannot specify
additional output in the same `PRINT` statement.)

If a `BYTE` value is used in the `PRINT` statement, the output will
show the "`<byte value>`" text for this element when the report output is
regular text. If the report output is XML, the `BYTE` value is converted to Base64
before it is written to the output stream.

`PRINT` statement output begins at
the current character position, sometimes called simply the current position. On each page of a
report, the initial default character position is the first character position in the first line.
This position can be offset horizontally and vertically by margin and header specifications and by
executing any of the following statements:

- The `SKIP` statement moves it down
  to the left margin of a new line.
- The `NEED` statement can conditionally
  move it to a new page.
- The `PRINT` statement moves it horizontally (and sometimes down).

Unless you use the keyword `CLIPPED`
or `USING`, values are displayed with
widths (including any sign) that depend on their declared data types.

| Data type | Default Print Width |
| --- | --- |
| `BYTE` | N/A |
| `CHAR` | Length of character data type declaration. |
| `DATE` | DBDATE dependent, 10 if DBDATE = "MDY4/" |
| `DATETIME` | From 2 to 25, as implied in the data type declaration. |
| `DECIMAL` | (2 + p + s), where p is the precision and s is the scale from the data type declaration. |
| `FLOAT` | 14 |
| `INTEGER` | 11 |
| `INTERVAL` | From 3 to 25, as implied in the data type declaration. |
| `MONEY` | (2 + c + p + s), where c is the length of the currency defined by DBMONEY and p is the precision and s is the scale from the data type declaration. |
| `NCHAR` | Length of character data type declaration. |
| `NVARCHAR` | Length current value in the variable. |
| `SMALLFLOAT` | 14 |
| `SMALLINT` | 6 |
| `STRING` | Length current value in the variable. |
| `TEXT` | Length current value in the variable. |
| `VARCHAR` | Length current value in the variable. |

Unless you specify the `FILE` or `WORDWRAP` option, each `PRINT`
statement displays output on a single line. For example, this fragment displays output on two
lines:

```
PRINT fname, lname 
PRINT city, ", ", state, " ", zip-code
```

If you terminate a `PRINT` statement with a semicolon. However, you suppress
the implicit LINEFEED character at the end of the line. The following example has the same effect
as the `PRINT` statements in the previous example:

```
PRINT fname;
PRINT lname 
PRINT city, ", ", state, " ", zip-code
```

The expression list of a `PRINT` statement returns one or more values that can
be displayed as printable characters. The expression list can contain report variables, built-in
functions, and operators. Some of these can appear only in a `REPORT` program block
such as [`PAGENO`](2498-pageno.md "Contains the current page number in a report."), [`LINENO`](2497-lineno.md "Contains the current line number in a report."), [`PERCENT`](2503-percent.md "Calculates the percentage of rows matching a condition.").

If the expression list applies the `USING` operator to format a
`DATE` or `MONEY` value, the format string of the
`USING` operator takes precedence over the [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values."), [DBMONEY](../07_configuration/0512-dbmoney.md "Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined."), and [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") environment variables.

The `PRINT FILE file-name` statement reads the contents of
the specified file into the report, beginning at the current character position. This statement
permits you to insert a multiple-line character string into the output of a report. If
file-name stores the value of a TEXT variable, the `PRINT FILE
file-name` statement has the same effect as specifying `PRINT
text-variable`. (But only PRINT variable can include the
`WORDWRAP` operator)

Aggregate report functions summarize data from several records in a report. The syntax and
effects of aggregates in a report resemble those of SQL aggregate functions but are not
identical.

The expression (in parentheses) that `SUM()`, `AVG()`, `MIN()`, or `MAX()` takes as an argument is typically of a
number or `INTERVAL` data type; `ARRAY`, `BYTE`,
`RECORD`, and `TEXT` are not valid. The `SUM()`,
`AVG()`, `MIN()` , and `MAX()` aggregates ignore
input records for which their arguments have null values, but each returns `NULL`
if every record has a null value for the argument.

The `GROUP` keyword is an optional keyword that causes the aggregate function
to include data only for a group of records that have the same value for a variable that you
specify in an `AFTER GROUP
OF` control block. An aggregate function can only include the `GROUP`
keyword within an `AFTER GROUP OF` control block.

The optional `WHERE` clause allows you to select among records passed to the
report, so that only records for which the boolean expression is `TRUE` are
included.

## Example

The following example is from the `FORMAT` section of a report definition that
displays both quoted strings and values from rows of the customer table:

```
FIRST PAGE HEADER
  PRINT COLUMN 30, "CUSTOMER LIST"
  SKIP 2 LINES
  PRINT "Listings for the State of ", thisstate 
SKIP 2 LINES
  PRINT "NUMBER", COLUMN 12, "NAME", COLUMN 35, "LOCATION",
        COLUMN 57, "ZIP", COLUMN 65, "PHONE"
  SKIP 1 LINE
PAGE HEADER
  PRINT "NUMBER", COLUMN 12, "NAME", COLUMN 35, "LOCATION",
        COLUMN 57, "ZIP", COLUMN 65, "PHONE"
  SKIP 1 LINE
ON EVERY ROW
  PRINT customer_num USING "###&", COLUMN 12, fname CLIPPED,
       1 SPACE, lname CLIPPED, COLUMN 35, city CLIPPED, ", ",
       state, COLUMN 57, zip-code, COLUMN 65, phone
```

## Related links

**Related concepts**  

[PRINTX](2492-printx.md "Prints an XML formatted row of data in a report, with an additional identifier for XML outputs.")

[COLUMN](../08_language-basics/0631-column.md "The COLUMN operator generates blanks.")
