---
title: "Install IBM Informix and create a database - database configuration/design tasks"
source: "fgl-topics/t_fgl_odiagifx_005.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Installation (Runtime Configuration) > Install IBM® Informix® and create a database - database configuration/design tasks"
type: "task"
description: "Install the IBM® Informix® database software (IDS for example) on your database server. Install the IBM Informix Software Development Kit (SDK) on your application server. With some IBM Informix ..."
---

# Install IBM Informix and create a database - database configuration/design tasks

1. Install the IBM® Informix® database software
   (IDS for example) on your database server.
2. Install the IBM Informix Software Development
   Kit (SDK) on your application server.

   With some IBM
   Informix distributions (IDS 11), this package is
   included in the server bundle. it is recommended that you check the IBM
   web site for SDK upgrades or patches. Genero BDL is certified with IBM
   Informix SDK version 3.70 or higher.
3. Setup the IDS server (onconfig file, etc)
   1. Starting with IDS version 11, the `TEMPTAB_NOLOG` is set to 1 by default.

      Consider setting the `TEMPTAB_NOLOG` parameter to 0, if you want to log
      temporary table changes. This can affect the behavior of programs expecting that a ROLLBACK WORK
      cancels changes done on a temporary table.
   2. Starting with IDS version 11, by default the precision of SQL statement timing is the second
      (`USEOSTIME` is 0). For example, CURRENT HOUR TO FRACTION(3) returns a fraction part
      of zero.

      If sub-second precision is required, set the `USEOSTIME` configuration
      parameter to 1.
   3. Starting with IDS version 15, in the onconfig file, the TABLE\_SIZE option is set to LARGE
      by default.

      With TABLE\_SIZE LARGE, for new created permanent SQL tables, ROWIDs will be 8-byte long and
      contain large integer numbers that will not fit into an FGL INTEGER variable. Consider changing the
      onconfig file to use TABLE\_SIZE SMALL, in order to get the same table size possiblities and 4-bytes
      ROWID as with IDS 14.
4. Define a database user dedicated to your application: the application administrator.

   This user will manage the database schema of the application (all tables will be owned by this
   user). With IBM
   Informix, database users reference Operating System
   users, and must be part of the IBM
   Informix group. See IBM
   Informix documentation.
5. Connect to the server as the "informix" user (for example with the dbaccess
   tool) and give all requested database administrator privileges to the application
   administrator.

   ```
   GRANT CONNECT TO appadmin;
   GRANT RESOURCE TO appadmin;
   GRANT DBA TO appadmin;
   ```
6. Login as the application administrator (appadmin), and setup the Informix
   client environment.
7. Define the client and database locale.

   The Informix
   environment variables defining locale settings are:
   - CLIENT\_LOCALE: defines the locale of the client application.
   - DB\_LOCALE: defines the locale of the database (matters at DB creation!)
   - SERVER\_LOCALE: defines the locale to be used by the server to interact with the OS file
     system.

   > **Important:**
   >
   > Always define the DB\_LOCALE environment variable,
   > before executing the CREATE DATABASE statement: If DB\_LOCALE is not set, it will default to the
   > DB\_LOCALE value set when starting the Informix engine. If no DB\_LOCALE was set when starting the
   > engine, it defaults to `en_us.8859-1` (i.e. `en_us.819`).
8. Create a database entity, for example with the following SQL statement:

   ```
   CREATE DATABASE dbname WITH BUFFERED LOG;
   ```

   > **Tip:**
   >
   > To check the locale of an Informix database, connect
   > to the `sysmaster` database, and run the following
   > query:
   >
   > ```
   > $ dbaccess sysmaster -
   > Database selected.
   >
   > > SELECT * FROM sysdbslocale WHERE dbs_dbsname = 'mydb'; 
   > dbs_dbsname  mydb
   > dbs_collate  en_US.819
   > ```
   >
   > The
   > charset code for ISO-8859-1 is 819, for UTF8 it is 57372.

   > **Tip:**
   >
   > To check the value of the SQL\_LOGICAL\_CHAR size
   > multiplier of an existing database, connect to that database (not to `sysmaster`!),
   > and execute the following
   > query:
   >
   > ```
   > $ dbaccess mydb -
   > Database selected.
   >
   > > SELECT MOD(flags,4)+1 FROM systables WHERE tabname = ' VERSION';
   >       4
   > ```
9. Create the application tables.
