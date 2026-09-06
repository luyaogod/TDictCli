---
title: "Specifying the table schema/owner"
source: "fgl-topics/c_fgl_DatabaseSchema_008.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > Specifying the table schema/owner"
type: "concept"
---

# Specifying the table schema/owner

> Providing the database schema of SQL tables is mandatory when extracting a schema from some databases.

If tables with the same name exist in different database schemas, you can get multiple
definitions for the same table name in a single .sch database schema file. The
key to find a column definition in a .sch file is the table name and the column
name. There is no concept of table schema/owner in the .sch file structure to
support different SQL table definitions using the same table name.

To prevent such duplicates, specify the database schema/owner of the tables with the
`-ow` option. If this option is not used, fgldbsch will use the
database login name passed with the `-un` option as the default schema/owner.

If the table schema/owner is mandatory, fgldbsch will ask for it. If no tables
are found for a given schema/owner, fgldbsch will report that no tables were
found and that you need to provide the correct table schema/owner.

For example, with Microsoft® SQL
Server, the owner of tables is usually `dbo`; therefore, you want to include
`-ow dbo` in the fgldbsch command:

```
fgldbsch -db test1 -un scott -up fourjs -ow dbo
```

## Related links

**Related concepts**  

[Database connections](../10_sql-support/1059-database-connections.md "Explains how to manage database connections in a program.")

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
