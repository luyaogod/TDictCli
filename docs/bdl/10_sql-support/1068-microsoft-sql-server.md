---
title: "Microsoft SQL Server"
source: "fgl-topics/c_fgl_Connections_033.html"
breadcrumb: "SQL support > Database connections > Database client environment > Microsoft™ SQL Server"
type: "concept"
description: "Make sure that ODBC data source is defined on database client and database server systems, with the correct ODBC driver. Note that Genero FGL provides different types of SQL Server drivers: dbmsnc : ..."
---

# Microsoft SQL Server

1. Make sure that ODBC data source is defined on database client and database server systems, with
   the correct ODBC driver. Note that Genero FGL provides different types of SQL Server drivers:
   - `dbmsnc`: For Microsoft ODBC driver for SQL Server (recommended)
   - `dbmesm`: For Easysoft ODBC driver for SQL Server
   - `dbmftm`: For FreeTDS ODBC
2. On Windows® platforms, the PATH environment variable
   must define the access path to database client programs (ODBC32.DLL). On UNIX platforms, check
   database client software documentation for environment settings (LD\_LIBRARY\_PATH, ldconfig).
3. On Windows, check the SQL Server Client configuration
   with the Client Network Utility tool. Verify that the ANSI to OEM conversion corresponds to the
   execution of applications in a CONSOLE environment.
4. Make sure the database client locale is properly defined. On UNIX platforms, check that the
   client character set parameter of the ODBC data source corresponds the locale used by the
   application (LANG/LC\_ALL).
5. On Windows, you can make a connection test with the
   Microsoft™ Query Analyzer tool. On UNIX, see client
   software documentation for available SQL command tools (isql command line tool
   for example).

## Related links

**Related concepts**  

[Database client settings](../09_advanced-features/0885-database-client-settings.md "This section describes the settings defining the locale for the database client.")

**Related tasks**  

[Prepare the runtime environment - connecting to the database](1260-prepare-the-runtime-environment-connecting-to-the-database.md "Prepare the runtime environment - connecting to the database")
