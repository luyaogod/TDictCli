---
title: "Prepare the runtime environment - connecting to the database"
source: "fgl-topics/t_fgl_odiagpgs_004.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Installation (Runtime Configuration) > Prepare the runtime environment - connecting to the database"
type: "task"
description: "In order to connect to PostgreSQL, you must have a PostgreSQL database driver \"dbmpgs\" in $FGLDIR/dbdrivers . The PostgreSQL client software is required to connect to a database server. Check whether ..."
---

# Prepare the runtime environment - connecting to the database

1. In order to connect to PostgreSQL, you must have a PostgreSQL database driver "dbmpgs" in
   $FGLDIR/dbdrivers.
2. The PostgreSQL client software is required to connect to a database server.

   Check whether the PostgreSQL client library (libpq.\*) is installed on the
   machine where the BDL programs run.
3. Make sure that the PostgreSQL client environment variables are properly set.

   Check, for example, PGDIR (the path to the installation directory), PGDATA (the path to the
   data files directory), etc. See the PostgreSQL documentation for more details.
4. Check the database client locale settings.

   The database client locale must match the locale used by the runtime system (LC\_ALL,
   LANG).

   With PostgreSQL, the default client encoding is the database encoding. If the application locale
   is different than the database locale, set the database client locale with
   `client_locale` in postgresql.conf, or the PGCLIENTENCODING
   environment variable.
5. Verify the environment variable defining the search path for the PostgreSQL database client
   shared libraries (libpq.so on UNIX™, LIBPQ.DLL on Windows™).

   | PostgreSQL version | Shared library environment setting |
| --- | --- |
| PostgreSQL 11 and higher | *UNIX*: Add $PGDIR/lib to LD\_LIBRARY\_PATH (or its equivalent).*Windows*: Add %PGDIR%\bin to PATH. |
6. To verify if the PostgreSQL client environment is correct, you can start the PostgreSQL command
   interpreter:

   ```
   $ psql dbname -U appadmin -W
   ```
7. Set up the FGLPROFILE entries for [database connections](1059-database-connections.md "Explains how to manage database connections in a program.").
   1. Define the PostgreSQL database driver:

      ```
      dbi.database.dbname.driver = "dbmpgs"
      ```
   2. The '**source**' parameter defines the name of the PostgreSQL database, as well as
      additional connection parameters if needed, such as the server host name, the TCP port and specific
      PostgresSQL connection options.

      ```
      dbi.database.dbname.source = "test1"
      ```

      The `source` parameter must have the following form:

      ```
      dbname[@host[:port]][?options]
      ```

      where:
      - dbname defines the name of the PostgreSQL database
      - host defines the server host name, or IP address (IPv6 host address needs to
        be enclosed it in square brackets)
      - port defines the TCP port
      - options is a URI-style query string defining PostgreSQL connection
        parameters

      For example:

      ```
      mydb@orion:5433?connect_timeout=10&application_name=myapp
      ```
   3. Define pre-fetch rows.

      To improve performances, you can define the number of result set rows that the driver must
      prefetch:

      ```
      dbi.database.dbname.pgs.prefetch.rows = integer
      ```

      This will be applied to all application cursors.

      The default is 50 rows. Do not change the default, except if it gives really better performances:
      This can blow up memory usage for each DB client process.
   4. If needed, define the FGLPROFILE entry to enable PostgreSQL client trace.

      PostgreSQL client trace file can be defined
      with:

      ```
      dbi.database.dsname.pgs.trace.file = "/tmp/pgstrace_%(PID).log"
      ```
   5. If needed, define the database schema search path with the following FGLPROFILE entry:

      Database schema search path (see [Name resolution of SQL objects](1438-name-resolution-of-sql-objects.md)):

      ```
      dbi.database.dsname.pgs.schema = "public,schema1,schema2
      ```

   For more details, see [PostgreSQL specific FGLPROFILE parameters](1084-postgresql-specific-fglprofile-parameters.md).
8. To configure a secure connection to PostgreSQL server using TLS, see [Server-side configuration for
   secure connections with PostgreSQL](1418-install-postgresql-and-create-a-database-database-configurat.md)
