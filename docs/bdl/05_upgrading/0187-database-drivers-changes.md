---
title: "Database drivers changes"
source: "fgl-topics/c_fgl_Migrate_to_310_database_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Database drivers changes"
type: "concept"
---

# Database drivers changes

> New and desupported database drivers.

## New database server versions supported in Genero 3.10:

Support for new database server releases, with existing ODI drivers:

- PostgreSQL 10 is supported with the existing `dbmpgs_9` driver (linked to
  libpq.so.5).

Support for new database server and client releases with new ODI drivers:

- Microsoft® ODBC for SQL Server
  v13 (`dbmsnc_13`) and v17 (`dbmsnc_17`).
- Oracle® MySQL 5.7
  (`dbmmys_5_7`) and MySQL 8.0 (`dbmmys_8_0`).
- Oracle 12c on macOS® with `dbmora_12`.
- MariaDB 10.2 with `dbmmdb_10_2`.

## Databases / ODI drivers desupported in Genero 3.10:

Database drivers for old database client versions are removed in accordance with DB vendor
de-support plans:

- SQL Server 2005 / SQL Native Client V9 (`dbmsnc_9`).
- Oracle MySQL 5.1.x
  (`dbmmys_5_1`)

## Related links

**Related concepts**  

[Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md "Database driver specification (driver)")
