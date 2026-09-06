---
title: "Creating a database from programs"
source: "fgl-topics/c_fgl_sql_programming_059.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Creating a database from programs"
type: "concept"
---

# Creating a database from programs

> Creating a database from within a program requires special consideration.

## Understanding database creation statements

The Genero language syntax supports database creation statements such as:

```
CREATE DATABASE mydb WITH BUFFERED LOG
```

Such instruction performs an implicit connection to the database server (this means that no
`CONNECT TO` or `DATABASE` is required before a `CREATE
DATABASE`), which leads to a default connection.

## Creating a database in a database server

When using a database server engine, the creation of a database entity is not a trivial
operation. The process usually requires additional tasks such as data storage configuration,
database user creation, data access policy, and so on. These tasks are typically left to the
database administrator.

Database creation statements such as `CREATE DATABASE`, `CREATE
DBSPACE`, and `DROP DATABASE` can be used in programs connected to an IBM®
Informix® server, but these statements are not portable.
Use database creation statements only for development or testing purpose.

## Creating a database on mobile devices (SQLite)

Mobile applications usually create their database at first execution. Database creation on a
mobile device is a much simpler operation than database creation on a database server. For example,
with SQLite, creating a database only requires creating an empty file.

The SQLite database file must be created in the application sandbox, in a writable directory.
This directory is specific to the type of mobile device, and can be found in programs with the
`os.Path.pwd` method.

To build the full path to the database file, get the current working directory
(`os.Path.pwd()`) and add this path to the database file name. This defines the [`source`](1072-database-source-specification-source.md) specification in the database
connection parameters, to build the string used for the `CONNECT` instruction.

```
IMPORT os
...
DEFINE dbfile, source, connstr VARCHAR(256)

FUNCTION init_connection_strings()
   LET dbfile = "contacts.dbs"
   LET source = os.Path.join(os.Path.pwd(), dbfile)
   LET connstr = SFMT("contacts+source='%1'", source)
   IF NOT base.Application.isMobile() THEN
      -- Add db driver spec when in development mode
      LET connstr = connstr, ",driver='dbmsqt'"
   END IF
END FUNCTION
```

If not specified, the `source` connection parameter (i.e., the path to the
database file) defaults to the database name specification in the `CONNECT`
instruction. Thus, the `source='dbpath'` parameter is usually
omitted, and dbpath is specified directly as the database name. In this case,
however, the identifier of the database connection is the complete path to the SQLite database file.
For more details about database connection parameters, see [Database connections](1059-database-connections.md "Explains how to manage database connections in a program.").

Before executing the `CONNECT` instruction, check if the database file already
exists with [`os.Path.exists(source)`](../15_library-reference/3715-os-path-exists.md "Checks if a file exists."). Create the database file and tables only if
needed:

```
IMPORT os
...
   CALL init_connection_strings()
   IF os.Path.exists(source) THEN
      CONNECT TO connstr AS "c1"
   ELSE
      CALL create_empty_file(source)
      CONNECT TO connstr AS "c1"
      CALL create_database_tables()
   END IF
   ...

FUNCTION create_empty_file(fn)
   DEFINE fn STRING
   DEFINE ch base.Channel
   LET ch = base.Channel.create()
   CALL ch.openFile(fn,"w")
   CALL ch.close()
END FUNCTION
```

Instead of creating an empty database file, it is also possible to prepare a template
(pre-configured) SQLite database file on the development platform, deploy the template database with
the other program files, and copy the template file from the program files directory (`base.Application.getProgramDir`) into the working directory
(`os.Path.pwd`) on the first application execution
(that is when the database file in the working directory does not yet exist):

```
IMPORT os
...
   CALL init_connection_strings()
   IF NOT prepare_database("template.dbs", source) THEN
      ERROR "Could not prepare database"
      EXIT PROGRAM 1
   END IF
   CONNECT TO connstr AS "c1"
   ...

FUNCTION prepare_database(template, target)
   DEFINE template, target STRING
   DEFINE tplpath STRING
   IF os.Path.exists(target) THEN
      RETURN TRUE
   END IF
   LET tplpath = os.Path.join(base.Application.getProgramDir(), template)
   IF NOT os.Path.exists(tplpath) THEN
      ERROR "Database template file not found"
      RETURN FALSE
   END IF
   RETURN os.Path.copy(tplpath, target)
END FUNCTION
```

> **Important:**
>
> When creating an initial database file in the working directory from a
> template file deployed in the program files directory, different file names should be used for the
> template and actual database file, as folders pointed to by `base.Application.getProgramDir` and `os.Path.pwd` may be the same on some devices.

## Related links

**Related concepts**  

[Directory structure for GMA apps](../17_mobile-applications/5104-directory-structure-for-gma-apps.md "Platform-specific rules need to be considered when deploying on Android devices (GMA).")

[Directory structure for GMI apps](../17_mobile-applications/5108-directory-structure-for-gmi-apps.md "Platform-specific rules need to be considered when deploying on iOS devices (GMI).")
