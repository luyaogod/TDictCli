---
title: "MariaDB support"
source: "fgl-topics/c_fgl_Migrate_to_300_mariadb.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > MariaDB support"
type: "concept"
---

# MariaDB support

> The MariaDB database is now supported by Genero 3.00.

MariaDB is the open source brand of Oracle's MySQL that has been adopted by several major
organizations.

The purpose of the MariaDB project is to be a drop-in replacement for MySQL.

MariaDB supported versions are 10.0 and higher.

To connect to MariaDB, use the MySQL database driver (dbmmys), and follow MySQL
adaptation guide for configuration and SQL portability issues.

Depending on the libmysqlclient library compatibility, you might need to
configure Genero to use a version-stamped driver. As of Genero version 3.00, the generic driver name
"dbmmys" can be used to connect to MariaDB 10.0. See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md) for
more details.

## Related links

**Related concepts**  

[Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md "Oracle MySQL / MariaDB")
