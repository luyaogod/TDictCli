---
title: "PostgreSQL 12 notes"
source: "fgl-topics/c_fgl_Migrate_to_320_postgresql_12.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > PostgreSQL 12 notes"
type: "concept"
---

# PostgreSQL 12 notes

> This topics contains notes about PostgreSQL 12 changes that affect Genero applications.

## Desupport of OID columns

Genero BDL 3.20 supports PostgreSQL 12.

Starting with PostgreSQL version 12, the OID columns (similar to Informix ROWID columns) are no
longer supported.

If your application is using ROWIDs, you must review the code.

For more details, see [ROWID columns](../10_sql-support/1435-rowid-columns.md).

## Related links

**Related concepts**  

[PostgreSQL](../10_sql-support/1415-postgresql.md "PostgreSQL")
