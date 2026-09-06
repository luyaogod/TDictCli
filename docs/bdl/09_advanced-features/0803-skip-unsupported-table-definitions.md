---
title: "Skip unsupported table definitions"
source: "fgl-topics/c_fgl_DatabaseSchema_021.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > Skip unsupported table definitions"
type: "concept"
description: "By default fgldbsch stops with an error, if a database table column is defined with an SQL data type that cannot be supported by Genero BDL. Use the -ie option, to skip such tables from the schema ..."
---

# Skip unsupported table definitions

By default fgldbsch stops with an error, if a database table column is defined
with an SQL data type that cannot be supported by Genero BDL.

Use the `-ie` option, to skip such tables from the schema extraction.

```
fgldbsch -db test1 -ie
```

When using the `-ie` option, the whole table definition is skipped when a column
type is not supported. Otherwise, the .sch file would contain a subset of the
column deinitions of the table, and a `SELECT * INTO record-defined-like.*
FROM tabname` would fail because the number of columns in the table and
into the record definition would not match.

## Related links

**Related concepts**  

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
