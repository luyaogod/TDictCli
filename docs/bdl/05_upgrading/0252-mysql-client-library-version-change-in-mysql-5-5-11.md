---
title: "MySQL client library version change in MySQL 5.5.11"
source: "fgl-topics/c_fgl_Migrate_to_240_libmysqlclient.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > MySQL client library version change in MySQL 5.5.11"
type: "concept"
---

# MySQL client library version change in MySQL 5.5.11

> Shared library version number of the MySQL client library must match the library used to link the ODI driver.

Starting with Oracle® MySQL 5.5.11,
the client library version number was changed from 16 to 18. In fact the
libmysqlclient.so.16 file was renamed to
libmysqlclient.so.18. From a cross-5.5.x compatibility point of view, this may
not be the best thing to do, since the major shared library version has changed, client applications
using the C API (such as Genero ODI MySQL drivers) need to be recompiled and re-linked in order to
use the latest library.

In Genero version 2.40, the dbmmys55x ODI driver is linked with
libmysqlclient.so.18 on the platforms where MySQL 5.5.11+ is available. That
is: Linux®, Solaris and macOS® platforms, when writing these lines.
On other UNIX™ platforms such as HP, the client library is
still libmysqlclient.so.16.
This may change in future Genero versions, following the availability of MySQL 5.5.11+ versions.

Therefore, you must pay attention to the MySQL 5.5 version you have installed. You need to
upgrade your MySQL 5.5 client software to match the client library used to build the
dbmmys55x.so shared library. On Linux,
you can run the ldd command to check what libmysqlclient.so
version is required. If it's not possible to upgrade your MySQL client software, please contact the
support channel.

## Related links

**Related concepts**  

[Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md "Oracle MySQL / MariaDB")
