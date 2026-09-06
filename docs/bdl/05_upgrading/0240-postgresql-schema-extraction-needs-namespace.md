---
title: "PostgreSQL schema extraction needs namespace"
source: "fgl-topics/c_fgl_Migrate_to_250_pgs_owner.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.50 upgrade guide > PostgreSQL schema extraction needs namespace"
type: "concept"
---

# PostgreSQL schema extraction needs namespace

> To extract a database schema from PostgreSQL, the fgldbsch tool now requires db namespace specification.

Starting with version 2.50, the fgldbsch database schema extractor will only
extract the schema from a PostgreSQL database if you specify the namespace with the
`-ow` option.

For more details, see [Extracting database schemas from PostgreSQL](../10_sql-support/1428-extracting-database-schemas-from-postgresql.md).
