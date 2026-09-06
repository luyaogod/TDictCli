---
title: "Scrollable cursors"
source: "fgl-topics/c_fgl_sql_programming_scroll_cursors.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Scrollable cursors"
type: "concept"
---

# Scrollable cursors

> How scrollable cursors can be supported on different databases.

Scrollable cursors can be used to go forward and backward in an SQL query result
set:

```
DEFINE cust_rec RECORD LIKE customer.*
DECLARE sc SCROLL CURSOR
   FOR SELECT * FROM customer
OPEN sc
FETCH NEXT sc INTO cust_rec.*
FETCH LAST sc INTO cust_rec.*
FETCH FIRST sc INTO cust_rec.*
CLOSE sc
```

This is a useful feature to implement record set navigation in applications. Scrollable
cursors are typically implemented in the database server. But not all database servers support
scrollable cursors.

When scrollable cursors are not supported by the target database server, the database
driver will emulate it with temporary files.

The temporary files are create in a temporary directory, that can be defined with the
DBTEMP environment variable. If DBTEMP is not defined, the default temporary directory
dependents from the platform used.

It is recommended that you avoid scrollable cursor usage if the target database does not
support this feature:

With emulated scrollable cursors, when scrolling to the last row, all rows will be
fetched into the temporary file. This can generate a lot of network traffic and can
produce a large temporary file if the result-set contains a lot of rows. Additionally,
programs are dependent on the file system resource allocated to the OS user
(ulimit).

Some databases do not support to use large objects data types (`TEXT/BYTE`) with
scrollable cursors: The `OPEN` statement will produce SQL error [-611](../15_library-reference/4483-genero-bdl-errors.md). To write portable SQL, use only
simple data types in the result set of the scrollable cursor, and use the primary key column to
fetch `TEXT/BYTE` data in a secondary `SELECT` statement.

The following table lists the native scrollable cursor availability for each supported
database:

| Database Server Type | Scrollable cursors support | TEXT/BYTE support (with scrollable cursors) |
| --- | --- | --- |
| IBM® Informix® | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) | No |
| Microsoft™ SQL Server | Yes, [see details](1307-scrollable-cursors.md) | Yes |
| Oracle® MySQL / MariadDB | Emulated, [see details](1354-scrollable-cursors.md) | No |
| Oracle Database Server | Yes, [see details](1410-scrollable-cursors.md) | Yes |
| PostgreSQL | Yes, [see details](1461-scrollable-cursors.md) | Yes |
| SQLite | Emulated, [see details](1503-scrollable-cursors.md) | No |
| Dameng® | Yes, [see details](1255-scrollable-cursors.md) | Yes |

## Related links

**Related concepts**  

[DECLARE (result set cursor)](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.")

[DBTEMP](../07_configuration/0517-dbtemp.md "Defines the directory for temporary files.")

[Optimizing scrollable cursors](1058-optimizing-scrollable-cursors.md "A programming pattern to get fresh data from scrollable cursors.")
