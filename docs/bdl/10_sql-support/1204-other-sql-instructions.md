---
title: "Other SQL instructions"
source: "fgl-topics/c_fgl_odiagifx_017.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > Other SQL instructions"
type: "concept"
description: "Genero BDL static SQL syntax implements common Data Manipulation Statements such as SELECT, INSERT, UPDATE and DELETE. Data Definition Language statements such as CREATE TABLE, CREATE INDEX, CREATE ..."
---

# Other SQL instructions

Genero BDL static SQL syntax implements common Data Manipulation
Statements such as SELECT, INSERT, UPDATE and DELETE. Data Definition
Language statements such as CREATE TABLE, CREATE INDEX, CREATE SEQUENCE
and their corresponding ALTER and DROP statements are also part of
the static SQL grammar. These are supported with a syntax limited
to the standard SQL clauses. For example, Genero BDL might not support
the most recent CREATE TABLE storage options supported by IBM® Informix® SQL.

Since the first days of the Informix 4GL language,
the SQL language has been extended, and it has become so large that it's impossible to embed
all the existing new statements without introducing grammar conflicts
with the I4GL language. In addition, each DB vendor has improved the
standard SQL language with proprietary SQL statements that are not
portable; it would not be a good idea to use these specific instructions
if you plan to make your application run with different types of database
engines.

However, the Genero BDL static SQL is constantly being improved with standard SQL syntax that
works with most types of database servers. For example, Genero BDL supports the ANSI outer join
syntax, constraints definition in DDL statements, sequence instructions, BIGINT and BOOLEAN data
types, and there is more to come.

If a statement is unsupported in static SQL, that does not mean
that you cannot execute it. If you want to execute an SQL instruction
that is not part of the static SQL grammar, you can use [Dynamic SQL](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.") as
follows:

- Use PREPARE + EXECUTE for statements that do not generated a result
  set
- Use (PREPARE/) DECLARE + OPEN for statements returning a result
  set
- Use EXECUTE IMMEDIATE if no SQL parameters are required and no
  result set is generated

Dynamic SQL instructions take a string as the input, so there is
no limitation regarding the SQL text you can execute. However, only
one statement can be executed at a time. Preferably, write your SQL
statements directly in static SQL when possible, because it makes
the code more readable and the syntax is checked at compiled time.

For more details about statements supported in the static SQL syntax,
see [Static
SQL](1115-static-sql-statements.md "Describes static SQL statements supported in the language.").

Below is a list of the IBM Informix SQL statements that
are not allowed in the static SQL syntax (last updated from IDS  **11.50** SQL
instructions). The IBM Informix SQL Syntax manual
includes ESQL/C specific statements such as ALLOCATE DESCRIPTOR, which
are not part of the basic SQL statements supported by the engines.
ESQL/C specific statements are not listed here:

```
ALTER ACCESS_METHOD 
ALTER FRAGMENT 
ALTER FUNCTION 
ALTER PROCEDURE 
ALTER ROUTINE 
ALTER SECURITY LABEL COMPONENT 
CREATE ACCESS_METHOD 
CREATE AGGREGATE 
CREATE CAST 
CREATE DISTINCT TYPE 
CREATE EXTERNAL TABLE Statement 
CREATE FUNCTION (with body)
CREATE OPAQUE TYPE 
CREATE OPCLASS 
CREATE PROCEDURE (with body)
CREATE ROLE 
CREATE ROUTINE FROM 
CREATE ROW TYPE 
CREATE SCHEMA 
CREATE SECURITY LABEL 
CREATE SECURITY LABEL COMPONENT 
CREATE SECURITY POLICY 
CREATE TRIGGER 
CREATE VIEW 
CREATE XADATASOURCE 
CREATE XADATASOURCE TYPE 
DROP ACCESS_METHOD 
DROP AGGREGATE 
DROP CAST 
DROP FUNCTION 
DROP OPCLASS 
DROP PROCEDURE 
DROP ROLE 
DROP ROUTINE 
DROP ROW TYPE 
DROP SECURITY 
DROP TRIGGER 
DROP TYPE 
DROP XADATASOURCE 
DROP XADATASOURCE TYPE 
EXECUTE FUNCTION 
EXECUTE PROCEDURE 
GRANT FRAGMENT 
INFO 
MERGE 
OUTPUT 
RELEASE SAVEPOINT 
RENAME COLUMN 
RENAME DATABASE 
RENAME SECURITY 
REVOKE FRAGMENT 
SAVE EXTERNAL DIRECTIVES 
SAVEPOINT 
SET AUTOFREE 
SET COLLATION 
SET CONSTRAINTS 
SET DATASKIP 
SET DEBUG FILE 
SET ENCRYPTION PASSWORD 
SET ENVIRONMENT 
SET INDEXES 
SET LOG 
SET OPTIMIZATION 
SET PDQPRIORITY 
SET ROLE 
SET SESSION AUTHORIZATION 
SET STATEMENT CACHE 
SET TRANSACTION 
SET TRIGGERS 
START VIOLATIONS TABLE 
STOP VIOLATIONS TABLE
```
