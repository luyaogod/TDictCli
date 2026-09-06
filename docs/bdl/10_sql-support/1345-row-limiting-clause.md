---
title: "Row limiting clause"
source: "fgl-topics/c_fgl_odiagmys_046.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data manipulation > Row limiting clause"
type: "concept"
description: "Informix® Informix SQL supports the SKIP and FIRST/LIMIT keywords to limit the number of rows of a result set: SELECT SKIP 10 FIRST 20 customer.* FROM customer ... ORDER BY cust_name This Informix SQL ..."
---

# Row limiting clause

## Informix®

Informix SQL supports the `SKIP` and
`FIRST/LIMIT` keywords to limit the number of rows of a result set:

```
SELECT SKIP 10 FIRST 20 customer.* FROM customer ... ORDER BY cust_name
```

This Informix SQL syntax is not portable.

Recent database engines support the row limiting clause syntax defined by the SQL
standard:

```
SELECT ... OFFSET n ROWS FETCH FIRST m ROWS ONLY
```

This should be the prefered syntax to be used, if all target database types support this
`SELECT` clause.

The ODI database drivers can convert
the Informix SQL `SKIP/FIRST` row limiting clause to a native SQL equivalent, if the
row limiting clause parameters are simple integer literals (the clause is not translated when using
SQL parameters / program variables).
> **Important:**
>
> In addition to the
> `SKIP/FIRST` clause of the projection clause, Informix SQL supports also a
> `LIMIT` clause after the `ORDER BY`
> clause:
>
> ```
> SELECT customer.* FROM customer ... ORDER BY cust_name LIMIT 10
> ```
>
> This
> Informix SQL syntax construction is not converted by the ODI drivers. To benefit from the
> conversion, review the code to use the Informix SQL `SKIP/FIRST` clause
> instead.

## Oracle® MySQL and MariaDB

Oracle MySQL and MariaDB support the
following row limiting
clause:

```
SELECT ... ORDER BY ... LIMIT m [ OFFSET n ]
```

## Solution

The Informix SQL row limiting clause can be converted by the Oracle MySQL and MariaDB drivers to the native SQL equivalent clause, when
the parameters are simple integer literals.

> **Note:**
>
> - Row limiting clauses using SQL parameters will not be converted: The `SKIP` and
>   `FIRST` keywords must be followed by an integer constaint.
> - When using nested SQL queries, only the row limiting clause of the main `SELECT`
>   will be converted. The row limiting clauses in subqueries will not be converted.

The translation of the Informix SQL
row limiting clause can be controlled with the following FGLPROFILE
entry:

```
dbi.database.dsname.ifxemul.rowlimiting = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[Row limiting clause (SELECT)](1051-row-limiting-clause-select.md "How to use the right clause to limit the number of rows produced by a SELECT statement?")
