---
title: "Database entities"
source: "fgl-topics/c_fgl_sql_programming_057.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Database entities"
type: "concept"
---

# Database entities

> The database entity concept across different database engines.

Most database servers can handle multiple database entities (you
can create multiple 'databases'), but this is not possible with all
engines:

| Database Server Type | Multiple Database support |
| --- | --- |
| IBM® Informix® | Yes, [see details](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | Yes, [see details](1265-database-concepts.md) |
| Oracle® MySQL / MariadDB | Yes, [see details](1320-database-concepts.md) |
| Oracle Database Server | Yes, [see details](1365-database-concepts.md) |
| PostgreSQL | Yes, [see details](1421-database-concepts.md) |
| SQLite | Yes, [see details](1472-database-concepts.md) |
| Dameng® | Yes, [see details](1220-database-concepts.md) |

When using a database server that does not support multiple database entities, you
can emulate different databases with schema entities, but this requires you to check
for the database user definition. Each database user must have privileges to access
any schema, and to see any table of any schema without needing to set a schema prefix
before table names in SQL statements.

Some database drivers allow to select a specific schema at connection with the following
FGLPROFILE entry:

```
dbi.database.dbname.dbtype.schema = "schema-name"
```

Some databases also allow you to define a default schema for each database user.
When the user connects to the database, the default schema is automatically selected.

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
