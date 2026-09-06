---
title: "MariaDB"
source: "fgl-topics/c_fgl_Connections_env_mdb.html"
breadcrumb: "SQL support > Database connections > Database client environment > MariaDB"
type: "concept"
description: "The MYSQL_HOME environment variable must define the MariaDB software installation path. The PATH environment variable must define the access path to database client programs. On UNIX™, LD_LIBRARY_PATH ..."
---

# MariaDB

1. The MYSQL\_HOME environment variable must define the MariaDB software installation path.
2. The PATH environment variable must define the access path to database
   client programs.
3. On UNIX™, LD\_LIBRARY\_PATH
   (or equivalent) must hold the path to `$MYSQL_HOME/lib`.
4. Make sure the database client locale is properly defined
   (`default-character-set`)
5. You can make a connection test with the mysql command line tool.

## Related links

**Related concepts**  

[Database client settings](../09_advanced-features/0885-database-client-settings.md "This section describes the settings defining the locale for the database client.")

**Related tasks**  

[Prepare the runtime environment - connecting to the database](1318-prepare-the-runtime-environment-connecting-to-the-database.md "Prepare the runtime environment - connecting to the database")
