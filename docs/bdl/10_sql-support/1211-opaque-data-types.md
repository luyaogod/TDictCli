---
title: "OPAQUE data types"
source: "fgl-topics/c_fgl_odiagifx_023.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Unsupported IBM® Informix® SQL features > OPAQUE data types"
type: "concept"
description: "Opaque User Defined Types can be implemented in IBM® Informix® with the CREATE OPAQUE TYPE statement. The storage structure of an OPAQUE type is unknown to the database server, data can only be ..."
---

# OPAQUE data types

Opaque User Defined Types can be implemented in IBM® Informix® with
the CREATE OPAQUE TYPE statement. The storage structure of an OPAQUE
type is unknown to the database server, data can only be accessed
through user-defined routines.

Genero BDL does not support the IBM Informix OPAQUE data types:

- It is not possible to define BDL variables with an opaque type.
- The static SQL syntax does not support OPAQUE-related syntax elements:
  - The DDL statements CREATE OPAQUE TYPE, DROP TYPE, CREATE CAST
    and DROP CAST are not allowed,
  - In CREATE TABLE / ALTER TABLE DDL statements, the data type must
    be a built-in type.
  - The :: cast operator is not supported. However, the `CAST()` expressions are
    allowed.
- The [fgldbsch](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.") schema
  extractor will report an invalid data type if you try to get the schema
  for a table with a column defined with a OPAQUE type.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
