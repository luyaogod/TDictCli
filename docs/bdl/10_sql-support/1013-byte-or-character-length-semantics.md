---
title: "Byte or Character Length semantics?"
source: "fgl-topics/c_fgl_sql_programming_064.html"
breadcrumb: "SQL support > SQL programming > SQL portability > CHAR and VARCHAR types > Byte or Character Length semantics?"
type: "concept"
---

# Byte or Character Length semantics?

> Length Semantics defines the unit used to express the length of a character string, the position of a given character, and the size of a character data type.

When defining a `CHAR`/`VARCHAR` database column or program
variable, you must specify a size. When using a multibyte character set, the unit of this size
matters: it can be specified in bytes or characters.

In programs, the size unit of `CHAR`/`VARCHAR` variables depends on
the length semantics defined by the FGL\_LENGTH\_SEMANTICS environment variable.

In databases, the size unit of the `CHAR`/`VARCHAR` columns can be
expressed in bytes or characters, depending on the database server and its configuration.

## Related links

**Related concepts**  

[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md "Length semantics settings")
