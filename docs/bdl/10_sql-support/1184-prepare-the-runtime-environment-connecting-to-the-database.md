---
title: "Prepare the runtime environment - connecting to the database"
source: "fgl-topics/t_fgl_odiagifx_006.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Installation (Runtime Configuration) > Prepare the runtime environment - connecting to the database"
type: "task"
description: "In order to connect to IBM® Informix®, you must have a database driver \" dbmifx \" in $FGLDIR/dbdrivers . Make sure the IBM Informix client environment variables are properly set. Check for example ..."
---

# Prepare the runtime environment - connecting to the database

1. In order to connect to IBM®
   Informix®, you must have a database driver
   "dbmifx" in $FGLDIR/dbdrivers.
2. Make sure the IBM Informix client environment
   variables are properly set.

   Check for example INFORMIXDIR (the path to the installation directory), INFORMIXSERVER (the
   name of the server defined in the sqlhosts list), etc. For more details, see
   the IBM
   Informix documentation.
3. In order to connect to an IBM
   Informix server, you must define a line in the
   $INFORMIXDIR/etc/sqlhosts file, referencing the server name specified in
   the INFORMIXSERVER environment variable.

   On Windows® platforms, the
   sqlhost entries are defined in the registry database. See IBM
   Informix documentation.
4. Verify the environment variable defining the search path
   for IBM Informix SDK database client shared libraries.

   | IBM Informix SDK version | Shared library environment setting |
| --- | --- |
| All versions | *UNIX™*: Add $INFORMIXDIR/lib, $INFORMIXDIR/lib/esql, $INFORMIXDIR/lib/tools and $INFORMIXDIR/lib/cli to LD\_LIBRARY\_PATH (or its equivalent).*Windows*: Add %INFORMIXDIR%\bin to PATH. |
5. Check the Informix database client locale settings.

   The Informix database client locale must match the locale used by the runtime system (LC\_ALL,
   LANG):

   The Informix
   environment variables defining locale settings are:
   - CLIENT\_LOCALE: defines the locale of the client application.
   - DB\_LOCALE: defines the locale of the database (matters at DB creation!)
   - SERVER\_LOCALE: defines the locale to be used by the server to interact with the OS file
     system.
6. To verify if the IBM Informix client environment
   is correct, you can start the SQL command interpreter:

   ```
   $ dbaccess - - 
   > CONNECT TO "dbname" USER "appadmin"; 
   ENTER PASSWORD: password
   ```
7. Set up the FGLPROFILE entries for [database connections](1059-database-connections.md "Explains how to manage database connections in a program.").

   > **Important:**
   >
   > Make sure that you are using the ODI driver corresponding to
   > the database client and server version.
