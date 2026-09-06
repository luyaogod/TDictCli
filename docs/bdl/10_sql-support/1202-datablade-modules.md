---
title: "DataBlade modules"
source: "fgl-topics/c_fgl_odiagifx_015.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > DataBlade® modules"
type: "concept"
description: "IBM® Informix® IDS provides several database extensions implemented with the DataBlade® Application Programming Interface, such as MQ Messaging, Large OBjects management, Text Search DataBlades, ..."
---

# DataBlade modules

IBM® Informix® IDS provides several database extensions
implemented with the DataBlade® Application
Programming Interface, such as MQ Messaging, Large OBjects management,
Text Search DataBlades, Spatial DataBlade Module,
etc.

Genero BDL partially supports DataBlade modules:

- DataBlade extensions
  are based on User Defined Functions and User Defined Types. It is
  not possible to define program variables with specific User Defined
  Types. For example, you cannot define a program variable with the
  ST\_Point type implemented by the Spatial DataBlade module.
- The static SQL grammar does not support DataBlade specific syntax. For example,
  it is not possible to create a Basic Text Search index with the USING
  bts clause of the CREATE INDEX statement.

However, as long as the syntax of the DataBlade
functions follows basic SQL expressions, it can be used in static SQL statements. For example, this
query uses the `bts_contains()` function of the Basic Text Search
extension:

```
SELECT id FROM products WHERE bts_contains( brands, 'standard' )
```

You can also use [Dynamic SQL](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.") to
perform queries with a syntax that is not allowed in the static SQL
grammar.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
