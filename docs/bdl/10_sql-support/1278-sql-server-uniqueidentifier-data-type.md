---
title: "SQL Server UNIQUEIDENTIFIER data type"
source: "fgl-topics/c_fgl_odiagmsv_007.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > SQL Server UNIQUEIDENTIFIER data type"
type: "concept"
description: "SQL Server supports a special type named UNIQUEIDENTIFIER, which can be used to store \"Globally Unique Identifiers\" (GUIDs). UNIQUEIDENTIFIER values can be generated with the NEWID() function. When ..."
---

# SQL Server UNIQUEIDENTIFIER data type

SQL Server supports a special type named UNIQUEIDENTIFIER, which can be used to store "Globally
Unique Identifiers" (GUIDs). UNIQUEIDENTIFIER values can be generated with the NEWID() function.
When creating a table, you typically define a UNIQUEIDENTIFIER column with a DEFAULT clause where
the value is produced from a NEWID() call:

```
CREATE TABLE mytab ( k INT, id UNIQUEIDENTIFIER DEFAULT NEWID(), c VARCHAR(10) )
```

The UNIQUEIDENTIFIER type is based on the BINARY(16) SQL Server
type. The Genero language does not have an equivalent type for BINARY(16).
However, BINARY values can be represented as hexadecimal strings in
CHAR or VARCHAR variables.

A UNIQUEIDENTIFIER value is usually represented as a GUID identifier, with the following
hexadecimal format:

```
XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
```

where X can be 0-9, A-F.

You typically fetch UNIQUEIDENTIFIER data into a CHAR(36) Genero
variable. The resulting value will be expressed in hexadecimal string
using the format. You can then reuse that value in an SQL statement,
but you have to convert the CHAR(36) hexadecimal string value back
to a UNIQUEIDENTIFIER value with the CONVERT() SQL Server function,
as shown in this example:

```
DEFINE pi CHAR(36)
CREATE TABLE mytab ( k INT, i UNIQUEIDENTIFIER DEFAULT NEWID(), c VARCHAR(10) )
INSERT INTO mytab ( k, c ) VALUES ( 1, 'aaa' )
SELECT i INTO pi FROM mytab WHERE k = 1
UPDATE mytab SET c = 'xxx' WHERE i = CONVERT(UNIQUEIDENTIFIER, pi)
```

When extracting a [database
schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions."), UNIQUEIDENTIFIER columns can be clearly distinguished
from BINARY(N) columns. The fgldbsch tool will produce a CHAR(36)
type code in the .sch file for UNIQUEIDENTIFIER columns.

You can also exclude the UNIQUEIDENTIFIER columns from the table
definition in the schema file, by using the x character at the appropriate
position of the string passed with the -cv data type conversion option
of [fgldbsch](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.").
