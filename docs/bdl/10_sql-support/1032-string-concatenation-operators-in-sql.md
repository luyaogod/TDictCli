---
title: "String concatenation operators in SQL"
source: "fgl-topics/c_fgl_sql_programming_099.html"
breadcrumb: "SQL support > SQL programming > SQL portability > String concatenation operators in SQL"
type: "concept"
---

# String concatenation operators in SQL

> The || operator is the standard to concatenate strings.

The ANSI SQL standards define the double-pipe as string concatenation operator, for
example:

```
SELECT city_id || '/' || city_name FROM customer
```

Produces the following result:

```
456/Paris
1234/London
```

However, some database engine types do not support the standard double-pipe concatenation
operator, or need some configuration setting to allow its use.

| Database Server Type | Double pipe operator |
| --- | --- |
| IBM® Informix® | Yes |
| Microsoft™ SQL Server | Yes (SQL Server 2025), [see details](1298-string-concatenation-operator.md) |
| Oracle® MySQL / MariadDB | Can be enabled, [see details](1346-string-concatenation-operator.md) |
| Oracle Database Server | Yes |
| PostgreSQL | Yes |
| SQLite | Yes |
| Dameng® | Yes |

If needed (typically, with Microsoft SQL Server
versions prior to 2025), ODI drivers will convert the `||` double-pipe operator to
the native concatenation operator.

## Related links

**Related concepts**  

[String literals in SQL statements](1031-string-literals-in-sql-statements.md "Single quotes is the standard for delimiting string literals in SQL.")

[fgl\_db\_driver\_type()](../15_library-reference/2737-fgl-db-driver-type.md "Returns the 3-letter identifier/code of the current database driver.")
