---
title: "SQL errors on PREPARE"
source: "fgl-topics/c_fgl_odiagpgs_014.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > BDL programming > SQL errors on PREPARE"
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

## PostgreSQL

The PostgreSQL database driver is implemented with the PostgreSQL libpq API. This library does
not provide a way to send SQL statements to the database server during the BDL
`PREPARE` instruction, like the Informix
interface does.

When preparing an SQL statement with
the BDL `PREPARE` or `DECLARE` instruction, no SQL error will be
returned if the SQL statement is invalid. However, an SQL error will occur after the
`OPEN` / `FOREACH` / `EXECUTE` instructions.

## Solution

Make sure your BDL programs do not test the `status` or
`sqlca.sqlcode` variable just after `PREPARE` instructions.

Change the program logic in order to handle the SQL errors when opening the cursors
(`OPEN`) or when executing SQL statements (`EXECUTE`).
