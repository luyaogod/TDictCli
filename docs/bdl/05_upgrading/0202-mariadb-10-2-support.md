---
title: "MariaDB 10.2 support"
source: "fgl-topics/c_fgl_Migrate_to_310_mariadb_10_2.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > MariaDB 10.2 support"
type: "concept"
---

# MariaDB 10.2 support

> The new ODI driver dbmmdb_10_2 is provided to connect to MariaDB 10.2.

MariaDB 10.2 comes with a new client library name libmariadb.so. In prior
versions of MariaDB, the client library name was libmysqlclient.so.

Starting with Genero 3.10, MariaDB 10.2 is supported by using the new `dbmmdb_10_2`
ODI driver. This driver is linked to libmariadb.so.3, the MariaDB 10.2 client
library.

> **Important:**
>
> The driver-specific FGLPROFILE entries must use the "`mdb`"
> code, for example:
>
> ```
> dbi.database.test1.mdb.config = "/opt/var/app/my.cnf"
> ```

To connect to older MariaDB version 10.0 and 10.1, you can still use the
dbmmys\_5\_5 ODI driver, linked to the libmysqlclient.so.18
shared library, provided in MariaDB 10.0 and 10.1 distributions.

> **Important:**
>
> The driver-specific FGLPROFILE entries must use the "`mys`"
> code, for example:
>
> ```
> dbi.database.test1.mys.config = "/opt/var/app/my.cnf"
> ```

## Related links

**Related concepts**  

[Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md "Database driver specification (driver)")

**Related tasks**  

[Prepare the runtime environment - connecting to the database](../10_sql-support/1318-prepare-the-runtime-environment-connecting-to-the-database.md "Prepare the runtime environment - connecting to the database")
