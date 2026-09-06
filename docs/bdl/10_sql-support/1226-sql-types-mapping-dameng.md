---
title: "SQL types mapping: Dameng"
source: "fgl-topics/c_fgl_odiagdmg_045.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data dictionary > SQL types mapping: Dameng®"
type: "concept"
description: "Table 1. SQL data types mapping for Dameng® Original data types Dameng data types CHAR(n) CHAR(n) (see note 1 ) VARCHAR(n[,m]) VARCHAR(n) (see note 1 ) LVARCHAR(n) VARCHAR(n) (see note 1 ) NCHAR(n) ..."
---

# SQL types mapping: Dameng

| Original data types | Dameng data types |
| --- | --- |
| `CHAR(n)` | `CHAR(n)` (see note 1) |
| `VARCHAR(n[,m])` | `VARCHAR(n)` (see note 1) |
| `LVARCHAR(n)` | `VARCHAR(n)` (see note 1) |
| `NCHAR(n)` | `CHAR(n)` (see note 1) |
| `NVARCHAR(n[,m])` | `VARCHAR(n[,m])` (note 1) |
| `BOOLEAN` | `BIT` |
| `SMALLINT` | `SMALLINT` |
| `INTEGER` | `INTEGER` |
| `BIGINT` | `BIGINT` |
| `INT8` | `BIGINT` |
| `SERIAL[(start)]` | `INTEGER` (see note 2) |
| `BIGSERIAL[(start)]` | `BIGINT` (see note 2) |
| `SERIAL8[(start)]` | `BIGINT` (see note 2) |
| `DOUBLE PRECISION / FLOAT[(n)]` | `DOUBLE` |
| `REAL / SMALLFLOAT` | `REAL` |
| `DECIMAL(p,s)` | `DECIMAL(p,s)` (max is 38 digits) |
| `DECIMAL[(p)]` | `DECIMAL` (floating point decimal) |
| `MONEY(p,s)` | `DECIMAL(p,s)` (max is 38 digits) |
| `MONEY(p)` | `DECIMAL(p,2)` (max is 38 digits) |
| `MONEY` | `DECIMAL(16,2)` |
| `DATE` | `DATE` |
| `DATETIME HOUR TO MINUTE` | `TIME(0)` |
| `DATETIME HOUR TO SECOND` | `TIME(0)` |
| `DATETIME HOUR TO FRACTION(n)` | `TIME(n)` |
| `DATETIME YEAR TO MONTH` | `TIMESTAMP(0)` |
| `DATETIME YEAR TO DAY` | `TIMESTAMP(0)` |
| `DATETIME YEAR TO HOUR` | `TIMESTAMP(0)` |
| `DATETIME YEAR TO MINUTE` | `TIMESTAMP(0)` |
| `DATETIME YEAR TO SECOND` | `TIMESTAMP(0)` |
| `DATETIME YEAR TO FRACTION(n)` | `TIMESTAMP(n)` |
| `INTERVAL YEAR(n) TO YEAR` | `INTERVAL YEAR(n)` |
| `INTERVAL YEAR(n) TO MONTH` | `INTERVAL YEAR(n) TO MONTH` |
| `INTERVAL MONTH(n) TO MONTH` | `INTERVAL MONTH(n)` |
| `INTERVAL DAY(n) TO DAY` | `INTERVAL DAY(n)` |
| `INTERVAL DAY(n) TO HOUR` | `INTERVAL DAY(n) TO HOUR` |
| `INTERVAL DAY(n) TO MINUTE` | `INTERVAL DAY(n) TO MINUTE` |
| `INTERVAL DAY(n) TO SECOND` | `INTERVAL DAY(n) TO SECOND(0)` |
| `INTERVAL DAY(n) TO FRACTION(p)` | `INTERVAL DAY(n) TO SECOND(p)` |
| `INTERVAL HOUR(n) TO HOUR` | `INTERVAL HOUR(n)` |
| `INTERVAL HOUR(n) TO MINUTE` | `INTERVAL HOUR(n) TO MINUTE` |
| `INTERVAL HOUR(n) TO SECOND` | `INTERVAL HOUR(n) TO SECOND(0)` |
| `INTERVAL HOUR(n) TO FRACTION(p)` | `INTERVAL HOUR(n) TO SECOND(p)` |
| `INTERVAL MINUTE(n) TO MINUTE` | `INTERVAL MINUTE(n)` |
| `INTERVAL MINUTE(n) TO SECOND` | `INTERVAL MINUTE(n) TO SECOND(0)` |
| `INTERVAL MINUTE(n) TO FRACTION(p)` | `INTERVAL MINUTE(n) TO SECOND(p)` |
| `INTERVAL SECOND(n) TO SECOND` | `INTERVAL SECOND(n)` |
| `INTERVAL SECOND(n) TO FRACTION(p)` | `INTERVAL SECOND(n,p)` |
| `INTERVAL FRACTION TO FRACTION(p)` | `INTERVAL SECOND(1,p)` |
| `TEXT` | `CLOB` (see note 3) |
| `BYTE` | `BLOB` (see note 3) |

Notes:

1. For character data types, see [CHAR and VARCHAR data types](1228-char-and-varchar-data-types.md)
2. For serial emulation, see [SERIAL and BIGSERIAL data types](1232-serial-and-bigserial-data-types.md).
3. For LOB data types, see [TEXT and BYTE (LOB) types](1234-text-and-byte-lob-types.md).
