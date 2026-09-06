---
title: "fglmkrtm tool removed"
source: "fgl-topics/c_fgl_Migrate_to_200_fglmkrtm_tool.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > fglmkrtm tool removed"
type: "concept"
---

# fglmkrtm tool removed

> The fglmkrtm tool has been removed, as database drivers are loaded dynamically.

Starting with version 2.00, database drivers are now always loaded dynamically. Thus the
fglmkrtm tool has been removed from the distribution. This tool was
provided in previous versions to create a fglrun runner with the correct
database driver.

Refer to [Database connections](../10_sql-support/1059-database-connections.md "Explains how to manage database connections in a program.") for more details about database driver
configuration.

## Related links

**Related concepts**  

[Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md "Database driver specification (driver)")

[Command reference](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.")
