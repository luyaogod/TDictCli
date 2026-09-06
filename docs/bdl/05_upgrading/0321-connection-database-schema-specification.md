---
title: "Connection database schema specification"
source: "fgl-topics/c_fgl_Migrate_to_200_dbi_schema.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Connection database schema specification"
type: "concept"
---

# Connection database schema specification

> Changes with FGLPROFILE entries to define the database schema at runtime.

Before version 2.00, an [FGLPROFILE](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.") entry was
specified to define the database schema at runtime:

```
dbi.database.dbname.schema = "schema-name"
```

This entry was used to select the native database schema after connecting to the server, for
example with Oracle® database.

Starting with version 2.00, this entry must now be specified by prefixing with the database
driver code, like "ora":

```
dbi.database.dbname.ora.schema = "schema-name"
```

For other database servers, this configuration parameter is not defined.

> **Important:**
>
> It is no longer possible to specify the "schema" parameter in the
> connection string (dbname+schema='name').

## Related links

**Related concepts**  

[Database connections](../10_sql-support/1059-database-connections.md "Explains how to manage database connections in a program.")
