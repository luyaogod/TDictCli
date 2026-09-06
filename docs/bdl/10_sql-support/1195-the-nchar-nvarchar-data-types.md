---
title: "The NCHAR / NVARCHAR data types"
source: "fgl-topics/c_fgl_odiagifx_008.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > The NCHAR / NVARCHAR data types"
type: "concept"
description: "IBM® Informix® supports the standard NCHAR and NVARCHAR data types. These types are equivalent to CHAR and VARCHAR (the same character set is used), except that the collation order is locale specific ..."
---

# The NCHAR / NVARCHAR data types

IBM® Informix®
supports the standard NCHAR and NVARCHAR data types. These types are equivalent to CHAR and VARCHAR
(the same character set is used), except that the collation order is locale specific with
NCHAR/NVARCHAR types.

Genero BDL syntax does allow to define program variables by using NCHAR / NVARCHAR keywords.
However, the character strings of Informix NCHAR/NVARCHAR database columns can be managed by program
variables defined with the CHAR/VARCHAR types.
> **Note:**
>
> Since the character set is identical for
> NCHAR/NVARCHAR and CHAR/VARCHAR columns in an Informix database, no specific consideration needs to
> be given for the "N" character types.

When extracting a database schema with [fgldbsch](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages."), NCHAR/NVARCHAR types will be identified in the .sch file by
the native Informix type codes 15 and 16. When compiling
.4gl or .per sources referencing NCHAR/NVARCHAR columns in
the schema file, the compilers will automatically use the CHAR/VARCHAR Genero BDL types for the type
codes 15 and 16.

## Related links

**Related concepts**  

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")
