---
title: "Database drivers changes"
source: "fgl-topics/c_fgl_Migrate_to_500_database_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.00 upgrade guide > Database drivers changes"
type: "concept"
---

# Database drivers changes

> New and desupported database drivers.

## New database server versions supported in Genero 5.00:

- Support for [Oracle® database 23c](../10_sql-support/1359-oracle-database.md) with the new `dbmora_23` ODI
  driver.
- Support for [Dameng®
  database server version 8](../10_sql-support/1214-dameng-database-server.md) with the new `dbmdmg_8` ODI driver.
- Support for [MariaDB 11](../10_sql-support/1314-oracle-mysql-mariadb.md) with the existing
  `dbmmdb_10_2` ODI driver.

## New database clients supported in Genero 5.00:

- The new `dbmdmg_8` ODI driver for Dameng
  DPI version 8.

For more details, see [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md).

## Full desupport of IBM DB2 LUW

Starting with Genero BDL 5.00, the IBM DB2 LUW database is no longer supported.

The name of the ODI driver was `dbmdb2_10`.

## Full desupport of IBM Netezza

Starting with Genero BDL 5.00, the IBM Netezza database is no longer supported.

The name of the ODI driver was `dbmntz_6`.

## Desupported databases server versions in Genero 5.00:

Database server versions no longer supported:

- Microsoft™ SQL Server 2016. For supported versions,
  see [Microsoft SQL Server](../10_sql-support/1256-microsoft-sql-server.md).
- Oracle MySQL 5.7. For supported
  versions, see [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md).

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Database drivers changes in BDL 4.01](0126-database-drivers-changes.md "New and desupported database drivers.").

Notable changes introduced in maintenance releases:

- No particular change to consider.
