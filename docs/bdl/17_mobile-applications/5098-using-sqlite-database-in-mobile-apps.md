---
title: "Using SQLite database in mobile apps"
source: "fgl-topics/c_fgl_mobile_using_sqlite.html"
breadcrumb: "Mobile applications > Database support on mobile devices > Using SQLite database in mobile apps"
type: "concept"
---

# Using SQLite database in mobile apps

> On the device, Genero Mobile uses the SQLite database only.

## Running an app in development mode

When running an app in development mode (where the app runs on a computer), you can
use any database server that Genero supports for the operating system of the
server-side app.

## Running an app on a mobile device

When running the application on the device, only SQLite can be used. The database
must be created at the first application execution, or it must be delivered as the
default database in the .ipa or .apk
package.

## Locale character set and length semantics

SQLite stores data in UTF-8 codeset, mobile apps will by default run in UTF-8 and
with character length semantics (`FGL_LENGTH_SEMANTICS=CHAR`).

## Creating the database

Mobile applications usually create their database at first execution. The SQLite
database file must be created in the application sandbox, in a writable directory.
If the database file does not exist in the current working directory
(`os.Path.pwd()`), create an empty file and then perform a
`CONNECT TO` instruction.

For more details, see [Creating a database from programs](../10_sql-support/1003-creating-a-database-from-programs.md "Creating a database from within a program requires special consideration.").

## Providing a default database

SQLite database file format is cross-platform. Instead of creating the database the
first time the application starts, you might want to prepare a default database file
in your development environment, and include it in the .ipa /
.apk package.

## Data types with SQLite

SQLite does not have strict data type checking as traditional databases. If you
define a table column as a `DECIMAL`, you can still store character
values in that column. Pay attention to this SQLite specific feature, to avoid
invalid storage and type conversion errors in your application.

Consider using the following data types for maximum portability, especially when data
needs to be synchronized with a central database server, where the data types must
match to the types used in the mobile application: `CHAR`,
`VARCHAR`, `DATE`, `DATETIME YEAR TO
MINUTE`, `DATETIME YEAR TO FRACTION(3)`,
`DECIMAL`, `SMALLINT`, `INTEGER`,
`BIGINT`, `BYTE`, `TEXT`.

## Optimizing data changes

SQLite can be slow at doing commits, due to the data integrity technique used for
transactions. Since each `INSERT` / `UPDATE` /
`DELETE` statements acts as an individual transaction (i.e.,
auto-commit), there will be as many transactions/commits as data manipulation
statements.

When executing code that modifies a lot of rows (for example, when inserting default
data at first application execution, or when doing synchronization with a central
database), enclose the SQL statements within a `BEGIN WORK` /
`COMMIT WORK` transaction block to speed up the
process:

```
BEGIN WORK
FOR i=1 TO mylog.getLength()
    -- INSERT / UPDATE / DELETE statements
END FOR
COMMIT WORK
```

## Enforcing foreign key constraints

SQLite 3.6.19 and + support foreign key constraints, with `ON DELETE
CASCADE` and `ON UPDATE CASCADE` options. By default,
however, foreign key constraints are not enforced. Each application must explicitly
turn on the feature with a `PRAGMA` command. Immediately after the
database connection, you can perform the `PRAGMA` command in an
`EXECUTE IMMEDIATE` statement:

```
CONNECT TO connstr AS "c1"
EXECUTE IMMEDIATE "PRAGMA foreign_keys = ON"
```

## Truncating the SQLite database file

By default, when deleting rows, SQLite keeps the unused database file pages for
future storage. As result, when deleting a large amount of data, the database file
might be larger than necessary. Consider truncating the database file with the
`VACUUM SQL` command (in an `EXECUTE IMMEDIATE`
statement), if disk space is limited and when a lot of database rows were
deleted.

Depending on the application, the `VACUUM` command can be
executed:

- when starting the application, just after connecting to the database,
- after doing a large database operation (such as a synchronization with a central
  database),
- as a manual option that the user can trigger.

For example, after connecting to the
database:

```
CONNECT TO connstr AS "c1"
EXECUTE IMMEDIATE "VACUUM"
```

## Sharing database files between Android apps

Two different Android™
apps (each packaged as a separate .apk) execute in their own
sandbox, but have access to the storage area (SD-CARD) and therefore can share a
common database file.

SQLite handles concurrent access to the same database file by setting a lock on the
entire db file when modifying data
(`INSERT`/`UPDATE`/`DELETE`). By
default, if a writer process locks the file, other processes must wait until the
lock owner process completes its transaction and releases the lock.

Because of Informix®
compatibility, Genero BDL uses a default lock timeout or zero (i.e., not waiting for
locks to be released). As result, when writing to a database file that is locked by
another process, if the isolation level is `SERIALIZATION` (the
default with SQLite), an application will get the SQL error -244.

To avoid this problem, you must change the default lock timeout with the `SET
LOCK MODE` instruction, after starting the database
session:

```
CONNECT TO connstr AS "c1"
SET LOCK MODE TO 5 -- seconds
```

The second process will then wait until the first process releases the lock. If
transactions are short (milliseconds), having processes waiting for each other is
transparent to the user.

## Related links

**Related concepts**  

[SQL programming](../10_sql-support/0984-sql-programming.md "Covers topics about interacting with a database server using SQL.")

[SQLite](../10_sql-support/1466-sqlite.md "SQLite")
