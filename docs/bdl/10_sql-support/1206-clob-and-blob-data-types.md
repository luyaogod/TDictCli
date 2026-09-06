---
title: "CLOB and BLOB data types"
source: "fgl-topics/c_fgl_odiagifx_018.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Unsupported IBM® Informix® SQL features > CLOB and BLOB data types"
type: "concept"
description: "In addition to the TEXT and BYTE data types (known as Simple Large Objects), IBM® Informix® servers support the CLOB and BLOB types to store large objects. CLOB/BLOB are known as Smart Large Objects. ..."
---

# CLOB and BLOB data types

In addition to the TEXT and BYTE data types (known as Simple Large Objects),
IBM® Informix®
servers support the CLOB and BLOB types to store large objects. CLOB/BLOB are
known as Smart Large Objects. The main difference is that Smart Large Objects
support random access to the data - seek, read and write through the LOB as
if it was an OS file.

Genero BDL does not support the CLOB and BLOB types:

- It is not possible to define BDL variables with the CLOB or BLOB
  types, so you cannot manipulate CLOB/BLOB objects within programs.
- Defining a TEXT / BYTE variable to hold CLOB / BLOB column data
  is not supported; you will get error -609 (Illegal attempt to use
  a Text/Byte host variable).
- The static SQL syntax for DDL statements like CREATE TABLE does
  not allow the CLOB / BLOB keywords for column types.
- The [fgldbsch](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.") schema
  extractor will report an invalid data type if you try to get the schema
  for a table with a CLOB or BLOB column.

You can, however:

- Create a table with CLOB/BLOB columns by using [Dynamic SQL](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.").
- Use the Smart Large Object functions FILETOBLOB(), FILETOCLOB(),
  LOCOPY(), LOTOFILE() in static SQL statements.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
