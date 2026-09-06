---
title: "Direct database specification method"
source: "fgl-topics/c_fgl_Connections_013.html"
breadcrumb: "SQL support > Database connections > Direct database specification method"
type: "concept"
---

# Direct database specification method

> Genero BDL applies direct database source specification when no FGLPROFILE entry corresponds to the database name used in programs.

Direct database specification method takes place when the database name used in a
`DATABASE` or `CONNECT TO` instruction is not defined in [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") with a
'`dbi.database.dbname.source`' entry. In this case, the database specification
used in the connection instruction is used as the data source.

This method is well known with IBM®
Informix®, for example to specify the database server:

```
MAIN
  DATABASE stores@orion
  ...
END MAIN
```

In the next example, the database server is PostgreSQL. The string
used in the connection instruction defines the PostgreSQL database
(stores), the host (localhost), and the TCP service (5432) the postmaster
is listening to. As PostgreSQL syntax is not allowed in the language,
a CHAR variable must be used:

```
MAIN
  DEFINE db CHAR(50)
  LET db = "stores@localhost:5432"
  DATABASE db
  ...
END MAIN
```

This technique makes the compiled program specific to a given database
server configuration. Consider using indirect database specification method
instead of direct database specification.

## Related links

**Related concepts**  

[Indirect database specification method](1078-indirect-database-specification-method.md "Genero BDL allows to define database connection parameters in FGLPROFILE, that can be referenced by a single identifier in programs.")
