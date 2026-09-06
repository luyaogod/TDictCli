---
title: "Reserved words"
source: "fgl-topics/c_fgl_odiagmsv_006.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data manipulation > Reserved words"
type: "concept"
description: "Informix® With Informix, it is possible to create database objects with reserved words. For example: CREATE TABLE table ( char CHAR(10) ); Indeed this is not good practice, but Informix SQL allows ..."
---

# Reserved words

## Informix®

With Informix, it is possible to create database
objects with reserved words.

For example:

```
CREATE TABLE table ( char CHAR(10) );
```

Indeed this is not good practice, but Informix SQL
allows this to be backward compatible when introducing a new keyword in the SQL syntax.

Most other database systems do not allow reserved words as database identifiers. If your legacy
code is using SQL reserved words of the target database SQL syntax, an error will be thrown at
`CREATE TABLE` execution.

## Microsoft™ SQL Server

Microsoft Transact-SQL does not allow you to use
reserved words as database object names (tables, columns, constraint, indexes, triggers, stored
procedures, ...).

An example of a common word which is part of SQL Server grammar is '`go`' (see the
'Reserved keywords' section in the SQL Server Documentation).

## Solution

Database objects having a name
which is a Transact-SQL reserved word must be renamed.

All BDL
application sources must be verified. To check if a given keyword
is used in a source, you can use UNIX™ 'grep'
or 'awk' tools. Most modifications can be automatically done with UNIX tools like 'sed' or 'awk'.

You can use `SET QUOTED_IDENTIFIER ON` with double-quotes to enforce the use of
keywords in the database objects naming, but it is not recommended.
