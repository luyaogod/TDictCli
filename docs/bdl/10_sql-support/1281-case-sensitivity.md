---
title: "Case sensitivity"
source: "fgl-topics/c_fgl_odiagmsv_041.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > Case sensitivity"
type: "concept"
description: "With Informix®, database object names like table and column names are not case sensitive: CREATE TABLE Customer ( CustNo INTEGER, CustName VARCHAR(50) ) SELECT custno FROM customer -- All lowercase In ..."
---

# Case sensitivity

With Informix®, database object names like table and
column names are not case sensitive:

```
CREATE TABLE Customer ( CustNo INTEGER, CustName VARCHAR(50) )
SELECT custno FROM customer   -- All lowercase
```

In Microsoft™ SQL Server (2019 and older),
case-sensitivity of database object names and character data comparison in SQL are strongly tied.

The case-sensitivity is defined by the "CS" or "CI" option in the collation. Collation can be
defined at different levels: server instance, database, and even column and expression level.

The server-level collation is defined at installation and matters for example when querying
system tables or when using temporary tables. Therefore, it is important to consider the character
case sensitivity when installing SQL Server.

See SQL Server documentation for more details and the "Collation" concept.

For example, with a case-sensitive SQL Server collation, execute the following SQL statements in
the SQL Server query tool:

```
CREATE TABLE Customer ( CustNo INTEGER, CustName VARCHAR(50) )
INSERT INTO Customer VALUES ( 1, 'TECHNOSOFT' )             -- Row is inserted
SELECT CustNo FROM Customer WHERE CustName = 'TECHNOSOFT'   -- Row is found
SELECT CustNo FROM Customer WHERE CustName = 'TechnoSoft'   -- Row is NOT found
INSERT INTO CUSTOMER VALUES ( 2, 'FOURJS' )                 -- SQL Error: Invalid table name
SELECT custno FROM Customer WHERE CustName = 'TECHNOSOFT'   -- SQL Error: Invalid column name
```

Genero compilers convert table and column names to lower case.
For example, when writing following static SQL statement:

```
SELECT COUNT(*) FROM CUSTOMER WHERE CUSTNAME LIKE 'S%'
```

The SQL text stored in the pcode module will be:

```
SELECT COUNT(*) FROM customer WHERE custname LIKE 'S%'
```

## Solution

Usually, it is preferable to have case-sensitive character comparison in SQL Server, for example
to make the string `'ABC'` different from `'abc'`.

Since SQL Server does not have an option to distinguish character case sensitivity in SQL
expressions and for database object names, the table and column names in SQL statements must
exactly match the names as defined in the database.

Select a "CS" case-sensitive collation when installing SQL Server and when creating databases,
to make queries case-sensitive, and define the database tables and columns in lower case only,
because Genero compilers convert them to lower case.

See also [Name resolution of SQL objects](1284-name-resolution-of-sql-objects.md).
