---
title: "The ROW data types"
source: "fgl-topics/c_fgl_odiagifx_022.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Unsupported IBM® Informix® SQL features > The ROW data types"
type: "concept"
description: "IBM® Informix® supports the named and unnamed ROW data types. A ROW type is a complex type that combines several table columns. You create a ROW type with the CREATE ROW TYPE instruction, and then you ..."
---

# The ROW data types

IBM®
Informix® supports the named and unnamed ROW data types.
A ROW type is a complex type that combines several table columns. You create a ROW type with the
CREATE ROW TYPE instruction, and then you can reuse the type definition for a table column.

Genero BDL does not support the IBM
Informix ROW data types:

- It is not possible to define BDL variables with a named ROW type. The equivalent would be a
  RECORD variable, but data is not mapped directly from a structured ROW column, you must list
  individual fields of the ROW column.
- The static SQL syntax does not support ROW-related syntax elements:
  - The DDL statements CREATE ROW TYPE, DROP ROW TYPE, CREATE CAST and DROP CAST are not
    allowed,
  - In CREATE TABLE / ALTER TABLE DDL statements, the data type must be a built-in type.
  - The :: cast operator is not supported when specifying a ROW() literal. However, the CAST()
    expressions are allowed.
- The [fgldbsch](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.") schema extractor will
  report an invalid data type if you try to get the schema for a table with a column defined with a
  ROW type.

However:

- Static SQL allows multilevel single-dot notation, so you can, for example, identify a ROW field
  as *employee.address.city*.
- Dynamic SQL can be used to insert or update rows with ROW type columns.
- Individual ROW column fields can be fetched to BDL program variables, as long as the basic types
  match.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")

[Dynamic SQL management](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")
