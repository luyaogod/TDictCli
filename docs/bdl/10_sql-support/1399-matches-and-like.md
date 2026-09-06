---
title: "MATCHES and LIKE"
source: "fgl-topics/c_fgl_odiagora_032.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > MATCHES and LIKE"
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

## ORACLE

Oracle® does not provide an equivalent
of the Informix
`MATCHES` operator.

The `LIKE` operator is supported.

> **Important:**
>
> With Oracle,
> columns defined as `CHAR(N)` are blank padded, and trailing blanks are significant in
> `LIKE` expressions. As a result, with a `CHAR(5)` value such as
> `'abc '` (with 2 trailing blanks), the expression `(colname LIKE
> 'ab_')` will not match. To workaround this behavior, you can use `(RTRIM(colname)
> LIKE 'pattern')`. However, consider adding the condition  `AND
> (colname LIKE 'patten%')` to force the DB server to optimize the query of
> the column as indexed. The `CONSTRUCT` instruction uses this technique when the
> entered criteria does not end with a `*` star wildcard.

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

## Related links

**Related concepts**  

[MATCHES and LIKE operators](1043-matches-and-like-operators.md "Use the standard LIKE operator instead of the MATCHES operator.")

[Query by example (CONSTRUCT)](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.")
