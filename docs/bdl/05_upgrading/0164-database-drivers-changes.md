---
title: "Database drivers changes"
source: "fgl-topics/c_fgl_Migrate_to_320_database_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Database drivers changes"
type: "concept"
---

# Database drivers changes

> New and desupported database drivers.

## New database server versions supported in Genero 3.20:

Support for new database server releases, with existing ODI drivers:

- PostgreSQL 11 and 12 is supported with the existing `dbmpgs_9` driver (linked to
  libpq.so.5) - PostgreSQL 13 and 14 support has been backported from FGL
  4.01.
- Microsoft SQL Server 2019 (v15) is supported with the existing `dbmsnc_17` driver
  (linked to libmsodbcsql-17.so).

Support for new database server and client releases with new ODI drivers:

- Oracle® 18c database driver
  (`dbmora_18`/`dbmora`) for Oracle 18c database server.
  > **Important:**
  >
  > The generic driver name `dbmora` maps now to
  > `dbmora_18`. If your application is using Oracle 12c or Oracle 11g, specify
  > respectively `dbmora_12` or `dbmora_11` drivers in the connection
  > parameters.
- Oracle 19c database server support, by
  using the (`dbmora_18`/`dbmora`) driver. On Unix platforms, the
  `dbmora_18` driver is linked with libclntsh.so.18.1, which is
  provided by Oracle Instant Client 19c.

ODI driver replacements:

- Oracle MySQL 5.6 driver
  (`dbmmys_5_6`) has to be used in place of the desupported `dbmmys_5_5`
  driver.

For more details, see [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md).

## Databases / ODI drivers desupported in Genero 3.20:

Database drivers for old database client versions are removed in accordance with DB vendor
de-support plans:

- Microsoft™ SQL Server 2008 / SQL Native Client V10
  (`dbmsnc_10`)
- Oracle MySQL 5.5.x
  (`dbmmys_5_5`) - for MySQL 5.6, use the new driver
  `dbmmys_5_6`.
  > **Note:**
  >
  > The `dbmmys_5_5` driver could also be used to
  > connect to MariadDB 10.0 and 10.1. These MariadDB versions are no longer supported: Use MariadDB
  > 10.2+ with `dbmmdb_10_2`.

## Minimum required FreeTDS version is 1.00.104

See [FreeTDS 1.00 for SQL Server](0166-freetds-1-00-for-sql-server.md "Genero 3.20 requires FreeTDS version 1.00+ to connect to SQL Server.") for more details.

## FGLSQLDEBUG levels and output

Starting with version 3.20, the [FGLSQLDEBUG](../07_configuration/0534-fglsqldebug.md "Defines the debug level for tracing SQL instructions.") and [`fgl_sqldebug()`](../15_library-reference/2783-fgl-sqldebug.md "Sets the SQL debug level from program code.") maximum debug level is 3. If you set the SQL debug level to
a higher value, it will be equivalent to level 3.

When FGLSQLDEBUG is enabled, the [`OPTIONS SQL
INTERRUPT ON/OFF`](../10_sql-support/0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.") is traced.
