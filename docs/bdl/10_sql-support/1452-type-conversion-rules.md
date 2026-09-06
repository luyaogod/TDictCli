---
title: "Type conversion rules"
source: "fgl-topics/c_fgl_odiagpgs_028.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data manipulation > Type conversion rules"
type: "concept"
description: "Informix® When using parameterized SQL statements (SQL with program variables), the Informix SQL engine supports various type conversions. For example, it is possible to compare a VARCHAR column to an ..."
---

# Type conversion rules

## Informix®

When using parameterized SQL statements (SQL with program variables), the Informix SQL engine
supports various type conversions.

For example, it is possible to compare a `VARCHAR` column to an
`INTEGER` program variable:

```
MAIN
    DEFINE v_int INTEGER
         , v_count INTEGER
    DATABASE test1
    CREATE TABLE tt1 ( col1 VARCHAR(10) )
    INSERT INTO tt1 VALUES ( '12345' )
    LET v_int = 12345
    SELECT COUNT(*) INTO v_count
       FROM tt1 WHERE col1 = v_int
    DISPLAY "count = ", v_count
END MAIN
```

## PostgreSQL

PostgreSQL is more strict as Informix when comparing or assigning SQL column values to program
variable, if the data types do not match.

With the above code example, PostgreSQL produces the following SQL error:

- SQLSTATE: `42883`
- SQL message: `operator does not exist: character varying = integer`

A similar error is raised for date vs string and decimal vs string types.

## Solution

A good coding convention is to always use a program variable type that corresponds to the
database column type.

For a complete list of Informix / PostgreSQL data type correspondance, see [SQL types mapping: PostgreSQL](1427-sql-types-mapping-postgresql.md).

> **Tip:**
>
> Check that the data stored in the database actually requires the current type of
> the column. For example, if a `VARCHAR` column contains (and should only contain)
> whole numbers, consider to modify the column type to `SMALLINT`,
> `INTEGER` or `BIGINT`.

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
