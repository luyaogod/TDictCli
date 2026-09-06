---
title: "String concatenation operator"
source: "fgl-topics/c_fgl_odiagmys_043.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data manipulation > String concatenation operator"
type: "concept"
description: "Informix® The Informix concatenation operator is the double pipe ( || ): SELECT firstname || ' ' || lastname FROM employee Oracle® MySQL and MariaDB By default, Oracle MySQL and MariaDB do support the ..."
---

# String concatenation operator

## Informix®

The Informix concatenation operator is the double pipe
( || ):

```
SELECT firstname || ' ' || lastname FROM employee
```

## Oracle® MySQL and MariaDB

By default, Oracle MySQL and MariaDB do
support the double-pipe operator, and requires the `concat()` function to concatenate
values:

```
SELECT concat(firstname, ' ',lastname) FROM employee
```

However, when the `sql_mode` configuration parameter contains the
`PIPES_AS_CONCAT` value, the DB engine will recognize `||` as
concatenation operator.

## Solution

The database interface for Oracle MySQL
and MariaDB does not convert double-pipe expressions to `concat()` function
calls.

In order to use the standard `||` double-pipe concatenation operator with MySQL
and MariaDB, set `PIPES_AS_CONCAT` in the `sql_mode` parameter. See
Oracle MySQL and MariaDB documentation for
more details about the `sql_mode` parameter.

## Related links

**Related concepts**  

[String concatenation operators in SQL](1032-string-concatenation-operators-in-sql.md "The || operator is the standard to concatenate strings.")
