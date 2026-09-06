---
title: "Database drivers changes"
source: "fgl-topics/c_fgl_Migrate_to_501_database_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.01 upgrade guide > Database drivers changes"
type: "concept"
---

# Database drivers changes

> New and desupported database drivers.

## Relax Microsoft ODBC for SQL Server Wide Char setting

The `dbmsnc*` ODI driver build with Microsoft ODBC for SQL Server uses
CHAR/VARCHAR type binding, or NCHAR/NVARCHAR type binding for ODBC SQL parameters. This feature is
controlled by the `dbi.database.dbname.snc.widechar` FGLPROFILE
setting, which had to be set manually in earlier Genero version. Since version 3.10, in a regular
configuration, there is no need to set the `snc.widechar` parameter: According to the
context (app locale and database collation), the ODI driver will automatically use the appropriate
character type for ODBC bindings. See [Wide Char mode of SNC driver](0201-wide-char-mode-of-snc-driver.md "The snc.widechar FGLPROFILE entry defaults to the right setting for the current application locale.").

Before Genero version 5.01.06, for an app using a UTF-8 locale, the ODI driver was reporting a
runtime error -6319 at connection time, if `snc.widechar=false` (assuming
CHAR/VARCHAR columns usage), but the SQL Server database collation for CHAR/VARCHAR columns is
not UTF-8. In fact, this configuration is unexpected, because UTF-8 characters manipulated by
the Genero application cannot be stored in a non-UTF-8 database.

Starting with Genero version 5.01.06 (change FGL-6509), to simplify migration from non-UTF-8 to
UTF-8, the ODI driver allows the connection, and only reports a warning in the FGLSQLDEBUG output,
when the app the locale is UTF-8, `snc.widechar=false`, and the DB collation is
not UTF-8.

> **Important:**
>
> Setting FGLPROFILE `dbi.database.dbname.snc.widechar` to
> `false` is strongly discouraged and done at your own risk, when using a UTF-8
> application locale with a non-UTF-8 database collation for CHAR/VARCHAR column storage: Any UTF-8
> characters that cannot be represented in the database's collation will be converted to a question
> mark (`?`) when strings are sent to the database.

For more details on UTF-8 support with Genero and SQL Server, read carefully [CHAR and VARCHAR data types](../10_sql-support/1273-char-and-varchar-data-types.md).

## Minimum versions of ODBC clients for SQL Server

Starting with Genero version 5.01.06, the SQL Server ODI drivers require the following ODBC
client versions:

- `dbmsnc_18`: Microsoft ODBC 18.5.2 or +
- `dbmesm_1`: Easysoft ODBC 2.3.0 or +
- `dbmftm_0`: FreeTDS ODBC 1.5.168 or +

For more details, see [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md), [Prepare the runtime environment - connecting to the database](../10_sql-support/1260-prepare-the-runtime-environment-connecting-to-the-database.md).

## New database server versions supported in Genero 5.01:

- Support for [Oracle® MySQL 8.4 LTS](../10_sql-support/1314-oracle-mysql-mariadb.md) with new `dbmmys_8_4` ODI
  driver.
- Support for [Oracle MySQL 9.7 LTS](../10_sql-support/1314-oracle-mysql-mariadb.md) with new `dbmmys_9_7` ODI
  driver.
- Support for [Oracle database 23ai/26ai](../10_sql-support/1359-oracle-database.md) with
  the `dbmora_23` ODI driver.
- Support for [PostgreSQL 18](../10_sql-support/1415-postgresql.md) with the
  `dbmpgs_9` ODI driver.
- Support for [SQL Server 2025](../10_sql-support/1256-microsoft-sql-server.md) with the
  `dbmsnc_18` ODI driver.
- Support for [MariaDB 12](../10_sql-support/1314-oracle-mysql-mariadb.md) with the
  `dbmmdb_10_2` ODI driver.

## New database clients supported in Genero 5.01:

- The new `dbmmys_8_4` ODI driver for Oracle MySQL 8.4 LTS:
  > **Note:**
  >
  > The MySQL 8.2 and 8.3 versions are Innovation Releases (IR), while MySQL versions 8.0 and 8.4 are
  > Long-Term Support releases (LTS). The `dbmmys_8_2` driver provided in previous Genero
  > versions has been desupported in favor to MySQL 8.4 with the new `dbmmys_8_4` ODI
  > driver, requiring libmysqlclient.so.24. If you need to access a MySQL 8.2 or
  > 8.3 server, install MySQL Connector client with libmysqlclient.so.24 and use
  > the `dbmmys_8_4` ODI driver.
- The new `dbmmys_9_7` ODI driver for Oracle MySQL 9.7 LTS (with VECTOR support)

For more details, see [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md).

## Deprecated ODI drivers in Genero 5.01:

The following ODI database drivers will be removed in a future version:

- `dbmesm*` : Easysoft ODBC for SQL Server.

  On Linux and Windows®, use Microsoft ODBC for SQL server. For other
  platforms, use FreeTDS.
- `dbmsnc_17`: ODI driver for Microsoft ODBC driver 17 for SQL Server

## Desupported databases server versions in Genero 5.01:

Desupported database server versions:

- MySQL 8.0 (end of support 2026-04-30), and consequently the `dbmmys_8_0` ODI
  driver for MySQL 8.0 client.

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Database drivers changes in BDL 5.00](0115-database-drivers-changes.md "New and desupported database drivers.").

Notable changes introduced in maintenance releases:

- No particular change to consider.
