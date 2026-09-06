---
title: "Name resolution of SQL objects"
source: "fgl-topics/c_fgl_odiagpgs_025.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > Name resolution of SQL objects"
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

## PostgreSQL

With PostgreSQL, a database object name takes the following
form:

```
[[database.]schema.]identifier
```

You can define a list of schemas to be searched when the database object is not qualified. The
list of schemas to look in is specified with `SET search_path TO
schema-1,schema-2,...` for an SQL session (it can also
be defined at the database level and user definition level).

When creating a table without a schema prefix, the table is created in the
"`public`" schema. This schema is listed in the search path by default.

## Solution

To write portable SQL, regarding database object names:

1. Use simple database object names (without any owner/schema prefix)
2. Do not use double quotes to surround database object identifiers.
3. If needed, define public synonyms to reference database objects in others databases/schema.
4. Specify database object identifiers in lowercase.

See also [Naming database objects](1034-naming-database-objects.md).

Use the following FGLPROFILE entry to define the schema search path for programs connecting to
PostgreSQL:

```
dbi.database.stores.pgs.schema = "\"$user\",public,stock"
```

For more details see [PostgreSQL specific FGLPROFILE parameters](1084-postgresql-specific-fglprofile-parameters.md).
