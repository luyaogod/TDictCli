---
title: "SQL identifiers syntax"
source: "fgl-topics/c_fgl_sql_programming_083.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Naming database objects > SQL identifiers syntax"
type: "concept"
---

# SQL identifiers syntax

> Database object naming conventions are different for each database engine.

The table below describes the naming conventions for database objects (tables,
sequences, stored procedures):

| Database Server Type | Naming Syntax |
| --- | --- |
| IBM® Informix® | [database[@dbservername]:][owner.]identifier |
| Microsoft™ SQL Server | [[[server.][database].][owner].]identifier |
| Oracle® MySQL / MariadDB | [database.]identifier |
| Oracle Database Server | [schema.]identifier[@database-link] |
| PostgreSQL | [[database.]schema.]identifier |
| SQLite | [database.]identifier |
| Dameng® | [schema.]identifier |

As a general rule, use simple unqualified database object names and use database vendor specific
configuration setting or SQL instruction to point to the expected schema.

When possible, Genero provides FGLPROFILE entries to define the database schema to be selected
after connecting to the database server,
with:

```
dbi.database.dbname.drivercode.schema = "schema-name"
```

For more details, see [Database type specific parameters in FGLPROFILE](1080-database-type-specific-parameters-in-fglprofile.md "Specific connection parameters can be configured with FGLPROFILE entries.").
