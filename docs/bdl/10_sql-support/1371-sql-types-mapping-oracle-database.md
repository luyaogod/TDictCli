---
title: "SQL types mapping: Oracle database"
source: "fgl-topics/c_fgl_odiagora_052.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > SQL types mapping: Oracle® database"
type: "concept"
description: "Table 1. SQL data types mapping for Oracle® database Original data types ORACLE data types CHAR(n) CHAR(n) (see note 1 ) VARCHAR(n[,m]) VARCHAR2(n) (see note 1 ) LVARCHAR(n) VARCHAR2(n) (see note 1 ) ..."
---

# SQL types mapping: Oracle database

| Original data types | ORACLE data types |
| --- | --- |
| `CHAR(n)` | `CHAR(n)` (see note 1) |
| `VARCHAR(n[,m])` | `VARCHAR2(n)` (see note 1) |
| `LVARCHAR(n)` | `VARCHAR2(n)` (see note 1) |
| `NCHAR(n)` | `NCHAR(n)` (see note 1) |
| `NVARCHAR(n[,m])` | `NVARCHAR2(n)` (see note 1) |
| `BOOLEAN` | `BOOLEAN` or `CHAR(1)` with Oracle 23c. Otherwise, `CHAR(1)` (see note 4) |
| `SMALLINT` | `NUMBER(5,0)` |
| `INTEGER` | `NUMBER(10,0)` |
| `BIGINT` | `NUMBER(20,0)` |
| `INT8` | `NUMBER(20,0)` |
| `SERIAL[(start)]` | `NUMBER(10,0)` (see note 2) |
| `BIGSERIAL[(start)]` | `NUMBER(20,0)` (see note 2) |
| `SERIAL8[(start)]` | `NUMBER(20,0)` (see note 2) |
| `DOUBLE PRECISION / FLOAT[(n)]` | `BINARY_DOUBLE` |
| `REAL / SMALLFLOAT` | `BINARY_FLOAT` |
| `DECIMAL(p,s)` | `NUMBER(p,s)` |
| `DECIMAL(p)` | `FLOAT(p*3.32193)` |
| `DECIMAL` (not recommended) | `FLOAT` |
| `MONEY(p,s)` | `NUMBER(p,s)` |
| `MONEY(p)` | `NUMBER(p,2)` |
| `MONEY` | `NUMBER(16,2)` |
| `TEXT` | `CLOB` (see note 3) |
| `BYTE` | `BLOB` (see note 3) |
| `DATE` | `DATE` |
| `DATETIME YEAR TO YEAR` | `DATE` |
| `DATETIME YEAR TO MONTH` | `DATE` |
| `DATETIME YEAR TO DAY` | `DATE` |
| `DATETIME YEAR TO HOUR` | `DATE` |
| `DATETIME YEAR TO MINUTE` | `DATE` |
| `DATETIME YEAR TO SECOND` | `DATE` |
| `DATETIME YEAR TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME MONTH TO MONTH` | `DATE` |
| `DATETIME MONTH TO DAY` | `DATE` |
| `DATETIME MONTH TO HOUR` | `DATE` |
| `DATETIME MONTH TO MINUTE` | `DATE` |
| `DATETIME MONTH TO SECOND` | `DATE` |
| `DATETIME MONTH TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME DAY TO DAY` | `DATE` |
| `DATETIME DAY TO HOUR` | `DATE` |
| `DATETIME DAY TO MINUTE` | `DATE` |
| `DATETIME DAY TO SECOND` | `DATE` |
| `DATETIME DAY TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME HOUR TO HOUR` | `DATE` |
| `DATETIME HOUR TO MINUTE` | `DATE` |
| `DATETIME HOUR TO SECOND` | `DATE` |
| `DATETIME HOUR TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME MINUTE TO MINUTE` | `DATE` |
| `DATETIME MINUTE TO SECOND` | `DATE` |
| `DATETIME MINUTE TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME SECOND TO SECOND` | `DATE` |
| `DATETIME SECOND TO FRACTION(n)` | `TIMESTAMP(n)` |
| `DATETIME FRACTION TO FRACTION(n)` | `TIMESTAMP(n)` |
| `INTERVAL YEAR[(p)] TO MONTH` | `INTERVAL YEAR[(p)] TO MONTH` |
| `INTERVAL MONTH[(p)] TO MONTH` | `INTERVAL YEAR(p-1) TO MONTH` |
| `INTERVAL DAY[(p)] TO FRACTION(n)` | `INTERVAL DAY[(p)] TO SECOND(n)` |
| `INTERVAL HOUR[(p)] TO HOUR` | `INTERVAL DAY(p-1) TO SECOND(0)` |
| `INTERVAL HOUR[(p)] TO MINUTE` | `INTERVAL DAY(p-1) TO SECOND(0)` |
| `INTERVAL HOUR[(p)] TO SECOND` | `INTERVAL DAY(p-1) TO SECOND(0)` |
| `INTERVAL HOUR[(p)] TO FRACTION(n)` | `INTERVAL DAY(p-1) TO SECOND(n)` |
| `INTERVAL MINUTE[(p)] TO MINUTE` | `INTERVAL DAY(p-3) TO SECOND(0)` |
| `INTERVAL MINUTE[(p)] TO SECOND` | `INTERVAL DAY(p-3) TO SECOND(0)` |
| `INTERVAL MINUTE[(p)] TO FRACTION(n)` | `INTERVAL DAY(p-3) TO SECOND(n)` |
| `INTERVAL SECOND[(p)] TO SECOND` | `INTERVAL DAY(p-4) TO SECOND(n)` |
| `INTERVAL SECOND[(p)] TO FRACTION(n)` | `INTERVAL DAY(p-4) TO SECOND(n)` |
| `INTERVAL FRACTION[(p)] TO FRACTION` | `INTERVAL DAY(0) TO SECOND(n)` |

Notes:

1. For character data types, see [CHAR and VARCHAR data types](1373-char-and-varchar-data-types.md).
2. For serial emulation, see [SERIAL and BIGSERIAL data types](1377-serial-and-bigserial-data-types.md).
3. For LOB data types, see [TEXT and BYTE (LOB) types](1380-text-and-byte-lob-types.md).
4. For BOOLEAN data type, see [BOOLEAN data type](1372-boolean-data-type.md).
