---
title: "Prepare the runtime environment - connecting to the database"
source: "fgl-topics/t_fgl_odiagsqt_004.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Installation (Runtime Configuration) > Prepare the runtime environment - connecting to the database"
type: "task"
description: "In order to connect to SQLite, the database driver \"dbmsqt\" must exist in $FGLDIR/dbdrivers . On most platforms, the dbmsqt SQLite driver is linked statically with the SQLite library, in other words, ..."
---

# Prepare the runtime environment - connecting to the database

1. In order to connect to SQLite, the database driver `"dbmsqt"` must exist in
   $FGLDIR/dbdrivers. On most platforms, the `dbmsqt` SQLite driver
   is linked statically with the SQLite library, in other words, SQLite is embedded in the ODI driver.
   On Linux®, macOS® and iOS platforms, the `dbmsqt` driver is linked dynamically with the
   SQLite library, and this library must be present on the system.

   To check the version of the SQLite library, issue the following SQL statement with the sqlite3
   command tool, of from a Genero program after
   setup:

   ```
   select sqlite_version()
   ```
2. Make sure that the SQLite environment variables are properly set. You may want to define an
   environment variable such as SQLITEDIR to hold the installation directory of SQLite, which can then
   be used to set PATH and LD\_LIBRARY\_PATH. See SQLite documentation for more details.
3. If the SQLite library is not embedded in the `dbmsqt` driver, the environment
   must be set to find the SQLite library. Verify the environment variable defining the search path for
   the SQLite shared library.

   | SQLite version | Shared library environment setting (if SQLite lib not built-in driver) |
| --- | --- |
| SQLite 3.6 and higher | *UNIX™*: Add $SQLITEDIR/lib to LD\_LIBRARY\_PATH (or its equivalent).*Windows®*: Add %SQLITEDIR%\bin to PATH. |
4. Make sure that all operating system users running the application
   have read/write access to the database file.
5. SQLite uses UTF-8 encoding. If the locale used by the runtime system (LANG/LC\_ALL) is not
   compatible to UTF-8 (for example, fr\_FR.iso88591), Genero will do the
   appropriate character set conversions.
6. Set up the FGLPROFILE entries for [database
   connections](1059-database-connections.md "Explains how to manage database connections in a program.").
   1. Define the SQLite database driver:

      ```
      dbi.database.dbname.driver = "dbmsqt"
      ```
   2. The "`source`" parameter defines the path to the
      SQLite database file. Note that the database file must reside on
      the local disk (SQLite does not support network file systems).
      SQLite also supports in-memory database creation with the
      `:memory:` db specification. See SQLite
      documentation (sqlite3\_open) for more details.

      ```
      dbi.database.dbname.source = "/opt/myapp/stock.dbs"
      ```
   3. If the "`source`" parameter defines a relative
      path or a simple filename and the SQLite database file does not
      reside in that location based on the current directory of the
      fglrun process, define the DBPATH environment variable to find
      the database file. See [DBPATH](../07_configuration/0513-dbpath.md "Defines a list of paths for Genero program resource files.")
      documentation for more details about this environment
      variable.

      ```
      DBPATH="/opt/myapp"
      ```

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
