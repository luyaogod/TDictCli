---
title: "Stored Procedures"
source: "fgl-topics/c_fgl_odiagifx_011.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > Stored Procedures"
type: "concept"
description: "With IBM® Informix® database servers, you can write stored procedures with the SPL (Stored Procedure Language) or with an external language in C or JAVA. If you plan to support different types of ..."
---

# Stored Procedures

With IBM® Informix® database servers, you can write
stored procedures with the SPL (Stored Procedure Language) or with
an external language in C or JAVA.

If you plan to support different types of database servers, you
must be aware that each DB vendor has defined its own stored procedure
language. In such cases, you may consider writing most of your business
logic in BDL, and implementing only some stored procedures in the
database, mainly to get better performance or to use database features
that only exist with stored procedures.

Genero BDL partially supports SP creation, but has full support
of SP invocation:

- The Genero BDL static SQL syntax does not include CREATE FUNCTION and CREATE PROCEDURE with a
  body block. However, you can create stored procedures with a body block by using dynamic SQL
  (EXECUTE IMMEDIATE), or with CREATE PROCEDURE and the FROM *filename* clause, which is
  supported by Genero BDL static SQL.
- The EXECUTE FUNCTION or EXECUTE PROCEDURE instruction is not allowed
  in the static SQL syntax. To invoke a stored procedure with Informix, you must use the
  PREPARE instruction, followed by EXECUTE or OPEN. The PREPARE instruction
  must initiate the EXECUTE FUNCTION/PROCEDURE instruction.

For more details about stored procedure invocation, see [SQL
Programming](0985-sql-basics.md "This section contains fundamental information to know about SQL programming with Genero BDL.").

## Related links

**Related concepts**  

[Stored procedure calls](1189-stored-procedure-calls.md "Stored procedure calls")

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")

[Dynamic SQL management](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")

[PREPARE (SQL statement)](1142-prepare-sql-statement.md "Prepares an SQL statement for execution.")
