---
title: "Name resolution of SQL objects"
source: "fgl-topics/c_fgl_odiagora_028.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > Name resolution of SQL objects"
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

## ORACLE

With Oracle®, an object name takes the following
form:

```
[(schema|"schema").](identifier|"identifier")[@database-link]
```

Oracle has separate namespaces for different
classes of objects (tables, views, triggers, indexes, clusters).

Object names are limited to
30 chars in ORACLE.

Unlike Informix, Oracle database object names are stored in
UPPERCASE in system catalogs. That means that `SELECT "col1" FROM "tab1"` will
produce an error because those objects are identified by "`COL1`" and
"`TAB1`" in Oracle system
catalogs.

An Oracle database schema is
owned by a user (usually, the application administrator) and this user must create `PUBLIC
SYNONYMS` to provide a global scope for his table names. `PUBLIC SYNONYMS`
can have the same name as the schema objects they point to.

## Solution

To write portable SQL, regarding database object names:

1. Use simple database object names (without any owner/schema prefix)
2. Do not use double quotes to surround database object identifiers.
3. If needed, define public synonyms to reference database objects in others databases/schema.
4. Specify database object identifiers in lowercase.

See also [Naming database objects](1034-naming-database-objects.md).

Without double quotes around the database object names, all names will be converted to uppercase
letters by ORACLE before executing the SQL.

Check that you do not use single-quoted or double-quoted table names or column names in your
source. Those quotes must be removed because the database interface automatically
converts double quotes to single quotes, and Oracle does not allow single quotes as database object
name delimiters.

See also the issue [Database Concepts](1365-database-concepts.md)
