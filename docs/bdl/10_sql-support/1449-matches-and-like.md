---
title: "MATCHES and LIKE"
source: "fgl-topics/c_fgl_odiagpgs_029.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data manipulation > MATCHES and LIKE"
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

## PostgreSQL

PostgreSQL supportes the `LIKE` operator, and the `~` operator wich
expects regular expressions as follows: `( col ~ 'a.*' )`

PostgreSQL provides the `SIMILAR TO` opertator, allowing character range
specification as the Informix `MATCHES`
operator:

```
( col SIMILAR TO '[Pp]ar%' )
```

> **Important:**
>
> With PostgreSQL, columns defined as `CHAR(N)` are blank
> padded, and trailing blanks are significant in the `LIKE` expressions. As result,
> with a `CHAR(5)` value such as `'abc '` (with 2 trailing blanks), the
> expression `(colname LIKE 'ab_')` will not match. To workaround this behavior, you
> can do `(RTRIM(colname) LIKE 'pattern')`. However, consider adding
> the condition  `AND (colname LIKE 'patten%')` to force the DB
> server to optimize the query of the column is indexed. The `CONSTRUCT` instruction
> uses this technique when the entered criteria does not end with a `*` star
> wildcard.

## Solution

The database driver converts Informix
`MATCHES` expressions to `LIKE` expressions, when no `[
]` bracket character ranges are used in the `MATCHES` operand. When character
ranges are used, the driver converts to a PostgreSQL `SIMILAR TO` expression, to find
the same values as with the Informix `MATCHES` operator.

> **Important:**
>
> Only `[NOT] MATCHES` followed by a search pattern provided as
> a string literal can be converted by ODI drivers. A `[NOT] MATCHES` followed by a ?
> question mark parameter place holder is not translated!

For maximum portability, consider replacing the `MATCHES` expressions to
`LIKE` expressions in all SQL statements of your programs.

With PostgreSQL, trailing blanks are significant when comparing `CHAR()` columns
with `LIKE` or `SIMILAR TO` expressions. Consider adding an ending \*
when comparing with `CHAR()` columns. This is not needs for
`VARCHAR()` columns.

Avoid using `CHAR(N)` types for variable length character data (such as name,
address).

## Related links

**Related concepts**  

[MATCHES and LIKE operators](1043-matches-and-like-operators.md "Use the standard LIKE operator instead of the MATCHES operator.")
