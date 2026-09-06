---
title: "Database drivers changes"
source: "fgl-topics/c_fgl_Migrate_to_400_database_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Database drivers changes"
type: "concept"
---

# Database drivers changes

> New and desupported database drivers.

## New database server versions supported in Genero 4.00:

- PostgreSQL 13 and 14 (using ODI driver `dbmpgs_9`)
- Oracle® 21c (using ODI driver
  `dbmora_18`)

For more details, see [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md).

## New embedded SQLite version: 3.27.2

On Operating Systems where the SQLite engine is embedded in the `dbmsqt_3` ODI
driver, SQLite has been upgraded to version 3.27.2.

## Full desupport of SAP® Adaptive Server Enterprise

Starting with Genero BDL 4.00, the SAP
Adaptive Server Enterprise database server is fully desupported.

The name of the ODI driver was `dbmase_16`.

## Desupported databases server versions in Genero 4.00:

Desupported database server versions:

- Oracle Database 11.x
  (`dbmora_11` ODI driver removed)
- Microsoft™ SQL Server 2012, 2014
  (`dbmsnc_11` ODI driver removed)
- Oracle MySQL 5.6
  (`dbmmys_5_6` ODI driver removed)
- PostgreSQL 9.x and 10 (`dbmpgs_9` remains for PostgreSQL versions > 10)

## Desupported ODI drivers in Genero 4.00:

Desupported database drivers:

- Microsoft SQL Native Client 11:
  `dbmsnc_11` (use `dbmsnc_17`)
- Microsoft ODBC for SQL Server 13:
  `dbmsnc_13` (use `dbmsnc_17`)
- Oracle MySQL 5.6 client:
  `dbmmys_5_6` (use `dbmmys_5_7` or `dbmmys_8_0`)

> **Note:**
>
> For Microsoft SQL Server, use Microsoft ODBC 17 and the `dbmsnc_17` ODI driver, on Linux
> and Windows platforms.

## Required SQL Server ODBC drivers for UTF-8 char/varchar

For \_UT8 database collation support in char/varchar columns of SQL Server, you need to install
recent Microsoft ODBC, Easysoft ODBC and FreeTDS ODBC drivers.

See [SQL Server UTF-8 support](0148-sql-server-utf-8-support.md "Support for UTF-8 collation in CHAR/VARCHAR columns with SQL Server 2019.") for more details.

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Database drivers changes in BDL 3.20](0164-database-drivers-changes.md "New and desupported database drivers.").
