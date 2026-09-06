---
title: "Prepare the runtime environment - connecting to the database"
source: "fgl-topics/t_fgl_odiagdmg_004.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Installation (Runtime Configuration) > Prepare the runtime environment - connecting to the database"
type: "task"
description: "In order to connect to Dameng®, you must have the database driver \" dbmdmg \" in $FGLDIR/dbdrivers . Get a license with Dameng database support. Important: In order to use the Dameng ODI driver, you ..."
---

# Prepare the runtime environment - connecting to the database

1. In order to connect to Dameng®, you must have
   the database driver "`dbmdmg`" in $FGLDIR/dbdrivers.
2. Get a license with Dameng database support.

   > **Important:**
   >
   > In order to use the Dameng ODI driver, you need a license with Dameng database option. If the
   > installed license does not have this option, you get the error
   > message:
   >
   > ```
   > FORMS statement error number -6000.
   >  FLM-135 - The license 'license-number' does not allow to use the Dameng DB.
   > Please contact your vendor.
   > ```
3. To connect to a remote Dameng server, the Dameng DPI client must be installed and configured on
   the computer running the BDL applications.

   Declare the Dameng data source in the dm\_svc.conf configuration
   file:

   ```
   test1=([::]:5237)

   [test1]
   PAGE_MODE=PG_UTF8
   # PAGE_MODE=PG_ISO_8859_1
   ```

   The DM\_SVC\_PATH environment variable defines the path to the directory where the
   dm\_svc.conf configuration file is located.

   See Dameng documentation for more details.
4. Check the database client locale settings.

   The database client locale must match the locale used by the runtime system (LC\_ALL, LANG).
   See [CHAR and VARCHAR data types](1228-char-and-varchar-data-types.md) for more details.
5. Verify the environment variable defining the search path for Dameng DPI database client shared
   libraries (libdmdpi.so on UNIX™).

   | Dameng version | Shared library environment setting |
| --- | --- |
| Dameng 8 and higher | *UNIX*: Add the following paths to LD\_LIBRARY\_PATH (or its equivalent):dmg\_install\_dir/bindmg\_install\_dir/drivers/dcidmg\_install\_dir/drivers/odbc |
6. To verify if the Dameng client environment is correct, make a connection to the server with the
   disql SQL command line tool:

   ```
   $ disql dmguser/fourjs@[::]:5237
   ```
7. Setup the FGLPROFILE entries for [database
   connections](1059-database-connections.md "Explains how to manage database connections in a program.").
   1. Define the Dameng database driver:

      ```
      dbi.database.dbname.driver = "dbmdbm"
      ```
   2. The "`source`" parameter defines the name of the Dameng database name as defined
      in the dm\_svc.conf file:

      ```
      dbi.database.dbname.source = "test1"
      ```
   3. Define the database schema selection if needed:

      Use the following entry to define the database schema to be used by the application. The database
      interface will automatically perform a `alter session set current_schema =
      name` instruction to switch to a specific schema:

      ```
      dbi.database.dbname.dmg.schema = 'name'
      ```

      Here dbname identifies the database name used in the BDL program (
      `DATABASE dbname` ) and name is the Dameng
      database schema name.
   4. By default, the tablespace for the default temporary table emulation is
      `TEMPTABS`.

      If required, define the tablespace to be used for temporary table
      emulations:

      ```
      dbi.database.dbname.dmg.temptables.tablespace = "mytemptabs"
      ```
   5. By default, temp table emulation uses the session id and the fraction of seconds of
      systimestamp to get a unique table name.

      If needed, change the SELECT statement that produces a unique identifier for temp table
      emulation, with the following FGLPROFILE
      entry:

      ```
      dbi.database.dbname.dmg.sid.command = "SELECT ... FROM dual"
      ```
