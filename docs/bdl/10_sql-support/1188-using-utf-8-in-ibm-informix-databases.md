---
title: "Using UTF-8 in IBM Informix databases"
source: "fgl-topics/c_fgl_odiagifx_003.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Fully supported IBM® Informix® SQL features > Using UTF-8 in IBM® Informix® databases"
type: "concept"
description: "UTF-8 in Genero programs and Informix databases Genero BDL programs can run in a UTF-8 locale, and use an IBM® Informix® database created with a UTF-8 locale. The locale of an Informix database is ..."
---

# Using UTF-8 in IBM Informix databases

## UTF-8 in Genero programs and Informix databases

Genero BDL programs can run in a UTF-8 locale, and use an IBM® Informix® database created with a
UTF-8 locale.

The locale of an Informix database is defined at creation by the DB\_LOCALE environment
variable.

At runtime, the Informix client locale is defined by the CLIENT\_LOCALE environment variable and
must match the [Genero application locale
(LANG/LC\_ALL)](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."), while DB\_LOCALE must match the database locale. If needed, Informix will make
charset conversions between CLIENT\_LOCALE and DB\_LOCALE. For more details about Informix client
setup, see [Prepare the runtime environment - connecting to the database](1184-prepare-the-runtime-environment-connecting-to-the-database.md).

When creating a database for UTF-8 storage:

- Define the DB\_LOCALE environment variable to a UTF-8 locale like
  `en_us.utf8`
- Check the `SQL_LOGICAL_CHAR` parameter in the onconfig
  Informix server file.

## Defining the Informix database locale

When you create a database with a `CREATE DATABASE` SQL statement, the current
value of the DB\_LOCALE environment variable defines the locale for the created database:

```
$ DB_LOCALE=en_US.utf8
$ export DB_LOCALE
$ dbaccess - -
> CREATE DATABASE mydb WITH BUFFERED LOG;
```

> **Important:**
>
> Always define the DB\_LOCALE environment variable,
> before executing the CREATE DATABASE statement: If DB\_LOCALE is not set, it will default to the
> DB\_LOCALE value set when starting the Informix engine. If no DB\_LOCALE was set when starting the
> engine, it defaults to `en_us.8859-1` (i.e. `en_us.819`).

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

## The `SQL_LOGICAL_CHAR` onconfig parameter

Informix (14.10) does not support CHAR length semantics, it uses only BYTE length semantics: By
default, in a `CHAR(10)` SQL type, the size represents a number of bytes. Since UTF-8
locale is a multi-byte, variable-length encoding, it is not possible to store 10 of all possible
UTF-8 characters in 10 bytes.

However, the Informix onconfig `SQL_LOGICAL_CHAR` parameter
can be defined as a size multiplier for `CHAR/VARCHAR` SQL types, when creating
tables. For example, when using a `CHAR(10)` SQL type and
`SQL_LOGICAL_CHAR=4`, the real SQL type will be `CHAR(40)`, where 40
is the size in bytes.

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

## Persistence of database settings

The DB\_LOCALE and `SQL_LOGICAL_CHAR` values used at database creation are
persistent properties for the lifetime of that database.

If you stop the Informix engine, change the `SQL_LOGICIAL_CHAR`, then restart the
engine, or set a different DB\_LOCALE, the locale settings of existing databases remain unchanged:
New created tables are based on the DB\_LOCALE and `SQL_LOGICAL_CHAR` settings when
the database was created.

## Database schema and CHAR/VARCHAR sizes

Program variables can be declared with [`DEFINE
LIKE`](../08_language-basics/0695-database-column-types.md "Simple variables and record structures can be defined from database columns types."), to use SQL types from .sch schema files.

Schema files are produced with the fgldbsch tool, reading Informix
(`syscolumns`) system tables. The sizes of `CHAR/VARCHAR` columns is
always expressed in bytes in `syscolumns.collength`.

For databases created with an `SQL_LOGICAL_CHAR` size multiplier, the
fgldbsch tool divides the `collength` by the
`SQL_LOGICAL_CHAR` value used by the database.

For example:

1. At database creation, onconfig server file has
   `SQL_LOGICAL_CHAR=4`
2. `CREATE TABLE tab1 ( col1 CHAR(10) )`
3. In the Informix system table: `syscolumns.collength = 40` (40 bytes)
4. fgldbsch -db mydb produces the .sch file (for
   `tab1.col1: 40/4=10`)
5. In .sch schema file, `tab1.col1` is defined as
   `CHAR(10)`
6. Source code declares: `DEFINE rec RECOR LIKE tab1.*`
7. The compiler reads schema file and defines `rec.col1` as
   `CHAR(10)`

At runtime, the actual storage capabilities of the `CHAR(10)` variable is then
defined by the FGL\_LENGTH\_SEMANTICS environment variable:

- When FGL\_LENGTH\_SEMANTICS=CHAR, a `CHAR(10)` variable can store 10 UTF-8
  characters, no matter the number of bytes used by these characters.
- When FGL\_LENGTH\_SEMANTICS=BYTE, a `CHAR(10)` variable can store 10 ASCII
  characters, 5 (5x2b=10b) UTF-8 accute characters, or 3 (3x3b=9b) UTF-8 Chinese characters.

## FGL\_LENGTH\_SEMANTICS: CHAR or BYTE ?

With UTF-8 in Genero programs, it is recommended to use char length semantics by setting the
[FGL\_LENGTH\_SEMANTICS=CHAR](../09_advanced-features/0881-length-semantics-settings.md) environment variable.

However, depending on the level of compatibility with Informix SQL length semantics, and the
complexity of character string operations in the program's code, you need to choose CHAR or BYTE
length semantics for Genero programs.

If Informix would support CHAR length semantics, there would be no question: We should just use
FGL\_LENGTH\_SEMANTICS=CHAR as well.

But Informix fundamentally uses BYTE length semantics (even when
`SQL_LOGICAL_CHAR` is defined). For example, in SQL statements:

- `LENGTH(expr)` returns a number of bytes
- `custname[2,5]` means the substrings from byte position 2 to 5

If the application code does not have complex string operations, you may consider using
FGL\_LENGTH\_SEMANTICS=BYTE, to get the same length semantics as in SQL statements.

But when the application code uses a lot of substring positions such as
`custcode[2,5]`, it is mandatory to use FGL\_LENGTH\_SEMANTICS=CHAR, to keep the source
code untouched.

> **Note:**
>
> The fgldbsch tool takes database's SQL\_LOGICAL\_CHAR setting into account: A
> column created as `CHAR(10)` will become a `CHAR(10)` variable from
> `DEFINE LIKE` (even if `syscolumns.collength` is defined as 40 bytes).
> Thus, the schema file is prepared for FGL\_LENGTH\_SEMANTICS=CHAR usage.

## The future of Informix UTF-8 support

Ideally, Informix should support CHAR length semantics like most other databases, to get the same
behavior with SQL and FGL expressions.

This feature request is known to Informix developers and should be available in a future version
(> 14.10).

## Related links

**Related concepts**  

[CHAR(size)](../08_language-basics/0557-char-size.md "The CHAR data type is a fixed-length character string data type.")

[VARCHAR(size)](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.")

[TEXT](../08_language-basics/0569-text.md "The TEXT data type stores large text data.")
