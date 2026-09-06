---
title: "The UNITS operator"
source: "fgl-topics/c_fgl_sql_programming_units.html"
breadcrumb: "SQL support > SQL programming > SQL portability > The UNITS operator"
type: "concept"
---

# The UNITS operator

> The UNITS SQL operator is specific to Informix SQL and needs to be considered.

## Basics about the UNITS operator

Informix SQL language provides the `UNITS` operator, to produce an
`INTERVAL` value from a single numeric
expression.

```
expr UNITS {YEAR|MONTH|DAY|HOUR|MINUTE|SECOND|FRACTION(n)}
```

This
operator is specific to the Informix SQL language and is not supported by database engine types such
as Oracle® Database Server or Microsoft™ SQL Server.

The `UNITS` operator in SQL statements is not the [BDL language `UNITS` operator](../08_language-basics/0669-units.md "The UNITS operator converts an integer to an interval."). The first is
processed by the SQL engine, while the BDL `UNITS` operator is processed by the
Genero runtime system and is independent from the database engine used by the program.

`UNITS` is used in date/time arithmetic expressions, to increase or decrease one
of the time units in an `INTERVAL`, `DATETIME` or
`DATE` value, or to compare a date/time
interval:

```
SELECT lead_time + 5 UNITS DAY FROM manufact;
SELECT * FROM cust_calls WHERE (TODAY - call_dtime) > 30 UNITS DAY;
```

## Adapting SQL statement using UNITS operator

The `UNITS` operator in SQL statements must be replaced by a portable SQL
expression.

The SQL expression can be modified to use a program variable as SQL parameter, that will be set
before executing the SQL statement.
> **Tip:**
>
> When reviewing such SQL statements, try to
> simplify the expressions, in order to ease the work of the SQL optimizer and use table indexes when
> possible.

The following example:

```
DEFINE cnt INTEGER
SELECT COUNT(*) INTO cnt FROM cust_calls
   WHERE (TODAY - call_dtime) > 30 UNITS DAY
```

Can be converted to compute the date limit in a program variable, before executing the SQL
statement:

```
DEFINE cnt INTEGER
DEFINE limit DATETIME YEAR TO SECOND
LET limit = CURRENT - 30 UNITS DAY
SELECT COUNT(*) INTO cnt FROM cust_calls
   WHERE call_dtime < limit
```
