---
title: "SQL errors on PREPARE"
source: "fgl-topics/c_fgl_odiagora_014.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > BDL programming > SQL errors on PREPARE"
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

## ORACLE

The Oracle® interface is implemented
with the Oracle Call Interface (OCI). This
library does not provide a way to send SQL statements to the database server during the BDL
`PREPARE` instruction, as in the Informix
interface.

When preparing an SQL statement with
the BDL `PREPARE` or `DECLARE` instruction, no SQL error will be
returned if the SQL statement is invalid. However, an SQL error will occur after the
`OPEN` / `FOREACH` / `EXECUTE` instructions.

## Solution

Make sure your BDL programs do not test the `status` or
`sqlca.sqlcode` variable just after `PREPARE` instructions.

Change the program logic in order to handle the SQL errors when opening the cursors
(`OPEN`) or when executing SQL statements (`EXECUTE`).

## Related links

**Related concepts**  

[The sqlca diagnostic record](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.")
