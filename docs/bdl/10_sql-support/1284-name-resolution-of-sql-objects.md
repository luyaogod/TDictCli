---
title: "Name resolution of SQL objects"
source: "fgl-topics/c_fgl_odiagmsv_025.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > Name resolution of SQL objects"
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

## Microsoft™ SQL Server

With Microsoft SQL Server, an object
name takes the following form:

```
[[database.]owner.]identifier
```

Object names are limited to 128 characters in SQL Server and cannot
start with one of the following characters: @ (local variable) # (temp
object).

To support double quotes as string delimiters in SQL Server, you can switch OFF the database
option "Use quoted identifiers" in the database properties panel. But quoted table and column names
are not supported when this option is OFF.

## Solution

To write portable SQL, regarding database object names:

1. Use simple database object names (without any owner/schema prefix)
2. Do not use double quotes to surround database object identifiers.
3. If needed, define public synonyms to reference database objects in others databases/schema.
4. Specify database object identifiers in lowercase.

See also [Naming database objects](1034-naming-database-objects.md).

With Microsoft SQL Server, use a "CS" case-sensitive
collation at server and database level. See [Case sensitivity](1281-case-sensitivity.md) for more
details.

Check
for single or double quoted table or column names in your source and
remove them.

If cross-database queries are required (using the `dbname:tabname` Informix
notation), consider to create views in the main database in SQL Server, to allow program to access
the distant table from the same database
connection:

```
CREATE VIEW myotherdb_customers AS SELECT * FROM myotherdb.dbo.customers
```

## Related links

**Related concepts**  

[Database concepts](1265-database-concepts.md "Database concepts")
