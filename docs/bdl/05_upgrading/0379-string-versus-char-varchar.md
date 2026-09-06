---
title: "STRING versus CHAR/VARCHAR"
source: "fgl-topics/c_fgl_MigI4GL_027.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > STRING versus CHAR/VARCHAR"
type: "concept"
---

# STRING versus CHAR/VARCHAR

> Genero BDL supports the STRING data type in addition to CHAR and VARCHAR. While the STRING data type is useful in certain situations, there are times when you should use CHAR or VARCHAR instead.

Genero Business Development Language (BDL) introduces a new data type named
`STRING`, which is similar to `VARCHAR` but without a size limit. The
`STRING` data type does not exist in IBM®
Informix® 4GL. The `STRING` data type
implementation is optimized for memory usage; unlike `CHAR`/`VARCHAR`,
BDL will only allocate the memory needed to hold the actual character string value in a
`STRING` variable.

A `STRING` variable is typically used within utility functions (for example, to
hold the path to a file). Another typical usage is with `CONSTRUCT`, to hold the SQL
condition. The `STRING` variable can then be completed to build the SQL text and
passed to the `PREPARE` or `DECLARE` instruction.

`STRING` variables can be used in SQL statements. However, it is recommended to
use the `CHAR`/`VARCHAR` variables that match the size of the
corresponding `CHAR`/`VARCHAR` database columns.

The `STRING` data type has a number of built-in methods that are very useful and
will reduce source code, such as `getLength()`.

## Related links

**Related concepts**  

[STRING](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.")
