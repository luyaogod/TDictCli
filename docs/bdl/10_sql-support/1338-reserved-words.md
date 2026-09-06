---
title: "Reserved words"
source: "fgl-topics/c_fgl_odiagmys_006.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data manipulation > Reserved words"
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

## Oracle® MySQL and MariaDB

SQL object names like table and column names cannot be SQL reserved words in MySQL and
MariaDB.

## Solution

Table or column names which are MySQL reserved words must be renamed.
