---
title: "Prepare the runtime environment - connecting to the database"
source: "fgl-topics/t_fgl_odiagmys_004.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Installation (Runtime Configuration) > Prepare the runtime environment - connecting to the database"
type: "task"
description: "Check that the required ODI database driver is available in $FGLDIR/dbdrivers . For Oracle® MySQL 8.4, you must have the MySQL database driver \" dbmmys_8_4 \". For Oracle MySQL 9.7, you must have the ..."
---

# Prepare the runtime environment - connecting to the database

1. Check that the required ODI database driver is available in
   $FGLDIR/dbdrivers.

   - For Oracle® MySQL 8.4, you must have
     the MySQL database driver "`dbmmys_8_4`".
   - For Oracle MySQL 9.7, you must have
     the MySQL database driver "`dbmmys_9_7`".
   - For MariaDB 10.2+, you must have the MariaDB database driver
     "`dbmmdb_10_2`".
2. Check that the MySQL/MariaDB database client software is available on the system.

   - For MySQL, check that the libmysqlclient.\* library is installed on the
     system. The shared library version of the MySQL client library must match the libmysqlclient library
     version linked to the dbmmys.so ODI driver.
   - For MariaDB 10.2 and +, check that the libmariadb.\* library is installed on
     the system. The shared library version of the MariaDB client library must match the libmariadb
     library version linked to the dbmmdb.so ODI driver.
3. Make sure that the MySQL/MariaDB client environment variables are properly set.

   Check for example MYSQL\_HOME (the path to the installation directory), DATADIR (the path to
   the data files directory), etc. See MySQL documentation for more details about
   client environment variables to be set.
4. Check the MySQL/MariaDB client configuration options in the my.cnf file on
   your MySQL client environment.

   The driver will read the options defined in the `[client]` configuration group
   of the my.cnf file defined for the client application. Note that you can
   specify a particular configuration file with the following FGLPROFILE
   parameter:

   ```
   dbi.database.dbname.mys.config = "/opt/var/myapp/my.cnf"
   ```

   or,
   for
   MariaDB:

   ```
   dbi.database.dbname.mdb.config = "/opt/var/myapp/my.cnf"
   ```
5. Check the database client locale settings
   (`default-character-set` option in the
   my.cnf configuration file).

   The database client locale must match the locale used by the runtime system
   (LC\_ALL, LANG). For example, in order to define (4-bytes) UTF-8 as MySQL client character
   set:

   ```
   [client]
   default-character-set="utf8mb4"
   ```
6. Verify the environment variable defining the search path for the database
   client shared library (libmysqlclient.so on UNIX™, LIBMYSQL.dll on Windows®).

   | MySQL/MariaDB version | Shared library environment setting |
| --- | --- |
| MySQL 8.0 and higher and MariaDB 10.2 and higher | UNIX : Add $MYSQL\_HOME/lib to LD\_LIBRARY\_PATH (or its equivalent).Windows : Add %MYSQL\_HOME%\bin to PATH. |
7. To verify if the MySQL/MariaDB client environment is correct, start the SQL command
   interpreter:

   ```
   $ mysql dbname -u appadmin -p
   ```
8. Set up the FGLPROFILE entries for [database
   connections](1059-database-connections.md "Explains how to manage database connections in a program.").
   1. Define the MySQL/MariaDB database driver:

      For MySQL:

      ```
      dbi.database.dbname.driver = "dbmmys"
      ```

      For
      MariaDB:

      ```
      dbi.database.dbname.driver = "dbmmdb"
      ```
   2. The "`source`" parameter defines the name of the MySQL/MariaDB database.

      ```
      dbi.database.dbname.source = "test1"
      ```
9. To configure a secure connection to MySQL server using TLS, see [Server-side configuration for
   secure connections with MySQL](1317-install-oracle-mysql-mariadb-and-create-a-database-database.md)
