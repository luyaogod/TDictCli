---
title: "IBM Informix Dynamic Server"
source: "fgl-topics/c_fgl_Connections_029.html"
breadcrumb: "SQL support > Database connections > Database client environment > IBM® Informix® Dynamic Server"
type: "concept"
description: "The INFORMIXDIR environment variable must define the IBM® Informix® software installation path. The PATH environment variable must define the access path to database client programs. On UNIX™, ..."
---

# IBM Informix Dynamic Server

1. The INFORMIXDIR environment variable must define the IBM® Informix® software installation path.
2. The PATH environment variable must define the access path to database
   client programs.
3. On UNIX™, LD\_LIBRARY\_PATH
   (or equivalent) must hold the path to `$INFORMIXDIR/lib:$INFORMIXDIR/lib/esql`.
4. The IBM Informix client libraries
   `'INFORMIXDIR/lib/*'` must be available.
5. The INFORMIXSERVER environment variable can be used to define
   the name of the database server.
6. The sqlhost file must define the database server identified by
   INFORMIXSERVER.
7. Make sure the database client locale is properly defined (CLIENT\_LOCALE)
8. You can make a connection test with the IBM Informix dbaccess command line tool.

## Related links

**Related concepts**  

[Installation (Runtime Configuration)](1181-installation-runtime-configuration.md "ODI adaptation guide Installation topics.")

[Database client settings](../09_advanced-features/0885-database-client-settings.md "This section describes the settings defining the locale for the database client.")
