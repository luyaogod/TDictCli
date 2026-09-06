---
title: "SQL types mapping: SQL Server"
source: "fgl-topics/c_fgl_odiagmsv_047.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > SQL types mapping: SQL Server"
type: "concept"
description: "Table 1. SQL data types mapping for Microsoft™ SQL Server Original data types Microsoft SQL Server data types CHAR(n) CHAR(n) or NCHAR(n) (see note 1 ) VARCHAR(n[,m]) VARCHAR(n) or NVARCHAR(n) (see ..."
---

# SQL types mapping: SQL Server

| Original data types | Microsoft SQL Server data types |
| --- | --- |
| `CHAR(n)` | `CHAR(n)` or `NCHAR(n)` (see note 1) |
| `VARCHAR(n[,m])` | `VARCHAR(n)` or `NVARCHAR(n)` (see note 1) |
| `LVARCHAR(n)` | `VARCHAR(n)` or `NVARCHAR(n)` (see note 1) |
| `NCHAR(n)` | `NCHAR(n)` (see note 1) |
| `NVARCHAR(n[,m])` | `NVARCHAR(n)` (see note 1) |
| `BOOLEAN` | `BIT` |
| `SMALLINT` | `SMALLINT` |
| `INTEGER` | `INTEGER` |
| `BIGINT` | `BIGINT` |
| `INT8` | `BIGINT` |
| `SERIAL[(start)]` | `INTEGER` (see note 2) |
| `BIGSERIAL[(start)]` | `BIGINT` (see note 2) |
| `SERIAL8[(start)]` | `BIGINT` (see note 2) |
| `DOUBLE PRECISION / FLOAT[(n)]` | `FLOAT(n)` |
| `REAL / SMALLFLOAT` | `REAL` |
| `DECIMAL(p,s)` | `DECIMAL(p,s)` |
| `DECIMAL(p<=16)` | `DECIMAL(2*p,p)` |
| `DECIMAL(p>16)` | N/A |
| `DECIMAL` | `DECIMAL(32,16)` |
| `MONEY(p,s)` | `DECIMAL(p,s)` |
| `MONEY(p)` | `DECIMAL(p,2)` |
| `MONEY` | `DECIMAL(16,2)` |
| `DATE` | `DATE` |
| `DATETIME HOUR TO MINUTE` | `TIME(0)` |
| `DATETIME HOUR TO SECOND` | `TIME(0)` |
| `DATETIME HOUR TO FRACTION(n)` | `TIME(n)` |
| `DATETIME YEAR TO MONTH` | `DATETIME2(0)` |
| `DATETIME YEAR TO DAY` | `DATETIME2(0)` |
| `DATETIME YEAR TO HOUR` | `DATETIME2(0)` |
| `DATETIME YEAR TO MINUTE` | `DATETIME2(0)` |
| `DATETIME YEAR TO SECOND` | `DATETIME2(0)` |
| `DATETIME YEAR TO FRACTION(n)` | `DATETIME2(n)` |
| `INTERVAL q1 TO q2` | `CHAR(50)` |
| `TEXT` | `VARCHAR(MAX)` or `NVARCHAR(MAX)` (see note 3) |
| `BYTE` | `VARBINARY(MAX)` |

Notes:

1. For character type conversions, see [CHAR and VARCHAR data types](1273-char-and-varchar-data-types.md).
2. For serial emulation, see [SERIAL and BIGSERIAL data types](1277-serial-and-bigserial-data-types.md).
3. For TEXT type conversions, see [TEXT and BYTE (LOB) types](1282-text-and-byte-lob-types.md).
