---
title: "Database concepts"
source: "fgl-topics/c_fgl_odiagsqt_009.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Database concepts > Database concepts"
type: "concept"
description: "Informix® servers can handle multiple database entities, while SQLite can manage several database files. Tip: If you have several Informix database entities, migrating from the Informix database to ..."
---

# Database concepts

Informix® servers can handle multiple database
entities, while SQLite can manage several database files.

> **Tip:**
>
> If you have several Informix database entities,
> migrating from the Informix database to another database it is a good opportunity to centralize all
> tables in a single database. To avoid conflicts with table names, use a prefix when needed.

## Solution

Map each Informix database to a SQLite database
file.

Consider creating the SQLite database file before using the connection instruction. The database
file can be created as an empty file, with a OS shell command (touch) or by
program by using the file utility classes ([`base.Channel`](../15_library-reference/2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")).

> **Important:**
>
> On Microsoft™ Windows®, the path to an SQLite database file must be specified by using
> forward slash characters as directory separator. Using backslash characters is not supported and
> would anyway bring confusion, as these must be escaped.

It is possible to specify an SQLite database filename in the database specification in
`CONNECT TO` or `DATABASE`
instructions:

```
CONNECT TO "mydb+driver='dbmsqt',source='/opt/myapp/database/stock1.dbs'"
```

However, it is recommended to use an indirection by providing an abstract name identifier in the
program, and by defining the real database file with the ["`source`"](1072-database-source-specification-source.md) connection parameter. The file defined by
"`source`" is then found directly (can be a relative or absolute path), or by using
[DBPATH](../07_configuration/0513-dbpath.md "Defines a list of paths for Genero program resource files.") settings, if not found from the current
directory of fglrun (when it's not an absolute path).

In the program:

```
DATABASE stock
```

In the FGLPROFILE configuration file, define the SQLite driver and the database file:

```
dbi.database.stock.driver = "dbmsqt"
dbi.database.stock.source = "/opt/myapp/database/stock1.dbf"
```

Or if in FGLPROFILE, you define the filename
only:

```
dbi.database.stock.source = "stock1.dbf"
```

The file is found by using
DBPATH:

```
DBPATH="/opt/myapp/database"
```

When specifying `:memory:` as database filename, an empty SQLite database is
created in memory. This can be useful if the persistence of the data is not required after the
program has terminated:

```
DATABASE ":memory:"
```
