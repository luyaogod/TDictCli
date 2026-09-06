---
title: "Database drivers changes"
source: "fgl-topics/c_fgl_Migrate_to_600_database_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 6.00 upgrade guide > Database drivers changes"
type: "concept"
---

# Database drivers changes

> New and desupported database drivers.

## New database server versions supported in Genero 6.00:

- Support for [Informix 15](../10_sql-support/1179-ibm-informix.md) by using the Informix CSDK
  15 client environment. In a previous version, Informix 15 could only be accessed with Informix CSDK
  4.50.

## New database clients supported in Genero 6.00:

- Support for [Informix CSDK 15](../10_sql-support/1182-supported-ibm-informix-server-and-csdk-versions.md) with
  new `dbmifx_15` ODI driver.
  > **Important:**
  >
  > For many Genero BDL versions, there was only one Informix ODI driver available, with the name
  > `dbmifx_9`, compatible with CSDK versions 4.10 and 4.50. Starting with Genero BDL
  > 6.00, as the CSDK 15 libraries are not compatible with older CSDK libraries, Genero provides now two
  > distinct Informix ODI drivers:
  > - `dbmifx_9`, compatible with Informix CSDK 4.10 - 4.50.
  > - `dbmifx_15`, compatible with Informix CSDK 15 - which is now the default in
  >   V6.00!Make sure that the ODI driver selected during the installation (or configured manually in
  > FGLPROFILE), matches your Informix CSDK version.
  >
  > The default ODI driver can be selected during the Genero BDL installation process.

  > **Important:**
  >
  > Starting with Genero version 6.00, the `sqlca.sqlerrd` member is now an
  > `ARRAY[6] OF BIGINT` (it was based on `INTEGER` in prior versions).
  > For more details, read [sqlca.sqlerrd[1-6] as BIGINT](0091-core-language-changes.md).

  > **Tip:**
  >
  > Informix 15 uses by default large tables (onconfig `TABLE_SIZE LARGE`). This
  > server parameter impacts the size of ROWIDs (8 bytes long with large tables). To get 4-bytes ROWIDs,
  > use `TABLE_SIZE SMALL`, as suggested in [Informix setup
  > recommendations](../10_sql-support/1183-install-ibm-informix-and-create-a-database-database-configur.md).

For more details about new database drivers, see [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md).

## Deprecated ODI drivers in Genero 6.00:

The following ODI database drivers will be removed in a future version:

- None

## Full desupport of SAP HANA

Starting with Genero BDL 6.00, the SAP HANA 2 database is no longer supported.

The name of the ODI driver was `dbmhdb_2`.

## Desupported ODI drivers in Genero 6.00:

The following ODI database drivers have been removed in current version:

- `dbmsnc_17`: ODI driver for Microsoft ODBC driver 17 for SQL Server

## Desupported databases server versions in Genero 6.00:

Database server versions no longer supported:

- none

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Database drivers changes in BDL 5.01](0104-database-drivers-changes.md "New and desupported database drivers.").

Notable changes introduced in maintenance releases:

- No particular change to consider.
