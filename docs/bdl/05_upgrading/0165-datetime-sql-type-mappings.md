---
title: "DATETIME SQL type mappings"
source: "fgl-topics/c_fgl_Migrate_to_320_datetime_storage.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > DATETIME SQL type mappings"
type: "concept"
---

# DATETIME SQL type mappings

> For some databases, the type mapping for DATETIME HOUR TO MINUTE has changed.

## Summary of DATETIME storage changes

Starting with Genero 3.20, the native types used to store some of the `DATETIME
qual1 TO qual2` data types have changed.

> **Important:**
>
> The native type mapping has not changed for the most common date/time types
> `DATETIME YEAR TO SECOND`, `DATETIME YEAR TO FRACTION(N)`,
> `DATETIME HOUR TO SECOND`, `DATETIME HOUR TO FRACTION(N)`. The
> `DATETIME` type concerned by this modification is `DATETIME HOUR TO
> MINUTE`, and only for the database brands listed in this topic.

`DATETIME HOUR TO MINUTE` data is now stored with the closest native time data
type of the target database. Prior to version 3.20, such data was stored in a timestamp
(`YYYY-MM-DD hh:mm:ss[.ffff]`), with the year/month/day parts set to 1900-01-01. This
was not consistent with `DATETIME HOUR TO SECOND` storage, using time
(`hh:mm:ss[.ffff]`) types.

Using the native time data type for time-only data simplifies interoperability with other
applications and database components (for example, with data load/unload tools).

> **Note:**
>
> Only the database brands listed in this page are concerned by this `DATETIME`
> storage change.

## Database schema update

Existing databases on development and production sites will need data type modification for
columns concerned by this change.

Depending on the type of database, an `ALTER TABLE` statement can be used. For
example, with SQL Server, it is possible to convert a `datetime2(0)` column to
`time(0)`:

```
CREATE TABLE t1 ( pk INT, dt DATETIME2(0) )
INSERT INTO t1 VALUES ( 101, '1900-01-01 12:33:44' )
SELECT * FROM t1
ALTER TABLE t1 ALTER COLUMN dt TIME(0)
SELECT * FROM t1
DROP TABLE t1
```

For databases that do not support such date/time type modification with `ALTER
TABLE`, create a new table with the new types, write some SQL or BDL code to fill the new
table with the old table rows by doing a `CAST()`, drop the old table and rename the
new table to the original table name.

## Microsoft SQL Server

| FGL Data Type | FGL 3.10 | FGL 3.20 |
| --- | --- | --- |
| `DATETIME HOUR TO HOUR` | `DATETIME2(0)` | `TIME(0)` |
| `DATETIME HOUR TO MINUTE` | `DATETIME2(0)` | `TIME(0)` |
| `DATETIME HOUR TO SECOND` | `TIME(0)` | No change |
| `DATETIME HOUR TO FRACTION(N)` | `TIME(N)` | No change |
| `DATETIME MINUTE TO MINUTE` | `DATETIME2(0)` | `TIME(0)` |
| `DATETIME MINUTE TO SECOND` | `DATETIME2(0)` | `TIME(0)` |
| `DATETIME MINUTE TO FRACTION(N)` | `DATETIME2(n)` | `TIME(n)` |
| `DATETIME SECOND TO SECOND` | `DATETIME2(0)` | `TIME(0)` |
| `DATETIME SECOND TO FRACTION(N)` | `DATETIME2(n)` | `TIME(n)` |
| `DATETIME FRACTION TO FRACTION(N)` | `DATETIME2(n)` | `TIME(n)` |

See also [Date/time support with Microsoft SQL Server](../10_sql-support/1275-date-and-datetime-data-types.md).

## Oracle MySQL and MariaDB

| FGL Data Type | FGL 3.10 | FGL 3.20 |
| --- | --- | --- |
| `DATETIME HOUR TO HOUR` | `TIME` | No change |
| `DATETIME HOUR TO MINUTE` | `TIME` | No change |
| `DATETIME HOUR TO SECOND` | `TIME` | No change |
| `DATETIME HOUR TO FRACTION(n)` | `TIME(n)` | No change |
| `DATETIME MINUTE TO MINUTE` | `DATETIME` | `TIME` |
| `DATETIME MINUTE TO SECOND` | `DATETIME` | `TIME` |
| `DATETIME MINUTE TO FRACTION(n)` | `DATETIME(n)` | `TIME(n)` |
| `DATETIME SECOND TO SECOND` | `DATETIME` | `TIME` |
| `DATETIME SECOND TO FRACTION(n)` | `DATETIME(n)` | `TIME(n)` |
| `DATETIME FRACTION TO FRACTION(n)` | `DATETIME(n)` | `TIME(n)` |

See also [Date/time support with Oracle MySQL and
MariaDB](../10_sql-support/1329-date-and-datetime-data-types.md).

## PostgreSQL

| FGL Data Type | FGL 3.10 | FGL 3.20 |
| --- | --- | --- |
| `DATETIME HOUR TO HOUR` | `TIMESTAMP(0)` | `TIME(0)` |
| `DATETIME HOUR TO MINUTE` | `TIME(0)` | No change |
| `DATETIME HOUR TO SECOND` | `TIME(0)` | No change |
| `DATETIME HOUR TO FRACTION(N)` | `TIME(N)` | No change |
| `DATETIME MINUTE TO MINUTE` | `TIMESTAMP(0)` | `TIME(0)` |
| `DATETIME MINUTE TO SECOND` | `TIMESTAMP(0)` | `TIME(0)` |
| `DATETIME MINUTE TO FRACTION(N)` | `TIMESTAMP(n)` | `TIME(n)` |
| `DATETIME SECOND TO SECOND` | `TIMESTAMP(0)` | `TIME(0)` |
| `DATETIME SECOND TO FRACTION(N)` | `TIMESTAMP(n)` | `TIME(n)` |
| `DATETIME FRACTION TO FRACTION(N)` | `TIMESTAMP(n)` | `TIME(n)` |

> **Note:**
>
> All PostgreSQL time data type mappings use `TIME` and
> `TIMESTAMP` without time zone.

See also [Date/time support with PostgreSQL](../10_sql-support/1375-date-and-datetime-data-types.md).

## SQLite

With SQLite, the mapping rules for FGL `DATETIME` use custom type names such as
`SMALLTIME` and `TINYDATETIME`, to keep type information, for more
details see [DATE and DATETIME data types](../10_sql-support/1481-date-and-datetime-data-types.md).

| FGL Data Type | FGL 3.10 | FGL 3.20 |
| --- | --- | --- |
| Time class types |  |  |
| `DATETIME HOUR TO HOUR` | `TIMESTAMP` | `SMALLTIME` |
| `DATETIME HOUR TO MINUTE` | `SMALLTIME` | No change |
| `DATETIME HOUR TO SECOND` | `TIME` | No change |
| `DATETIME HOUR TO FRACTION(n)` | `TIME(n)` | No change |
| `DATETIME MINUTE TO MINUTE` | `TIMESTAMP` | `SMALLTIME` |
| `DATETIME MINUTE TO SECOND` | `TIMESTAMP` | `TIME` |
| `DATETIME MINUTE TO FRACTION(n)` | `TIMESTAMP(n)` | `TIME(n)` |
| `DATETIME SECOND TO SECOND` | `TIMESTAMP` | `TIME` |
| `DATETIME SECOND TO FRACTION(n)` | `TIMESTAMP(n)` | `TIME(n)` |
| `DATETIME FRACTION TO FRACTION(n)` | `TIMESTAMP(n)` | `TIME(n)` |
| Date/time class types |  |  |
| `DATETIME YEAR TO YEAR` | `TIMESTAMP` | `TINYDATETIME` |
| `DATETIME YEAR TO MONTH` | `TIMESTAMP` | `TINYDATETIME` |
| `DATETIME YEAR TO DAY` | `TINYDATETIME` | No change |
| `DATETIME YEAR TO HOUR` | `TIMESTAMP` | `SMALLDATETIME` |
| `DATETIME YEAR TO MINUTE` | `SMALLDATETIME` | No change |
| `DATETIME YEAR TO SECOND` | `DATETIME` | No change |
| `DATETIME YEAR TO FRACTION(n)` | `DATETIME(n)` | No change |
| `DATETIME MONTH TO MONTH` | `TIMESTAMP` | `TINYDATETIME` |
| `DATETIME MONTH TO DAY` | `TIMESTAMP` | `TINYDATETIME` |
| `DATETIME MONTH TO HOUR` | `TIMESTAMP` | `SMALLDATETIME` |
| `DATETIME MONTH TO MINUTE` | `TIMESTAMP` | `SMALLDATETIME` |
| `DATETIME MONTH TO SECOND` | `TIMESTAMP` | `DATETIME` |
| `DATETIME MONTH TO FRACTION(n)` | `TIMESTAMP` | `DATETIME(n)` |
| `DATETIME DAY TO DAY` | `TIMESTAMP` | `TINYDATETIME` |
| `DATETIME DAY TO HOUR` | `TIMESTAMP` | `SMALLDATETIME` |
| `DATETIME DAY TO MINUTE` | `TIMESTAMP` | `SMALLDATETIME` |
| `DATETIME DAY TO SECOND` | `TIMESTAMP` | `DATETIME` |
| `DATETIME DAY TO FRACTION(n)` | `TIMESTAMP` | `DATETIME(n)` |

Notes:

1. To be consistent with `TIME` class storage, types that can be stored in
   `TINYDATETIME`, `SMALLDATETIME`, `DATETIME` and
   `DATETIME(n)` will be stored with such type.

See also [Date/time support with SQLite](../10_sql-support/1481-date-and-datetime-data-types.md).
