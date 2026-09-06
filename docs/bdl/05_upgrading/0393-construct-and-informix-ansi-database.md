---
title: "CONSTRUCT and Informix ANSI database"
source: "fgl-topics/c_fgl_MigI4GL_053.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > CONSTRUCT and Informix ANSI database"
type: "concept"
---

# CONSTRUCT and Informix ANSI database

> I4GL CONSTRUCT generates the ANSI-mode SQL table owner in the WHERE clause, while Genero BDL does not.

With IBM® Informix® 4GL, when connected to an Informix database create with ANSI mode, the WHERE clause
produced by `CONSTRUCT` will contain table owners in the
form:

```
"owner-name".table-name.column-name
```

Genero BDL will not produce the table owner
prefix:

```
table-name.column-name
```

As result, the current Informix user must be the owner of the SQL table used by a
`SELECT` statement and `WHERE` clause, since it does not specify the
table owner.

## Related links

**Related concepts**  

[Query by example (CONSTRUCT)](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.")
