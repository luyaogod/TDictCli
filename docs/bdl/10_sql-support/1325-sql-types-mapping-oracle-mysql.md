---
title: "SQL types mapping: Oracle MySQL"
source: "fgl-topics/c_fgl_odiagmys_035.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data dictionary > SQL types mapping: Oracle® MySQL"
type: "concept"
description: "Table 1. SQL data types mapping Oracle® MySQL Original data types Oracle MySQL data types CHAR(n) CHAR(n) or TEXT (see note 1 ) VARCHAR(n[,m]) VARCHAR(n) (see note 1 ) LVARCHAR(n) VARCHAR(n) (see note ..."
---

# SQL types mapping: Oracle MySQL

| Original data types | Oracle MySQL data types |
| --- | --- |
| `CHAR(n)` | `CHAR(n)` or `TEXT` (see note 1) |
| `VARCHAR(n[,m])` | `VARCHAR(n)` (see note 1) |
| `LVARCHAR(n)` | `VARCHAR(n)` (see note 1) |
| `NCHAR(n)` | `NCHAR(n)` (see note 1) |
| `NVARCHAR(n[,m])` | `NVARCHAR(n)` (see note 1) |
| `BOOLEAN` | `BOOLEAN` |
| `SMALLINT` | `SMALLINT` |
| `INTEGER` | `INTEGER` |
| `BIGINT` | `BIGINT` |
| `INT8` | `BIGINT` |
| `SERIAL[(start)]` | `INTEGER` (see note 2) |
| `BIGSERIAL[(start)]` | `BIGINT` (see note 2) |
| `SERIAL8[(start)]` | `BIGINT` (see note 2) |
| `DOUBLE PRECISION / FLOAT[(n)]` | `DOUBLE` |
| `REAL / SMALLFLOAT` | `FLOAT` |
| `DECIMAL(p,s)` | `DECIMAL(p,s)` |
| `DECIMAL(p<=16)` | `DECIMAL(p*2,p)` |
| `DECIMAL(p>16)` | N/A |
| `DECIMAL` | `DECIMAL(32,16)` |
| `MONEY(p,s)` | `DECIMAL(p,s)` |
| `MONEY(p)` | `DECIMAL(p,2)` |
| `MONEY` | `DECIMAL(16,2)` |
| `DATE` | `DATE` |
| `DATETIME HOUR TO MINUTE` | `TIME` |
| `DATETIME HOUR TO SECOND` | `TIME` |
| `DATETIME HOUR TO FRACTION(p)` | `TIME(p)` |
| `DATETIME YEAR TO MONTH` | `DATETIME` |
| `DATETIME YEAR TO DAY` | `DATETIME` |
| `DATETIME YEAR TO HOUR` | `DATETIME` |
| `DATETIME YEAR TO MINUTE` | `DATETIME` |
| `DATETIME YEAR TO SECOND` | `DATETIME` |
| `DATETIME YEAR TO FRACTION(p)` | `DATETIME(p)` |
| `INTERVAL q1 TO q2` | `CHAR(50)` |
| `TEXT` | `MEDIUMTEXT / LONGTEXT` (max is 2Gb) |
| `BYTE` | `MEDIUMBLOB / LONGBLOB` (max is 2Gb) |

Notes:

1. The `CHAR` types with a size > 255 are converted `TEXT` types. For
   more details, see [CHAR and VARCHAR data types](1327-char-and-varchar-data-types.md).
2. For serial emulation, see [SERIAL and BIGSERIAL data type](1331-serial-and-bigserial-data-type.md).

## Related links

**Related tasks**  

[Install Oracle MySQL/MariaDB and create a database - database configuration/design tasks](1317-install-oracle-mysql-mariadb-and-create-a-database-database.md "Install Oracle MySQL/MariaDB and create a database - database configuration/design tasks")
