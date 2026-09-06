---
title: "Install SQLite and create a database - database configuration/design tasks"
source: "fgl-topics/t_fgl_odiagsqt_003.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Installation (Runtime Configuration) > Install SQLite and create a database - database configuration/design tasks"
type: "task"
description: "If you are tasked with installing and configuring the database, here is a list of steps to be taken: The core of SQlite is a shared library (or DLL). For Linux®, macOS® and iOS platforms,the dbmsqt ..."
---

# Install SQLite and create a database - database configuration/design tasks

If you are tasked with installing and configuring the database, here is
a list of steps to be taken:

1. The core of SQlite is a shared library (or DLL). For Linux®, macOS® and iOS platforms,the
   `dbmsqt` database driver is linked dynamically with the SQLite library. In this case,
   the SQLite shared library must be present on the system. On other platforms such as Microsoft™ Windows® and Android™, the `dbmsqt` driver
   has an embedded version of the SQLite library, and no SQLite library needs to be installed for
   Genero applications.

   The minimum required version is SQLite 3.6.
2. Create a new SQLite database.

   To create a new database with tables, start the sqlite3 command line tool
   and execute SQL statements:

   ```
   $ sqlite3 /var/data/stores.db
   sqlite> CREATE TABLE customer ( cust_id INT PRIMARY KEY, ... );
   $ .exit
   ```

   To create an empty database, you can also issue the following
   command:

   ```
   $ sqlite3 /var/data/stores.db ""
   ```

   or create an empty file with
   operating system command:

   ```
   $ touch /var/data/stores.db
   ```

   An empty file can
   also be created from a program by using a `base.Channel`
   object:

   ```
   DEFINE ch base.Channel
   LET ch = base.Channel.create()
   CALL ch.openFile("/var/data/stores.db","w")
   CALL ch.close
   ```

## Related links

**Related concepts**  

[SQL programming](0984-sql-programming.md "Covers topics about interacting with a database server using SQL.")
