---
title: "Database Triggers"
source: "fgl-topics/c_fgl_odiagifx_012.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > Database Triggers"
type: "concept"
description: "Triggers can be created for IBM® Informix® database tables with the CREATE TRIGGER instruction. If you plan to support different types of database servers, you must be aware that each DB vendor has ..."
---

# Database Triggers

Triggers can be created for IBM® Informix® database tables with
the CREATE TRIGGER instruction.

If you plan to support different types of database servers, you must be aware that each DB vendor
has defined its own trigger creation syntax and stored procedure language. In such cases, you may
consider writing most of your business logic in BDL, and implementing only some triggers in the
database, mainly to get better performance or use database features that only exist with stored
procedures.

Genero BDL partially supports trigger creation:

- The Genero BDL static SQL syntax does not include the CREATE TRIGGER
  and DROP TRIGGER instructions. However, you can create database triggers
  by using dynamic SQL (EXECUTE IMMEDIATE).

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")

[EXECUTE IMMEDIATE](1145-execute-immediate.md "Performs a simple SQL execution without SQL parameters or result set.")
