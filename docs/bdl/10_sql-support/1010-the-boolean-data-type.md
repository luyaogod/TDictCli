---
title: "The BOOLEAN data type"
source: "fgl-topics/c_fgl_sql_programming_boolean.html"
breadcrumb: "SQL support > SQL programming > SQL portability > The BOOLEAN data type"
type: "concept"
---

# The BOOLEAN data type

> SQL implementation of the BOOLEAN data type varies on the database type.

## Basics about BOOLEAN SQL type

BDL [`BOOLEAN`](../08_language-basics/0556-boolean.md "The BOOLEAN data type stores a logical value, TRUE or FALSE.") values can be
`NULL`, `TRUE(1)` or `FALSE(0)`.

SQL Database vendor specific implementation of the boolean SQL type may not correspond exactly to
the BDL `BOOLEAN` values.

For example, IBM® Informix® boolean SQL data type uses the `'t'` and `'f'` values
respectively to represent the true and false values.

However, BDL `BOOLEAN` variables can be used in SQL statements: The conversion
from the BDL `BOOLEAN` values `TRUE/1` and `FALSE/0`
to/from the SQL boolean values will be done by the database client or by the ODI driver.

| Database Server Type | Native boolean type |
| --- | --- |
| IBM Informix | Yes, [see details](1185-fully-supported-ibm-informix-sql-features.md "Fully supported IBM Informix SQL features."). |
| Microsoft™ SQL Server | Yes, [as `BIT` data type](1272-boolean-data-type.md). |
| Oracle® MySQL / MariaDB | Yes, [see details](1326-boolean-data-type.md). |
| Oracle Database Server | Yes, [see details](1372-boolean-data-type.md). |
| PostgreSQL | Yes, [see details](1429-boolean-data-type.md). |
| SQLite | Yes, [see details](1478-boolean-data-type.md). |
| Dameng® | Yes, [as `BIT` data type](1227-boolean-data-type.md). |

## The TRUE and FALSE constants in SQL

The `TRUE`/`FALSE` constants are Genero language
constants:

```
DEFINE b1, b2 BOOLEAN
LET b1 = TRUE
LET b2 = FALSE
```

The SQL syntax of the target database engine may not support the `TRUE/FALSE`
keywords, for example in an statement such
as:

```
INSERT INTO mytable (key,bcol) VALUES (455,TRUE)
```

In such case, consider using a `BOOLEAN` program variable for maximum SQL
portability:

```
DEFINE rec RECORD
       key INTEGER,
       bcol BOOLEAN
   END RECORD
LET rec.key = 455
LET rec.bcol = TRUE
INSERT INTO mytable (key,bcol) VALUES (rec.*)
```

## Related links

**Related concepts**  

[BOOLEAN](../08_language-basics/0556-boolean.md "The BOOLEAN data type stores a logical value, TRUE or FALSE.")
