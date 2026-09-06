---
title: "SQL types mapping: PostgreSQL"
source: "fgl-topics/c_fgl_odiagpgs_043.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > SQL types mapping: PostgreSQL"
type: "concept"
description: "Table 1. SQL data types mapping for PostgreSQL Original data types PostgreSQL data types CHAR(n) CHAR(n) VARCHAR(n[,m]) VARCHAR(n) LVARCHAR(n[,m]) VARCHAR(n) NCHAR(n) N/A NVARCHAR(n[,m]) N/A BOOLEAN ..."
---

# SQL types mapping: PostgreSQL

| Original data types | PostgreSQL data types |
| --- | --- |
| `CHAR(n)` | `CHAR(n)` |
| `VARCHAR(n[,m])` | `VARCHAR(n)` |
| `LVARCHAR(n[,m])` | `VARCHAR(n)` |
| `NCHAR(n)` | N/A |
| `NVARCHAR(n[,m])` | N/A |
| `BOOLEAN` | `BOOLEAN` |
| `SMALLINT` | `INT2` |
| `INTEGER` | `INT4` |
| `BIGINT` | `BIGINT` |
| `INT8` | `BIGINT` |
| `SERIAL[(start)]` | `INTEGER` (see note 1) |
| `BIGSERIAL[(start)]` | `BIGINT` (see note 1) |
| `SERIAL8[(start)]` | `BIGINT` (see note 1) |
| `DOUBLE PRECISION / FLOAT[(n)]` | `FLOAT4` |
| `REAL / SMALLFLOAT` | `FLOAT8` |
| `DECIMAL(p,s)` | `DECIMAL(p,s)` |
| `DECIMAL(p)` | `DECIMAL` (no precision = floating point) |
| `DECIMAL` | `DECIMAL` (no precision = floating point) |
| `MONEY(p,s)` | `DECIMAL(p,s)` |
| `MONEY(p)` | `DECIMAL(p,2)` |
| `MONEY` | `DECIMAL(16,2)` |
| `DATE` | `DATE` |
| `DATETIME HOUR TO MINUTE` | `TIME(0) WITHOUT TIME ZONE` |
| `DATETIME HOUR TO SECOND` | `TIME(0) WITHOUT TIME ZONE` |
| `DATETIME HOUR TO FRACTION(p)` | `TIME(p) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO MONTH` | `TIMESTAMP(0) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO DAY` | `TIMESTAMP(0) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO HOUR` | `TIMESTAMP(0) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO MINUTE` | `TIMESTAMP(0) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO SECOND` | `TIMESTAMP(0) WITHOUT TIME ZONE` |
| `DATETIME YEAR TO FRACTION(p)` | `TIMESTAMP(p) WITHOUT TIME ZONE` |
| `INTERVAL YEAR[(p)] TO MONTH` | `INTERVAL YEAR TO MONTH` |
| `INTERVAL YEAR[(p)] TO YEAR` | `INTERVAL YEAR` |
| `INTERVAL MONTH[(p)] TO MONTH` | `INTERVAL MONTH` |
| `INTERVAL DAY[(p)] TO FRACTION(n)` | `INTERVAL DAY TO SECOND(n)` |
| `INTERVAL DAY[(p)] TO SECOND` | `INTERVAL DAY TO SECOND(0)` |
| `INTERVAL DAY[(p)] TO MINUTE` | `INTERVAL DAY TO MINUTE` |
| `INTERVAL DAY[(p)] TO HOUR` | `INTERVAL DAY TO HOUR` |
| `INTERVAL DAY[(p)] TO DAY` | `INTERVAL DAY` |
| `INTERVAL HOUR[(p)] TO FRACTION(n)` | `INTERVAL HOUR TO SECOND(n)` |
| `INTERVAL HOUR[(p)] TO SECOND` | `INTERVAL HOUR TO SECOND(0)` |
| `INTERVAL HOUR[(p)] TO MINUTE` | `INTERVAL HOUR TO MINUTE` |
| `INTERVAL HOUR[(p)] TO HOUR` | `INTERVAL HOUR` |
| `INTERVAL MINUTE[(p)] TO FRACTION(n)` | `INTERVAL MINUTE TO SECOND(n)` |
| `INTERVAL MINUTE[(p)] TO SECOND` | `INTERVAL MINUTE TO SECOND(0)` |
| `INTERVAL MINUTE[(p)] TO MINUTE` | `INTERVAL MINUTE` |
| `INTERVAL SECOND[(p)] TO FRACTION(n)` | `INTERVAL SECOND(n)` |
| `INTERVAL SECOND[(p)] TO SECOND` | `INTERVAL SECOND(0)` |
| `INTERVAL FRACTION TO FRACTION(n)` | `INTERVAL SECOND(n)` |
| `TEXT` | `TEXT` |
| `BYTE` | `BYTEA` |

Notes:

1. For serial emulation, see [SERIAL and BIGSERIAL data types](1434-serial-and-bigserial-data-types.md).
