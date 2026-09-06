---
title: "The LIST data type"
source: "fgl-topics/c_fgl_odiagifx_019.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Unsupported IBM® Informix® SQL features > The LIST data type"
type: "concept"
description: "With IBM® Informix®, the LIST type is a collection type that can store ordered elements of a specific base type. Unlike the MULTISET type, the elements of a LIST have ordinal positions. Elements can ..."
---

# The LIST data type

With IBM® Informix®, the LIST type is a collection
type that can store ordered elements of a specific base type. Unlike
the MULTISET type, the elements of a LIST have ordinal positions.
Elements can be duplicated.

Genero BDL does not support the IBM Informix LIST data type.

- It is not possible to define BDL variables with the LIST type.
- The static SQL syntax does not support collection-related syntax
  elements:
  - DDL statements like CREATE TABLE cannot use the LIST keyword for
    column types,
  - The collection-derived notation `TABLE()` is not
    allowed,
  - The `INSERT AT position` instruction
    is not supported,
  - The `LIST { }` literal syntax is not allowed.
  - The `value IN identifier` syntax
    is not allowed.
- The [fgldbsch](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.") schema
  extractor will report an invalid data type if you try to get the schema
  for a table with a LIST column.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
