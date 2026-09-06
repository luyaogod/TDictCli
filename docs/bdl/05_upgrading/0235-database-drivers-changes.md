---
title: "Database drivers changes"
source: "fgl-topics/c_fgl_Migrate_to_250_database_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.50 upgrade guide > Database drivers changes"
type: "concept"
---

# Database drivers changes

> Desupported database drivers.

## Databases / ODI drivers desupported in versions 2.50:

- SQL Server MDAC drivers (Code MSV, name: `dbmmsv*`):

  On a Microsoft™ Windows®
  platform, use the SQL Server Native Client driver instead (Code SNC).

  With the SNC drivers,
  set the `dbi.database.dbname.snc.widechar` FGLPROFILE entry to
  `false` when using CHAR/VARCHAR/TEXT in the SQL Server database. Note that this is
  only required for versions up to 3.00: Starting with version 3.10, the char mode of SNC driver is
  selected automatically, depending on the current application locale.
- Oracle® MySQL 5.4
  (`dbmmys54x`)
- Oracle Database 8.1
  (`dbmora81x`)
- Oracle Database 9.0
  (`dbmora90x`)

## Related links

**Related concepts**  

[Microsoft SQL Server](../10_sql-support/1256-microsoft-sql-server.md "Microsoft SQL Server")

[Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md "Oracle MySQL / MariaDB")

[Oracle Database](../10_sql-support/1359-oracle-database.md "Oracle Database")
