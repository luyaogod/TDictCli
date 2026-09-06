---
title: "The MULTISET data type"
source: "fgl-topics/c_fgl_odiagifx_020.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Unsupported IBM® Informix® SQL features > The MULTISET data type"
type: "concept"
description: "The MULTISET IBM® Informix® data type is a collection type that can store non-ordered elements of a specific base type. Unlike the LIST type, the elements of a MULTISET have no ordinal positions. ..."
---

# The MULTISET data type

The MULTISET IBM®
Informix® data type is a collection type that can store
non-ordered elements of a specific base type. Unlike the LIST type, the elements of a MULTISET have
no ordinal positions. Elements can be duplicated.

Genero BDL does not support the IBM
Informix MULTISET data type:

- It is not possible to define BDL variables with the MULTISET type.
- The static SQL syntax does not support collection-related syntax elements:
  - DDL statements like CREATE TABLE cannot use the MULTISET keyword for column types,
  - The collection-derived notation `TABLE()` is not allowed,
  - The `MULTISET { }` literal syntax is not allowed.
  - The `value IN identifier` syntax is not
    allowed.
- The [fgldbsch](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.") schema extractor will
  report an invalid data type if you try to get the schema for a table with a MULTISET column.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
