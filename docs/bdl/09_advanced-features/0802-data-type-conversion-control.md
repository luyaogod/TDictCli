---
title: "Data type conversion control"
source: "fgl-topics/c_fgl_DatabaseSchema_007.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > Data type conversion control"
type: "concept"
description: "The fglcomp and fglform compilers expect known language data types (FGL types) in the schema file. While most data types correspond to IBM® Informix® SQL data types, some databases (including ..."
---

# Data type conversion control

The fglcomp and fglform compilers expect known language
data types (FGL types) in the schema file. While most data types correspond to IBM®
Informix® SQL data types, some databases (including Informix) can use specific types that do not map to an FGL
type. Therefore, data types in the schema file are generated from the system catalog tables based on
some conversion rules.

Type conversion can be controlled with the `-cv` option.
Each character position of the string passed by this option represents
a line in the conversion table of the corresponding source database.
Give a conversion code for each data type (for example: `-cv
AABAAAB`).

When using `X` as conversion code, the columns using the corresponding data types
will be ignored and not written to the .sch file. This is particularly useful
in the case of auto-generated columns like SQL Server's *uniqueidentifier* data type, when
using a `DEFAULT NEWID()` clause.

Run the tool with the `-ct` option to see all the
data type conversion tables, or use the `-cx dbtype` option
to display the conversion table for a given database type (*dbtype* must
be ifx, ora, msv, pgs, mys, ...).

```
fgldbsch -cx ifx
...
------------------------------------------------------------------
Informix        Informix A          Informix B
------------------------------------------------------------------
1 BOOLEAN       BOOLEAN (t=45)      CHAR(1)
2 INT8          INT8                DECIMAL(19,0)
3 SERIAL8       SERIAL8             DECIMAL(19,0)
4 LVARCHAR(m)   VARCHAR2(m)         VARCHAR2(m)
5 BIGINT        BIGINT              DECIMAL(19,0)
6 BIGSERIAL     BIGSERIAL           DECIMAL(19,0)
------------------------------------------------------------------
(ns) = Not supported in 4gl.
...
 
fgldbsch -db test1 -cv BAAABB
                       123456
```

In the above example, the -`cv` option instructs fgldbsch to
use the types of the "Informix A" column for all original
column types except for BOOLEAN, BIGINT and BIGSERIAL, which must be converted to a
`VARCHAR2(m)` FGL type.

The IBM Informix LVARCHAR(m) type can be converted
by default to a `VARCHAR2(m)` pseudo type (code
201), which will be identified as a VARCHAR(m) by compilers.

In schema files, `VARCHAR2(m)` (type code 201) is
equivalent to `VARCHAR(m)` (type code 13), without
the 255 bytes limitation of the original Informix VARCHAR type.

Not all native data types can be converted to FGL types. For example, user-defined types or
spatial types are not supported by the language. When a table column with such unsupported data type
is found, fgldbsch stops and displays an error to bring the problem to your
attention. Use the -`ie` option of fgldbsch to ignore the database
tables having columns with unsupported types. When this option is used, none of the table columns
definition will be written to the schema file.

## Related links

**Related concepts**  

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")

[Column Definition File (.sch)](0795-column-definition-file-sch.md "The .sch database schema file contains the data types of database table columns.")
