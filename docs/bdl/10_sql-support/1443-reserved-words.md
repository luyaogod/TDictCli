---
title: "Reserved words"
source: "fgl-topics/c_fgl_odiagpgs_006.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data manipulation > Reserved words"
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

## PostgreSQL

SQL object names like table and column names can be SQL reserved words in PostgreSQL.

However, this is not good practice and must be avoided.

## Solution

Table or column names which are
PostgreSQL reserved words must be renamed.

For a list of reserved PostgreSQL keywords, see <https://www.postgresql.org/docs/current/sql-keywords-appendix.html>.
