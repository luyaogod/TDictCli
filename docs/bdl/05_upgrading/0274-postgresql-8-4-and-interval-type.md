---
title: "PostgreSQL 8.4 and INTERVAL type"
source: "fgl-topics/c_fgl_Migrate_to_221_postgresql_interval.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.21 upgrade guide > PostgreSQL 8.4 and INTERVAL type"
type: "concept"
---

# PostgreSQL 8.4 and INTERVAL type

> The dbmpgs84x database driver requires your database schema use the INTERVAL type, rather than a CHAR(50) type.

Version 2.21 introduced support for PostgreSQL 8.4 with the new database driver
`dbmpgs84x`. This version of PostgreSQL implements a native INTERVAL data type
that is similar to the Genero Business Development Language INTERVAL type.

When using the `dbmpgs84x` (and higher) driver, Informix-style INTERVAL types will
be mapped / translated to native PostgreSQL INTERVALs. Prior drivers will keep using the CHAR(50)
replacement. If your application is storing INTERVAL in a PostgreSQL database, you will have to
modify you database schema to replace the existing CHAR(50) column with the native INTERVAL data
type of PostgreSQL 8.4. If you cannot migrate the database, you can still use the older dbmpgs83x
driver using CHAR(50) for INTERVALs, but that driver requires a PostgreSQL client version
8.3.

## Related links

**Related concepts**  

[PostgreSQL](../10_sql-support/1415-postgresql.md "PostgreSQL")
