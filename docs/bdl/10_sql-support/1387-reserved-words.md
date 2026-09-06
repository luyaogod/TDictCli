---
title: "Reserved words"
source: "fgl-topics/c_fgl_odiagora_006.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > Reserved words"
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

## ORACLE

SQL object names like table and column names cannot be SQL reserved words in ORACLE.

An example of a common word which is part of the ORACLE SQL grammar is 'level'.

## Solution

Table or column names which are
ORACLE reserved words must be renamed.

ORACLE reserved keywords
are listed in the ORACLE documentation, or Oracle 8i provides the
V$RESERVED\_WORDS view to track Oracle reserved words. All BDL application
sources must be verified. To check if a given keyword is used in a
source, you can use UNIX™ 'grep'
or 'awk' tools. Most modifications can be done automatically with UNIX tools like 'sed' or 'awk'.
