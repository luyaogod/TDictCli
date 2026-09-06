---
title: "Extracting database schemas from PostgreSQL"
source: "fgl-topics/c_fgl_odiagpgs_044.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > Extracting database schemas from PostgreSQL"
type: "concept"
description: "The fgldbsch database schema extractor can only extract the schema from a PostgreSQL database if you specify the namespace of tables with the -ow option. PostgreSQL distinguishes table owners from ..."
---

# Extracting database schemas from PostgreSQL

The fgldbsch database schema extractor can only extract the schema from a
PostgreSQL database if you specify the namespace of tables with the `-ow` option.

PostgreSQL distinguishes table owners from table schemas, with the concept of
namespaces.

> **Tip:**
>
> In PostgreSQL system catalogs, the namespace of a table is defined by the
> `pg_class.relnamespace` column. This column contains the `oid` that
> references a namespace in the `pg_namespace` system table.

For PostgreSQL, the fgldbsch `-ow` option will specify the
PostgreSQL namespace, instead of the owner of the table, because a database user can
create several namespaces and use the same table name in those different namespaces. Filtering
schema extraction on the table owner would mix up table definitions from different
schemas/namespaces.

When extracting a database schema from a PostgreSQL database, you must specify the namespace of
tables with the `-ow` option. If no `-ow` option is provided and the
`-un` option is specified, fgldbsch will use the login name of the
`-un` option as namespace. If neither `-ow`, nor `-up`
options are specified, fgldbsch will use the PostgreSQL "public" namespace by
default.

Since database tables are usually created in the "public" namespace, you typically specify this
namespace with the `-ow`
option:

```
fgldbsch -db test1 -dv dbmpgs -un pgsuser -up fourjs -v -ow public
```

## Related links

**Related concepts**  

[Specifying the table schema/owner](../09_advanced-features/0804-specifying-the-table-schema-owner.md "Providing the database schema of SQL tables is mandatory when extracting a schema from some databases.")

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
