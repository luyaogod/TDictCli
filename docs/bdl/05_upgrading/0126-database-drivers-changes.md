---
title: "Database drivers changes"
source: "fgl-topics/c_fgl_Migrate_to_401_database_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.01 upgrade guide > Database drivers changes"
type: "concept"
---

# Database drivers changes

> New and desupported database drivers.

## New database server versions supported in Genero 4.01:

Starting with FGL 4.01.03:

- Microsoft™ SQL Server 2022
- PostgreSQL 15

Starting with FGL 4.01.06:

- PostgreSQL 16
- Oracle® MySQL 8.2 (with new ODI driver
  `dbmmys_8_2`) - desupported in FGL 5.01.00
  > **Note:**
  >
  > The MySQL 8.2 and 8.3 versions are Innovation Releases (IR), while MySQL versions 8.0 and 8.4 are
  > Long-Term Support releases (LTS). The `dbmmys_8_2` driver provided in previous Genero
  > versions has been desupported in favor to MySQL 8.4 with the new `dbmmys_8_4` ODI
  > driver, requiring libmysqlclient.so.24. If you need to access a MySQL 8.2 or
  > 8.3 server, install MySQL Connector client with libmysqlclient.so.24 and use
  > the `dbmmys_8_4` ODI driver.

## New database clients supported in Genero 4.01:

- Microsoft ODBC 18 for SQL Server (new ODI driver
  `dbmsnc_18`):

  In this version, the generic `dbmsnc` driver name is an alias for
  `dbmsnc_18`. If you don't plan to install Microsoft ODBC 18 and keep using ODBC 17,
  you must use the `dbmsnc_17` ODI driver for Microsoft ODBC 17
  (libmsodbcsql-17.so).

  By default, Microsoft ODBC 18 for SQL Server enables connection encryption, that can lead to the following
  ODBC error when TLS/SSL certificates are not properly
  configured:

  ```
  [Microsoft][ODBC Driver 18 for SQL Server]SSL Provider: [error:1416F086:SSL routines:
  tls_process_server_certificate:certificate verify failed:self signed certificate]
  ```

  To
  disable encryption on Linux platforms, set the following ODBC option in the data source definition
  file:

  ```
  Encrypt = No
  ```

  On Microsoft Windows®, set "Connection Encryption" to "Optional" in the 4th
  panel of the ODBC data source configuration application.

- Oracle MySQL 8.2 client
  (libmysqlclient.so.22) with new ODI driver `dbmmys_8_2`.

  In
  this version, the generic `dbmmys` driver name is an alias for
  `dbmmys_8_2`. If you don't plan to install MySQL 8.2 and keep using MySQL 8.0, you
  must use the `dbmmys_8_0` ODI driver for MySQL 8.0 client lib
  (libmysqlclient.so.21).

  > **Warning:**
  >
  > The `dbmmys_8_2` ODI driver is desupported starting with FGL
  > 5.01.00

For more details, see [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md).

## Desupported databases server versions in Genero 4.01:

Desupported database server versions:

- Oracle Database 12.x
  (`dbmora_12` ODI driver is no longer provided)
- Generic ODBC ODI driver (`dbmodc_3`)

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Database drivers changes in BDL 4.00](0146-database-drivers-changes.md "New and desupported database drivers.").

Notable changes introduced in maintenance releases:

- No particular change to consider.
