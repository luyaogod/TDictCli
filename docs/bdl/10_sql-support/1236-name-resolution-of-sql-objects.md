---
title: "Name resolution of SQL objects"
source: "fgl-topics/c_fgl_odiagdmg_023.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data dictionary > Name resolution of SQL objects"
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

## Dameng®

Dameng database object names case-sensitivity is defined
by the `CASE_SENSITIVE` database parameter.

In a Dameng database, tables always belong to a database schema. When executing an
SQL statement, a schema name must be used as the high-order part of a two-part object name, unless
the current schema corresponds to the table's schema.

The default (implicit) schema is the current user's name but it can be changed with the
`ALTER SESSION SET CURRENT_SCHEMA = name` instruction.

## Solution

To write portable SQL, regarding database object names:

1. Use simple database object names (without any owner/schema prefix)
2. Do not use double quotes to surround database object identifiers.
3. If needed, define public synonyms to reference database objects in others databases/schema.
4. Specify database object identifiers in lowercase.

See also [Naming database objects](1034-naming-database-objects.md).

The database parameter `CASE_SENSITIVE` defines if character string comparison is
case-sensitive (Y) or case-insensitive (N). Depending on the application needs for data search, this
parameter will condition the way table and column names must be written. As a general portable SQL
pattern, always use lowercase table and column names, without double quote delimiters.

The Dameng schema concept:

After a connection, the database interface can automatically execute a `ALTER SESSION SET
CURRENT_SCHEMA name` instruction if the following FGLPROFILE entry is
defined:

```
dbi.database.dbname.dmg.schema= "name"
```

Here dbname identifies the database name used in the BDL program
(`DATABASE dbname`) and name is the schema name
to be used during this SQL session. If this entry is not defined, no `ALTER SESSION`
instruction is executed and the current schema defaults to the user's name.

Examples:

```
dbi.database.stores.dmg.schema= "STORES1"
dbi.database.accnts.dmg.schema= "ACCSCH"
```
