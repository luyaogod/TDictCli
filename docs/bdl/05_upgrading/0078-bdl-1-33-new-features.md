---
title: "BDL 1.33 new features"
source: "fgl-topics/fgl_whatsnew_133.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 1.33 new features"
type: "topic"
---

# BDL 1.33 new features

> Features added in 1.33 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 1.33 upgrade guide](0327-bdl-1-33-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 1.33.").

Prior new features guide: [BDL 1.32 new features](0079-bdl-1-32-new-features.md "Features added in 1.32 release of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| New `base.TypeInfo` built-in class to serialize program variables. | See [The TypeInfo class](../15_library-reference/3082-the-typeinfo-class.md). |
| The `base.Channel` class now supports a binary mode with the `'b'` option, to control CR/LF translation when using DOS files. | See [Line terminators on Windows and UNIX](../15_library-reference/3003-line-terminators-on-windows-and-unix.md). |

| Overview | Reference |
| --- | --- |
| Up to three accelerators can now be defined for an action in actions defaults files or in the `ACTION DEFAULTS` section of form files. | See [Defining keyboard accelerators for actions](../11_user-interface/2265-defining-keyboard-accelerators-for-actions.md). |

| Overview | Reference |
| --- | --- |
| Generic ODBC database driver is now available (code is generic ODBC database driver is now available (code is `odc`). | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md). |
| MySQL version 5.0.x is now supported. | See [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| PostgreSQL version 8.1.x is now supported. | See [PostgreSQL](../10_sql-support/1415-postgresql.md). |
| Microsoft™ SQL Server 2005 is now supported. | See [Microsoft SQL Server](../10_sql-support/1256-microsoft-sql-server.md). |
| Pre-fetch rows by block with SQL Server to get better performance. Use the following FGLPROFILE entry to specify the maximum number of rows the driver can pre-fetch:dbi.database.dbname.msv.prefetch.rows = countSee ["Database vendor specific parameters" in Connections](https://4js.com/online_documentation/fjs-fgl-manual-html/User/Connections.html#DS_ODI_DBVSPEC) for more details. | See [SQL Server (MS ODBC) specific FGLPROFILE parameters](../10_sql-support/1085-sql-server-ms-odbc-specific-fglprofile-parameters.md). |
| Upgrade notes for database drivers. | See [Database drivers changes](0328-database-drivers-changes.md "Desupported database drivers."). |
