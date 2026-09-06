---
title: "Specific CREATE INDEX clauses"
source: "fgl-topics/c_fgl_odiagifx_016.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > Specific CREATE INDEX clauses"
type: "concept"
description: "In addition to the standard index-key specification using a column list, the CREATE INDEX statement supported by IBM® Informix® SQL allows specific clauses, for example to define storage options. ..."
---

# Specific CREATE INDEX clauses

In addition to the standard index-key specification using a column
list, the CREATE INDEX statement supported by IBM® Informix®
SQL allows specific clauses, for example to define storage options.

Genero BDL partially supports the CREATE INDEX statement; the following
are not supported in static SQL grammar:

- The IF NOT EXISTS clause.
- Functional index specification is not allowed in the index-key list.
- Storage options such as IN *dbspace*, EXTEND SIZE, NEXT SIZE.
- The index mode clauses such as FILTERING WITH/WITHOUT ERROR.
- The USING clause.
- The HASH ON clause.
- The FILLFACTOR clause.

You can use [Dynamic SQL](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.") to
execute CREATE INDEX statements with clauses that are not allowed in the
static SQL grammar.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
