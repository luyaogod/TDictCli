---
title: "SQL types mapping: SQLite"
source: "fgl-topics/c_fgl_odiagsqt_027.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Data dictionary > SQL types mapping: SQLite"
type: "concept"
description: "Table 1. SQL data types mapping for SQLite Original data types SQLite data types CHAR(n) CHAR(n) COLLATE RTRIM VARCHAR(n[,m]) VARCHAR(n) COLLATE RTRIM LVARCHAR(n) VARCHAR(n) COLLATE RTRIM NCHAR(n) ..."
---

# SQL types mapping: SQLite

| Original data types | SQLite data types |
| --- | --- |
| `CHAR(n)` | `CHAR(n) COLLATE RTRIM` |
| `VARCHAR(n[,m])` | `VARCHAR(n) COLLATE RTRIM` |
| `LVARCHAR(n)` | `VARCHAR(n) COLLATE RTRIM` |
| `NCHAR(n)` | `NCHAR(n)` |
| `NVARCHAR(n)` | `NVARCHAR(n)` |
| `BOOLEAN` | `BOOLEAN` |
| `SMALLINT` | `SMALLINT` |
| `INTEGER` | `INTEGER` |
| `BIGINT` | `BIGINT` |
| `INT8` | `BIGINT` |
| `SERIAL[(start)]` | `INTEGER` (see note 1) |
| `BIGSERIAL[(start)]` | N/A (see note 1) |
| `INT8[(start)]` | N/A (see note 1) |
| `DOUBLE PRECISION / FLOAT[(n)]` | `FLOAT` |
| `REAL / SMALLFLOAT` | `SMALLFLOAT` |
| `DECIMAL(p,s)` | `DECIMAL(p,s)` |
| `DECIMAL(p)` | `DECIMAL(p)` |
| `DECIMAL` | `DECIMAL` |
| `MONEY(p,s)` | `DECIMAL(p,s)` |
| `MONEY(p)` | `DECIMAL(p,2)` |
| `MONEY` | `DECIMAL(16,2)` |
| `TEXT` | `TEXT` |
| `BYTE` | `BLOB` |
| `DATE` | `DATE` |
| `DATETIME HOUR TO HOUR` | `SMALLTIME` |
| `DATETIME HOUR TO MINUTE` | `SMALLTIME` |
| `DATETIME HOUR TO SECOND` | `TIME` |
| `DATETIME HOUR TO FRACTION(n)` | `TIME(n)` |
| `DATETIME MINUTE TO MINUTE` | `SMALLTIME` |
| `DATETIME MINUTE TO SECOND` | `TIME` |
| `DATETIME MINUTE TO FRACTION(n)` | `TIME(n)` |
| `DATETIME SECOND TO SECOND` | `TIME` |
| `DATETIME SECOND TO FRACTION(n)` | `TIME(n)` |
| `DATETIME FRACTION TO FRACTION(n)` | `TIME(n)` |
| `DATETIME YEAR TO YEAR` | `TINYDATETIME` |
| `DATETIME YEAR TO MONTH` | `TINYDATETIME` |
| `DATETIME YEAR TO DAY` | `TINYDATETIME` |
| `DATETIME YEAR TO HOUR` | `SMALLDATETIME` |
| `DATETIME YEAR TO MINUTE` | `SMALLDATETIME` |
| `DATETIME YEAR TO SECOND` | `DATETIME` |
| `DATETIME YEAR TO FRACTION(n)` | `DATETIME(n)` |
| `DATETIME MONTH TO MONTH` | `TINYDATETIME` |
| `DATETIME MONTH TO DAY` | `TINYDATETIME` |
| `DATETIME MONTH TO HOUR` | `SMALLDATETIME` |
| `DATETIME MONTH TO MINUTE` | `SMALLDATETIME` |
| `DATETIME MONTH TO SECOND` | `DATETIME` |
| `DATETIME MONTH TO FRACTION(n)` | `DATETIME(n)` |
| `DATETIME DAY TO DAY` | `TINYDATETIME` |
| `DATETIME DAY TO HOUR` | `SMALLDATETIME` |
| `DATETIME DAY TO MINUTE` | `SMALLDATETIME` |
| `DATETIME DAY TO SECOND` | `DATETIME` |
| `DATETIME DAY TO FRACTION(n)` | `DATETIME(n)` |
| `INTERVAL q1 TO q2` | `CHAR(50)` |

Notes:

1. For serial emulation, see [SERIAL and BIGSERIAL data types](1483-serial-and-bigserial-data-types.md).
