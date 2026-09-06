---
title: "The SET data type"
source: "fgl-topics/c_fgl_odiagifx_021.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Unsupported IBM® Informix® SQL features > The SET data type"
type: "concept"
description: "The SET IBM® Informix® data type is a collection type that stores non-ordered unique elements of a specific base type. Unlike the LIST type, the elements of a LIST have no ordinal positions. Elements ..."
---

# The SET data type

The SET IBM®
Informix® data type is a collection type that stores
non-ordered unique elements of a specific base type. Unlike the LIST type, the elements of a LIST
have no ordinal positions. Elements cannot be duplicated.

Genero BDL does not support the IBM
Informix SET data type:

- It is not possible to define BDL variables with the SET type.
- The static SQL syntax does not support collection-related syntax elements:
  - DDL statements like CREATE TABLE cannot use the SET keyword for column types,
  - The collection-derived notation `TABLE()` is not allowed,
  - The `SET { }` literal syntax is not allowed.
  - The `value IN identifier` syntax is not
    allowed.
- The [fgldbsch](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.") schema extractor will
  report an invalid data type if you try to get the schema for a table with a SET column.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
