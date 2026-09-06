---
title: "Name resolution of SQL objects"
source: "fgl-topics/c_fgl_odiagsqt_028.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Data dictionary > Name resolution of SQL objects"
type: "concept"
description: "Informix® Informix uses the following form to identify an SQL object: database[@dbservername]:][{owner|\"owner\"}.]identifier The ANSI convention is to use double quotes for identifier delimiters (For ..."
---

# Name resolution of SQL objects

## Informix®

Informix uses the following form to identify an SQL
object:

```
database[@dbservername]:][{owner|"owner"}.]identifier
```

The ANSI convention is to use double quotes for identifier delimiters (For example:
`"customer"."cust_name"`).

Informix
database object names are not case-sensitive in non-ANSI databases. When using double-quoted
identifiers, Informix becomes case sensitive.

With non-ANSI Informix databases, you do not have to
give a schema name before the tables when executing an SQL
statement:

```
SELECT ... FROM customer WHERE ...
```

In
Informix ANSI compliant databases:

- The table name must include "owner", unless the connected user is the owner of the database
  object.
- The database server shifts the owner name to uppercase letters before the statement executes,
  unless the owner name is enclosed in double quotes.

## SQLite

SQLite database object names are case-insensitive. Using double quotes to surround table names in
possible. However, the letter case is kept even without double
quotes:

```
sqlite> CREATE TABLE tab1 ( pk INT );
sqlite> CREATE TABLE "TAB2" ( pk INT );
sqlite> CREATE TABLE Tab3 ( pk INT );
sqlite> .tables
TAB2  Tab3  tab1
sqlite> .schema "TAB3" 
CREATE TABLE Tab3 ( pk INT );
```

In an SQLite, if a prefix is specified as part of an object reference, it must be either "main",
or "temp" or the schema-name of an attached database. There is no such concept as user schema in
SQLite.

## Solution

To write portable SQL, regarding database object names:

1. Use simple database object names (without any owner/schema prefix)
2. Do not use double quotes to surround database object identifiers.
3. If needed, define public synonyms to reference database objects in others databases/schema.
4. Specify database object identifiers in lowercase.

See also [Naming database objects](1034-naming-database-objects.md).
