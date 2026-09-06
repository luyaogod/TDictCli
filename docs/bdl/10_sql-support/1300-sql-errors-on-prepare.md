---
title: "SQL errors on PREPARE"
source: "fgl-topics/c_fgl_odiagmsv_039.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > BDL programming > SQL errors on PREPARE"
type: "concept"
description: "Informix® With Informix, a PREPARE instruction returns an SQL error in case of problem: TRY PREPARE stmt FROM \"SELECT * FROM WHERE pk=1\" -- table is missing! CATCH DISPLAY \"SQL ERROR:\", sqlca.sqlcode ..."
---

# SQL errors on PREPARE

## Informix®

With Informix, a `PREPARE` instruction
returns an SQL error in case of
problem:

```
TRY
   PREPARE stmt FROM "SELECT * FROM WHERE pk=1"  -- table is missing!
CATCH
   DISPLAY "SQL ERROR:", sqlca.sqlcode
END TRY
```

## Microsoft™ SQL Server

The Genero database drivers for Microsoft SQL Server
are based on ODBC drivers (SQL Native Client, MS ODBC SQL, FreeTDS or Easysoft drivers). These ODBC
drivers use system stored procedures to prepare and execute SQL statements (you can check this with
the SQL Server Profiler).

With SQL server, `PREPARE` or `DECLARE` instructions do not return
an SQL error, because the statement preparation is deferred to the execution, improving overall
performances.

When preparing an SQL statement with
the BDL `PREPARE` or `DECLARE` instruction, no SQL error will be
returned if the SQL statement is invalid. However, an SQL error will occur after the
`OPEN` / `FOREACH` / `EXECUTE` instructions.

## Solution

Make sure your BDL programs do not test the `status` or
`sqlca.sqlcode` variable just after `PREPARE` instructions.

Change the program logic in order to handle the SQL errors when opening the cursors
(`OPEN`) or when executing SQL statements (`EXECUTE`).
