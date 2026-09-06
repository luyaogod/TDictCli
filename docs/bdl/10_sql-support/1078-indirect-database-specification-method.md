---
title: "Indirect database specification method"
source: "fgl-topics/c_fgl_Connections_014.html"
breadcrumb: "SQL support > Database connections > Indirect database specification method"
type: "concept"
---

# Indirect database specification method

> Genero BDL allows to define database connection parameters in FGLPROFILE, that can be referenced by a single identifier in programs.

Indirect database specification method takes place when the database name used in
the [`DATABASE`](1096-database.md "Opens a new database connection in unique-session mode.") or [`CONNECT TO`](1099-connect-to.md "Opens a new database session in multi-session mode.") instruction corresponds to a
'`dbi.database.dbname.source`' entry defined in the [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") configuration file. In this case, the
dbname database specification is used as a key to read the connection information
from the configuration file.

In FGLPROFILE, the entries starting with '`dbi.database`' group information
defining connection parameters for indirect database specification:

```
dbi.database.dbname.source   = "value" 
dbi.database.dbname.driver   = "value" 
dbi.database.dbname.username = "value" 
dbi.database.dbname.password = "value"  
-- Warning: Password is not encrypted, do not use in production!
```

FGLPROFILE entry names are converted to lower case when loaded by the runtime system. In order
to avoid database name matching mistakes, it is mandatory to write FGLPROFILE entry names and
program database names in lower case.

In the next example, the program specifies the name `mydb` in the `CONNECT
TO` instruction, referecing the FGLPROFILE `dbi.database.mydb.source` and
`dbi.database.mydb.driver` entries, to connect to the `stores`
database controlled by a PostgreSQL server running on the `localhost`, using TCP port
`5432`:

Program:

```
MAIN
  DEFINE un, up STRING
  ...
  CONNECT TO "mydb" USER un USING up
  ...
END MAIN
```

FGLPROFILE:

```
dbi.database.mydb.source   = "stores@localhost:5432"
dbi.database.mydb.driver   = "dbmpgs"
```

The indirect database specification technique is a flexible technique to define the database
source: The database name in programs is a kind of alias for the real data source defined in an
external FGLPROFILE configuration file, where entries can be easily changed on production sites
without needing program recompilation.

## Related links

**Related concepts**  

[Connection parameters](1071-connection-parameters.md "This section describes the different parameters which need to be specified in order to connect to a database.")

[Direct database specification method](1077-direct-database-specification-method.md "Genero BDL applies direct database source specification when no FGLPROFILE entry corresponds to the database name used in programs.")
