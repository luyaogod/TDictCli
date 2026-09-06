---
title: "SQL Server ROWVERSION data type"
source: "fgl-topics/c_fgl_odiagmsv_008.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > SQL Server ROWVERSION data type"
type: "concept"
description: "SQL Server provides a special type named ROWVERSION, to stamp row modifications. The ROWVERSION data type replaces the old TIMESTAMP column definition. When you define a column with the ROWVERSION, ..."
---

# SQL Server ROWVERSION data type

SQL Server provides a special type named ROWVERSION, to stamp row modifications. The ROWVERSION
data type replaces the old TIMESTAMP column definition. When you define a column with the
ROWVERSION, SQL Server will automatically increment the version column when the row is modified.
ROWVERSION is just an incrementing number, it does not preserve date or time information. It can be
used to control concurrent access to the same rows.

The ROWVERSION type is based on the BINARY(8) SQL Server type. The Genero language does not have
an equivalent type for BINARY(8). Therefore, you must fetch ROWVERSION data into a CHAR(16)
variable. The resulting value will be expressed in hexadecimal. You can then reuse that value in an
UPDATE statement to check that the row was not modified by another process, but you have to convert
the CHAR(16) hexadecimal value back to a BINARY(8) value with the CONVERT() SQL Server function, as
shown in this
example:

```
DEFINE pv CHAR(16) 
CREATE TABLE mytab ( k INT, v ROWVERSION, c VARCHAR(10) ) 
INSERT INTO mytab VALUES ( 1, NULL, 'aaa' ) 
SELECT v INTO pv FROM mytab WHERE k = 1 
UPDATE mytab SET c = 'xxx' WHERE k = 1 AND v = CONVERT(BINARY(8), pv, 2)
```

Since ROWVERSION is a synonym for BINARY(8), ROWVERSION columns
cannot be clearly identified in ODBC. Therefore, the following conversion
rule applies when fetching data from the server:

- If the column is defined as BINARY(N), with N<=128, the data
  will be fetched as a CHAR(N\*2), as an hexadecimal string.
- If the column is defined as BINARY(N), with N>128, the data
  will be fetched as a BYTE, as a regular binary value.

When extracting a [database
schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions."), ROWVERSION columns are identified as TIMESTAMP columns and can be
clearly distinguished from BINARY(N) columns. The fgldbsch tool will
produce a CHAR(16) type code in the .sch file for ROWVERSION or
TIMESTAMP columns.
