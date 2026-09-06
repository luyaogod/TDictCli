---
title: "MATCHES and LIKE"
source: "fgl-topics/c_fgl_odiagmsv_029.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data manipulation > MATCHES and LIKE"
type: "concept"
description: "Informix® Informix supports MATCHES and LIKE operators in SQL statements. MATCHES expects * and ? wild-card characters, while LIKE uses the % and _ wild-cards as equivalents. ( col MATCHES 'Smi*' AND ..."
---

# MATCHES and LIKE

## Informix®

Informix supports `MATCHES` and
`LIKE` operators in SQL statements.

`MATCHES` expects `*` and `?` wild-card characters,
while `LIKE` uses the `%` and `_` wild-cards as
equivalents.

```
( col MATCHES 'Smi*' AND col NOT MATCHES 'R?x' )
( col LIKE 'Smi%' AND col NOT LIKE 'R_x' )
```

`MATCHES` accepts also brackets notation, to specify a set of matching characters
at a given position:

```
( col MATCHES '[Pp]aris' )
( col MATCHES '[0-9][a-z]*' )
```

## Microsoft™ SQL Server

Microsoft SQL Server does not provide an equivalent of
the Informix `MATCHES` operator.

The `LIKE` operator is supported.

> **Important:**
>
> The `LIKE` operator of SQL Server does not evaluate to true with
> `CHAR/NCHAR` columns, if the `LIKE` pattern is provided as a UNICODE
> string literal (with the N prefix), and the search pattern matches the value in the column (without
> an ending % wildcard for example).
>
> See the following test:
>
> ```
> CREATE TABLE mytable ( k INT, nc NCHAR(20) )
> INSERT INTO mytable VALUES ( 1, N'abc' )
> SELECT * FROM mytable WHERE nc = 'abc' -- one row is returned
> SELECT * FROM mytable WHERE nc = N'abc' -- one row is returned
> SELECT * FROM mytable WHERE nc LIKE 'abc' -- one row is returned
> SELECT * FROM mytable WHERE nc LIKE N'abc' -- no rows are found
> SELECT * FROM mytable WHERE nc LIKE N'abc%' -- one row is returned
> ```
>
> This can be an issue because the SQL Server driver will by default automatically add an N prefix
> before all string literals in SQL statements.
>
> See Microsoft SQL Server documentation for more
> details about the `LIKE` semantics regarding blank padding and see also [CHARACTER data types](1273-char-and-varchar-data-types.md) for the N prefix usage
> and single-char or wide-char mode usage.

## Solution

The database driver is able to translate Informix
`MATCHES` expressions to `LIKE` expressions, when no `[
]` bracket character ranges are used in the `MATCHES` operand.

The `MATCHES` to `LIKE` expression translation is controlled by the
following FGLPROFILE
entry:

```
dbi.database.dbname.ifxemul.matches = { true | false }
```

> **Important:**
>
> Only `[NOT] MATCHES` followed by a search pattern provided as
> a string literal can be converted by ODI drivers. A `[NOT] MATCHES` followed by a ?
> question mark parameter place holder is not translated!

For maximum portability, consider replacing the `MATCHES` expressions with
`LIKE` expressions in all SQL statements.

Avoid using `CHAR(N)` types for variable length character data (such as name,
address).

## LIKE with UNICODE string literals on CHAR/NCHAR columns

Pay attention to UNICODE string prefixes N'...' in the `LIKE` expressions when
used with `CHAR/NCHAR` columns: You might want to always add a `%`
wildcard at the end of the `LIKE` expression, or use the equal operator when doing a
query with exact values.

## Related links

**Related concepts**  

[MATCHES and LIKE operators](1043-matches-and-like-operators.md "Use the standard LIKE operator instead of the MATCHES operator.")
