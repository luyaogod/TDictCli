---
title: "MATCHES and LIKE"
source: "fgl-topics/c_fgl_odiagdmg_027.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data manipulation > MATCHES and LIKE"
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

## Dameng®

Dameng does not provide an equivalent of the
Informix `MATCHES` operator.

The `LIKE` operator is supported.

With Dameng, columns defined as `CHAR(N)`
are blank padded, but trailing blanks are not significant in the `LIKE`
expressions:

```
SQL> CREATE TABLE tab1 ( c5 CHAR(5) );
SQL> INSERT INTO tab1 VALUES ( 'abc' );
SQL> INSERT INTO tab1 VALUES ( 'def' );

SQL> SELECT * FROM tab1 WHERE c5 LIKE 'abc%';
LINEID     c5   
---------- -----
1          abc  

SQL> SELECT * FROM tab1 WHERE c5 LIKE 'abc_';
LINEID     c5   
---------- -----
1          abc
```

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
