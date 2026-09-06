---
title: "PostgreSQL"
source: "fgl-topics/c_fgl_Connections_032.html"
breadcrumb: "SQL support > Database connections > Database client environment > PostgreSQL"
type: "concept"
description: "The PGDIR environment variable must define the PostgreSQL software installation path. The PATH environment variable must define the access path to database client programs. On UNIX™, LD_LIBRARY_PATH ..."
---

# PostgreSQL

1. The PGDIR environment variable must define the PostgreSQL software
   installation path.
2. The PATH environment variable must define the access path to database
   client programs.
3. On UNIX™, LD\_LIBRARY\_PATH
   (or equivalent) must hold the path to `$PGDIR/lib`.
4. The PostgreSQL client library `'PGDIR/lib/libpq*'` must
   be available.
5. On the database server, the `pg_hba.conf` file
   must define security policies.
6. Make sure the database client locale is properly defined (PGCLIENTENCODING)
7. You can make a connection test with the PostgreSQL psql command line
   tool.

## Related links

**Related concepts**  

[Database client settings](../09_advanced-features/0885-database-client-settings.md "This section describes the settings defining the locale for the database client.")

**Related tasks**  

[Prepare the runtime environment - connecting to the database](1419-prepare-the-runtime-environment-connecting-to-the-database.md "Prepare the runtime environment - connecting to the database")
