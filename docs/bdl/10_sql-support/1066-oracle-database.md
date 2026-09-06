---
title: "Oracle database"
source: "fgl-topics/c_fgl_Connections_031.html"
breadcrumb: "SQL support > Database connections > Database client environment > Oracle database"
type: "concept"
description: "On UNIX™, LD_LIBRARY_PATH (or equivalent) must hold the path to find the Oracle® client library. On Microsoft Windows, the PATH environment variable must contain the path to Oracle client library. In ..."
---

# Oracle database

1. On UNIX™, LD\_LIBRARY\_PATH (or equivalent) must hold the
   path to find the Oracle® client library. On
   Microsoft Windows, the PATH environment variable must contain the path to Oracle client
   library.
2. In an Oracle Database Server context, make sure that ORACLE\_HOME and ORACLE\_SID environment
   variables are properly set.
3. With an Oracle Instant Client environment, define the TNS\_ADMIN variable to find the
   configutation files.
4. The TNSNAMES.ORA file must define the database server identifiers for remote connections (the
   Oracle Listener must be started on the
   database server to allow remote connections).
5. The SQLNET.ORA file must define network settings for remote connections.
6. Make sure the database client locale is properly defined (NLS\_LANG).
7. Do a connection test with the Oracle
   sqlplus command line tool.

## Related links

**Related concepts**  

[Database client settings](../09_advanced-features/0885-database-client-settings.md "This section describes the settings defining the locale for the database client.")

**Related tasks**  

[Prepare the runtime environment - connecting to the database](1363-prepare-the-runtime-environment-connecting-to-the-database.md "Prepare the runtime environment - connecting to the database")
