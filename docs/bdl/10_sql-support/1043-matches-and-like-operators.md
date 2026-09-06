---
title: "MATCHES and LIKE operators"
source: "fgl-topics/c_fgl_sql_programming_091.html"
breadcrumb: "SQL support > SQL programming > SQL portability > MATCHES and LIKE operators"
type: "concept"
---

# MATCHES and LIKE operators

> Use the standard LIKE operator instead of the MATCHES operator.

The `MATCHES` operator is specific to IBM®
Informix® SQL, at allows to compare a character string
column to a search pattern:

```
SELECT * FROM customer WHERE customer_name MATCHES "A*"
```

The Genero language supports a [`MATCHES`](../08_language-basics/0608-matches.md "The MATCHES operator returns TRUE if a string matches a given mask.") operator. Do not confuse the language `MATCHES`
operator (used in BDL instructions such as `IF custname MATCHES "S*"`), with the SQL
`MATCHES` operator (used in SQL statements). There is no problem in using the
`MATCHES` operator of BDL.

The standard SQL operator for pattern search is
`LIKE`:

```
SELECT * FROM customer WHERE customer_name LIKE "A%"
```

When using a non-Informix driver, the `MATCHES` expressions using a string
constant are replaced by a `LIKE` expression.
> **Important:**
>
> Only
> `MATCHES` expressions with a string constant can be converted to
> `LIKE` expressions, if the `MATCHES` uses a `?` SQL
> parameter place holder, no translation is done.

| Database Server Type | SQL MATCHES support |
| --- | --- |
| IBM Informix | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | Emulated, [see details](1295-matches-and-like.md) |
| Oracle® MySQL / MariadDB | Emulated, [see details](1344-matches-and-like.md) |
| Oracle Database Server | Emulated, [see details](1399-matches-and-like.md) |
| PostgreSQL | Emulated, [see details](1449-matches-and-like.md) |
| SQLite | Emulated, [see details](1493-matches-and-like.md) |
| Dameng® | Emulated, [see details](1245-matches-and-like.md) |

For maximum portability, replace SQL `MATCHES` expressions by `LIKE`
expression. `MATCHES` uses `*` and `?` as wildcards.
The equivalent wildcards in the `LIKE` operator are `%` and
`_`.

> **Important:**
>
> `MATCHES` character ranges such as `[a-z]`
> cannot be converted for the `LIKE` operator.

Pay attention to blank padding semantics of the target database when using a program variable in
static SQL or when using the variable for an `?` SQL parameter place holder after the
`LIKE` operator: If the program variable is defined as a `CHAR(N)`, it
is filled by the runtime system with trailing blanks, in order to have a size of N, and the pattern
for the `LIKE` operator will contain trailing blanks. For example, when a
`CHAR(10)` variable is assigned with `"ABC%"`, it will in fact contain
`"ABC%<6 blanks>"`. When used as SQL parameter for a `LIKE`
expression, the database server will search for column values matching `"ABC"` + some
characters + 6 blanks. To avoid this, use a `VARCHAR(N)` data type instead of
`CHAR(N)` to hold `LIKE` patterns.

Some database engines have specific semantics for the `LIKE` operator, especially
when using `CHAR(N)` data types. For example, with Oracle DB, the expression `custname LIKE '%h'`: If
`custname` is defined as `CHAR(30)`, Oracle will only find the rows when the `custname` values end
with a `'h'` at the last character position (30). Values such as
`'Smith'` will not match. Similarly, when doing `custname LIKE 'ab_'`,
rows where the column type is `CHAR(N>3)`, values such as `'abc'` will
not match in Oracle and PostgreSQL, because
of the significant trailing blanks.

As a general advice, use the `VARCHAR` type for variable string data, and leave
`CHAR` usage for fixed-length character string data such as codes.

PostgreSQL provides the `SIMILAR TO` operator, allowing
`[start-end]` character range specification as
in `MATCHES`.

## Related links

**Related concepts**  

[MATCHES](../08_language-basics/0608-matches.md "The MATCHES operator returns TRUE if a string matches a given mask.")
