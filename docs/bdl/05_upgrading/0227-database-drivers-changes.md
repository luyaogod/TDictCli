---
title: "Database drivers changes"
source: "fgl-topics/c_fgl_Migrate_to_251_database_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.51 upgrade guide > Database drivers changes"
type: "concept"
---

# Database drivers changes

> Desupported database drivers.

## Databases / ODI drivers desupported in version 2.51:

- Genero DB is no longer supported (`dbmads*`).
- Oracle® MySQL 4.1 and 5.0
  (`dbmmys41x`, `dbmmys50x`)
- PostgreSQL 8.3, 8.4 (`dbmora83x`, `dbmpgs84x`)
- Oracle Database 9.2
  (`dbmora92x`)

Note also that database driver naming convention has changed in 2.51, for more details see
[New database driver name specification](0228-new-database-driver-name-specification.md "Allows database driver specification without target database version information.").

## Related links

**Related concepts**  

[Microsoft SQL Server](../10_sql-support/1256-microsoft-sql-server.md "Microsoft SQL Server")

[Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md "Database driver specification (driver)")
